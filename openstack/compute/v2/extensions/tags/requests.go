package tags

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// List all tags on a server.
func List(client *ktvpcsdk.ServiceClient, serverID string) (r ListResult) {
	url := listURL(client, serverID)
	resp, err := client.Get(url, &r.Body, &ktvpcsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Check if a tag exists on a server.
func Check(client *ktvpcsdk.ServiceClient, serverID, tag string) (r CheckResult) {
	url := checkURL(client, serverID, tag)
	resp, err := client.Get(url, nil, &ktvpcsdk.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// ReplaceAllOptsBuilder allows to add additional parameters to the ReplaceAll request.
type ReplaceAllOptsBuilder interface {
	ToTagsReplaceAllMap() (map[string]interface{}, error)
}

// ReplaceAllOpts provides options used to replace Tags on a server.
type ReplaceAllOpts struct {
	Tags []string `json:"tags" required:"true"`
}

// ToTagsReplaceAllMap formats a ReplaceALlOpts into the body of the ReplaceAll request.
func (opts ReplaceAllOpts) ToTagsReplaceAllMap() (map[string]interface{}, error) {
	return ktvpcsdk.BuildRequestBody(opts, "")
}

// ReplaceAll replaces all Tags on a server.
func ReplaceAll(client *ktvpcsdk.ServiceClient, serverID string, opts ReplaceAllOptsBuilder) (r ReplaceAllResult) {
	b, err := opts.ToTagsReplaceAllMap()
	url := replaceAllURL(client, serverID)
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(url, &b, &r.Body, &ktvpcsdk.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Add adds a new Tag on a server.
func Add(client *ktvpcsdk.ServiceClient, serverID, tag string) (r AddResult) {
	url := addURL(client, serverID, tag)
	resp, err := client.Put(url, nil, nil, &ktvpcsdk.RequestOpts{
		OkCodes: []int{201, 204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Delete removes a tag from a server.
func Delete(client *ktvpcsdk.ServiceClient, serverID, tag string) (r DeleteResult) {
	url := deleteURL(client, serverID, tag)
	resp, err := client.Delete(url, &ktvpcsdk.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// DeleteAll removes all tag from a server.
func DeleteAll(client *ktvpcsdk.ServiceClient, serverID string) (r DeleteResult) {
	url := deleteAllURL(client, serverID)
	resp, err := client.Delete(url, &ktvpcsdk.RequestOpts{
		OkCodes: []int{204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
