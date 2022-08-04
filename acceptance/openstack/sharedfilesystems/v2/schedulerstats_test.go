//go:build acceptance
// +build acceptance

package v2

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/sharedfilesystems/v2/schedulerstats"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
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
