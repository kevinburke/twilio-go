package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetActivity(t *testing.T) {
	t.Parallel()
	client, server := getServer(taskRouterActivityResponse)
	defer server.Close()

	sid := "WAc74e6c39eb3080f8211d049a8b95611c"
	name := "NewAvailableActivity"

	activity, err := client.TaskRouter.Workspace("WS58f1e8f2b1c6b88ca90a012a4be0c279").Activities.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if activity.Sid != sid {
		t.Errorf("activity: got sid %q, want %q", activity.Sid, sid)
	}

	if activity.FriendlyName != name {
		t.Errorf("activity: got sid %q, want %q", activity.FriendlyName, name)
	}
}
func TestCreateActivity(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()

	data := url.Values{}
	newname := "Some Activity"
	data.Set("FriendlyName", newname)

	acct, err := envClient.TaskRouter.Workspace("WS58f1e8f2b1c6b88ca90a012a4be0c279").Activities.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if acct.FriendlyName != newname {
		t.Errorf("FriendlyName don't match")
	}
}
