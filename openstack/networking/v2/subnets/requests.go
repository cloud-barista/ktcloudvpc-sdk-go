// ### KT Cloud D1 platform > 'Tier' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/tier

package subnets

import (
	// "fmt"
    // "github.com/davecgh/go-spew/spew"
	cblog "github.com/cloud-barista/cb-log"	
	"github.com/sirupsen/logrus"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"  
    "github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

var cblogger *logrus.Logger

func init() {
	cblogger = cblog.GetLogger("KTCloud Subnet Client")
}

// ListOptsBuilder allows extensions to add additional parameters to the ListSubnet request.
type ListOptsBuilder interface {
    ToSubnetListQuery() (string, error)
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
    ToSubnetCreateMap() (map[string]interface{}, error)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
    ToSubnetUpdateMap() (map[string]interface{}, error)
}

// ListOpts allows filtering and sorting of paginated collections.
type ListOpts struct {                  // Modified
    Page            int    `q:"page"`               // [default: 1] page number
    Size            int    `q:"size"`               // [default: 20] page size
    NetworkType     string `q:"networkType"`        // [default: trust] TRUST / UNTRUST / ALL (External connection type of Tier.)
                                                    // TRUST: Internal connection, UNTRUST: External connection
    IsPrivateSubnet *bool  `q:"isPrivateSubnet"`    // Whether it is a private subnet network
    NetworkID       string `q:"networkId"`          // KT Cloud Tier's UUID
    RefID           string `q:"refId"`              // KT Cloud Tier ID based on OpenStack Neutron
}

// ToSubnetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSubnetListQuery() (string, error) {
    q, err := gophercloud.BuildQueryString(opts)
    return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of Subnets.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
    url := listURL(c)
    if opts != nil {
        query, err := opts.ToSubnetListQuery()
        if err != nil {
            return pagination.Pager{Err: err}
        }
        url += query
    }
    return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
        return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
    })
}

// Get retrieves a specific Subnet based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
    cblogger.Infof("# Get Subnet URL : %s\n", getURL(c, id))

    resp, err := c.Get(getURL(c, id), &r.Body, nil)
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}

// SubnetDetail represents the detailed configuration for subnet creation.
type SubnetDetail struct {                  // Modified
    CIDR          string `json:"cidr" required:"true"`
    GatewayIP     string `json:"gatewayIp" required:"true"`
    StartIP       string `json:"startIp" required:"true"`
    EndIP         string `json:"endIp" required:"true"`
    LBStartIP     string `json:"lbStartIp,omitempty"`
    LBEndIP       string `json:"lbEndIp,omitempty"`
    BMStartIP     string `json:"bmStartIp,omitempty"`
    BMEndIP       string `json:"bmEndIp,omitempty"`
    IscsiStartIP  string `json:"iscsiStartIp,omitempty"`
    IscsiEndIP    string `json:"iscsiEndIp,omitempty"`
}

// CreateOpts represents the options for creating a new Subnet.
type CreateOpts struct {                    // Modified
    Name     string       `json:"name" required:"true"`
    Type     string       `json:"type" required:"true"`
    IsCustom bool         `json:"isCustom" required:"true"`
    Detail   SubnetDetail `json:"detail" required:"true"`
}

// ToSubnetCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
    b, err := gophercloud.BuildRequestBody(opts, "")  // Caution!!) Error msg
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Create accepts a CreateOpts struct and creates a new Subnet.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
    b, err := opts.ToSubnetCreateMap()
    if err != nil {
        r.Err = err
        return
    }
    resp, err := c.Post(createURL(c), b, &r.Body, &gophercloud.RequestOpts{ // Modified
        OkCodes: []int{200, 201},
    })
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}

// UpdateOpts represents the options for updating an existing Subnet.
type UpdateOpts struct {
    Name        *string `json:"name,omitempty"`
    Description *string `json:"description,omitempty"`
    Gateway     *string `json:"gateway,omitempty"`
    Netmask     *string `json:"netmask,omitempty"`
}

// ToSubnetUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToSubnetUpdateMap() (map[string]interface{}, error) {
    b, err := gophercloud.BuildRequestBody(opts, "")  // Caution!!) Error msg
	if err != nil {
		return nil, err
	}
	return b, nil
}

// UpdateSubnet accepts a UpdateOpts struct and updates an existing Subnet.
func UpdateSubnet(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r commonResult) {
    b, err := opts.ToSubnetUpdateMap()
    if err != nil {
        r.Err = err
        return
    }
    resp, err := c.Put(updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
        OkCodes: []int{200, 202},
    })
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}

// Delete accepts a unique ID and deletes the Subnet associated with it.
// ### Need NetworkId, not TierId.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
    resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{ // Modified  
        OkCodes: []int{200, 202, 204},
    })
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}
