package external

import (
	"net/url"
	"strconv"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/networks"
)

// ListOptsExt adds the external network options to the base ListOpts.
type ListOptsExt struct {
	networks.ListOptsBuilder
	External *bool `q:"router:external"`
}

// ToNetworkListQuery adds the router:external option to the base network
// list options.
func (opts ListOptsExt) ToNetworkListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts.ListOptsBuilder)
	if err != nil {
		return "", err
	}

	params := q.Query()
	if opts.External != nil {
		v := strconv.FormatBool(*opts.External)
		params.Add("router:external", v)
	}

	q = &url.URL{RawQuery: params.Encode()}
	return q.String(), err
}

// CreateOptsExt is the structure used when creating new external network
// resources. It embeds networks.CreateOpts and so inherits all of its required
// and optional fields, with the addition of the External field.
type CreateOptsExt struct {
	networks.CreateOptsBuilder
	External *bool `json:"router:external,omitempty"`
}

// ToNetworkCreateMap adds the router:external options to the base network
// creation options.
func (opts CreateOptsExt) ToVPCCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToVPCCreateMap()
	if err != nil {
		return nil, err
	}

	if opts.External == nil {
		return base, nil
	}

	networkMap := base["network"].(map[string]interface{})
	networkMap["router:external"] = opts.External

	return base, nil
}

// UpdateOptsExt is the structure used when updating existing external network
// resources. It embeds networks.UpdateOpts and so inherits all of its required
// and optional fields, with the addition of the External field.
type UpdateOptsExt struct {
	networks.UpdateOptsBuilder
	External *bool `json:"router:external,omitempty"`
}

// ToNetworkUpdateMap casts an UpdateOpts struct to a map.
func (opts UpdateOptsExt) ToVPCUpdateMap() (map[string]interface{}, error) {
	base, err := opts.UpdateOptsBuilder.ToVPCUpdateMap()
	if err != nil {
		return nil, err
	}

	if opts.External == nil {
		return base, nil
	}

	networkMap := base["network"].(map[string]interface{})
	networkMap["router:external"] = opts.External

	return base, nil
}
