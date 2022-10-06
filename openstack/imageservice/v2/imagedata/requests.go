package imagedata

import (
	"io"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// Upload uploads an image file.
func Upload(client *ktvpcsdk.ServiceClient, id string, data io.Reader) (r UploadResult) {
	resp, err := client.Put(uploadURL(client, id), data, nil, &ktvpcsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Stage performs PUT call on the existing image object in the Imageservice with
// the provided file.
// Existing image object must be in the "queued" status.
func Stage(client *ktvpcsdk.ServiceClient, id string, data io.Reader) (r StageResult) {
	resp, err := client.Put(stageURL(client, id), data, nil, &ktvpcsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Download retrieves an image.
func Download(client *ktvpcsdk.ServiceClient, id string) (r DownloadResult) {
	resp, err := client.Get(downloadURL(client, id), nil, &ktvpcsdk.RequestOpts{
		KeepResponseBody: true,
	})
	r.Body, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
