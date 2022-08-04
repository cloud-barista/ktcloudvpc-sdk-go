//go:build acceptance || networking || lbaas || vip
// +build acceptance networking lbaas vip

package lbaas

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	networking "github.com/innodreamer/ktvpc-sdk_poc/acceptance/openstack/networking/v2"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/networking/v2/extensions/lbaas/vips"
)

func TestVIPsList(t *testing.T) {
	t.Skip("Neutron LBaaS was replaced by Octavia and the API will be removed in a future release")
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	allPages, err := vips.List(client, vips.ListOpts{}).AllPages()
	if err != nil {
		t.Fatalf("Unable to list vips: %v", err)
	}

	allVIPs, err := vips.ExtractVIPs(allPages)
	if err != nil {
		t.Fatalf("Unable to extract vips: %v", err)
	}

	for _, vip := range allVIPs {
		tools.PrintResource(t, vip)
	}
}

func TestVIPsCRUD(t *testing.T) {
	t.Skip("Neutron LBaaS was replaced by Octavia and the API will be removed in a future release")
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	network, err := networking.CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer networking.DeleteNetwork(t, client, network.ID)

	subnet, err := networking.CreateSubnet(t, client, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer networking.DeleteSubnet(t, client, subnet.ID)

	pool, err := CreatePool(t, client, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create pool: %v", err)
	}
	defer DeletePool(t, client, pool.ID)

	vip, err := CreateVIP(t, client, subnet.ID, pool.ID)
	if err != nil {
		t.Fatalf("Unable to create vip: %v", err)
	}
	defer DeleteVIP(t, client, vip.ID)

	tools.PrintResource(t, vip)

	connLimit := 100
	updateOpts := vips.UpdateOpts{
		ConnLimit: &connLimit,
	}

	_, err = vips.Update(client, vip.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update vip: %v", err)
	}

	newVIP, err := vips.Get(client, vip.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get vip: %v", err)
	}

	tools.PrintResource(t, newVIP)
}
