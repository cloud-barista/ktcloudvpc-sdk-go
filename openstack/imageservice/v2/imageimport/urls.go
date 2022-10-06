package imageimport

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "images"
	infoPath     = "info"
	resourcePath = "import"
)

func infoURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(infoPath, resourcePath)
}

func importURL(c *ktvpcsdk.ServiceClient, imageID string) string {
	return c.ServiceURL(rootPath, imageID, resourcePath)
}
