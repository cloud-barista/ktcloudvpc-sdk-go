package extensions

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	common "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/common/extensions"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// ExtractExtensions interprets a Page as a slice of Extensions.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	return common.ExtractExtensions(page)
}

// Get retrieves information for a specific extension using its alias.
func Get(c *ktvpcsdk.ServiceClient, alias string) common.GetResult {
	return common.Get(c, alias)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *ktvpcsdk.ServiceClient) pagination.Pager {
	return common.List(c)
}
