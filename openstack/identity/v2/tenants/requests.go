package tenants

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// ListOpts filters the Tenants that are returned by the List call.
type ListOpts struct {
	// Marker is the ID of the last Tenant on the previous page.
	Marker string `q:"marker"`

	// Limit specifies the page size.
	Limit int `q:"limit"`
}

// List enumerates the Tenants to which the current token has access.
func List(client *ktvpcsdk.ServiceClient, opts *ListOpts) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		q, err := ktvpcsdk.BuildQueryString(opts)
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q.String()
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TenantPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOpts represents the options needed when creating new tenant.
type CreateOpts struct {
	// Name is the name of the tenant.
	Name string `json:"name" required:"true"`

	// Description is the description of the tenant.
	Description string `json:"description,omitempty"`

	// Enabled sets the tenant status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// CreateOptsBuilder enables extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToTenantCreateMap() (map[string]interface{}, error)
}

// ToTenantCreateMap assembles a request body based on the contents of
// a CreateOpts.
func (opts CreateOpts) ToTenantCreateMap() (map[string]interface{}, error) {
	return ktvpcsdk.BuildRequestBody(opts, "tenant")
}

// Create is the operation responsible for creating new tenant.
func Create(client *ktvpcsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTenantCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), b, &r.Body, &ktvpcsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Get requests details on a single tenant by ID.
func Get(client *ktvpcsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToTenantUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the base attributes that may be updated on an existing
// tenant.
type UpdateOpts struct {
	// Name is the name of the tenant.
	Name string `json:"name,omitempty"`

	// Description is the description of the tenant.
	Description *string `json:"description,omitempty"`

	// Enabled sets the tenant status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// ToTenantUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToTenantUpdateMap() (map[string]interface{}, error) {
	return ktvpcsdk.BuildRequestBody(opts, "tenant")
}

// Update is the operation responsible for updating exist tenants by their TenantID.
func Update(client *ktvpcsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTenantUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(updateURL(client, id), &b, &r.Body, &ktvpcsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Delete is the operation responsible for permanently deleting a tenant.
func Delete(client *ktvpcsdk.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
