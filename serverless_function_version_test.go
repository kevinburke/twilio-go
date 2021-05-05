package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/kevinburke/twilio-go/testdata"
)

func TestGetFunctionVersion(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionResponse)
	defer server.Close()

	ctx := context.Background()
	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	functionSid := "ZH691419f6825741bb96d2ea9af301d055"
	sid := "ZN684765f345f046b5aff4f6d29eea30d4"

	function, err := client.Serverless.Service(serviceSid).Functions.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}

	versionClient, versionServer := getServer(testdata.ServerlessFunctionVersionResponse)
	defer versionServer.Close()

	version, err := versionClient.Serverless.Service(serviceSid).Function(function.Sid).Versions.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}

	pathName := "/test-path"
	if version.Path != pathName {
		t.Errorf("function version: got path %q, want %q", version.Path, pathName)
	}

	if len(versionServer.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(versionServer.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions/" + functionSid + "/Versions/" + sid
	if versionServer.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", versionServer.URLs[0], want)
	}
}

func TestPagedFunctionVersion(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionResponse)
	defer server.Close()

	ctx := context.Background()
	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	functionSid := "ZH691419f6825741bb96d2ea9af301d055"
	sid := "ZN684765f345f046b5aff4f6d29eea30d4"

	function, err := client.Serverless.Service(serviceSid).Functions.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}

	versionClient, versionServer := getServer(testdata.ServerlessFunctionVersionsResponse)
	defer versionServer.Close()

	iterator := versionClient.Serverless.Service(serviceSid).Function(function.Sid).Versions.GetPageIterator(url.Values{"PageSize": []string{"50"}})
	count := 0
	for {
		page, err := iterator.Next(context.Background())
		if err == NoMoreResults {
			break
		}
		count += len(page.FunctionVersions)
	}

	if count != 2 {
		t.Errorf("FunctionVersions length is %d, want 2", count)
	}
	if len(versionServer.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(versionServer.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions/" + functionSid + "/Versions?PageSize=50"
	if versionServer.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", versionServer.URLs[0], want)
	}
}

func TestGetFunctionVersionContent(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.ServerlessFunctionResponse)
	defer server.Close()

	ctx := context.Background()
	serviceSid := "ZSc6c0e43c485bfd439d6e076abb51aaa6"
	functionSid := "ZH691419f6825741bb96d2ea9af301d055"
	sid := "ZN684765f345f046b5aff4f6d29eea30d4"

	function, err := client.Serverless.Service(serviceSid).Functions.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}

	versionClient, versionServer := getServer(testdata.ServerlessFunctionVersionContentResponse)
	defer versionServer.Close()

	versionContent, err := versionClient.Serverless.Service(serviceSid).Function(function.Sid).Versions.GetContent(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "exports.handler = function (context, event, callback) { console.log(event) };"
	if versionContent.Content != expectedContent {
		t.Errorf("function version content: got content %q, want %q", versionContent.Content, expectedContent)
	}

	if len(versionServer.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(versionServer.URLs))
	}
	want := "/v1/Services/" + serviceSid + "/Functions/" + functionSid + "/Versions/" + sid + "/Content"
	if versionServer.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", versionServer.URLs[0], want)
	}
}
