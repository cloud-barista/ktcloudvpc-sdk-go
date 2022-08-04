package apiversions

import (
	"github.com/innodreamer/ktvpc-sdk_poc"
	"github.com/innodreamer/ktvpc-sdk_poc/pagination"
)

// List lists all the load balancer API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
