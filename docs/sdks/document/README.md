# Document

## Overview

Methods related to document

### Available Operations

* [Get](#get) - Get document details
* [List](#list) - List team documents

## Get

Get document details

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
        writerclientsdkgo.WithOrganizationID(289406),
    )

    ctx := context.Background()
    res, err := s.Document.Get(ctx, operations.GetDocumentDetailsRequest{
        DocumentID: 264730,
        TeamID: 183191,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Document != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetDocumentDetailsRequest](../../models/operations/getdocumentdetailsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetDocumentDetailsResponse](../../models/operations/getdocumentdetailsresponse.md), error**


## List

List team documents

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
        writerclientsdkgo.WithOrganizationID(397821),
    )

    ctx := context.Background()
    res, err := s.Document.List(ctx, operations.ListTeamDocumentsRequest{
        Limit: writerclientsdkgo.Int(586513),
        Offset: writerclientsdkgo.Int64(552822),
        Search: writerclientsdkgo.String("perferendis"),
        SortField: operations.ListTeamDocumentsSortFieldTitle.ToPointer(),
        SortOrder: operations.ListTeamDocumentsSortOrderDesc.ToPointer(),
        TeamID: 369808,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BriefDocuments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListTeamDocumentsRequest](../../models/operations/listteamdocumentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListTeamDocumentsResponse](../../models/operations/listteamdocumentsresponse.md), error**

