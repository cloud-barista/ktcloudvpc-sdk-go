//go:build acceptance || blockstorage
// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/blockstorage/extensions/schedulerhints"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/blockstorage/v3/volumes"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestSchedulerHints(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volumeName := tools.RandomString("ACPTTEST", 16)
	createOpts := volumes.CreateOpts{
		Size: 1,
		Name: volumeName,
	}

	volume1, err := volumes.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, volume1.ID, "available", 60)
	th.AssertNoErr(t, err)
	defer volumes.Delete(client, volume1.ID, volumes.DeleteOpts{})

	volumeName = tools.RandomString("ACPTTEST", 16)
	base := volumes.CreateOpts{
		Size: 1,
		Name: volumeName,
	}

	schedulerHints := schedulerhints.SchedulerHints{
		SameHost: []string{
			volume1.ID,
		},
	}

	createOptsWithHints := schedulerhints.CreateOptsExt{
		VolumeCreateOptsBuilder: base,
		SchedulerHints:          schedulerHints,
	}

	volume2, err := volumes.Create(client, createOptsWithHints).Extract()
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, volume2.ID, "available", 60)
	th.AssertNoErr(t, err)

	err = volumes.Delete(client, volume2.ID, volumes.DeleteOpts{}).ExtractErr()
	th.AssertNoErr(t, err)
}
