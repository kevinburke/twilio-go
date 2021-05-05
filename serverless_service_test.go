package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/kevinburke/twilio-go/testdata"
)

func TestGetService(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessServiceResponse)
	defer server.Close()

	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	name := "ServiceGet"

	srv, err := client.Serverless.Services.Get(context.Background(), serviceSid)
	if err != nil {
		t.Fatal(err)
	}
	if srv.Sid != serviceSid {
		t.Errorf("service: got sid %q, want %q", srv.Sid, serviceSid)
	}

	if srv.FriendlyName != name {
		t.Errorf("service: got friendly name  %q, want %q", srv.FriendlyName, name)
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services/" + serviceSid
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}

func TestCreateService(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessServiceCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Some Service"
	data.Set("FriendlyName", newname)

	srv, err := client.Serverless.Services.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if srv.FriendlyName != newname {
		t.Errorf("FriendlyNames don't match")
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}

func TestPagedService(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessServicePageResponse)
	defer server.Close()

	iterator := client.Serverless.Services.GetPageIterator(url.Values{"PageSize": []string{"50"}})
	count := 0
	for {
		page, err := iterator.Next(context.Background())
		if err == NoMoreResults {
			break
		}
		count += len(page.Services)
	}

	if count != 2 {
		t.Errorf("Services length is %d, want 2", count)
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services?PageSize=50"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
