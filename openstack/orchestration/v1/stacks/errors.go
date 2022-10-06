package stacks

import (
	"fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

type ErrInvalidEnvironment struct {
	ktvpcsdk.BaseError
	Section string
}

func (e ErrInvalidEnvironment) Error() string {
	return fmt.Sprintf("Environment has wrong section: %s", e.Section)
}

type ErrInvalidDataFormat struct {
	ktvpcsdk.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return fmt.Sprintf("Data in neither json nor yaml format.")
}

type ErrInvalidTemplateFormatVersion struct {
	ktvpcsdk.BaseError
	Version string
}

func (e ErrInvalidTemplateFormatVersion) Error() string {
	return fmt.Sprintf("Template format version not found.")
}

type ErrTemplateRequired struct {
	ktvpcsdk.BaseError
}

func (e ErrTemplateRequired) Error() string {
	return fmt.Sprintf("Template required for this function.")
}
