package portforwarding

import (
	"encoding/json"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

type PortForwarding struct {
	// ID uniquely identifies the port forwarding rule
	ID string `json:"portForwardingId"`
	
	// Name is the human-readable name of the rule
	Name string `json:"name"`
	
	// MappedIP is the private IP address that traffic is forwarded to
	MappedIP string `json:"mappedIp"`
	
	// StartPrivatePort is the starting port on the private network
	StartPrivatePort string `json:"startPrivatePort"`
	
	// StartPublicPort is the starting port on the public network
	StartPublicPort string `json:"startPublicPort"`
	
	// EndPrivatePort is the ending port on the private network
	EndPrivatePort string `json:"endPrivatePort"`
	
	// EndPublicPort is the ending port on the public network
	EndPublicPort string `json:"endPublicPort"`
	
	// Protocol specifies the network protocol (TCP, UDP)
	Protocol string `json:"protocol"`
	
	// PublicIPID is the ID of the public IP address
	PublicIPID string `json:"publicIpId"`
	
	// PublicIP is the public IP address
	PublicIP string `json:"publicIp"`
	
	// StackType indicates the stack type (usually "OS" for OpenStack)
	StackType string `json:"stackType"`
	
	// CreateDate is the timestamp when the port forwarding rule was created
	CreateDate string `json:"createDate"`
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a PortForwarding.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a PortForwarding.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// type DeleteResult struct {															// Modified
// 	commonResult
// }

// Extract will extract a Port Forwarding resource from a result.
func (r commonResult) Extract() (*PortForwarding, error) {
	var s PortForwarding
	err := r.ExtractInto(&s)
	return &s, err
}


// func (r commonResult) ExtractInto(v interface{}) error {
// 	return r.Result.ExtractIntoStructPtr(v, "job_id")
// }


// PFResponse represents the full response structure for P/F list.
type PFResponse struct {                    // Modified
    HTTPStatus int              `json:"httpStatus"`
    Meta       interface{}      `json:"meta"`
    Pagination PaginationInfo   `json:"pagination"`
    Data       []PortForwarding `json:"data"`
}

// PaginationInfo represents pagination information in API responses.
type PaginationInfo struct {                // Added
    Size   int `json:"size"`
    Total  int `json:"total"`
    Offset int `json:"offset"`
}

// PortForwardingPage is the page returned by a pager when traversing over a
// collection of port forwardings.
type PortForwardingPage struct {
	pagination.LinkedPageBase
}

type CreatePFData struct {
    PortForwardingID string `json:"portForwardingId"`
}

type CreatePFResponse struct {
    HTTPStatus int          `json:"httpStatus"`
    Data       CreatePFData `json:"data"`
}

func ExtractPortForwardingID(result CreateResult) (string, error) {
    var resp CreatePFResponse
    err := result.ExtractInto(&resp)
    if err != nil {
        return "", err
    }
    return resp.Data.PortForwardingID, nil
}

// type DellPortforwardingResponse struct {											// Added
// 	JopID string `json:"job_id"`
// }

// func (r commonResult) ExtractDelJobInfo() (*DellPortforwardingResponse, error) {   	// Added
// 	var s struct {
// 		DellPortforwardingResponse *DellPortforwardingResponse `json:"nc_deleteportforwardingruleresponse"`
// 	}
// 	err := r.ExtractInto(&s)
// 	return s.DellPortforwardingResponse, err
// }

// IsEmpty checks whether a PortForwardingPage struct is empty.
func (r PortForwardingPage) IsEmpty() (bool, error) {
	is, err := ExtractPFs(r)
	return len(is) == 0, err
}

// NextPageURL returns the next page URL for traversing over VPC pages.
func (r PortForwardingPage) NextPageURL() (string, error) {
    // Pagination URL generation logic
    // Currently, the API does not provide pagination links, so manual implementation is required.
    return "", nil
}

// ExtractPFs extracts 'a slice of PortForwarding' from a PortForwardingPage.
func ExtractPFs(r pagination.Page) ([]PortForwarding, error) {
    var s PFResponse
    err := r.(PortForwardingPage).ExtractInto(&s)
    return s.Data, err
}

// ExtractPortForwarding extracts 'a PortForwarding' from a GetResult.
func (r GetResult) ExtractPortForwarding() (*PortForwarding, error) { // Added
    var response PFResponse
    err := r.ExtractInto(&response)
    if err != nil {
        return nil, err
    }
    
    if len(response.Data) == 0 {
        return nil, gophercloud.ErrDefault404{}
    }
    
    return &response.Data[0], nil
}

// ExtractPortForwardingList extracts a PFResponse from the result.
func (r commonResult) ExtractPFList() (*PFResponse, error) { // Added
    var s PFResponse
    err := r.Result.ExtractInto(&s)
    return &s, err
}

// ExtractPFResponse extracts the full PFResponse from a result.
func (r commonResult) ExtractPFResponse() (*PFResponse, error) { // Added
    var s PFResponse
    err := r.ExtractInto(&s)
    return &s, err
}

// ErrorResponse represents the error response structure for port forwarding API.
type ErrorResponse struct {
    HTTPStatus int    `json:"httpStatus"`
    Title      string `json:"title"`
    Detail     string `json:"detail"`
    Instance   string `json:"instance"`
    ServiceID  string `json:"serviceId"`
    Status     int    `json:"status"`
}

// ExtractErrorResponse extracts ErrorResponse from a CreateResult, UpdateResult, or commonResult.
func (r commonResult) ExtractErrorResponse() (*ErrorResponse, error) {
    var e ErrorResponse
    err := r.Result.ExtractInto(&e)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

// For DeleteResult (which embeds ErrResult), add similar helper:
func (r DeleteResult) ExtractErrorResponse() (*ErrorResponse, error) {
    var e ErrorResponse
    err := r.ErrResult.ExtractInto(&e)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

// Helper method to extract error detail from Result
func (r commonResult) ExtractErrorDetail() string {
    // Try to extract "detail" field from the response body if it's a map
    if r.Body == nil {
        return ""
    }
    if m, ok := r.Body.(map[string]interface{}); ok {
        if detail, ok := m["detail"].(string); ok {
            return detail
        }
    }
    // Try to marshal and unmarshal into a struct with Detail field
    type errorDetail struct {
        Detail string `json:"detail"`
    }
    var ed errorDetail
    b, err := json.Marshal(r.Body)
    if err == nil {
        if json.Unmarshal(b, &ed) == nil {
            return ed.Detail
        }
    }
    return ""
}
