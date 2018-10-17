// +build integration

//Package codecommit provides gucumber integration tests support.
package codecommit

import (
	"github.com/kzmake/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/kzmake/nifcloud-sdk-go/service/codecommit"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codecommit", func() {
		gucumber.World["client"] = codecommit.New(smoke.Session)
	})
}
