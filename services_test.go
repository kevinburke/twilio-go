package twilio

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	friendlyName := fmt.Sprintf("test-messaging-service-%d", rand.Int())
	data := url.Values{"FriendlyName": []string{friendlyName}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	service, err := envClient.Message.Services.Create(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(service.Sid) == 0 {
		t.Error("expected to create a messaging services, got back 0")
	}
	if service.FriendlyName != friendlyName {
		t.Error("friendly name was not set correctly")
	}
}

func TestService_GetPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	data := url.Values{"PageSize": []string{"1000"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	services, err := envClient.Message.Services.GetPage(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(services.Services) == 0 {
		t.Error("expected to get a list of messaging services, got back 0")
	}
}

func TestServiceService_PhoneNumbers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	data := url.Values{"PageSize": []string{"5"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	numbers, err := envClient.IncomingNumbers.GetPage(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(numbers.IncomingPhoneNumbers) == 0 {
		t.Error("expected to get a list of phone numbers, got back 0")
	}
	services, err := envClient.Message.Services.GetPage(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(services.Services) == 0 {
		t.Error("expected to get a messaging services, got back 0")
	}

	number := numbers.IncomingPhoneNumbers[0]
	service := services.Services[0]
	if _, err := envClient.Message.Services.CreatePhoneNumber(ctx, service.Sid, number.Sid); err != nil {
		t.Error("expected to CreatePhoneNumber but got error", err)
	}

	if _, err := envClient.Message.Services.GetPhoneNumber(ctx, service.Sid, number.Sid); err != nil {
		t.Error("expected to GetPhoneNumber after CreatePhoneNumber but got error", err)
	}

	if err := envClient.Message.Services.DeletePhoneNumber(ctx, service.Sid, number.Sid); err != nil {
		t.Error("expected to DeletePhoneNumber after CreatePhoneNumber but got error", err)
	}
}
