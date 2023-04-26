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
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(100226),
    )

    ctx := context.Background()    
    req := operations.PageDetailsRequest{
        PageID: 99569,
    }

    res, err := s.Styleguide.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PageWithSectionResponse != nil {
        // handle response
    }
}
```

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
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(919483),
    )

    ctx := context.Background()    
    req := operations.ListPagesRequest{
        Limit: writer.Int64(352312),
        Offset: writer.Int64(714242),
        Status: operations.ListPagesStatusEnumLive.ToPointer(),
    }

    res, err := s.Styleguide.ListPages(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultPagePublicAPIResponse != nil {
        // handle response
    }
}
```
