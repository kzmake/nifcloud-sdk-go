// +build integration

//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/kzmake/nifcloud-sdk-go/nifcloud"
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/efs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@efs", func() {
		// FIXME remove custom region
		gucumber.World["client"] = efs.New(smoke.Session,
			nifcloud.NewConfig().WithRegion("us-west-2"))
	})
}
