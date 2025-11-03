package rules

import (
	// "github.com/davecgh/go-spew/spew"
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

type (
	Protocol string
)

const (
	// ProtocolAny is to allow any protocol
	ProtocolAny Protocol = "any"

	// ProtocolTCP is to allow the TCP protocol
	ProtocolTCP Protocol = "TCP"

	// ProtocolUDP is to allow the UDP protocol
	ProtocolUDP Protocol = "UDP"

	// ProtocolICMP is to allow the ICMP protocol
	ProtocolICMP Protocol = "ICMP"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
    ToRuleListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the rule attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	Page     int    `q:"page"`
    Size     int    `q:"size"`
	PolicyID string `q:"policyId"`
}

// ToRuleListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRuleListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// firewall rules. It accepts a ListOptsBuilder, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToRuleListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return FirewallRulePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToRuleCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new firewall rule.
type CreateOpts struct {
	// [Required] Pool name creation flag
	SrcNat bool `json:"srcNat,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with dstAddress and staticNatId
	PortForwardingId string `json:"portForwardingId,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with dstAddress and portForwardingId
	StaticNatId string `json:"staticNatId,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with srcNetwork
	// Optional string parameter for source VLAN Name
	SrcInterface string `json:"srcInterface,omitempty"`
	
	// ### Caution!! : [Optional] Cannot be entered simultaneously with dstNetwork
	// Optional string parameter for destination VLAN Name
	DstInterface string `json:"dstInterface,omitempty"`

	// [Optional] [Default: true] true: accept / false : deny
	Action bool `json:"action,omitempty"` // Not 'string' type

	// [Optional] TCP/UDP/ICMP/FTP/ALL
	Protocol string `json:"protocol,omitempty"`
	
	// [Optional] 'Required' when the protocol is TCP/UDP.
	StartPort string `json:"startPort,omitempty"`

	// [Optional] 'Required' when the protocol is TCP/UDP.
	EndPort string `json:"endPort,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with srcInterface
	SrcNetwork []string `json:"srcNetwork,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with dstInterface
	DstNetwork []string `json:"dstNetwork,omitempty"`

	// [Optional] Array parameter for IP subnet format or domain name format addresses
	SrcAddress []string `json:"srcAddress,omitempty"`

	// ### Caution!! : [Optional] Cannot be entered simultaneously with portForwardingId
	DstAddress []string `json:"dstAddress,omitempty"`

	// [Optional] Description
	Comment string `json:"comment,omitempty"`
}

// ToRuleCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "") // Not "firewall_rule" but ""
}

// Create accepts a CreateOpts struct and uses the values to create a new
// firewall rule.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Post(createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular firewall rule based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// Delete will permanently delete a particular firewall rule based on its
// unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// InsertRuleOpts contains the values used when inserting a rule into a policy.
type InsertRuleOpts struct {
	// FirewallRuleID is the ID of the rule to insert.
	FirewallRuleID string `json:"firewall_rule_id"`

	// InsertAfter is the ID of the rule to insert after.
	// If omitted, the rule is inserted at the top of the policy.
	InsertAfter string `json:"insert_after,omitempty"`

	// InsertBefore is the ID of the rule to insert before.
	// If omitted, the rule is inserted at the bottom of the policy.
	InsertBefore string `json:"insert_before,omitempty"`
}

// ToInsertRuleMap casts an InsertRuleOpts struct to a map.
func (opts InsertRuleOpts) ToInsertRuleMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// RemoveRuleOpts contains the values used when removing a rule from a policy.
type RemoveRuleOpts struct {
	// FirewallRuleID is the ID of the rule to remove.
	FirewallRuleID string `json:"firewall_rule_id"`
}

// ToRemoveRuleMap casts a RemoveRuleOpts struct to a map.
func (opts RemoveRuleOpts) ToRemoveRuleMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}


// // ValidateCreateOpts validates the CreateOpts structure.
// func ValidateCreateOpts(opts CreateOpts) error {
// 	if opts.Protocol == "" {
// 		return fmt.Errorf("protocol is required")
// 	}
// 	if opts.Action == "" {
// 		return fmt.Errorf("action is required")
// 	}

// 	// Validate action
// 	validActions := map[string]bool{"allow": true, "deny": true, "reject": true}
// 	if !validActions[opts.Action] {
// 		return fmt.Errorf("invalid action: %s. Valid actions are: allow, deny, reject", opts.Action)
// 	}

// 	// Validate IP version
// 	if opts.IPVersion != 0 && opts.IPVersion != 4 && opts.IPVersion != 6 {
// 		return fmt.Errorf("invalid IP version: %d. Valid versions are: 4, 6", opts.IPVersion)
// 	}

// 	return nil
// }

/*

// BatchCreateOpts contains options for creating multiple rules in a single request.
type BatchCreateOpts struct {
	// Rules is a slice of CreateOpts for batch creation.
	Rules []CreateOpts `json:"firewall_rules"`
}

// ToRuleBatchCreateMap casts a BatchCreateOpts struct to a map.
func (opts BatchCreateOpts) ToRuleBatchCreateMap() (map[string]interface{}, error) {
	// Validate each rule before creating the map
	for i, rule := range opts.Rules {
		if err := ValidateCreateOpts(rule); err != nil {
			return nil, fmt.Errorf("validation error for rule %d: %v", i, err)
		}
	}
	
	return gophercloud.BuildRequestBody(opts, "")
}

// BatchCreate accepts a BatchCreateOpts struct and uses the values to create
// multiple firewall rules in a single request.
func BatchCreate(c *gophercloud.ServiceClient, opts BatchCreateOpts) (r BatchCreateResult) {
	b, err := opts.ToRuleBatchCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(batchURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

*/
