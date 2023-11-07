# PagePublicAPIResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CreatedAt`                                               | [time.Time](https://pkg.go.dev/time#Time)                 | :heavy_check_mark:                                        | N/A                                                       |
| `ID`                                                      | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `Order`                                                   | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `Section`                                                 | [*shared.SectionInfo](../../models/shared/sectioninfo.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `Status`                                                  | [shared.Status](../../models/shared/status.md)            | :heavy_check_mark:                                        | N/A                                                       |
| `Title`                                                   | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `UpdatedAt`                                               | [time.Time](https://pkg.go.dev/time#Time)                 | :heavy_check_mark:                                        | N/A                                                       |
| `UpdatedBy`                                               | [*shared.SimpleUser](../../models/shared/simpleuser.md)   | :heavy_minus_sign:                                        | N/A                                                       |
| `URL`                                                     | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |