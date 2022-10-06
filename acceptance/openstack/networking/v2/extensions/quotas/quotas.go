package quotas

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/networking/v2/extensions/quotas"
)

var updateOpts = quotas.UpdateOpts{
	FloatingIP:        ktvpcsdk.IntToPointer(10),
	Network:           ktvpcsdk.IntToPointer(-1),
	Port:              ktvpcsdk.IntToPointer(11),
	RBACPolicy:        ktvpcsdk.IntToPointer(0),
	Router:            ktvpcsdk.IntToPointer(-1),
	SecurityGroup:     ktvpcsdk.IntToPointer(12),
	SecurityGroupRule: ktvpcsdk.IntToPointer(13),
	Subnet:            ktvpcsdk.IntToPointer(14),
	SubnetPool:        ktvpcsdk.IntToPointer(15),
}
