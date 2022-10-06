package testing

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := ktvpcsdk.EndpointOpts{Availability: ktvpcsdk.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := ktvpcsdk.EndpointOpts{Availability: ktvpcsdk.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = ktvpcsdk.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = ktvpcsdk.EndpointOpts{Availability: ktvpcsdk.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
