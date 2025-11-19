// ### KT Cloud D1 platform > 'Tier' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/tier

package subnets

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
    "github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

type Subnet struct {                    // Modified
    AccountID        string  `json:"accountId"`
    NetworkID        string  `json:"networkId"`   // KT Cloud Tier's UUID (Note: differs from the Tier ID based on OpenStack Neutron.) 
    ZoneID           string  `json:"zoneId"`      // Zone UUID  
    NetworkType      string  `json:"networkType"` // External connection type of tier. TRUST: Internal connection, UNTRUST: External connection
    StackType        string  `json:"stackType"`   // OS : OpenStack, CS : CloudStack
    Type             string  `json:"type"`
    VlanID           int     `json:"vlanId"`
    VpcID            string  `json:"vpcId"`
    ProjectID        string  `json:"projectId"`
    IsPrivateSubnet  bool    `json:"isPrivateSubnet"`
    IsCustom         bool    `json:"isCustom"`
    GatewayIP        string  `json:"gatewayIp"`
    CIDR             string  `json:"cidr"`
    StartIP          string  `json:"startIp"`
    EndIP            string  `json:"endIp"`
    LbStartIP        *string `json:"lbStartIp"`
    LbEndIP          *string `json:"lbEndIp"`
    BmStartIP        *string `json:"bmStartIp"`
    BmEndIP          *string `json:"bmEndIp"`
    IscsiStartIP     string  `json:"iscsiStartIp"`
    IscsiEndIP       string  `json:"iscsiEndIp"`
    NetworkName      string  `json:"networkName"` // KT Cloud Tier Name
    RefID            string  `json:"refId"` 	  // KT Cloud Tier ID based on OpenStack Neutron
    RefName          string  `json:"refName"`     // KT Cloud Tier Name based on OpenStack Neutron
    NetmaskIP        string  `json:"netmaskIp"`
    PortalZoneID     string  `json:"portalZoneId"`// Zone ID
    Shared           bool    `json:"shared"`
    IsAllocate       bool    `json:"isAllocate"`
    Interface        string  `json:"interface"`
    Status           string  `json:"status"`      // Tier status
}

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Subnet.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Subnet.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Subnet.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// SubnetResponse represents the full response structure for Subnet list.
type SubnetResponse struct {                        // Added
    HTTPStatus int              `json:"httpStatus"`
    Meta       interface{}      `json:"meta"`
    Pagination PaginationInfo   `json:"pagination"`
    Data       []Subnet         `json:"data"`
}

// CreateSubnetResponse represents the response structure for subnet creation.
type CreateSubnetResponse struct {                  // Modified
    HTTPStatus int    `json:"httpStatus"`
    JobID      string `json:"jobId,omitempty"`
    Data       struct {
        NetworkID string `json:"networkId"`
        VlanID    int    `json:"vlanId"`
    } `json:"data"`
}

// PaginationInfo represents pagination information in API responses.
type PaginationInfo struct {
    Size   int `json:"size"`
    Total  int `json:"total"`
    Offset int `json:"offset"`
}

// SubnetPage is the page returned by a pager when traversing over a collection of Subnets.
type SubnetPage struct {
    pagination.LinkedPageBase
}

// IsEmpty checks whether a SubnetPage struct is empty.
func (r SubnetPage) IsEmpty() (bool, error) {
    subnets, err := ExtractSubnets(r)
    return len(subnets) == 0, err
}

// NextPageURL returns the next page URL for traversing over Subnet pages.
func (r SubnetPage) NextPageURL() (string, error) {
    // Pagination URL generation logic
    // Currently, the API does not provide pagination links, so manual implementation is required.
    return "", nil
}

// ExtractSubnets extracts 'a slice of Subnet' from a SubnetPage.
func ExtractSubnets(r pagination.Page) ([]Subnet, error) { // Modified
    var s SubnetResponse
    err := r.(SubnetPage).ExtractInto(&s)
    return s.Data, err
}

// ExtractSubnet extracts 'a Subnet' from a GetSubnetResult.
func (r GetResult) ExtractSubnet() (*Subnet, error) { // Added
    var response SubnetResponse
    err := r.ExtractInto(&response)
    if err != nil {
        return nil, err
    }
    
    if len(response.Data) == 0 {
        return nil, gophercloud.ErrDefault404{}
    }
    
    return &response.Data[0], nil
}

// ExtractSubnetList extracts a SubnetResponse from the result.
func (r commonResult) ExtractSubnetList() (*SubnetResponse, error) { // Added
    var s SubnetResponse
    err := r.Result.ExtractInto(&s)
    return &s, err
}

// ExtractSubnetResponse extracts the full SubnetResponse from a result.
func (r commonResult) ExtractSubnetResponse() (*SubnetResponse, error) { // Added
    var s SubnetResponse
    err := r.ExtractInto(&s)
    return &s, err
}

// ExtractCreate extracts a CreateSubnetResponse from a CreateResult.
func (r CreateResult) ExtractCreate() (*CreateSubnetResponse, error) { // Modified
    var s CreateSubnetResponse
    err := r.ExtractInto(&s)
    return &s, err
}
