package suspendresume

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// SuspendResult is the response from a Suspend operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type SuspendResult struct {
	ktvpcsdk.ErrResult
}

// UnsuspendResult is the response from an Unsuspend operation. Call
// its ExtractErr method to determine if the request succeeded or failed.
type UnsuspendResult struct {
	ktvpcsdk.ErrResult
}
