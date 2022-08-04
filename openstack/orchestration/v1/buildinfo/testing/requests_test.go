package testing

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/openstack/orchestration/v1/buildinfo"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
	fake "github.com/innodreamer/ktvpc-sdk_poc/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := buildinfo.Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
