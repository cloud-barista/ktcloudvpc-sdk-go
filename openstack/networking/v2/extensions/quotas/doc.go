/*
Package quotas provides the ability to retrieve and manage Networking quotas through the Neutron API.

Example to Get project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"
    quotasInfo, err := quotas.Get(networkClient, projectID).Extract()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)

Example to Get a Detailed Quota Set

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"
    quotasInfo, err := quotas.GetDetail(networkClient, projectID).Extract()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)

Example to Update project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"

    updateOpts := quotas.UpdateOpts{
        FloatingIP:        ktvpcsdk.IntToPointer(0),
        Network:           ktvpcsdk.IntToPointer(-1),
        Port:              ktvpcsdk.IntToPointer(5),
        RBACPolicy:        ktvpcsdk.IntToPointer(10),
        Router:            ktvpcsdk.IntToPointer(15),
        SecurityGroup:     ktvpcsdk.IntToPointer(20),
        SecurityGroupRule: ktvpcsdk.IntToPointer(-1),
        Subnet:            ktvpcsdk.IntToPointer(25),
        SubnetPool:        ktvpcsdk.IntToPointer(0),
        Trunk:             ktvpcsdk.IntToPointer(0),
    }
    quotasInfo, err := quotas.Update(networkClient, projectID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)
*/
package quotas
