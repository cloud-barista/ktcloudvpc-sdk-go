package limits

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// Get returns the limits about the currently scoped tenant.
func Get(client *ktvpcsdk.ServiceClient) (r GetResult) {
	url := getURL(client)
	resp, err := client.Get(url, &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
