// ### KT Cloud D1 platfrom > 'VPC' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/networking

package networks

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
    "github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Network.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Network.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Network.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type VPC struct {                   // Modified
    AccountID     string  `json:"accountId"`
    VpcID         string  `json:"vpcId"`
    ZoneID        string  `json:"zoneId"`
    SessionCount  int     `json:"sessionCount"`
    Name          string  `json:"name"`
    BizareaType   string  `json:"bizareaType"`
    VpcOfferingID string  `json:"vpcOfferingId"`
    PortalZoneID  string  `json:"portalZoneId"`
    Description   *string `json:"description"`
    State         string  `json:"state"`
    VdomInfo      string  `json:"vdomInfo"`
    Type          string  `json:"type"`
    DeviceID      string  `json:"deviceId"`
    VdomName      string  `json:"vdomName"`
    CreateDate    string  `json:"createDate"`
}

// VPCResponse represents the full response structure for VPC list.
type VPCResponse struct {                    // Modified
    HTTPStatus int              `json:"httpStatus"`
    Meta       interface{}      `json:"meta"`
    Pagination PaginationInfo   `json:"pagination"`
    Data       []VPC            `json:"data"`
}

// PaginationInfo represents pagination information in API responses.
type PaginationInfo struct {                // Added
    Size   int `json:"size"`
    Total  int `json:"total"`
    Offset int `json:"offset"`
}

// VPCPage is the page returned by a pager when traversing over a collection of VPCs.
type VPCPage struct {
    pagination.LinkedPageBase
}

// IsEmpty checks whether a VPCV2Page struct is empty.
func (r VPCPage) IsEmpty() (bool, error) {
    vpcs, err := ExtractVPCs(r)
    return len(vpcs) == 0, err
}

// NextPageURL returns the next page URL for traversing over VPC pages.
func (r VPCPage) NextPageURL() (string, error) {
    // Pagination URL generation logic
    // Currently, the API does not provide pagination links, so manual implementation is required.
    return "", nil
}

// ExtractVPCs extracts 'a slice of VPC' from a VPCPage.
func ExtractVPCs(r pagination.Page) ([]VPC, error) {
    var s VPCResponse
    err := r.(VPCPage).ExtractInto(&s)
    return s.Data, err
}

// ExtractVPC extracts 'a VPC' from a GetResult.
func (r GetResult) ExtractVPC() (*VPC, error) { // Added
    var response VPCResponse
    err := r.ExtractInto(&response)
    if err != nil {
        return nil, err
    }
    
    if len(response.Data) == 0 {
        return nil, gophercloud.ErrDefault404{}
    }
    
    return &response.Data[0], nil
}

// ExtractVPCList extracts a VPCResponse from the result.
func (r commonResult) ExtractVPCList() (*VPCResponse, error) { // Added
    var s VPCResponse
    err := r.Result.ExtractInto(&s)
    return &s, err
}

// ExtractVPCResponse extracts the full VPCResponse from a result.
func (r commonResult) ExtractVPCResponse() (*VPCResponse, error) { // Added
    var s VPCResponse
    err := r.ExtractInto(&s)
    return &s, err
}
