//go:build acceptance || networking || loadbalancer || quotas
// +build acceptance networking loadbalancer quotas

package v2

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/loadbalancer/v2/quotas"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func TestQuotasGet(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewLoadBalancerV2Client()
	th.AssertNoErr(t, err)

	quotasInfo, err := quotas.Get(client, os.Getenv("OS_PROJECT_NAME")).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotasInfo)
}

func TestQuotasUpdate(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewLoadBalancerV2Client()
	th.AssertNoErr(t, err)

	originalQuotas, err := quotas.Get(client, os.Getenv("OS_PROJECT_NAME")).Extract()
	th.AssertNoErr(t, err)

	var quotaUpdateOpts = quotas.UpdateOpts{
		Loadbalancer:  gophercloud.IntToPointer(25),
		Listener:      gophercloud.IntToPointer(45),
		Member:        gophercloud.IntToPointer(205),
		Pool:          gophercloud.IntToPointer(25),
		Healthmonitor: gophercloud.IntToPointer(5),
	}
	// L7 parameters are only supported in microversion v2.19 introduced in victoria
	if clients.IsReleasesAbove(t, "stable/ussuri") {
		quotaUpdateOpts.L7Policy = gophercloud.IntToPointer(55)
		quotaUpdateOpts.L7Rule = gophercloud.IntToPointer(105)
	}

	newQuotas, err := quotas.Update(client, os.Getenv("OS_PROJECT_NAME"), quotaUpdateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newQuotas)

	if reflect.DeepEqual(originalQuotas, newQuotas) {
		log.Fatal("Original and New Loadbalancer Quotas are the same")
	}

	var restoredQuotaUpdate = quotas.UpdateOpts{
		Loadbalancer:  &originalQuotas.Loadbalancer,
		Listener:      &originalQuotas.Listener,
		Member:        &originalQuotas.Member,
		Pool:          &originalQuotas.Pool,
		Healthmonitor: &originalQuotas.Healthmonitor,
	}
	// L7 parameters are only supported in microversion v2.19 introduced in victoria
	if clients.IsReleasesAbove(t, "stable/ussuri") {
		restoredQuotaUpdate.L7Policy = &originalQuotas.L7Policy
		restoredQuotaUpdate.L7Rule = &originalQuotas.L7Rule
	}

	// Restore original quotas.
	restoredQuotas, err := quotas.Update(client, os.Getenv("OS_PROJECT_NAME"), restoredQuotaUpdate).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, originalQuotas, restoredQuotas)

	tools.PrintResource(t, restoredQuotas)
}
