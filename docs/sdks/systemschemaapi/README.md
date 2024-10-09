# SystemSchemaAPI
(*SystemSchemaAPI*)

## Overview

System Schema API - System Schema related endpoints

### Available Operations

* [UpsertSystemSchema](#upsertsystemschema) - Creates/Updates a schema in development status

## UpsertSystemSchema

Creates a new schema or updates an already existing schema with status `DEVELOPMENT` in the schema repository. If a user tries to create a schema with status other than `DEVELOPMENT`, API will not throw an exception. <p>The update of schema without `DEVELOPMENT` status would cause error. Any schema instance with the same schemaIdentity is replaced. A schema state can also be changed from `DEVELOPMENT` to `PUBLISHED` or `OBSOLETE` while updating schema content or by providing the same schema content.</p> <p>**Note:** The schema may refer to other schema definitions in `DEVELOPMENT` state. If those schemas are updated themselves, it is the developer's responsibility to PUT the dependent schemas again to update the schema. Scope for a schema will be SHARED for all the schemas created using this API.</p><p>Service principal authorization is required to call thi API.</p>

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/components"
	"openapi/types"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SystemSchemaAPI.UpsertSystemSchema(ctx, components.SchemaRequest{
        SchemaInfo: components.SchemaInfo{
            SchemaIdentity: components.SchemaIdentity{
                Authority: "osdu",
                Source: "wks",
                EntityType: "wellbore",
                SchemaVersionMajor: 1,
                SchemaVersionMinor: 1,
                SchemaVersionPatch: 0,
                ID: openapi.String("osdu:wks:wellbore:1.0.0"),
            },
            CreatedBy: openapi.String("user@opendes.com"),
            DateCreated: types.MustNewTimeFromString("2019-05-23T11:16:03Z"),
            Status: components.SchemaStatusPublished,
            Scope: components.SchemaScopeInternal.ToPointer(),
            SupersededBy: &components.SchemaIdentity{
                Authority: "osdu",
                Source: "wks",
                EntityType: "wellbore",
                SchemaVersionMajor: 1,
                SchemaVersionMinor: 1,
                SchemaVersionPatch: 0,
                ID: openapi.String("osdu:wks:wellbore:1.0.0"),
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

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SchemaRequest](../../models/components/schemarequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.UpsertSystemSchemaResponse](../../models/operations/upsertsystemschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |