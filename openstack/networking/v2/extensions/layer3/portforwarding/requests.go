package portforwarding

import (
	// "github.com/davecgh/go-spew/spew"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToPortForwardingListQuery() (string, error)
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
    ToPortForwardingCreateMap() (map[string]interface{}, error)
}

// ListOpts allows the filtering and sorting of paginated collections through the API. 
type ListOpts struct {                   
    // Page number for pagination (default: 1)
    Page int `q:"page"`

    // Number of items per page (default: 20)
    Size int `q:"size"`

    // Private IP address mapped to the server
    MappedIP string `q:"mappedIp"`

    // Port forwarding protocol (e.g., TCP, UDP)
    Protocol string `q:"protocol"`

    // Starting private IP port number
    StartPrivatePort string `q:"startPrivatePort"`

    // Ending private IP port number
    EndPrivatePort string `q:"endPrivatePort"`

    // Starting public IP port number
    StartPublicPort string `q:"startPublicPort"`

    // Ending public IP port number
    EndPublicPort string `q:"endPublicPort"`

    // Port forwarding rule ID
    PortForwardingID string `q:"portForwardingId"`
}

// ToPortForwardingListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPortForwardingListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Port Forwarding resources. It accepts a ListOpts struct, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToPortForwardingListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PortForwardingPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a particular port forwarding resource based on its unique ID.
func Get(c *gophercloud.ServiceClient, floatingIpId string, pfId string) (r GetResult) {	
	resp, err := c.Get(getURL(c, pfId), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOpts contains all the values needed to create a new port forwarding resource. All attributes are required.
type CreateOpts struct {																		
	// [Required] PublicIPID is the ID of the public IP to use
	PublicIpID 		string `json:"publicIpId" required:"true"`
	
	// [Required] MappedIP is the private IP address to forward to
	MappedIP 		string `json:"mappedIp" required:"true"`
	
	// [Required] Protocol specifies the protocol (TCP, UDP)
	Protocol 		string `json:"protocol" required:"true"`

	// [Required] StartPrivatePort is the starting private port number
	StartPrivatePort string `json:"startPrivatePort" required:"true"`
	
	// [Required] EndPrivatePort is the ending private port number
	EndPrivatePort 	string `json:"endPrivatePort" required:"true"`

	// [Required] StartPublicPort is the starting public port number
	StartPublicPort string `json:"startPublicPort" required:"true"`
	
	// [Required] EndPublicPort is the ending public port number
	EndPublicPort 	string `json:"endPublicPort" required:"true"`
}

// ToPortForwardingCreateMap allows CreateOpts to satisfy the CreateOptsBuilder
// interface
func (opts CreateOpts) ToPortForwardingCreateMap() (map[string]interface{}, error) {			
	return gophercloud.BuildRequestBody(opts, "")  // ""	
}

// Create accepts a CreateOpts struct and uses the values provided to create a new port forwarding for an existing floating IP.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {		
	b, err := opts.ToPortForwardingCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Post(createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete will permanently delete a particular port forwarding for a given floating ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {								
	resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
