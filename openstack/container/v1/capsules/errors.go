package capsules

import (
	"fmt"

	"github.com/innodreamer/ktvpc-sdk_poc"
)

type ErrInvalidDataFormat struct {
	gophercloud.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return fmt.Sprintf("Data in neither json nor yaml format.")
}
