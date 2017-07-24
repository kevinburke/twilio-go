package twilio_test

import (
	"context"
	"fmt"

	twilio "github.com/kevinburke/twilio-go"
)

func ExampleWirelessService_Get() {
	client := twilio.NewClient("AC123", "123", nil)
	sim, _ := client.Wireless.Sims.Get(context.TODO(), "DE123")
	fmt.Println(sim.Status)
}
