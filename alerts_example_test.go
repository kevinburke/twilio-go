package twilio_test

import (
	"context"
	"fmt"
	"log"
	"net/url"

	twilio "github.com/appian/twilio-go"
)

func ExampleAlertService_GetPage() {
	client := twilio.NewClient("AC123", "123", nil)
	data := url.Values{}
	data.Set("ResourceSid", "SM123")
	page, err := client.Monitor.Alerts.GetPage(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	for _, alert := range page.Alerts {
		fmt.Println(alert.Sid)
	}
}
