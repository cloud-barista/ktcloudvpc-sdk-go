package floatingips

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFloatingIPListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Page     	int    `q:"page"`
    Size     	int    `q:"size"`
	CIDR 	 	string `q:"cidr"`
	PupblicIpID string `q:"publicIpId"`
}

// ToNetworkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFloatingIPListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// floating IP resources. It accepts a ListOpts struct, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToFloatingIPListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return FloatingIPPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToFloatingIPCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies a Floating IP allocation request.
type CreateOpts struct {
}

// ToFloatingIPCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToFloatingIPCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new Floating IP.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFloatingIPCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get returns data about a previously created Floating IP.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete requests the deletion of a previous allocated Floating IP.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
