package token_test

import (
	"fmt"
	"time"

	"github.com/appian/twilio-go/token"
)

func Example() {
	t := token.New("AC123", "456bef", "secretkey", "test@example.com", time.Hour)
	grant := token.NewConversationsGrant("a-conversation-sid")
	t.AddGrant(grant)
	jwt, _ := t.JWT()
	fmt.Println(jwt) // A string encoded with the given values.
}
