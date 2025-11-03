package floatingips

import (
    "github.com/rs/zerolog/log"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

const resourcePath = "publicIp"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

// getURL constructs a URL for a specific PublicIp with query parameter
// Ex) https://api.ucloudbiz.olleh.com/d1/nsm/v1/publicIp?publicIpId=b6c6dae2-7c2d-497e-8117-989dce873fae
func getURL(c *gophercloud.ServiceClient, id string) string {
    type PublicIpQueryParams struct {
        PublicIpID string `q:"publicIpId"`
    }

    params := PublicIpQueryParams{
        PublicIpID: id,
    }
    q, err := gophercloud.BuildQueryString(params)
    if err != nil {
		log.Error().Msgf("\n\nFailed to Create Query String : %v\n", err)
    }

    baseURL := c.ServiceURL(resourcePath)
    return baseURL + q.String()
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
