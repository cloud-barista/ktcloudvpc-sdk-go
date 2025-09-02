/*
// File path: ./ktcloudvpc-sdk-go/openstack/networking/v2/extensions/fwaas_v2/rules/doc.go
// Example: List firewall rules with filtering and pagination

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/extensions/fwaas_v2/rules"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // Set up filter options (example: filter by protocol and enabled status)
    listOpts := rules.ListOpts{
        Protocol: "TCP",
        Enabled:  true,
        Limit:    10, // Pagination: limit to 10 results per page
    }

    // Call List() to get a Pager
    pager := rules.List(client, listOpts)

    // Iterate through all pages and print rule info
    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        ruleList, err := rules.ExtractRules(page)
        if err != nil {
            return false, err
        }
        for _, rule := range ruleList {
            fmt.Printf("Rule ID: %s, Name: %s, Protocol: %s, Action: %s, Enabled: %v\n",
                rule.ID, rule.Name, rule.Protocol, rule.Action, rule.Enabled)
        }
        return true, nil
    })
    if err != nil {
        log.Fatalf("Failed to list firewall rules: %v", err)
    }
}

// ==============================================================================================
// Example: Delete a firewall rule using rules.Delete()

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/extensions/fwaas_v2/rules"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // The ID of the firewall rule to delete
    ruleID := "your-firewall-rule-id"

    // Call Delete to remove the rule
    result := rules.Delete(client, ruleID)

    // Check for errors
    if result.Err != nil {
        log.Fatalf("Failed to delete firewall rule: %v", result.Err)
    } else {
        fmt.Printf("Firewall rule %s deleted successfully.\n", ruleID)
    }
}

// ============================================================================================

// Example: Create a firewall rule using rules.Create()

package main

import (
    "fmt"
    "log"

    "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/extensions/fwaas_v2/rules"
    "github.com/cloud-barista/ktcloudvpc-sdk-go"
)

func main() {
    // Assume you have a valid gophercloud.ServiceClient named 'client'
    var client *gophercloud.ServiceClient // Initialize with your auth/session

    // Define the firewall rule creation options
    createOpts := rules.CreateOpts{
        SrcNat:           false,
        Action:           rules.ActionAllow,
        Protocol:         string(rules.ProtocolTCP),
        StartPort:        "80",
        EndPort:          "80",
        SrcAddress:       []string{"0.0.0.0/0"},
        DstAddress:       []string{"10.0.0.10"},
        Comment:          "Allow HTTP traffic to 10.0.0.10",
    }

    // Call Create to create the firewall rule
    result := rules.Create(client, createOpts)

    // Check for errors
    if result.Err != nil {
        log.Fatalf("Failed to create firewall rule: %v", result.Err)
    }

    // Extract the created rule (assuming Extract() is implemented)
    rule, err := result.Extract()
    if err != nil {
        log.Fatalf("Failed to extract created rule: %v", err)
    }

    fmt.Printf("Created Firewall Rule: ID=%s, Protocol=%s, Action=%v\n", rule.ID, rule.Protocol, rule.Action)
}



*/
package rules
