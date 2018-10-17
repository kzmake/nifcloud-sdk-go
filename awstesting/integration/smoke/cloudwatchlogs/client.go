// +build integration

//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/cloudwatchlogs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudwatchlogs", func() {
		gucumber.World["client"] = cloudwatchlogs.New(smoke.Session)
	})
}
