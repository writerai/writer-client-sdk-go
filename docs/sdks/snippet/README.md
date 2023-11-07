# Snippet
(*.Snippet*)

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
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(545907),
    )


    var teamID int64 = 841399

    var xRequestID *string = "string"

    ids := []string{
        "string",
    }

    var organizationID *int64 = 698486

    ctx := context.Background()
    res, err := s.Snippet.Delete(ctx, teamID, xRequestID, ids, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `teamID`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `xRequestID`                                          | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `ids`                                                 | []*string*                                            | :heavy_minus_sign:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


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
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(40141),
    )

    ctx := context.Background()
    res, err := s.Snippet.Find(ctx, operations.FindSnippetsRequest{
        Shortcuts: []string{
            "string",
        },
        Tags: []string{
            "string",
        },
        TeamID: 326883,
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
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(857478),
    )


    var teamID int64 = 24555

    requestBody := []shared.SnippetUpdate{
        shared.SnippetUpdate{
            ID: "<ID>",
            Snippet: "string",
            Tags: []shared.SnippetTagV2{
                shared.SnippetTagV2{
                    Tag: "string",
                },
            },
        },
    }

    var xRequestID *string = "string"

    var organizationID *int64 = 597129

    ctx := context.Background()
    res, err := s.Snippet.Update(ctx, teamID, requestBody, xRequestID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `teamID`                                                       | *int64*                                                        | :heavy_check_mark:                                             | N/A                                                            |
| `requestBody`                                                  | [][shared.SnippetUpdate](../../models/shared/snippetupdate.md) | :heavy_minus_sign:                                             | N/A                                                            |
| `xRequestID`                                                   | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `organizationID`                                               | **int64*                                                       | :heavy_minus_sign:                                             | N/A                                                            |


### Response

**[*operations.UpdateSnippetsResponse](../../models/operations/updatesnippetsresponse.md), error**

