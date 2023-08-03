# Snippet

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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(988374),
    )

    ctx := context.Background()
    res, err := s.Snippet.Delete(ctx, operations.DeleteSnippetsRequest{
        XRequestID: writer.String("sapiente"),
        Ids: []string{
            "mollitia",
        },
        TeamID: 208876,
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(635059),
    )

    ctx := context.Background()
    res, err := s.Snippet.Find(ctx, operations.FindSnippetsRequest{
        Limit: writer.Int64(161309),
        Offset: writer.Int64(995300),
        Search: writer.String("mollitia"),
        Shortcuts: []string{
            "numquam",
            "commodi",
            "quam",
        },
        SortField: operations.FindSnippetsSortFieldCreationTime.ToPointer(),
        SortOrder: operations.FindSnippetsSortOrderAsc.ToPointer(),
        Tags: []string{
            "quia",
            "quis",
            "vitae",
        },
        TeamID: 674752,
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(656330),
    )

    ctx := context.Background()
    res, err := s.Snippet.Update(ctx, operations.UpdateSnippetsRequest{
        RequestBody: []shared.SnippetUpdate{
            shared.SnippetUpdate{
                Description: writer.String("odit"),
                ID: "c3f5ad01-9da1-4ffe-b8f0-97b0074f1547",
                Shortcut: writer.String("dicta"),
                Snippet: "harum",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "accusamus",
                    },
                    shared.SnippetTagV2{
                        Tag: "commodi",
                    },
                },
            },
            shared.SnippetUpdate{
                Description: writer.String("repudiandae"),
                ID: "13b99d48-8e1e-491e-850a-d2abd4426980",
                Shortcut: writer.String("magni"),
                Snippet: "assumenda",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "alias",
                    },
                    shared.SnippetTagV2{
                        Tag: "fugit",
                    },
                },
            },
        },
        XRequestID: writer.String("dolorum"),
        TeamID: 569618,
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

