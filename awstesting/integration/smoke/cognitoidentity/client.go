// +build integration

//Package cognitoidentity provides gucumber integration tests support.
package cognitoidentity

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/cognitoidentity"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cognitoidentity", func() {
		gucumber.World["client"] = cognitoidentity.New(smoke.Session)
	})
}
