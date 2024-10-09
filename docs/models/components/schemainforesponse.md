# SchemaInfoResponse

The response for a GET schema request


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `SchemaInfos`                                                    | [][components.SchemaInfo](../../models/components/schemainfo.md) | :heavy_minus_sign:                                               | N/A                                                              |
| `Offset`                                                         | **int*                                                           | :heavy_minus_sign:                                               | The offset for the next query                                    |
| `Count`                                                          | **int*                                                           | :heavy_minus_sign:                                               | The number of schema versions in this response                   |
| `TotalCount`                                                     | **int*                                                           | :heavy_minus_sign:                                               | The total number of entity type codes in the repositories        |