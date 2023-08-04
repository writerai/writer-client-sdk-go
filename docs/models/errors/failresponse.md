# FailResponse

Bad Request


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `RawResponse`                                              | [*http.Response](https://pkg.go.dev/net/http#Response)     | :heavy_minus_sign:                                         | N/A                                                        |
| `Errors`                                                   | [][shared.FailMessage](../../models/shared/failmessage.md) | :heavy_minus_sign:                                         | N/A                                                        |
| `Extras`                                                   | *interface{}*                                              | :heavy_check_mark:                                         | N/A                                                        |
| `Tpe`                                                      | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |