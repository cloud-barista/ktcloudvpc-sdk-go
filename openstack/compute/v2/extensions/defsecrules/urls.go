package defsecrules

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const rulepath = "os-security-group-default-rules"

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rulepath, id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rulepath)
}
