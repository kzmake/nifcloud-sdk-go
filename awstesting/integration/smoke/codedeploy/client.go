// +build integration

//Package codedeploy provides gucumber integration tests support.
package codedeploy

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/codedeploy"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codedeploy", func() {
		gucumber.World["client"] = codedeploy.New(smoke.Session)
	})
}
