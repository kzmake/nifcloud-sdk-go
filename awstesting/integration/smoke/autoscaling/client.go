// +build integration

//Package autoscaling provides gucumber integration tests support.
package autoscaling

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/autoscaling"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@autoscaling", func() {
		gucumber.World["client"] = autoscaling.New(smoke.Session)
	})
}
