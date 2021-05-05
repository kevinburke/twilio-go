package twilio

// ServerlessService allows for interaction with a Serverless Service
type ServerlessService struct {
	Functions *FunctionService
	Function  func(sid string) *FunctionService
}
