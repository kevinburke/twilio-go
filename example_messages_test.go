package twilio_test

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	twilio "github.com/appian/twilio-go"
)

func ExampleMessageService_Get() {
	client := twilio.NewClient("AC123", "123", nil)
	message, _ := client.Messages.Get(context.TODO(), "SM123")
	fmt.Println(message.Status)
}

func ExampleMessageService_GetMessagesInRange() {
	// Get all messages between 10:34:00 Oct 26 and 19:25:59 Oct 27, NYC time.
	nyc, _ := time.LoadLocation("America/New_York")
	start := time.Date(2016, 10, 26, 22, 34, 00, 00, nyc)
	end := time.Date(2016, 10, 27, 19, 25, 59, 00, nyc)

	client := twilio.NewClient("AC123", "123", nil)
	iter := client.Messages.GetMessagesInRange(start, end, url.Values{})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
		page, err := iter.Next(ctx)
		if err == twilio.NoMoreResults {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for i, message := range page.Messages {
			fmt.Printf("%d: %s (%s)", i, message.Sid, message.DateCreated.Time)
		}
	}
}
