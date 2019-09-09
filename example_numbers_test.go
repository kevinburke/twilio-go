package twilio_test

import (
	"context"
	"fmt"
	"net/url"

	twilio "github.com/appian/twilio-go"
)

func Example_buyNumber() {
	client := twilio.NewClient("AC123", "123", nil)
	ctx := context.TODO()

	// Find a number in the 925 area code
	var data url.Values
	data.Set("AreaCode", "925")
	numbers, _ := client.AvailableNumbers.Local.GetPage(ctx, "US", data)

	// Buy the first one (if there's at least one to buy)
	if len(numbers.Numbers) == 0 {
		fmt.Println("No phone numbers available")
		return
	}
	number := numbers.Numbers[0]
	var buyData url.Values
	buyData.Set("PhoneNumber", string(number.PhoneNumber))
	boughtNumber, _ := client.IncomingNumbers.Create(ctx, buyData)
	fmt.Println("bought number!", boughtNumber.PhoneNumber)
}
