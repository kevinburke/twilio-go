package twilio

import (
	"context"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestAccountGet(t *testing.T) {
	t.Parallel()
	client, server := getServer(accountInstance)
	defer server.Close()
	sid := "AC58f1e8f2b1c6b88ca90a012a4be0c279"
	acct, err := client.Accounts.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if acct.Sid != sid {
		t.Errorf("wrong sid")
	}
	if acct.Type != "Full" {
		t.Errorf("wrong type")
	}
}

func TestAccountCreate(t *testing.T) {
	t.Parallel()
	client, server := getServer(accountCreateResponse)
	defer server.Close()
	data := url.Values{}
	newname := "new account name 1478105087"
	data.Set("FriendlyName", newname)
	acct, err := client.Accounts.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if acct.FriendlyName != newname {
		t.Errorf("new account name incorrect")
	}
}

func TestAccountUpdateLive(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	sid := "ACdd54a711c3d4031ac500c5236ab121d7"
	data := url.Values{}
	newname := "new account name " + strconv.FormatInt(time.Now().UTC().Unix(), 10)
	data.Set("FriendlyName", newname)
	acct, err := envClient.Accounts.Update(context.Background(), sid, data)
	if err != nil {
		t.Fatal(err)
	}
	if acct.FriendlyName != newname {
		t.Errorf("new account name incorrect")
	}
}

func TestAccountUpdateInMemory(t *testing.T) {
	t.Parallel()
	client, server := getServer(accountInstance)
	defer server.Close()
	sid := "AC58f1e8f2b1c6b88ca90a012a4be0c279"
	data := url.Values{}
	data.Set("FriendlyName", "kevin account woo")
	acct, err := client.Accounts.Update(context.Background(), sid, data)
	if err != nil {
		t.Fatal(err)
	}
	if acct.FriendlyName != "kevin account woo" {
		t.Errorf("new account name incorrect")
	}
}

func TestAccountGetPage(t *testing.T) {
	t.Parallel()
	client, server := getServer(accountList)
	defer server.Close()
	data := url.Values{}
	data.Set("PageSize", "2")
	page, err := client.Accounts.GetPage(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Accounts) != 2 {
		t.Errorf("expected accounts len to be 2")
	}
	if page.Accounts[0].Sid != "AC0cd9be8fd5e6e4fa0a04f50ac1caca4e" {
		t.Errorf("wrong sid")
	}
}

func TestAccountGetPageIterator(t *testing.T) {
	t.Skip()

	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	data := url.Values{}
	data.Set("PageSize", "2")
	iter := envClient.Accounts.GetPageIterator(data)
	page, err := iter.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Accounts) != 2 {
		t.Errorf("expected accounts len to be 2")
	}
	page, err = iter.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Accounts) != 2 {
		t.Errorf("expected accounts len to be 2")
	}
}
