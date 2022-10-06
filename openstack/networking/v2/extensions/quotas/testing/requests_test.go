package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	fake "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/networking/v2/common"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/networking/v2/extensions/quotas"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponseRaw)
	})

	q, err := quotas.Get(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &GetResponse)
}

func TestGetDetail(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76/details.json", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetDetailedResponseRaw)
	})

	q, err := quotas.GetDetail(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &GetDetailResponse)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateRequestResponseRaw)
	})

	q, err := quotas.Update(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76", quotas.UpdateOpts{
		FloatingIP:        ktvpcsdk.IntToPointer(0),
		Network:           ktvpcsdk.IntToPointer(-1),
		Port:              ktvpcsdk.IntToPointer(5),
		RBACPolicy:        ktvpcsdk.IntToPointer(10),
		Router:            ktvpcsdk.IntToPointer(15),
		SecurityGroup:     ktvpcsdk.IntToPointer(20),
		SecurityGroupRule: ktvpcsdk.IntToPointer(-1),
		Subnet:            ktvpcsdk.IntToPointer(25),
		SubnetPool:        ktvpcsdk.IntToPointer(0),
		Trunk:             ktvpcsdk.IntToPointer(5),
	}).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &UpdateResponse)
}
