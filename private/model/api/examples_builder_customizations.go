// +build codegen

package api

import (
	"bytes"
	"fmt"
)

type wafregionalExamplesBuilder struct {
	defaultExamplesBuilder
}

func (builder wafregionalExamplesBuilder) Imports(a *API) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`"fmt"
	"strings"
	"time"

	"github.com/kzmake/nifcloud-sdk-go/nifcloud"
	"github.com/kzmake/nifcloud-sdk-go/nifcloud/awserr"
	"github.com/kzmake/nifcloud-sdk-go/nifcloud/session"
	"github.com/kzmake/nifcloud-sdk-go/service/waf"
	`)

	buf.WriteString(fmt.Sprintf("\"%s/%s\"", "github.com/kzmake/nifcloud-sdk-go/service", a.PackageName()))
	return buf.String()
}
