package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/kevinburke/twilio-go/testdata"
)

func TestGetFunction(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionResponse)
	defer server.Close()

	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	sid := "ZH691419f6825741bb96d2ea9af301d055"
	name := "FunctionGet"

	function, err := client.Serverless.Service(serviceSid).Functions.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if function.ServiceSid != serviceSid {
		t.Errorf("function: got service sid %q, want %q", function.ServiceSid, serviceSid)
	}

	if function.FriendlyName != name {
		t.Errorf("function: got friendly name %q, want %q", function.FriendlyName, name)
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions/" + sid
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}

func TestCreateFunction(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Some Function"
	data.Set("FriendlyName", newname)

	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	function, err := client.Serverless.Service(serviceSid).Functions.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if function.ServiceSid != serviceSid {
		t.Errorf("function: got service sid %q, want %q", function.ServiceSid, serviceSid)
	}
	if function.FriendlyName != newname {
		t.Errorf("FriendlyNames don't match")
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}

func TestPagedFunction(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionPageResponse)
	defer server.Close()

	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"

	iterator := client.Serverless.Service(serviceSid).Functions.GetPageIterator(url.Values{"PageSize": []string{"50"}})
	count := 0
	for {
		page, err := iterator.Next(context.Background())
		if err == NoMoreResults {
			break
		}
		count += len(page.Functions)
	}

	if count != 2 {
		t.Errorf("Functions length is %d, want 2", count)
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions?PageSize=50"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
