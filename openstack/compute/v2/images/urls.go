// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Updated by ETRI, 2022.07.

package images

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {   // Added by B.T. Oh.
	return client.ServiceURL("images", "")
}

func listDetailURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}
