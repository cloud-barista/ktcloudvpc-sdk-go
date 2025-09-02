/*
package portforwarding enables management and retrieval of port forwarding resources for Floating IPs from the
OpenStack Networking service.

### Example to Create a Port Forwarding for a floating IP

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/extensions/layer3/portforwarding"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize this with your auth/session

    // Define the port forwarding creation options
    createOpts := portforwarding.CreateOpts{
        PublicIPID:      "your-public-ip-id",      // required
        MappedIP:        "192.168.0.10",           // required
        Protocol:        "TCP",                    // required ("TCP" or "UDP")
        StartPrivatePort: "8080",                  // required
        EndPrivatePort:   "8080",                  // required
        StartPublicPort:  "80",                    // required
        EndPublicPort:    "80",                    // required
    }

    // Call the Create function
    result := portforwarding.Create(client, createOpts)

    // Check for errors
    if result.Err != nil {
        log.Fatalf("Failed to create port forwarding: %v", result.Err)
    }

    // Extract the created PortForwarding resource
    pf, err := result.Extract()
    if err != nil {
        log.Fatalf("Failed to extract port forwarding result: %v", err)
    }

    // Print the created port forwarding rule details
    fmt.Printf("Created Port Forwarding Rule:\n")
    fmt.Printf("ID: %s\n", pf.ID)
    fmt.Printf("Name: %s\n", pf.Name)
    fmt.Printf("Mapped IP: %s\n", pf.MappedIP)
    fmt.Printf("Protocol: %s\n", pf.Protocol)
    fmt.Printf("Public IP: %s\n", pf.PublicIP)
    fmt.Printf("Public IP ID: %s\n", pf.PublicIPID)
    fmt.Printf("Stack Type: %s\n", pf.StackType)
    fmt.Printf("Created At: %s\n", pf.CreateDate)
}


// =============================================================================================================

// Example: List port forwarding rules using portforwarding.List()

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/extensions/layer3/portforwarding"
    "github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // Set up filter options (example: filter by protocol and public IP)
    listOpts := portforwarding.ListOpts{
        Protocol:  "TCP",
        PublicIP:  "203.0.113.10",
        Page:      1,
        Size:      10,
    }

    // Call List() to get a Pager
    pager := portforwarding.List(client, listOpts)

    // Iterate through all pages and print rule info
    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        pfList, err := portforwarding.ExtractPFs(page)
        if err != nil {
            return false, err
        }
        for _, pf := range pfList {
            fmt.Printf("PortForwarding ID: %s, Name: %s, Protocol: %s, PublicIP: %s, MappedIP: %s\n",
                pf.ID, pf.Name, pf.Protocol, pf.PublicIP, pf.MappedIP)
        }
        return true, nil
    })
    if err != nil {
        log.Fatalf("Failed to list port forwarding rules: %v", err)
    }
}


*/
package portforwarding
