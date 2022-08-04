//go:build acceptance
// +build acceptance

package v2

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/sharedfilesystems/v2/schedulerstats"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestSchedulerStatsList(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	client.Microversion = "2.23"
	th.AssertNoErr(t, err)

	allPages, err := schedulerstats.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allPools, err := schedulerstats.ExtractPools(allPages)
	th.AssertNoErr(t, err)

	for _, recordset := range allPools {
		tools.PrintResource(t, &recordset)
	}
}
