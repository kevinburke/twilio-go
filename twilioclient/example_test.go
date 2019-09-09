package twilioclient_test

import (
	"fmt"
	"log"
	"time"

	"github.com/appian/twilio-go/twilioclient"
)

func Example() {
	cap := twilioclient.NewCapability("AC123", "123")
	cap.AllowClientIncoming("client-name")
	tok, err := cap.GenerateToken(time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tok)
}
