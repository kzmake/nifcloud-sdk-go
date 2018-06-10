package corehandlers

import (
	"os"
	"runtime"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
)

// SDKVersionUserAgentHandler is a request handler for adding the SDK Version
// to the user agent.
var SDKVersionUserAgentHandler = request.NamedHandler{
	Name: "core.SDKVersionUserAgentHandler",
	Fn: request.MakeAddToUserAgentHandler(nifcloud.SDKName, nifcloud.SDKVersion,
		runtime.Version(), runtime.GOOS, runtime.GOARCH),
}

const execEnvVar = `AWS_EXECUTION_ENV`
const execEnvUAKey = `exec_env`

// AddHostExecEnvUserAgentHander is a request handler appending the SDK's
// execution environment to the user agent.
//
// If the environment variable AWS_EXECUTION_ENV is set, its value will be
// appended to the user agent string.
var AddHostExecEnvUserAgentHander = request.NamedHandler{
	Name: "core.AddHostExecEnvUserAgentHander",
	Fn: func(r *request.Request) {
		v := os.Getenv(execEnvVar)
		if len(v) == 0 {
			return
		}

		request.AddToUserAgent(r, execEnvUAKey+"/"+v)
	},
}