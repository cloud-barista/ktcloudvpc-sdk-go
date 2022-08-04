//go:build acceptance || networking || fwaas
// +build acceptance networking fwaas

package fwaas

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/networking/v2/extensions/fwaas/rules"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestRuleCRUD(t *testing.T) {
	t.Skip("Skip this test, FWAAS v1 is old and will be removed from Gophercloud")
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	rule, err := CreateRule(t, client)
	th.AssertNoErr(t, err)
	defer DeleteRule(t, client, rule.ID)

	tools.PrintResource(t, rule)

	ruleDescription := "Some rule description"
	updateOpts := rules.UpdateOpts{
		Description: &ruleDescription,
	}

	_, err = rules.Update(client, rule.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newRule, err := rules.Get(client, rule.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRule)

	allPages, err := rules.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allRules, err := rules.ExtractRules(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, rule := range allRules {
		if rule.ID == newRule.ID {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
