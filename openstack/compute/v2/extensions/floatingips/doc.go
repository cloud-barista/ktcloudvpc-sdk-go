/*
Package floatingips provides the ability to manage floating ips through the
Nova API.

// Usage example:
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
// filepath: /home/sean/go/src/github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/compute/v2/extensions/floatingips/

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
        Size:       10,
        CIDR:       "",
        PupblicIpID: "",
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


*/
package floatingips
