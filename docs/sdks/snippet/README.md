# Snippet
(*Snippet*)

## Overview

Methods related to Snippets

### Available Operations

* [Delete](#delete) - Delete snippets
* [Find](#find) - Find snippets
* [Update](#update) - Update snippets

## Delete

Delete snippets

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
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(545907),
    )

    ctx := context.Background()
    res, err := s.Snippet.Delete(ctx, operations.DeleteSnippetsRequest{
        XRequestID: writerclientsdkgo.String("Designer complexity gladiolus"),
        Ids: []string{
            "ick",
        },
        TeamID: 604931,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteSnippetsRequest](../../models/operations/deletesnippetsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteSnippetsResponse](../../models/operations/deletesnippetsresponse.md), error**


## Find

Find snippets

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
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(40141),
    )

    ctx := context.Background()
    res, err := s.Snippet.Find(ctx, operations.FindSnippetsRequest{
        Limit: writerclientsdkgo.Int64(326883),
        Offset: writerclientsdkgo.Int64(488098),
        Search: writerclientsdkgo.String("whoa West Credit"),
        Shortcuts: []string{
            "Granite",
        },
        SortField: operations.FindSnippetsSortFieldCreationTime.ToPointer(),
        SortOrder: operations.FindSnippetsSortOrderAsc.ToPointer(),
        Tags: []string{
            "Iowa",
        },
        TeamID: 785366,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultSnippetWithUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.FindSnippetsRequest](../../models/operations/findsnippetsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.FindSnippetsResponse](../../models/operations/findsnippetsresponse.md), error**


## Update

Update snippets

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
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(857478),
    )

    ctx := context.Background()
    res, err := s.Snippet.Update(ctx, operations.UpdateSnippetsRequest{
        RequestBody: []shared.SnippetUpdate{
            shared.SnippetUpdate{
                Description: writerclientsdkgo.String("Ameliorated methodical access"),
                ID: "<ID>",
                Shortcut: writerclientsdkgo.String("male Metal"),
                Snippet: "Arizona Cotton extend",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "bifurcated",
                    },
                },
            },
        },
        XRequestID: writerclientsdkgo.String("silver immediately"),
        TeamID: 302461,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetWithUsers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateSnippetsRequest](../../models/operations/updatesnippetsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateSnippetsResponse](../../models/operations/updatesnippetsresponse.md), error**

