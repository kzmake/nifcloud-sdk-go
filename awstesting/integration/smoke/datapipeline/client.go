// +build integration

//Package datapipeline provides gucumber integration tests support.
package datapipeline

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/datapipeline"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@datapipeline", func() {
		gucumber.World["client"] = datapipeline.New(smoke.Session)
	})
}
