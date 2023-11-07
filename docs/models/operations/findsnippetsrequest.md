# FindSnippetsRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Limit`                                                       | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `Offset`                                                      | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `OrganizationID`                                              | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `Search`                                                      | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Shortcuts`                                                   | []*string*                                                    | :heavy_minus_sign:                                            | N/A                                                           |
| `SortField`                                                   | [*operations.SortField](../../models/operations/sortfield.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `SortOrder`                                                   | [*operations.SortOrder](../../models/operations/sortorder.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `Tags`                                                        | []*string*                                                    | :heavy_minus_sign:                                            | N/A                                                           |
| `TeamID`                                                      | *int64*                                                       | :heavy_check_mark:                                            | N/A                                                           |