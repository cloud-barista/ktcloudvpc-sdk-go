package testing

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/openstack/sharedfilesystems/v2/services"
	"github.com/innodreamer/ktvpc-sdk_poc/pagination"
	"github.com/innodreamer/ktvpc-sdk_poc/testhelper"
	"github.com/innodreamer/ktvpc-sdk_poc/testhelper/client"
)

func TestListServices(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleListSuccessfully(t)

	pages := 0
	err := services.List(client.ServiceClient(), services.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := services.ExtractServices(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 services, got %d", len(actual))
		}
		testhelper.CheckDeepEquals(t, FirstFakeService, actual[0])
		testhelper.CheckDeepEquals(t, SecondFakeService, actual[1])

		return true, nil
	})

	testhelper.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}
