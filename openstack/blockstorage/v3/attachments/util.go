package attachments

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// WaitForStatus will continually poll the resource, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForStatus(c *ktvpcsdk.ServiceClient, id, status string, secs int) error {
	return ktvpcsdk.WaitFor(secs, func() (bool, error) {
		current, err := Get(c, id).Extract()
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}
