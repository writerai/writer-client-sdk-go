# Styleguide

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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(100226),
    )

    ctx := context.Background()
    res, err := s.Styleguide.Get(ctx, operations.PageDetailsRequest{
        PageID: 99569,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PageWithSectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PageDetailsRequest](../../models/operations/pagedetailsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(919483),
    )

    ctx := context.Background()
    res, err := s.Styleguide.ListPages(ctx, operations.ListPagesRequest{
        Limit: writer.Int64(352312),
        Offset: writer.Int64(714242),
        Status: operations.ListPagesStatusLive.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultPagePublicAPIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListPagesRequest](../../models/operations/listpagesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.ListPagesResponse](../../models/operations/listpagesresponse.md), error**
