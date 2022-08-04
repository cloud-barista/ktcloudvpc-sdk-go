package resetstate

import (
	"github.com/innodreamer/ktvpc-sdk_poc"
)

// ResetResult is the response of a ResetState operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type ResetResult struct {
	gophercloud.ErrResult
}
