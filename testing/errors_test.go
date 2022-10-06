package testing

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func returnsUnexpectedResp(code int) ktvpcsdk.ErrUnexpectedResponseCode {
	return ktvpcsdk.ErrUnexpectedResponseCode{
		URL:            "http://example.com",
		Method:         "GET",
		Expected:       []int{200},
		Actual:         code,
		Body:           nil,
		ResponseHeader: nil,
	}
}

func TestGetResponseCode404(t *testing.T) {
	var err404 error = ktvpcsdk.ErrDefault404{ErrUnexpectedResponseCode: returnsUnexpectedResp(404)}

	err, ok := err404.(ktvpcsdk.StatusCodeError)
	th.AssertEquals(t, true, ok)
	th.AssertEquals(t, err.GetStatusCode(), 404)
}

func TestGetResponseCode502(t *testing.T) {
	var err502 error = ktvpcsdk.ErrDefault502{ErrUnexpectedResponseCode: returnsUnexpectedResp(502)}

	err, ok := err502.(ktvpcsdk.StatusCodeError)
	th.AssertEquals(t, true, ok)
	th.AssertEquals(t, err.GetStatusCode(), 502)
}

func TestGetResponseCode504(t *testing.T) {
	var err504 error = ktvpcsdk.ErrDefault504{ErrUnexpectedResponseCode: returnsUnexpectedResp(504)}

	err, ok := err504.(ktvpcsdk.StatusCodeError)
	th.AssertEquals(t, true, ok)
	th.AssertEquals(t, err.GetStatusCode(), 504)
}
