// Package datausage has utilities for retrieving data usage from a SIM card.
package datausage

import (
	"context"
	"errors"
	"net/url"
	"time"

	"golang.org/x/sync/errgroup"

	types "github.com/kevinburke/go-types"
	twilio "github.com/appian/twilio-go"
)

// GetUsage gets usage for the given sim, starting at start and ending at end.
// The return value is a list of usage information. The first item in the list
// is the interval starting at start and so on until end is reached.
//
// sim may be either a Sim SID or a UniqueName.
func GetUsage(ctx context.Context, client *twilio.Client, sim string, start, end time.Time, interval time.Duration) ([]types.Bits, error) {
	if interval <= 0 {
		panic("invalid interval: " + interval.String())
	}
	iter := start
	count := 0
	for {
		count++
		iter = iter.Add(interval)
		if !iter.Before(end) {
			break
		}
	}
	list := make([]types.Bits, count)
	group, errctx := errgroup.WithContext(ctx)
	for i := 0; i < count; i++ {
		i := i
		group.Go(func() error {
			intervalStart := start.Add(interval * time.Duration(i)).Format(time.RFC3339)
			intervalEnd := start.Add(interval * time.Duration(i+1)).Format(time.RFC3339)
			page, err := client.Wireless.Sims.GetUsageRecords(errctx, sim, url.Values{
				"Start": []string{intervalStart},
				"End":   []string{intervalEnd},
			})
			if err != nil {
				return err
			}
			if len(page.UsageRecords) == 0 {
				return errors.New("no usage records for date range")
			}
			if len(page.UsageRecords) > 1 {
				return errors.New("too many usage records for date range")
			}
			list[i] = page.UsageRecords[0].Data.Total
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return list, nil
}
