# Styleguide
(*Styleguide*)

## Overview

Methods related to Styleguide

### Available Operations

* [Get](#get) - Page details
* [ListPages](#listpages) - List your styleguide pages

## Get

Page details

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(700347),
    )


    var pageID int64 = 90065

    ctx := context.Background()
    res, err := s.Styleguide.Get(ctx, pageID)
    if err != nil {
        log.Fatal(err)
    }
    if res.PageWithSectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `pageID`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.PageDetailsResponse](../../pkg/models/operations/pagedetailsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## ListPages

List your styleguide pages

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(763372),
    )


    var limit *int64 = writerclientsdkgo.Int64(760116)

    var offset *int64 = writerclientsdkgo.Int64(303332)

    var status *operations.Status = operations.StatusLive.ToPointer()

    ctx := context.Background()
    res, err := s.Styleguide.ListPages(ctx, limit, offset, status)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedResultPagePublicAPIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `limit`                                                     | **int64*                                                    | :heavy_minus_sign:                                          | N/A                                                         |
| `offset`                                                    | **int64*                                                    | :heavy_minus_sign:                                          | N/A                                                         |
| `status`                                                    | [*operations.Status](../../pkg/models/operations/status.md) | :heavy_minus_sign:                                          | N/A                                                         |


### Response

**[*operations.ListPagesResponse](../../pkg/models/operations/listpagesresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
