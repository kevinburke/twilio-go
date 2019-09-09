package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/appian/twilio-go/testdata"
)

func TestGetActivity(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.TaskRouterActivityResponse)
	defer server.Close()

	workspaceSid := "WS58f1e8f2b1c6b88ca90a012a4be0c279"
	sid := "WAc74e6c39eb3080f8211d049a8b95611c"
	name := "NewAvailableActivity"

	activity, err := client.TaskRouter.Workspace(workspaceSid).Activities.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if activity.Sid != sid {
		t.Errorf("activity: got sid %q, want %q", activity.Sid, sid)
	}

	if activity.FriendlyName != name {
		t.Errorf("activity: got sid %q, want %q", activity.FriendlyName, name)
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Workspaces/" + workspaceSid + "/Activities/" + sid
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}

func TestCreateActivity(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.TaskRouterActivityCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Some Activity"
	data.Set("FriendlyName", newname)

	workspaceSid := "WS7a2aa7d8acc191786ad3c647c5fc3110"
	act, err := client.TaskRouter.Workspace(workspaceSid).Activities.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if act.WorkspaceSid != workspaceSid {
		t.Errorf("WorkspaceSid not correct")
	}
	if act.FriendlyName != "twilio-go-activity-client-testing" {
		t.Errorf("FriendlyNames don't match")
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Workspaces/" + workspaceSid + "/Activities"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
