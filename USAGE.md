<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
	"openapi/types"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.SystemSchemaAPI.UpsertSystemSchema(ctx, components.SchemaRequest{
		SchemaInfo: components.SchemaInfo{
			SchemaIdentity: components.SchemaIdentity{
				Authority:          "osdu",
				Source:             "wks",
				EntityType:         "wellbore",
				SchemaVersionMajor: 1,
				SchemaVersionMinor: 1,
				SchemaVersionPatch: 0,
				ID:                 openapi.String("osdu:wks:wellbore:1.0.0"),
			},
			CreatedBy:   openapi.String("user@opendes.com"),
			DateCreated: types.MustNewTimeFromString("2019-05-23T11:16:03Z"),
			Status:      components.SchemaStatusPublished,
			Scope:       components.SchemaScopeInternal.ToPointer(),
			SupersededBy: &components.SchemaIdentity{
				Authority:          "osdu",
				Source:             "wks",
				EntityType:         "wellbore",
				SchemaVersionMajor: 1,
				SchemaVersionMinor: 1,
				SchemaVersionPatch: 0,
				ID:                 openapi.String("osdu:wks:wellbore:1.0.0"),
			},
		},
		Schema: components.Schema{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->