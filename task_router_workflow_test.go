package twilio

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/appian/twilio-go/testdata"
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

type TestFilter struct {
	Queue string `json:"queue"`
}

type TestTaskRouting struct {
	DefaultFilter TestFilter `json:"default_filter"`
}

type TestConfig struct {
	TaskRouting TestTaskRouting `json:"task_routing"`
}

func TestDecodeWorkflow(t *testing.T) {
	t.Parallel()
	msg := new(Workflow)

	if err := json.Unmarshal(testdata.WorkflowResponse, &msg); err != nil {
		t.Fatal(err)
	}

	cfg := new(TestConfig)
	if err := json.Unmarshal([]byte(msg.Configuration), &cfg); err != nil {
		t.Fatal(err)
	}

	if cfg.TaskRouting.DefaultFilter.Queue != "WQ0c1274082082355320d8a41f94eb57aa" {
		t.Errorf("Queue sid not correct")
	}
}

func TestCreateWorkflow(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.WorkflowCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Sales, Marketing, Support Workflow"
	data.Set("FriendlyName", newname)
	data.Set("Configuration", "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQ0c1274082082355320d8a41f94eb57aa\"}}}")
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
