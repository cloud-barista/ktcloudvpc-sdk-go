package quotas

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToQuotaCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	ProjectID string `json:"project_id"`
	Resource  string `json:"resource"`
	HardLimit int    `json:"hard_limit"`
}

// ToQuotaCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToQuotaCreateMap() (map[string]interface{}, error) {
	return ktvpcsdk.BuildRequestBody(opts, "")
}

// Create requests the creation of a new quota.
func Create(client *ktvpcsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToQuotaCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), b, &r.Body, &ktvpcsdk.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
