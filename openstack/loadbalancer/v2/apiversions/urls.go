package apiversions

import (
	"strings"

	"github.com/innodreamer/ktvpc-sdk_poc"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}
