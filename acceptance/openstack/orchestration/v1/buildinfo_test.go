//go:build acceptance
// +build acceptance

package v1

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/orchestration/v1/buildinfo"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestBuildInfo(t *testing.T) {
	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	bi, err := buildinfo.Get(client).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved build info: %+v\n", bi)
}
