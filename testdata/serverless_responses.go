package testdata

var ServerlessServiceResponse = []byte(`
{
    "sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "friendly_name": "ServiceGet",
    "unique_name": "twilio-go-service-client-testing-app",
    "include_credentials": true,
    "ui_editable": false,
    "date_created": "2020-11-18T16:52:30Z",
    "date_updated": "2020-11-18T16:52:30Z",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "links": {
      "environments": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Environments",
      "functions": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions",
      "assets": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Assets",
      "builds": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Builds"
    }
}
`)

var ServerlessServiceCreateResponse = []byte(`
{
    "sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "friendly_name": "Some Service",
    "unique_name": "twilio-go-service-client-testing-app",
    "include_credentials": true,
    "ui_editable": false,
    "date_created": "2020-11-18T16:52:30Z",
    "date_updated": "2020-11-18T16:52:30Z",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "links": {
      "environments": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Environments",
      "functions": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions",
      "assets": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Assets",
      "builds": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Builds"
    }
}
`)

var ServerlessServicePageResponse = []byte(`
{
  "services": [` + string(ServerlessServiceResponse) + `,` + string(ServerlessServiceResponse) + `],
  "meta": {
    "first_page_url": "https://serverless.twilio.com/v1/Services?PageSize=50&Page=0",
    "key": "services",
    "next_page_url": "https://serverless.twilio.com/v1/Services?PageSize=50&Page=1",
    "page": 0,
    "page_size": 50,
    "previous_page_url": "https://serverless.twilio.com/v1/Services?PageSize=50&Page=0",
    "url": "https://serverless.twilio.com/v1/Services?PageSize=50&Page=0"
  }
}
`)

var ServerlessFunctionResponse = []byte(`
{
    "sid": "ZH691419f6825741bb96d2ea9af301d055",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "service_sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "friendly_name": "FunctionGet",
    "date_created": "2020-11-10T20:00:00Z",
    "date_updated": "2020-11-10T20:00:00Z",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055",
    "links": {
      "function_versions": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions"
    }
}
`)

var ServerlessFunctionPageResponse = []byte(`
{
  "functions": [` + string(ServerlessFunctionResponse) + `,` + string(ServerlessFunctionResponse) + `],
  "meta": {
    "first_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions?PageSize=50&Page=0",
    "key": "functions",
    "next_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions?PageSize=50&Page=1",
    "page": 0,
    "page_size": 50,
    "previous_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions?PageSize=50&Page=0",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions?PageSize=50&Page=0"
  }
}
`)

var ServerlessFunctionCreateResponse = []byte(`
{
    "sid": "ZH20ac9c3583cb49eb80f833857ad8f696",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "service_sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "friendly_name": "Some Function",
    "date_created": "2020-11-10T20:00:00Z",
    "date_updated": "2020-11-10T20:00:00Z",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH20ac9c3583cb49eb80f833857ad8f696",
    "links": {
      "function_versions": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH20ac9c3583cb49eb80f833857ad8f696/Versions"
    }
}
`)

var ServerlessFunctionVersionResponse = []byte(`
{
    "sid": "ZN684765f345f046b5aff4f6d29eea30d4",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "service_sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "function_sid": "ZH691419f6825741bb96d2ea9af301d055",
    "path": "/test-path",
    "visibility": "public",
    "date_created": "2020-11-10T20:00:00Z",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions/ZN684765f345f046b5aff4f6d29eea30d4",
    "links": {
      "function_version_content": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions/ZN684765f345f046b5aff4f6d29eea30d4/Content"
    }
}
`)

var ServerlessFunctionVersionsResponse = []byte(`
{
    "function_versions": [` + string(ServerlessFunctionVersionResponse) + `,` + string(ServerlessFunctionVersionResponse) + `],
    "meta": {
      "first_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions?PageSize=50&Page=0",
      "key": "function_versions",
      "next_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions?PageSize=50&Page=1",
      "page": 0,
      "page_size": 50,
      "previous_page_url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions?PageSize=50&Page=0",
      "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions?PageSize=50&Page=0"
    }
}
`)

var ServerlessFunctionVersionContentResponse = []byte(`
{
    "sid": "ZN684765f345f046b5aff4f6d29eea30d4",
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "service_sid": "ZSc6c0e43c485bfd439d6e076abb51aaa6",
    "function_sid": "ZH691419f6825741bb96d2ea9af301d055",
    "content": "exports.handler = function (context, event, callback) { console.log(event) };",
    "url": "https://serverless.twilio.com/v1/Services/ZSc6c0e43c485bfd439d6e076abb51aaa6/Functions/ZH691419f6825741bb96d2ea9af301d055/Versions/ZN684765f345f046b5aff4f6d29eea30d4/Content"
}
`)
