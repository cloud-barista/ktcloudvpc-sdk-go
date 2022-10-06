package members

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func imageMembersURL(c *ktvpcsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID, "members")
}

func listMembersURL(c *ktvpcsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func createMemberURL(c *ktvpcsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func imageMemberURL(c *ktvpcsdk.ServiceClient, imageID string, memberID string) string {
	return c.ServiceURL("images", imageID, "members", memberID)
}

func getMemberURL(c *ktvpcsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func updateMemberURL(c *ktvpcsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func deleteMemberURL(c *ktvpcsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}
