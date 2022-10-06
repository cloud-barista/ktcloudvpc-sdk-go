package pauseunpause

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// PauseResult is the response from a Pause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type PauseResult struct {
	ktvpcsdk.ErrResult
}

// UnpauseResult is the response from an Unpause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type UnpauseResult struct {
	ktvpcsdk.ErrResult
}
