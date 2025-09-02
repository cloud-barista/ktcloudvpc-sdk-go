package portforwarding

import (
	cblog "github.com/cloud-barista/cb-log"	
	"github.com/sirupsen/logrus"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

var cblogger *logrus.Logger

func init() {
	cblogger = cblog.GetLogger("KTCloud VPC Client")
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToPortForwardingListQuery() (string, error)
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
    ToPortForwardingCreateMap() (map[string]interface{}, error)
}

// ListOpts allows the filtering and sorting of paginated collections through the API. 
// ListOpts allows filtering and sorting of paginated collections.
type ListOpts struct {                   // Modified
	// Pagination options	
    Page     int    `q:"page"`
    Size     int    `q:"size"`

	// MappedIP filters by the mapped private IP
	MappedIP string `q:"mappedIp"`

	// Protocol filters by protocol (TCP, UDP)
	Protocol string `q:"protocol"`	

	// // ID filters by the port forwarding rule ID
	// ID string `q:"id"`
	
	// // Name filters by the port forwarding rule name
	// Name string `q:"name"`	

	// // PublicIP filters by public IP address
	// PublicIP string `q:"public_ip"`
	
	// // PublicIPID filters by public IP ID
	// PublicIPID string `q:"public_ip_id"`	

	// // StackType filters by stack type
	// StackType string `q:"stack_type"`
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
	cblogger.Infof("# List URL : %s\n", listURL(c))

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
func Get(c *gophercloud.ServiceClient, floatingIpId string, pfId string) (r GetResult) {	// Modified
	cblogger.Infof("# Get URL : %s\n", getURL(c, pfId))

	resp, err := c.Get(getURL(c, pfId), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOpts contains all the values needed to create a new port forwarding resource. All attributes are required.
type CreateOpts struct {																		// Modified
	// PublicIPID is the ID of the public IP to use
	PublicIPID 		string `json:"publicipid" required:"true"`
	
	// MappedIP is the private IP address to forward to
	MappedIP 		string `json:"mappedip" required:"true"`
	
	// Protocol specifies the protocol (TCP, UDP)
	Protocol 		string `json:"protocol" required:"true"`

	// StartPrivatePort is the starting private port number
	StartPrivatePort string `json:"startPrivatePort" required:"true"`
	
	// EndPrivatePort is the ending private port number
	EndPrivatePort 	string `json:"endPrivatePort" required:"true"`

	// StartPublicPort is the starting public port number
	StartPublicPort string `json:"startPublicPort" required:"true"`
	
	// EndPublicPort is the ending public port number
	EndPublicPort 	string `json:"endPublicPort" required:"true"`
}

// ToPortForwardingCreateMap allows CreateOpts to satisfy the CreateOptsBuilder
// interface
func (opts CreateOpts) ToPortForwardingCreateMap() (map[string]interface{}, error) {			// Modified
	return gophercloud.BuildRequestBody(opts, "")	
}

// Create accepts a CreateOpts struct and uses the values provided to create a new port forwarding for an existing floating IP.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {		// Modified
	cblogger.Infof("# Create URL : %s\n", createURL(c))

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
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {								// Modified
	resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
