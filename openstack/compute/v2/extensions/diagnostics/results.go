package diagnostics

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

type serverDiagnosticsResult struct {
	ktvpcsdk.Result
}

// Extract interprets any diagnostic response as a map
func (r serverDiagnosticsResult) Extract() (map[string]interface{}, error) {
	var s map[string]interface{}
	err := r.ExtractInto(&s)
	return s, err
}
