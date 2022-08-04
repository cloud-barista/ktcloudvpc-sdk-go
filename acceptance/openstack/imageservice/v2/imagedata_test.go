package v2

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/tools"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func TestImageStage(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	imageFileName := tools.RandomString("image_", 8)
	imageFilepath := "/tmp/" + imageFileName
	imageURL := ImportImageURL

	err = DownloadImageFileFromURL(t, imageURL, imageFilepath)
	th.AssertNoErr(t, err)
	defer DeleteImageFile(t, imageFilepath)

	err = StageImage(t, client, imageFilepath, image.ID)
	th.AssertNoErr(t, err)
}
