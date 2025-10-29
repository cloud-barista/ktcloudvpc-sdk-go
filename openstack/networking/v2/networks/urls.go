// ### KT Cloud D1 platfrom > 'VPC' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/networking

package networks

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

const resourcePath = "vpc"									    // Modified

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)  // Modified
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)  // Modified
}

// getURL constructs a URL for a specific VPC with query parameter
// Ex) https://api.ucloudbiz.olleh.com/d1/nsm/v1/vpc?vpcId=60e5d9da-55cd-47be-a0d9-6cf67c54f15c
func getURL(c *gophercloud.ServiceClient, id string) string {  // Modified
    type VPCQueryParams struct {
        VPCID string `q:"vpcId"`
    }

    params := VPCQueryParams{
        VPCID: id,
    }
    q, err := gophercloud.BuildQueryString(params)
    if err != nil {
		cblogger.Errorf("\n\nFailed to Ceate Query String : %v\n", err)
    }

    baseURL := c.ServiceURL(resourcePath)        
    return baseURL + q.String()
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
