package job

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {  			// Added by B.T. Oh
	return sc.ServiceURL("Etc?command=queryAsyncJob&jobid=", jobId)
}
