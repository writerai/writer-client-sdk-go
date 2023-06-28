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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(885338),
    )

    ctx := context.Background()
    res, err := s.Document.Get(ctx, operations.GetDocumentDetailsRequest{
        DocumentID: 185636,
        TeamID: 679880,
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(952792),
    )

    ctx := context.Background()
    res, err := s.Document.List(ctx, operations.ListTeamDocumentsRequest{
        Limit: writer.Int(456130),
        Offset: writer.Int64(687488),
        Search: writer.String("iusto"),
        SortField: operations.ListTeamDocumentsSortFieldCreationTime.ToPointer(),
        SortOrder: operations.ListTeamDocumentsSortOrderDesc.ToPointer(),
        TeamID: 947371,
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

