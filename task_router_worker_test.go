package twilio

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/appian/twilio-go/testdata"
)

func TestGetWorker(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.WorkerResponse)
	defer server.Close()

	sid := "WK7a2aa7d8acc191786ad3c647c5fc3119"

	worker, err := client.TaskRouter.Workspace("WS58f1e8f2b1c6b88ca90a012a4be0c279").Workers.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if worker.Sid != sid {
		t.Errorf("task router worker: got sid %q, want %q", worker.Sid, sid)
	}

	if worker.FriendlyName != "NewWorker3" {
		t.Errorf("Incorrect FriendlyName")
	}
}

type WorkerTestConfig struct {
	Type string `json:"type"`
}

func TestDecodeWorker(t *testing.T) {
	t.Parallel()
	msg := new(Worker)

	if err := json.Unmarshal(testdata.WorkerResponse, &msg); err != nil {
		t.Fatal(err)
	}

	cfg := new(WorkerTestConfig)
	if err := json.Unmarshal([]byte(msg.Attributes), &cfg); err != nil {
		t.Fatal(err)
	}

	if cfg.Type != "support" {
		t.Errorf("Type not correct")
	}
}
func TestCreateWorker(t *testing.T) {
	t.Parallel()
	client, server := getServer(testdata.WorkerCreateResponse)
	defer server.Close()

	data := url.Values{}
	newname := "Support Worker 1"
	data.Set("FriendlyName", newname)
	workspaceSid := "WS7a2aa7d8acc191786ad3c647c5fc3110"

	worker, err := client.TaskRouter.Workspace(workspaceSid).Workers.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	if worker.WorkspaceSid != workspaceSid {
		t.Errorf("WorkspaceSid not correct")
	}

	if worker.FriendlyName != newname {
		t.Errorf("FriendlyNames don't match")
	}

	if len(server.URLs) != 1 {
		t.Errorf("URL length is %d, want 1", len(server.URLs))
	}
	want := "/v1/Workspaces/" + workspaceSid + "/Workers"
	if server.URLs[0].String() != want {
		t.Errorf("request URL:\ngot  %q\nwant %q", server.URLs[0], want)
	}
}
