package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestNotifyCredentialsGetPage(t *testing.T) {
	t.Parallel()

	client, server := getServer(credentialsPage)
	defer server.Close()

	page, err := client.Notify.Credentials.GetPage(context.Background(), url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	if len(page.Credentials) != 2 {
		t.Fatalf("expected len(credentials) to be 2, got %d", len(page.Credentials))
	}
}

func TestNotifyCredentialsGet(t *testing.T) {
	t.Parallel()

	client, server := getServer(notifyCredential)
	defer server.Close()

	cred, err := client.Notify.Credentials.Get(context.Background(), "123")
	if err != nil {
		t.Fatal(err)
	}

	expected := "AC6bc21af903cc765a9d7f7e0467ec812a"
	if cred.AccountSid != expected {
		t.Errorf("expected AccountSid to be %s, got %s", expected, cred.Sid)
	}

	expected = "MyFCMCredential"
	if cred.FriendlyName != expected {
		t.Errorf("expected FriendlyName to be %s, got %s", expected, cred.FriendlyName)
	}

	expected = string(TypeFCM)
	if cred.Type != TypeFCM {
		t.Errorf("expected Type to be %s, got %s", expected, cred.Type)
	}
}

func TestCreateNotifyCredentials(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}

	values := url.Values{}
	values.Set("FriendlyName", "Test FCM cred")
	values.Set("Type", string(TypeFCM))
	values.Set("Secret", "secret here")

	cred, err := envClient.Notify.Credentials.Create(context.Background(), values)
	if err != nil {
		t.Fatal(err)
	}

	if cred.Sid == "" {
		t.Error("Sid should not be empty")
	}

	if cred.AccountSid != envClient.AccountSid {
		t.Errorf("expected AccountSid to be %s, got %s", envClient.AccountSid, cred.AccountSid)
	}

	if cred.Type != TypeFCM {
		t.Errorf("expected Type to be %s, got %s", TypeFCM, cred.Type)
	}
}
