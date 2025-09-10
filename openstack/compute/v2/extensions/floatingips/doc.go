/*
Package floatingips provides the ability to manage floating ips

// Example: Create a floating IP using floatingips.Create()
// filepath: ./github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/compute/v2/extensions/floatingips/

result := floatingips.Create(client, opts)
if result.Err != nil {
    fmt.Printf("Create error: %v\n", result.Err)
} else {
    publicIpId, err := ExtractPublicIPID(result)
    if err != nil {
        fmt.Printf("Failed to extract publicIpId: %v\n", err)
    } else {
        fmt.Printf("Created publicIpId: %s\n", publicIpId)
    }
}

// ################################################################################

// Example: List floating IPs using floatingips.List()
// filepath: ./github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/compute/v2/extensions/floatingips/

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/compute/v2/extensions/floatingips"
    "github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // Set up filter options (example: filter by CIDR and page size)
    listOpts := floatingips.ListOpts{
        Page:       1,
        Size:       20,
        CIDR:       "",
        PublicIpID: "",
    }

    // Call List() to get a Pager
    pager := floatingips.List(client, listOpts)

    // Iterate through all pages and print floating IP info
    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        fips, err := floatingips.ExtractFloatingIPs(page)
        if err != nil {
            return false, err
        }
        for _, fip := range fips {
            fmt.Printf("FloatingIP ID: %s, PublicIP: %s, VPC ID: %s, ZoneID: %s\n",
                fip.PublicIpID, fip.PublicIP, fip.VpcID, fip.ZoneID)
        }
        return true, nil
    })
    if err != nil {
        log.Fatalf("Failed to list floating IPs: %v", err)
    }
}

// ################################################################################

// Example: Get a Floating IP by publicIpId using floatingips.Get()

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/compute/v2/extensions/floatingips"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // The public IP ID you want to query
    publicIpId := "69d66144-0e40-46ca-9d17-3b08bea226c9"

    // Call Get to retrieve the floating IP details
    result := floatingips.Get(client, publicIpId)

    // Extract the FloatingIP information
    fip, err := result.ExtractFloatingIP()
    if err != nil {
        log.Fatalf("Failed to get floating IP: %v", err)
    }

    fmt.Printf("Public IP ID: %s\n", fip.PublicIpID)
    fmt.Printf("Public IP Address: %s\n", fip.PublicIP)
    fmt.Printf("VPC ID: %s\n", fip.VpcID)
    fmt.Printf("Zone ID: %s\n", fip.ZoneID)
    fmt.Printf("Type: %s\n", fip.Type)
}

// Example: Delete a Floating IP by Floating IP Address

	publicIpAddr := "XXX.XX.XXX.XXX"

    // 1. Find the public IP ID by IP address
    publicIPID, err := vmHandler.FindPublicIPIDByIP(publicIpAddr)
    if err != nil {		
		return nil, err
    }

    // 2. Delete the public IP using the found ID
    result := ips.Delete(client, publicIPID)
    if result.Err != nil {
        return nil, result.Err
    }
    fmt.Printf("Successfully deleted public IP: %s (ID: %s)", publicIpAddr, publicIPID)
    
*/
package floatingips
