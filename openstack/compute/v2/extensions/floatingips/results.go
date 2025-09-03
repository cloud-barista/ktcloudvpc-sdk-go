package floatingips

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// FloatingIP represents a floating IP resource with detailed fields.
// Each field is documented above the field as per the provided JSON structure.
type FloatingIP struct {
    // Account ID for the created public IP
    AccountID string `json:"accountId"`

    // VPC ID
    VpcID string `json:"vpcId"`

    // Zone location of the IP
    ZoneID string `json:"zoneId"`

    // Public IP ID
    PublicIpID string `json:"publicIpId"`

    // IP usage type (SRCNAT, STATICNAT, PORTFORWARDING, ASSOCIATE)
    Type string `json:"type"`

    // Public IP CIDR
    CIDR string `json:"cidr"`

    // Public IP address
    PublicIP string `json:"publicIp"`

    // Zone ID where IP can be used
    PortalZoneID string `json:"portalZoneId"`

    // Serial IP (e.g., subnet range)
    SerialIP string `json:"serialIp"`

    // VLAN ID
    VlanID int `json:"vlanId"`

    // Public IP Pool ID
    PublicIPPoolID string `json:"publicIpPoolId"`

    // Allocation status (true if allocated)
    IsAllocate bool `json:"isAllocate"`

    // Static NAT settings list
    StaticNats []StaticNatInfo `json:"staticNats"`

    // Port forwarding settings list
    PortForwardings []PortForwardingInfo `json:"portForwardings"`
}

// StaticNatInfo represents a static NAT configuration.
// Extend fields as needed according to API response.
type StaticNatInfo struct {
    // TODO: Define fields if staticNats are used
}

// PortForwardingInfo represents a port forwarding configuration.
type PortForwardingInfo struct {
    // Port forwarding rule ID
    PortForwardingID string `json:"portForwardingId"`

    // Port forwarding rule name
    Name string `json:"name"`

    // VPC ID associated with the rule
    VpcID string `json:"vpcId"`

    // Mapped private IP address
    MappedIP string `json:"mappedIp"`

    // Network ID associated with the rule
    NetworkID string `json:"networkId"`

    // Public IP ID associated with the rule
    PublicIpID string `json:"publicIpId"`

    // Public IP address associated with the rule
    PublicIP string `json:"publicIp"`

    // Port forwarding protocol (e.g., TCP, UDP)
    Protocol string `json:"protocol"`

    // Starting private IP port number
    StartPrivatePort string `json:"startPrivatePort"`

    // Ending private IP port number
    EndPrivatePort string `json:"endPrivatePort"`

    // Starting public IP port number
    StartPublicPort string `json:"startPublicPort"`

    // Ending public IP port number
    EndPublicPort string `json:"endPublicPort"`
}

type CreateFloatingIPResponse struct {
    HTTPStatus int `json:"httpStatus"`
    Data struct {
        PublicIpID string `json:"publicIpId"`
    } `json:"data"`
}

// ExtractCreate extracts a CreateResponse from a CreateResult.
func (r CreateResult) ExtractCreate() (*CreateFloatingIPResponse, error) {
    var s CreateFloatingIPResponse
    err := r.ExtractInto(&s)
    return &s, err
}

func ExtractPublicIPID(result CreateResult) (string, error) {
    var resp CreateFloatingIPResponse
    err := result.ExtractInto(&resp)
    if err != nil {
        return "", err
    }
    return resp.Data.PublicIpID, nil
}

// PFResponse represents the full response structure for P/F list.
type IPResponse struct {                    // Modified
    HTTPStatus int              `json:"httpStatus"`
    Meta       interface{}      `json:"meta"`
    Pagination PaginationInfo   `json:"pagination"`
    Data       []FloatingIP 	`json:"data"`
}

// PaginationInfo represents pagination information in API responses.
type PaginationInfo struct {                // Added
    Size   int `json:"size"`
    Total  int `json:"total"`
    Offset int `json:"offset"`
}

// FloatingIPPage is the page returned by a pager when traversing over a
// collection of floating IPs.
type FloatingIPPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a FloatingIPsPage is empty.
func (page FloatingIPPage) IsEmpty() (bool, error) {
	va, err := ExtractFloatingIPs(page)
	return len(va) == 0, err
}

// NextPageURL returns the next page URL for traversing over FloatingIP pages.
func (page FloatingIPPage) NextPageURL() (string, error) {
    // Pagination URL generation logic
    // Currently, the API does not provide pagination links, so manual implementation is required.
    return "", nil
}

// ExtractFloatingIPs interprets a page of results as a slice of FloatingIPs.
func ExtractFloatingIPs(r pagination.Page) ([]FloatingIP, error) {
    var s IPResponse
    err := r.(FloatingIPPage).ExtractInto(&s)
    return s.Data, err
}

// ExtractFloatingIP extracts 'a FloatingIP' from a GetResult.
func (r GetResult) ExtractFloatingIP() (*FloatingIP, error) {
    var response IPResponse
    err := r.ExtractInto(&response)
    if err != nil {
        return nil, err
    }
    
    if len(response.Data) == 0 {
        return nil, gophercloud.ErrDefault404{}
    }
    
    return &response.Data[0], nil
}

// FloatingIPResult is the raw result from a FloatingIP request.
type FloatingIPResult struct {
	gophercloud.Result
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a FloatingIP.
type CreateResult struct {
	FloatingIPResult
}

// GetResult is the response from a Get operation. Call its Extract method to
// interpret it as a FloatingIP.
type GetResult struct {
	FloatingIPResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
