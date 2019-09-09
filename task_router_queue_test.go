package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/appian/twilio-go/testdata"
)

func TestGetTaskQueue(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.TaskQueueResponse)
	defer server.Close()

	sid := "WQ63868a235fc1cf3987e6a2b67346273f"
	name := "English"
	queue, err := client.TaskRouter.Workspace("WS58f1e8f2b1c6b88ca90a012a4be0c279").Queues.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if queue.Sid != sid {
		t.Errorf("task router queue: got sid %q, want %q", queue.Sid, sid)
	}

	if queue.FriendlyName != name {
		t.Errorf("ask router queue: got sid %q, want %q", queue.FriendlyName, name)
	}
}

func TestCreateTaskQueue(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.TaskQueueCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "English"
	data.Set("FriendlyName", newname)
	data.Set("AssignmentActivitySid", "WA086126500699bba752b0485c185013d1")
	data.Set("ReservationActivitySid", "WAd9ad5e7b1cb9c8327cf7eb14a8a31131")

	workspaceSid := "WS7a2aa7d8acc191786ad3c647c5fc3110"
	act, err := client.TaskRouter.Workspace(workspaceSid).Queues.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if act.WorkspaceSid != workspaceSid {
		t.Errorf("WorkspaceSid not correct")
	}
	if act.FriendlyName != newname {
		t.Errorf("FriendlyNames don't match")
	}
	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Workspaces/" + workspaceSid + "/TaskQueues"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
