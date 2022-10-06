package projectendpoints

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

type CreateOptsBuilder interface {
	ToEndpointCreateMap() (map[string]interface{}, error)
}

// Create inserts a new Endpoint association to a project.
func Create(client *ktvpcsdk.ServiceClient, projectID, endpointID string) (r CreateResult) {
	resp, err := client.Put(createURL(client, projectID, endpointID), nil, nil, &ktvpcsdk.RequestOpts{OkCodes: []int{204}})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// List enumerates endpoints in a paginated collection, optionally filtered
// by ListOpts criteria.
func List(client *ktvpcsdk.ServiceClient, projectID string) pagination.Pager {
	u := listURL(client, projectID)
	return pagination.NewPager(client, u, func(r pagination.PageResult) pagination.Page {
		return EndpointPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Delete removes an endpoint from the service catalog.
func Delete(client *ktvpcsdk.ServiceClient, projectID string, endpointID string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, projectID, endpointID), &ktvpcsdk.RequestOpts{OkCodes: []int{204}})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
