package testing

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/openstack/common/extensions"
	"github.com/innodreamer/ktvpc-sdk_poc/pagination"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
	"github.com/innodreamer/ktvpc-sdk_poc/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListExtensionsSuccessfully(t)

	count := 0

	extensions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := extensions.ExtractExtensions(page)
		th.AssertNoErr(t, err)
		th.AssertDeepEquals(t, ExpectedExtensions, actual)

		return true, nil
	})

	th.CheckEquals(t, 1, count)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetExtensionSuccessfully(t)

	actual, err := extensions.Get(client.ServiceClient(), "agent").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SingleExtension, actual)
}
