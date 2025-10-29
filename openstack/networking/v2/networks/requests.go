// ### KT Cloud D1 platfrom > 'VPC' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/networking

package networks

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
	cblogger = cblog.GetLogger("KTCloud VPC Client")
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
    ToVPCListQuery() (string, error)
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
    ToVPCCreateMap() (map[string]interface{}, error)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
    ToVPCUpdateMap() (map[string]interface{}, error)
}

// ListOpts allows filtering and sorting of paginated collections.
type ListOpts struct {                   // Modified
    Name     string `q:"name"`
    VpcID    string `q:"vpcId"`
    ZoneID   string `q:"zoneId"`
    Page     int    `q:"page"`
    Size     int    `q:"size"`
    Type     string `q:"type"`
    State    string `q:"state"`
}

// ToVPCListQuery formats a ListVPCV2Opts into a query string.
func (opts ListOpts) ToVPCListQuery() (string, error) {
    q, err := gophercloud.BuildQueryString(opts)
    return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of VPCs.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
    url := listURL(c)
    if opts != nil {
        query, err := opts.ToVPCListQuery()
        if err != nil {
            return pagination.Pager{Err: err}
        }
        url += query
    }
    return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
        return VPCPage{pagination.LinkedPageBase{PageResult: r}}
    })
}

// Get retrieves a specific VPC based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
    cblogger.Infof("# Get URL : %s\n", getURL(c, id))

    resp, err := c.Get(getURL(c, id), &r.Body, nil)
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}

// CreateOpts represents the options for creating a new VPC.
type CreateOpts struct {                    // Modified
    Name          string  `json:"name" required:"true"`
    ZoneID        string  `json:"zoneId" required:"true"`
    BizareaType   string  `json:"bizareaType,omitempty"`
    VpcOfferingID string  `json:"vpcOfferingId,omitempty"`
    Description   *string `json:"description,omitempty"`
    Type          string  `json:"type,omitempty"`
    SessionCount  *int    `json:"sessionCount,omitempty"`
}

// ToVPCCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToVPCCreateMap() (map[string]interface{}, error) {
    return gophercloud.BuildRequestBody(opts, "") // ""
}

// Create accepts a CreateOpts struct and creates a new VPC.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
    b, err := opts.ToVPCCreateMap()
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

// Update represents the options for updating an existing VPC.
type UpdateOpts struct {
    Name         *string `json:"name,omitempty"`
    Description  *string `json:"description,omitempty"`
    SessionCount *int    `json:"sessionCount,omitempty"`
    State        *string `json:"state,omitempty"`
}

// ToVPCUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToVPCUpdateMap() (map[string]interface{}, error) {
    return gophercloud.BuildRequestBody(opts, "")
}

// Update accepts a UpdateOpts struct and updates an existing VPC.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r commonResult) {
    b, err := opts.ToVPCUpdateMap()
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

// Delete accepts a unique ID and deletes the VPC associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
    resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
        OkCodes: []int{200, 204},
    })
    _, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
    return
}
