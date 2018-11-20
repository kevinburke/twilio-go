package testdata

var TaskRouterActivityCreateResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "available": false,
    "date_created": "2018-11-18T16:52:30Z",
    "date_updated": "2018-11-18T16:52:30Z",
    "friendly_name": "twilio-go-activity-client-testing",
    "links": {
        "workspace": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110"
    },
    "sid": "WAc6c0e43c485bfd439d6e076abb51aaa6",
    "url": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Activities/WAc6c0e43c485bfd439d6e076abb51aaa6",
    "workspace_sid": "WS7a2aa7d8acc191786ad3c647c5fc3110"
}
`)

var TaskRouterActivityResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "available": true,
    "date_created": "2014-05-14T10:50:02Z",
    "date_updated": "2014-05-14T23:26:06Z",
    "friendly_name": "NewAvailableActivity",
    "sid": "WAc74e6c39eb3080f8211d049a8b95611c",
    "url": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Activities/WAc74e6c39eb3080f8211d049a8b95611c",
    "workspace_sid": "WS58f1e8f2b1c6b88ca90a012a4be0c279"
  }
`)

var TaskQueueResponse = []byte(`
  {
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "assignment_activity_name": "817ca1c5-3a05-11e5-9292-98e0d9a1eb73",
    "assignment_activity_sid": "WAc74e6c39eb3080f8211d049a8b95612d",
    "date_created": "2015-08-04T01:31:41Z",
    "date_updated": "2015-08-04T01:31:41Z",
    "friendly_name": "English",
    "max_reserved_workers": 1,
    "links": {
      "assignment_activity": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Activities/WQ63868a235fc1cf3987e6a2b67346273f",
      "reservation_activity": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Activities/WQ63868a235fc1cf3987e6a2b67346273f",
      "workspace": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279",
      "statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/TaskQueues/WQ63868a235fc1cf3987e6a2b67346273f/Statistics",
      "real_time_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/TaskQueues/WQ63868a235fc1cf3987e6a2b67346273f/RealTimeStatistics",
      "cumulative_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/TaskQueues/WQ63868a235fc1cf3987e6a2b67346273f/CumulativeStatistics",
      "list_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/TaskQueues/Statistics"
    },
    "reservation_activity_name": "80fa2beb-3a05-11e5-8fc8-98e0d9a1eb73",
    "reservation_activity_sid": "WAc74e6c39eb3080f8211d049a8b95611c",
    "sid": "WQ63868a235fc1cf3987e6a2b67346273f",
    "target_workers": "languages HAS \"english\"",
    "task_order": "FIFO",
    "url": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/TaskQueues/WQ63868a235fc1cf3987e6a2b67346273f",
    "workspace_sid": "WS58f1e8f2b1c6b88ca90a012a4be0c279"
  }
`)

var TaskQueueCreateResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "assignment_activity_name": "817ca1c5-3a05-11e5-9292-98e0d9a1eb73",
    "assignment_activity_sid": "WAc6c0e43c485bfd439d6e076abb51aaa6",
    "date_created": "2015-08-04T01:31:41Z",
    "date_updated": "2015-08-04T01:31:41Z",
    "friendly_name": "English",
    "max_reserved_workers": 1,
    "links": {
      "assignment_activity": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Activities/WA7a2aa7d8acc191786ad3c647c5fc3110",
      "reservation_activity": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Activities/WA7a2aa7d8acc191786ad3c647c5fc3110",
      "workspace": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110",
      "statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/TaskQueues/WQ7a2aa7d8acc191786ad3c647c5fc3110/Statistics",
      "real_time_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/TaskQueues/WQ7a2aa7d8acc191786ad3c647c5fc3110/RealTimeStatistics",
      "cumulative_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/TaskQueues/WQ7a2aa7d8acc191786ad3c647c5fc3110/CumulativeStatistics",
      "list_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/TaskQueues/Statistics"
    },
    "reservation_activity_name": "80fa2beb-3a05-11e5-8fc8-98e0d9a1eb73",
    "reservation_activity_sid": "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "sid": "WQ7a2aa7d8acc191786ad3c647c5fc3110",
    "target_workers": "languages HAS \"english\"",
    "task_order": "FIFO",
    "url": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/TaskQueues/WQ7a2aa7d8acc191786ad3c647c5fc3110",
    "workspace_sid": "WS7a2aa7d8acc191786ad3c647c5fc3110"
  }
`)

var WorkflowResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "assignment_callback_url": "http://example.com",
    "configuration": "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQ0c1274082082355320d8a41f94eb57aa\"}}}",
    "date_created": "2014-05-14T10:50:02Z",
    "date_updated": "2014-05-14T23:26:06Z",
    "document_content_type": "application/json",
    "fallback_assignment_callback_url": null,
    "friendly_name": "Default Fifo Workflow",
    "sid": "WW63868a235fc1cf3987e6a2b67346273f",
    "task_reservation_timeout": 120,
    "url": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Workflows/WF63868a235fc1cf3987e6a2b67346273f",
    "workspace_sid": "WS58f1e8f2b1c6b88ca90a012a4be0c279",
    "links": {
      "statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Workflows/WF63868a235fc1cf3987e6a2b67346273f/Statistics",
      "real_time_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Workflows/WF63868a235fc1cf3987e6a2b67346273f/RealTimeStatistics",
      "cumulative_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS58f1e8f2b1c6b88ca90a012a4be0c279/Workflows/WF63868a235fc1cf3987e6a2b67346273f/CumulativeStatistics"
    }
  }
`)

var WorkflowCreateResponse = []byte(`
{
  "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
  "assignment_callback_url": "http://example.com",
  "configuration": "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQ0c1274082082355320d8a41f94eb57aa\"}}}",
  "date_created": "2014-05-14T10:50:02Z",
  "date_updated": "2014-05-14T23:26:06Z",
  "document_content_type": "application/json",
  "fallback_assignment_callback_url": "http://example2.com",
  "friendly_name": "Sales, Marketing, Support Workflow",
  "sid": "WW7a2aa7d8acc191786ad3c647c5fc3110",
  "task_reservation_timeout": 30,
  "url": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Workflows/WF7a2aa7d8acc191786ad3c647c5fc3110",
  "workspace_sid": "WS7a2aa7d8acc191786ad3c647c5fc3110",
  "links": {
    "statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Workflows/WF7a2aa7d8acc191786ad3c647c5fc3110/Statistics",
    "real_time_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Workflows/WF7a2aa7d8acc191786ad3c647c5fc3110/RealTimeStatistics",
    "cumulative_statistics": "https://taskrouter.twilio.com/v1/Workspaces/WS7a2aa7d8acc191786ad3c647c5fc3110/Workflows/WF7a2aa7d8acc191786ad3c647c5fc3110/CumulativeStatistics"
  }
}
`)
