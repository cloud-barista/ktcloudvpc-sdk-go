//go:build acceptance || networking || loadbalancer || pools
// +build acceptance networking loadbalancer pools

package lbaas_v2

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/networking/v2/extensions/lbaas_v2/pools"
)

func TestPoolsList(t *testing.T) {
	t.Skip("Neutron LBaaS v2 was replaced by Octavia and the API will be removed in a future release")
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	allPages, err := pools.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list pools: %v", err)
	}

	allPools, err := pools.ExtractPools(allPages)
	if err != nil {
		t.Fatalf("Unable to extract pools: %v", err)
	}

	for _, pool := range allPools {
		tools.PrintResource(t, pool)
	}
}
