//go:build acceptance || identity
// +build acceptance identity

package v2

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/identity/v2/extensions"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestExtensionsList(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2Client()
	th.AssertNoErr(t, err)

	allPages, err := extensions.List(client).AllPages()
	th.AssertNoErr(t, err)

	allExtensions, err := extensions.ExtractExtensions(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, extension := range allExtensions {
		tools.PrintResource(t, extension)
		if extension.Name == "OS-KSCRUD" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestExtensionsGet(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2Client()
	th.AssertNoErr(t, err)

	extension, err := extensions.Get(client, "OS-KSCRUD").Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, extension)
}
