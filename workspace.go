package twilio

import (
	"context"
	"net/url"
)

const WorkspacePath = "Workspaces/"

type WorkspaceCreator struct {
	client *Client
}

type Workspace struct {
	AccountSid           string            `json:"account_sid"`
	DateCreated          TwilioTime        `json:"date_created"`
	DateUpdated          TwilioTime        `json:"date_updated"`
	DefaultActivityName  string            `json:"default_activity_name"`
	DefaultActivitySid   string            `json:"default_activity_sid"`
	EventCallbackUrl     string            `json:"event_callback_url"`
	EventsFilter         string            `json:"friendly_name"`
	FriendlyName         string            `json:"friendly_name"`
	Links                map[string]string `json:"links"`
	MultiTaskEnabled     bool              `json:"multi_task_enabled"`
	PrioritizeQueueOrder map[string]string `json:"prioritize_queue_order"`
	Sid                  string            `json:"sid"`
	TimeoutActivityName  string            `json:"timeout_activity_name"`
	TimeoutActivitySid   string            `json:"timeout_activity_name"`
	URL                  string            `json:"url"`
}

// Get retrieves a Workspace by its sid.
//
// See https://www.twilio.com/docs/taskrouter/api/workspaces#action-create for more.
func (r *WorkspaceCreator) Get(ctx context.Context, sid string) (*Workspace, error) {
	workspace := new(Workspace)
	err := r.client.WorkspaceClient.GetResource(ctx, WorkspacePath, sid, workspace)
	return workspace, err
}

// Create creates a new Workspace.
//
// For a list of valid parameters see
// https://www.twilio.com/docs/taskrouter/api/workspaces#action-create
func (r *WorkspaceCreator) Create(ctx context.Context, data url.Values) (*Workspace, error) {
	workspace := new(Workspace)
	err := r.client.WorkspaceClient.CreateResource(ctx, WorkspacePath, data, workspace)
	return workspace, err
}

// Delete deletes a Workspace.
//
// See https://www.twilio.com/docs/taskrouter/api/workers#action-delete for more.
func (r *WorkspaceCreator) Delete(ctx context.Context, sid string) error {
	return r.client.WorkspaceClient.DeleteResource(ctx, WorkspacePath, sid)
}

// Update updates a Workspace.
//
// See https://www.twilio.com/docs/taskrouter/api/workspaces#action-update for more.
func (r *WorkspaceCreator) Update(ctx context.Context, sid string, data url.Values) (*Workspace, error) {
	worker := new(Workspace)
	err := r.client.WorkspaceClient.UpdateResource(ctx, WorkspacePath, sid, data, worker)
	return worker, err
}
