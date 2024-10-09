// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetSchemaRequest struct {
	// The system id of the schema
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSchemaRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetSchemaResponseBody - Search results matching criteria
type GetSchemaResponseBody struct {
}

type GetSchemaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Body     []byte
}

func (o *GetSchemaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSchemaResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
