package shelveunshelve

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// ShelveResult is the response from a Shelve operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type ShelveResult struct {
	ktvpcsdk.ErrResult
}

// ShelveOffloadResult is the response from a Shelve operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type ShelveOffloadResult struct {
	ktvpcsdk.ErrResult
}

// UnshelveResult is the response from Stop operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type UnshelveResult struct {
	ktvpcsdk.ErrResult
}
