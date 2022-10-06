package errors

import (
	"encoding/json"
	"fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

type ManilaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type ErrorDetails map[string]ManilaError

// error types from provider_client.go
func ExtractErrorInto(rawError error, errorDetails *ErrorDetails) (err error) {
	switch e := rawError.(type) {
	case ktvpcsdk.ErrDefault400:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault401:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault403:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault404:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault405:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault408:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault429:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault500:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	case ktvpcsdk.ErrDefault503:
		err = json.Unmarshal(e.ErrUnexpectedResponseCode.Body, errorDetails)
	default:
		err = fmt.Errorf("Unable to extract detailed error message")
	}

	return err
}
