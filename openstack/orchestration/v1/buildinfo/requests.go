package buildinfo

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// Get retreives data for the given stack template.
func Get(c *ktvpcsdk.ServiceClient) (r GetResult) {
	resp, err := c.Get(getURL(c), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
