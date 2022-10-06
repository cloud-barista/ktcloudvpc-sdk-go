package certificates

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

type commonResult struct {
	ktvpcsdk.Result
}

// GetResult is the response of a Get operations.
type GetResult struct {
	commonResult
}

// CreateResult is the response of a Create operations.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response of an Update operations.
type UpdateResult struct {
	ktvpcsdk.ErrResult
}

// Extract is a function that accepts a result and extracts a certificate resource.
func (r commonResult) Extract() (*Certificate, error) {
	var s *Certificate
	err := r.ExtractInto(&s)
	return s, err
}

// Represents a Certificate
type Certificate struct {
	ClusterUUID string             `json:"cluster_uuid"`
	BayUUID     string             `json:"bay_uuid"`
	Links       []ktvpcsdk.Link `json:"links"`
	PEM         string             `json:"pem"`
	CSR         string             `json:"csr"`
}
