// Binary report-data-usage-server sends info to your SIM about how much data
// has been used.
//
// The server is primarily designed to respond to incoming webhooks from Twilio.
// Make a POST request to the server and specify a Command parameter and
// a "SimSid" parameter. The only supported Command is "Usage". The SimSid
// should be a sim for your account.
//
// The server will send a Command to your phone using the Commands API. If
// successful, the server will respond with a 204.
//
// Example report format:
//
//     $ report-data-usage
//     iPhone v13
//     ----------
//     2017-06-19: 0
//     2017-06-20: 0
//     2017-06-21: 11.281MB
//     2017-06-22: 93.341MB
//     2017-06-23: 94.422MB
//     2017-06-24: 159.461MB
//     2017-06-25: 50.062MB
//     total (last 7 days): 408.568MB
//
// Your Twilio credentials are loaded from the TWILIO_ACCOUNT_SID and
// TWILIO_AUTH_TOKEN environment variables. There are several flags:
//
//     --days int
//         Change the number of days to report usage for (default 7)
//     --location string
//         Use a different timezone for day boundaries (example "America/Los_Angeles")
//     --sim string
//         Only fetch usage for this sim
//     --dry-run bool
//         Dry run mode (don't send commands back to the device)
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	types "github.com/kevinburke/go-types"
	"github.com/kevinburke/handlers"
	"github.com/kevinburke/rest"
	twilio "github.com/appian/twilio-go"
	"github.com/appian/twilio-go/datausage"
	"golang.org/x/sync/errgroup"
)

const Version = "0.1"

var duration = flag.Uint("days", 7, "Number of days to get usage for")
var location = flag.String("location", "", "Timezone to use (defaults to system location/TZ env var)")
var sim = flag.String("sim", "", "Load data about a specific Sim/unique name only")
var dryRun = flag.Bool("dry-run", false, "Dry run mode")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Start a server that prints information about how much data you used on each day.
`)
		flag.PrintDefaults()
	}
}

func create(ctx context.Context, c *twilio.Client, simSid string, txt string) error {
	if *dryRun {
		fmt.Fprintf(os.Stderr, "would have sent command:\n%s\n\nto sid %q\n\n", txt, simSid)
		return nil
	}
	_, err := c.Wireless.Commands.Create(ctx, url.Values{
		"Sim":     []string{simSid},
		"Command": []string{txt},
	})
	return err
}

func main() {
	flag.Parse()
	var loc *time.Location
	if *location == "" {
		loc = time.Local
	} else {
		var err error
		loc, err = time.LoadLocation(*location)
		if err != nil {
			log.Fatal(err)
		}
	}
	c := twilio.NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)
	iter := c.Wireless.Sims.GetPageIterator(nil)
	sims := make([]string, 0)
	if *sim != "" {
		sims = []string{*sim}
	} else {
		// TODO, maybe need to look this up on a loop every hour, or something.
		// for now just running this as a demo.
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 31*time.Second)
			page, err := iter.Next(ctx)
			cancel()
			if err == twilio.NoMoreResults {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(page.Sims); i++ {
				sims = append(sims, page.Sims[i].UniqueName)
			}
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			rest.NotAllowed(w, r)
			return
		}
		// https://www.twilio.com/docs/api/wireless/rest-api/command#list-post
		cmd := r.PostFormValue("Command")
		if strings.ToLower(cmd) != "usage" {
			// TODO figure out how to multiplex with other commands here.
			rest.BadRequest(w, r, &rest.Error{
				Title: fmt.Sprintf("unknown command %q", cmd),
			})
			return
		}
		simSid := r.PostFormValue("SimSid")
		if !strings.HasPrefix(simSid, "DE") {
			rest.BadRequest(w, r, &rest.Error{
				Title: fmt.Sprintf("unknown sim sid %q", simSid),
			})
			return
		}
		days := *duration
		if d := r.URL.Query().Get("days"); d != "" {
			daysInt, err := strconv.ParseUint(d, 10, 64)
			if err != nil {
				rest.BadRequest(w, r, &rest.Error{
					Title: err.Error(),
				})
				return
			}
			days = uint(daysInt)
		}
		end := time.Now().Add(24 * time.Hour).In(loc)
		end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
		start := end.Add(-24 * time.Hour * time.Duration(days+1))
		var mu sync.Mutex
		ctx, cancel := context.WithTimeout(r.Context(), 31*time.Second)
		defer cancel()
		group, errctx := errgroup.WithContext(ctx)
		usage := make(map[string][]types.Bits)
		for _, name := range sims {
			name := name
			group.Go(func() error {
				u, err := datausage.GetUsage(errctx, c, name, start, end, 24*time.Hour)
				if err != nil {
					return err
				}
				mu.Lock()
				usage[name] = u
				mu.Unlock()
				return nil
			})
		}
		if err := group.Wait(); err != nil {
			rest.ServerError(w, r, err)
		}
		buf := new(bytes.Buffer)
		for _, name := range sims {
			total := types.Bits(0)
			fmt.Fprintf(buf, "%s\n%s\n", name, strings.Repeat("-", len(name)))
			for i := 0; i < len(usage[name]); i++ {
				t := start.Add(time.Duration(i) * 24 * time.Hour)
				fmt.Fprintf(buf, "%s: %s\n", t.Format("2006-01-02"), usage[name][i])
				total += usage[name][i]
			}
			fmt.Fprintf(buf, "total (last %d days): %s", days, total)
		}
		out := buf.String()
		var err error
		for i := 0; i < len(out); i++ {
			if len(out)-i < 160 {
				// send the whole thing
				err = create(ctx, c, simSid, out[i:])
				i = len(out)
			} else {
				idx := strings.LastIndexByte(out[i:i+160], '\n')
				if idx == -1 || idx == 0 {
					idx = 160
				}
				err = create(ctx, c, simSid, out[i:i+idx])
				i = i + idx
			}
			if err != nil {
				rest.ServerError(w, r, err)
				return
			}
		}
		w.WriteHeader(204)
	})
	addr := ":4425"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		handlers.Logger.Error("Error listening", "addr", addr, "err", err)
		os.Exit(2)
	}
	mux := handlers.UUID(http.DefaultServeMux) // add UUID header
	mux = handlers.Server(mux,
		"report-data-usage-server/"+Version) // add Server header
	mux = handlers.Log(mux)      // log requests/responses
	mux = handlers.Duration(mux) // add Duration header
	handlers.Logger.Info("Started server", "protocol", "http", "port", addr)
	http.Serve(ln, mux)
}
