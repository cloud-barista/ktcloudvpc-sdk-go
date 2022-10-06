/*
Package quotas provides the ability to retrieve and manage Load Balancer quotas

Example to Get project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"
    quotasInfo, err := quotas.Get(networkClient, projectID).Extract()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)

Example to Update project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"

    updateOpts := quotas.UpdateOpts{
		Loadbalancer:  ktvpcsdk.IntToPointer(20),
		Listener:      ktvpcsdk.IntToPointer(40),
		Member:        ktvpcsdk.IntToPointer(200),
		Pool:          ktvpcsdk.IntToPointer(20),
		Healthmonitor: ktvpcsdk.IntToPointer(1),
		L7Policy:      ktvpcsdk.IntToPointer(50),
		L7Rule:        ktvpcsdk.IntToPointer(100),
    }
    quotasInfo, err := quotas.Update(networkClient, projectID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)
*/
package quotas
