package extensions

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// Get retrieves information for a specific extension using its alias.
func Get(c *ktvpcsdk.ServiceClient, alias string) (r GetResult) {
	resp, err := c.Get(ExtensionURL(c, alias), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *ktvpcsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, ListExtensionURL(c), func(r pagination.PageResult) pagination.Page {
		return ExtensionPage{pagination.SinglePageBase(r)}
	})
}
