# SchemaAPI
(*SchemaAPI*)

## Overview

Schema API - Core Schema related endpoints

### Available Operations

* [GetSchemaInfoList](#getschemainfolist) - Searches SchemaInfo repository
* [UpsertSchema](#upsertschema) - Creates/Updates a schema in development status
* [CreateSchema](#createschema) - Adds a schema to the schema repository.
* [GetSchema](#getschema) - Gets schema from the schema repository.

## GetSchemaInfoList

Searches for information of available schema (SchemaInfo) in schema repository. Support options to filter out the search contents. <p>Required roles:  `service.schema-service.viewers` groups to get the schema.</p>

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SchemaAPI.GetSchemaInfoList(ctx, operations.GetSchemaInfoListRequest{
        Authority: openapi.String("osdu"),
        Source: openapi.String("wks"),
        EntityType: openapi.String("wellbore"),
        SchemaVersionMajor: openapi.String("1"),
        SchemaVersionMinor: openapi.String("1"),
        SchemaVersionPatch: openapi.String("0"),
        Status: openapi.String("PUBLISHED"),
        Scope: openapi.String("INTERNAL"),
        LatestVersion: openapi.String("True"),
        Limit: openapi.String("10"),
        Offset: openapi.String("0"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSchemaInfoListRequest](../../models/operations/getschemainfolistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetSchemaInfoListResponse](../../models/operations/getschemainfolistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertSchema

Creates a new schema or updates an already existing schema with status `DEVELOPMENT` in the schema repository. If a user tries to create/update a schema with status other than `DEVELOPMENT`, API will throw an exception. <p>Any schema instance with the same schemaIdentity is replaced (in contrast to the immutability of `PUBLISHED` or `OBSOLETE` schemas). A schema state can also be changed from `DEVELOPMENT` to `PUBLISHED` or `OBSOLETE` while updating schema content or by providing the same schema content.</p> <p>**Note:** The schema may refer to other schema definitions in `DEVELOPMENT` state. If those schemas are updated themselves, it is the developer's responsibility to PUT the dependent schemas again to update the schemas. Scope for a schema can't be updated, its a system defined value.</p> <p>Required roles:  `service.schema-service.editors` groups to update schema.</p>

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
    res, err := s.SchemaAPI.UpsertSchema(ctx, components.SchemaRequest{
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

**[*operations.UpsertSchemaResponse](../../models/operations/upsertschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSchema

Adds a schema to the schema repository. The schemaIdentity must be unique. The `authority`, `source` and `entityType` will be registered if not present. <p>If lower minor versions are registered the service validates the new schema against breaking changes; if breaking changes are discovered the request fails.</p> <p>**Note:** The schema must not reference other schemas with status `DEVELOPMENT`. Scope to a schema will be set by system based on partition id (`SHARED` for common tenant and `INTERNAL` for private tenant). </p><p>Required roles : `service.schema-service.editors` groups to create schema.</p>

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
    res, err := s.SchemaAPI.CreateSchema(ctx, components.SchemaRequest{
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

**[*operations.CreateSchemaResponse](../../models/operations/createschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSchema

Retrieve a schema using its system defined id. Required roles:  `service.schema-service.viewers` groups to get the schema.

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SchemaAPI.GetSchema(ctx, "osdu:wks:wellbore:1.0.0")
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The system id of the schema                              | osdu:wks:wellbore:1.0.0                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSchemaResponse](../../models/operations/getschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |