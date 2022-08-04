package resetnetwork

import (
	"github.com/innodreamer/ktvpc-sdk_poc"
)

// ResetResult is the response of a ResetNetwork operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type ResetResult struct {
	gophercloud.ErrResult
}
