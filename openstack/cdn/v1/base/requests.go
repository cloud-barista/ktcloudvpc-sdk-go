package base

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// Get retrieves the home document, allowing the user to discover the
// entire API.
func Get(c *ktvpcsdk.ServiceClient) (r GetResult) {
	resp, err := c.Get(getURL(c), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Ping retrieves a ping to the server.
func Ping(c *ktvpcsdk.ServiceClient) (r PingResult) {
	resp, err := c.Get(pingURL(c), nil, &ktvpcsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Accept": ""},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
