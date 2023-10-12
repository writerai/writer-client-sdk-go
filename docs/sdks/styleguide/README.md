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
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
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

**[*operations.PageDetailsResponse](../../models/operations/pagedetailsresponse.md), error**


## ListPages

List your styleguide pages

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(763372),
    )


    var limit *int64 = 760116

    var offset *int64 = 303332

    var status *operations.ListPagesStatus = operations.ListPagesStatusLive

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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `limit`                                                                   | **int64*                                                                  | :heavy_minus_sign:                                                        | N/A                                                                       |
| `offset`                                                                  | **int64*                                                                  | :heavy_minus_sign:                                                        | N/A                                                                       |
| `status`                                                                  | [*operations.ListPagesStatus](../../models/operations/listpagesstatus.md) | :heavy_minus_sign:                                                        | N/A                                                                       |


### Response

**[*operations.ListPagesResponse](../../models/operations/listpagesresponse.md), error**

