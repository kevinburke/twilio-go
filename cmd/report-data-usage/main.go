// Binary report-data-usage reports information about how much data your sim cards have used.
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
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	types "github.com/kevinburke/go-types"
	twilio "github.com/appian/twilio-go"
	"github.com/appian/twilio-go/datausage"
	"golang.org/x/sync/errgroup"
)

var duration = flag.Uint("days", 7, "Number of days to get usage for")
var location = flag.String("location", "", "Timezone to use (defaults to system location/TZ env var)")
var sim = flag.String("sim", "", "Load data about a specific Sim/unique name only")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Print information about how much data you used on each day.
`)
		flag.PrintDefaults()
	}
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
	end := time.Now().Add(24 * time.Hour).In(loc)
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	start := end.Add(-time.Duration(*duration+1) * 24 * time.Hour)
	iter := c.Wireless.Sims.GetPageIterator(nil)
	sims := make([]string, 0)
	if *sim != "" {
		sims = []string{*sim}
	} else {
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
	// TODO we could probably get fancy here and send the sims on channels when
	// we get them above, which would give us a head start on fetching the data
	//
	// we could also guard data fetches with a semaphore:
	// github.com/kevinburke/semaphore
	var mu sync.Mutex
	ctx, cancel := context.WithTimeout(context.Background(), 31*time.Second)
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
		log.Fatal(err)
	}
	for _, name := range sims {
		total := types.Bits(0)
		fmt.Printf("%s\n%s\n", name, strings.Repeat("-", len(name)))
		for i := 0; i < len(usage[name]); i++ {
			t := start.Add(time.Duration(i) * 24 * time.Hour)
			fmt.Printf("%s: %s\n", t.Format("2006-01-02"), usage[name][i])
			total += usage[name][i]
		}
		fmt.Printf("total (last %d days): %s\n\n", *duration, total)
	}
}
