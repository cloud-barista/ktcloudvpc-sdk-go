package portforwarding

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
)

const resourcePath = "portforwarding"

// URL builders
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

// getURL constructs a URL for a specific VPC with query parameter
// Ex) https://api.ucloudbiz.olleh.com/d1/nsm/v1/vpc?vpcId=60e5d9da-55cd-47be-a0d9-6cf67c54f15c
func getURL(c *gophercloud.ServiceClient, id string) string {

    type VPCQueryParams struct {
        PortForwardingID string `q:"portForwardingId"`
    }

    params := VPCQueryParams{
       PortForwardingID: id,
    }
    q, err := gophercloud.BuildQueryString(params)
    if err != nil {
		cblogger.Errorf("\n\nFailed to Ceate Query String : %v\n", err)
    }

    baseURL := c.ServiceURL(resourcePath)        
    return baseURL + q.String()
}


func createURL(c *gophercloud.ServiceClient) string {
	// cblogger.Errorf("\n\nCreate URL : %v\n", c.ServiceURL(resourcePath))
	return c.ServiceURL(resourcePath)
}

// func getURL(c *gophercloud.ServiceClient, id string) string {
// 	return c.ServiceURL(resourcePath, id)
// }

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
