// +build integration

//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/kzmake/nifcloud-sdk-go/nifcloud"
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/devicefarm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@devicefarm", func() {
		// FIXME remove custom region
		gucumber.World["client"] = devicefarm.New(smoke.Session,
			nifcloud.NewConfig().WithRegion("us-west-2"))
	})
}
