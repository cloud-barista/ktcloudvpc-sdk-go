package messages

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// commonResult is the response of a base result.
type commonResult struct {
	ktvpcsdk.Result
}

// CreateResult is the response of a Create operations.
type CreateResult struct {
	ktvpcsdk.Result
}

// MessagePage contains a single page of all clusters from a ListDetails call.
type MessagePage struct {
	pagination.LinkedPageBase
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	ktvpcsdk.ErrResult
}

// CreateResult is the response of a Create operations.
type PopResult struct {
	ktvpcsdk.Result
}

// GetMessagesResult is the response of a GetMessages operations.
type GetMessagesResult struct {
	ktvpcsdk.Result
}

// GetResult is the response of a Get operations.
type GetResult struct {
	ktvpcsdk.Result
}

// Message represents a message on a queue.
type Message struct {
	Body     map[string]interface{} `json:"body"`
	Age      int                    `json:"age"`
	Href     string                 `json:"href"`
	ID       string                 `json:"id"`
	TTL      int                    `json:"ttl"`
	Checksum string                 `json:"checksum"`
}

// PopMessage represents a message returned from PopMessages.
type PopMessage struct {
	Body       map[string]interface{} `json:"body"`
	Age        int                    `json:"age"`
	ID         string                 `json:"id"`
	TTL        int                    `json:"ttl"`
	ClaimCount int                    `json:"claim_count"`
	ClaimID    string                 `json:"claim_id"`
}

// ResourceList represents the result of creating a message.
type ResourceList struct {
	Resources []string `json:"resources"`
}

// Extract interprets any CreateResult as a ResourceList.
func (r CreateResult) Extract() (ResourceList, error) {
	var s ResourceList
	err := r.ExtractInto(&s)
	return s, err
}

// Extract interprets any PopResult as a list of PopMessage.
func (r PopResult) Extract() ([]PopMessage, error) {
	var s struct {
		PopMessages []PopMessage `json:"messages"`
	}
	err := r.ExtractInto(&s)
	return s.PopMessages, err
}

// Extract interprets any GetMessagesResult as a list of Message.
func (r GetMessagesResult) Extract() ([]Message, error) {
	var s struct {
		Messages []Message `json:"messages"`
	}
	err := r.ExtractInto(&s)
	return s.Messages, err
}

// Extract interprets any GetResult as a Message.
func (r GetResult) Extract() (Message, error) {
	var s Message
	err := r.ExtractInto(&s)
	return s, err
}

// ExtractMessage extracts message into a  list of Message.
func ExtractMessages(r pagination.Page) ([]Message, error) {
	var s struct {
		Messages []Message `json:"messages"`
	}
	err := (r.(MessagePage)).ExtractInto(&s)
	return s.Messages, err
}

// IsEmpty determines if a MessagePage contains any results.
func (r MessagePage) IsEmpty() (bool, error) {
	s, err := ExtractMessages(r)
	return len(s) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r MessagePage) NextPageURL() (string, error) {
	var s struct {
		Links []ktvpcsdk.Link `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}

	next, err := ktvpcsdk.ExtractNextURL(s.Links)
	if err != nil {
		return "", err
	}
	return nextPageURL(r.URL.String(), next)
}
