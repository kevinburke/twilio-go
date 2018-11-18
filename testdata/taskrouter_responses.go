package testdata

var TaskRouterCreateResponse = []byte(`
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

var TaskRouterQueueCreateResponse = []byte(`
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
