package twilio

import (
	"context"
	"net/url"
)

const taskRouterQueuePathPart = "TaskQueues"

type TaskRouterQueueService struct {
	client       *Client
	workspaceSid string
}

type TaskRouterQueue struct {
	Sid                     string     `json:"sid"`
	AccountSid              string     `json:"account_sid"`
	FriendlyName            string     `json:"friendly_name"`
	AssignmentActivityName  string     `json:"assignment_activity_name"`
	AssignmentActivitySid   string     `json:"assignment_activity_sid"`
	ReservationActivityName string     `json:"reservation_activity_name"`
	ReservationActivitySid  string     `json:"reservation_activity_sid"`
	TargetWorkers           string     `json:"target_workers"`
	TaskOrder               string     `json:"task_order"`
	DateCreated             TwilioTime `json:"date_created"`
	DateUpdated             TwilioTime `json:"date_updated"`
	URL                     string     `json:"url"`
	WorkspaceSid            string     `json:"workspace_sid"`
	MaxReservedWorkers      int        `json:"max_reserved_workers"`
}

type TaskRouterQueuePage struct {
	Page
	TaskRouterQueues []*TaskRouterQueue `json:"task_queues"`
}

func (r *TaskRouterQueueService) Get(ctx context.Context, sid string) (*TaskRouterQueue, error) {
	queue := new(TaskRouterQueue)
	err := r.client.GetResource(ctx, "Workspaces/"+r.workspaceSid+"/"+taskRouterQueuePathPart, sid, queue)
	return queue, err
}

func (r *TaskRouterQueueService) Create(ctx context.Context, data url.Values) (*TaskRouterQueue, error) {
	queue := new(TaskRouterQueue)
	err := r.client.CreateResource(ctx, "Workspaces/"+r.workspaceSid+"/"+taskRouterQueuePathPart, data, queue)
	return queue, err
}

func (r *TaskRouterQueueService) Delete(ctx context.Context, sid string) error {
	return r.client.DeleteResource(ctx, "Workspaces/"+r.workspaceSid+"/"+taskRouterQueuePathPart, sid)
}

func (ipn *TaskRouterQueueService) Update(ctx context.Context, sid string, data url.Values) (*TaskRouterQueue, error) {
	queue := new(TaskRouterQueue)
	err := ipn.client.UpdateResource(ctx, "Workspaces/"+ipn.workspaceSid+"/"+taskRouterQueuePathPart, sid, data, queue)
	return queue, err
}

// GetPage retrieves an TaskRouterQueuePage, filtered by the given data.
func (ins *TaskRouterQueueService) GetPage(ctx context.Context, data url.Values) (*TaskRouterQueuePage, error) {
	iter := ins.GetPageIterator(data)
	return iter.Next(ctx)
}

type TaskRouterQueuePageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (c *TaskRouterQueueService) GetPageIterator(data url.Values) *TaskRouterQueuePageIterator {
	iter := NewPageIterator(c.client, data, "Workspaces/"+c.workspaceSid+"/"+taskRouterQueuePathPart)
	return &TaskRouterQueuePageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *TaskRouterQueuePageIterator) Next(ctx context.Context) (*TaskRouterQueuePage, error) {
	cp := new(TaskRouterQueuePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
