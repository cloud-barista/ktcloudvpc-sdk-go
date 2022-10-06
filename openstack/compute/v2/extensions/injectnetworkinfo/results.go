package injectnetworkinfo

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// InjectNetworkResult is the response of a InjectNetworkInfo operation. Call
// its ExtractErr method to determine if the request suceeded or failed.
type InjectNetworkResult struct {
	ktvpcsdk.ErrResult
}
