package ec2tokens

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func ec2tokensURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("ec2tokens")
}

func s3tokensURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("s3tokens")
}
