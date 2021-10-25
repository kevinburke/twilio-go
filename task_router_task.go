package twilio

import (
	"context"
	"net/url"
)

const TaskPathPart = "Task"

type TaskService struct {
	client       *Client
	workspaceSid string
}

type Links struct {
	TaskQueue    string `json:"task_queue"`
	Workflow     string `json:"workflow"`
	Workspace    string `json:"workspace"`
	Reservations string `json:"reservations"`
}

type Task struct {
	Sid                   string     `json:"sid"`
	AccountSid            string     `json:"account_sid"`
	Age                   int        `json:"age"`
	AssignmentStatus      string     `json:"assignment_status"`
	Attributes            string     `json:"attributes"`
	Addons                string     `json:"addons"`
	DateCreated           TwilioTime `json:"date_created"`
	DateUpdated           TwilioTime `json:"date_updated"`
	TaskQueueEnteredDate  TwilioTime `json:"task_queue_entered_date"`
	Priority              int        `json:"priority"`
	Reason                string     `json:"reason"`
	TaskQueueSid          string     `json:"task_queue_sid"`
	TaskChannelSid        string     `json:"task_channel_sid"`
	TaskChannelUniqueName string     `json:"task_channel_unique_name"`
	Timeout               int        `json:"timeout"`
	WorkflowSid           string     `json:"workflow_sid"`
	WorkflowFriendlyName  string     `json:"workflow_friendly_name"`
	WorkspaceSid          string     `json:"workspace_sid"`
	URL                   string     `json:"url"`
	Links                 Links      `json:"links"`
}

type TaskPage struct {
	Page
	Tasks []*Task `json:"task"`
}

func (r *TaskService) Get(ctx context.Context, sid string) (*Task, error) {
	task := new(Task)
	err := r.client.GetResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskPathPart, sid, task)
	return task, err
}

func (r *TaskService) Create(ctx context.Context, data url.Values) (*Task, error) {
	task := new(Task)
	err := r.client.CreateResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskPathPart, data, task)
	return task, err
}

func (r *TaskService) Delete(ctx context.Context, sid string) error {
	return r.client.DeleteResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskPathPart, sid)
}

func (ipn *TaskService) Update(ctx context.Context, sid string, data url.Values) (*Task, error) {
	task := new(Task)
	err := ipn.client.UpdateResource(ctx, "Workspaces/"+ipn.workspaceSid+"/"+TaskPathPart, sid, data, task)
	return task, err
}

// GetPage retrieves an TaskPage, filtered by the given data.
func (ins *TaskService) GetPage(ctx context.Context, data url.Values) (*TaskPage, error) {
	iter := ins.GetPageIterator(data)
	return iter.Next(ctx)
}

type TaskPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (c *TaskService) GetPageIterator(data url.Values) *TaskPageIterator {
	iter := NewPageIterator(c.client, data, "Workspaces/"+c.workspaceSid+"/"+TaskPathPart)
	return &TaskPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *TaskPageIterator) Next(ctx context.Context) (*TaskPage, error) {
	cp := new(TaskPage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
