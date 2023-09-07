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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(716327),
    )

    ctx := context.Background()
    res, err := s.Document.Get(ctx, operations.GetDocumentDetailsRequest{
        DocumentID: 841386,
        TeamID: 289406,
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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(264730),
    )

    ctx := context.Background()
    res, err := s.Document.List(ctx, operations.ListTeamDocumentsRequest{
        Limit: writer.Int(183191),
        Offset: writer.Int64(397821),
        Search: writer.String("cupiditate"),
        SortField: operations.ListTeamDocumentsSortFieldModificationTime.ToPointer(),
        SortOrder: operations.ListTeamDocumentsSortOrderAsc.ToPointer(),
        TeamID: 164940,
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

