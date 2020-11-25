package twilio

import "context"

const TaskChannelPathPart = "TaskChannels"

type TaskChannelService struct {
	client       *Client
	workspaceSid string
}

type TaskChannel struct {
	AccountSid              string            `json:"account_sid"`
	DateCreated             TwilioTime        `json:"date_created"`
	DateUpdated             TwilioTime        `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	Sid                     string            `json:"sid"`
	UniqueName              string            `json: "unique_name"`
	WorkspaceSid            string            `json:"workspace_sid"`
	ChannelOptimizedRouting bool              `json:"channel_optimized_routing"`
	URL                     string            `json:"url"`
	Links                   map[string]string `json:"links"`
}

type TaskChannelPage struct {
	Page
	TaskChannels []*Worker `json:"task_channels"`
}

func (r *TaskChannelService) Get(ctx context.Context, sid string) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := r.client.GetResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskChannelPathPart, sid, taskChannel)
	return taskChannel, err
}
