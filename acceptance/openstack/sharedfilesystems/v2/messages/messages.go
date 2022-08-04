package messages

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/sharedfilesystems/v2/messages"
)

// DeleteMessage will delete a message. An error will occur if
// the message was unable to be deleted.
func DeleteMessage(t *testing.T, client *gophercloud.ServiceClient, message *messages.Message) {
	err := messages.Delete(client, message.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Failed to delete message %s: %v", message.ID, err)
	}

	t.Logf("Deleted message: %s", message.ID)
}
