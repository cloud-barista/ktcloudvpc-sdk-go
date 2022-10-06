package limits

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToLimitListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Filters the response by a region ID.
	RegionID string `q:"region_id"`

	// Filters the response by a project ID.
	ProjectID string `q:"project_id"`

	// Filters the response by a domain ID.
	DomainID string `q:"domain_id"`

	// Filters the response by a service ID.
	ServiceID string `q:"service_id"`

	// Filters the response by a resource name.
	ResourceName string `q:"resource_name"`
}

// ToLimitListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLimitListQuery() (string, error) {
	q, err := ktvpcsdk.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the limits.
func List(client *ktvpcsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToLimitListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return LimitPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
