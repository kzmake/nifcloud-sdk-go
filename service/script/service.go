// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package script

import (
	"github.com/alice02/nifcloud-sdk-go/aws"
	"github.com/alice02/nifcloud-sdk-go/aws/client"
	"github.com/alice02/nifcloud-sdk-go/aws/client/metadata"
	"github.com/alice02/nifcloud-sdk-go/aws/request"
	"github.com/alice02/nifcloud-sdk-go/aws/signer/v4"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/script"
)

// Script provides the API operation methods for making requests to
// NIFCLOUD Script. See this package's package overview docs
// for details on the service.
//
// Script methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Script struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "script"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Script client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Script client from just a session.
//     svc := script.New(mySession)
//
//     // Create a Script client with additional configuration
//     svc := script.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Script {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Script {
	svc := &Script{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-09-01",

				TargetPrefix: "2015-09-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(script.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(script.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(script.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(script.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Script operation and runs any
// custom request initialization.
func (c *Script) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
