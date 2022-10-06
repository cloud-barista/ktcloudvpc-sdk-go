package v1

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/placement/v1/resourceproviders"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func CreateResourceProvider(t *testing.T, client *ktvpcsdk.ServiceClient) (*resourceproviders.ResourceProvider, error) {
	name := tools.RandomString("TESTACC-", 8)
	t.Logf("Attempting to create resource provider: %s", name)

	createOpts := resourceproviders.CreateOpts{
		Name: name,
	}

	client.Microversion = "1.20"
	resourceProvider, err := resourceproviders.Create(client, createOpts).Extract()
	if err != nil {
		return resourceProvider, err
	}

	t.Logf("Successfully created resourceProvider: %s.", resourceProvider.Name)
	tools.PrintResource(t, resourceProvider)

	th.AssertEquals(t, resourceProvider.Name, name)

	return resourceProvider, nil
}

func CreateResourceProviderWithParent(t *testing.T, client *ktvpcsdk.ServiceClient, parentUUID string) (*resourceproviders.ResourceProvider, error) {
	name := tools.RandomString("TESTACC-", 8)
	t.Logf("Attempting to create resource provider: %s", name)

	createOpts := resourceproviders.CreateOpts{
		Name:               name,
		ParentProviderUUID: parentUUID,
	}

	client.Microversion = "1.20"
	resourceProvider, err := resourceproviders.Create(client, createOpts).Extract()
	if err != nil {
		return resourceProvider, err
	}

	t.Logf("Successfully created resourceProvider: %s.", resourceProvider.Name)
	tools.PrintResource(t, resourceProvider)

	th.AssertEquals(t, resourceProvider.Name, name)
	th.AssertEquals(t, resourceProvider.ParentProviderUUID, parentUUID)

	return resourceProvider, nil
}

// DeleteResourceProvider will delete a resource provider with a specified ID.
// A fatal error will occur if the delete was not successful. This works best when
// used as a deferred function.
func DeleteResourceProvider(t *testing.T, client *ktvpcsdk.ServiceClient, resourceProviderID string) {
	t.Logf("Attempting to delete resourceProvider: %s", resourceProviderID)

	err := resourceproviders.Delete(client, resourceProviderID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to delete resourceProvider %s: %v", resourceProviderID, err)
	}

	t.Logf("Deleted resourceProvider: %s.", resourceProviderID)
}
