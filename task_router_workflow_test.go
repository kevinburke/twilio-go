package twilio

import (
	"context"
	"net/url"
	"testing"

	"github.com/kevinburke/twilio-go/testdata"
)

func TestGetWorkflow(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.WorkflowResponse)
	defer server.Close()

	sid := "WW63868a235fc1cf3987e6a2b67346273f"
	taskReservationTimeout := 120

	workflow, err := client.TaskRouter.Workspace("WS58f1e8f2b1c6b88ca90a012a4be0c279").Workflows.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if workflow.Sid != sid {
		t.Errorf("task router workflow: got sid %q, want %q", workflow.Sid, sid)
	}

	if workflow.TaskReservationTimeout != taskReservationTimeout {
		t.Errorf("task router workflow: got sid %q, want %q", workflow.TaskReservationTimeout, taskReservationTimeout)
	}
}

func TestCreateWorkflow(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.WorkflowCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Sales, Marketing, Support Workflow"
	data.Set("FriendlyName", newname)
	workspaceSid := "WS7a2aa7d8acc191786ad3c647c5fc3110"

	workflow, err := client.TaskRouter.Workspace(workspaceSid).Workflows.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	if workflow.WorkspaceSid != workspaceSid {
		t.Errorf("WorkspaceSid not correct")
	}

	if workflow.FriendlyName != newname {
		t.Errorf("FriendlyNames don't match")
	}

	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Workspaces/" + workspaceSid + "/Workflows"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
