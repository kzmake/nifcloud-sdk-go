// +build integration

//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/workspaces"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@workspaces", func() {
		gucumber.World["client"] = workspaces.New(smoke.Session)
	})
}
