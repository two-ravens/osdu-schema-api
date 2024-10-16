// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type UpsertSystemSchemaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Body     []byte
}

func (o *UpsertSystemSchemaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpsertSystemSchemaResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
