// ### KT Cloud D1 platform > 'Tier' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/computing/tier

package subnets

import (
    "github.com/rs/zerolog/log"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("network", id)  // Modified
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("network")  // Modified
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// getURL constructs a URL for a specific VPC with query parameter
// Ex) https://api.ucloudbiz.olleh.com/d1/nsm/v1/vpc?vpcId=60e5d9da-55cd-47be-a0d9-6cf67c54f15c
func getURL(c *gophercloud.ServiceClient, id string) string {  // Modified

    type SubnetQueryParams struct {
        SUBNETID string `q:"subnetId"`
    }

    params := SubnetQueryParams{
        SUBNETID: id,
    }
    q, err := gophercloud.BuildQueryString(params)
    if err != nil {
		log.Error().Msgf("\n\nFailed to Ceate Query String : %v\n", err)
    }

    baseURL := c.ServiceURL("network")        
    return baseURL + q.String()
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
