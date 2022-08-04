package apiversions

import (
	"github.com/innodreamer/ktvpc-sdk_poc"
	"github.com/innodreamer/ktvpc-sdk_poc/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users
func ListVersions(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
