package profiletypes

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

func Get(client *ktvpcsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body,
		&ktvpcsdk.RequestOpts{OkCodes: []int{200}})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// List makes a request against the API to list profile types.
func List(client *ktvpcsdk.ServiceClient) pagination.Pager {
	url := listURL(client)
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ProfileTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func ListOps(client *ktvpcsdk.ServiceClient, id string) pagination.Pager {
	url := listOpsURL(client, id)
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return OperationPage{pagination.SinglePageBase(r)}
	})
}
