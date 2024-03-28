# Document
(*Document*)

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


    var documentID int64 = 90065

    var teamID int64 = 558834

    var organizationID *int64 = writerclientsdkgo.Int64(844199)

    ctx := context.Background()
    res, err := s.Document.Get(ctx, documentID, teamID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Document != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `documentID`                                          | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `teamID`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.GetDocumentDetailsResponse](../../pkg/models/operations/getdocumentdetailsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## List

List team documents

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(768578),
    )

    ctx := context.Background()
    res, err := s.Document.List(ctx, operations.ListTeamDocumentsRequest{
        TeamID: 99895,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTeamDocumentsRequest](../../pkg/models/operations/listteamdocumentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListTeamDocumentsResponse](../../pkg/models/operations/listteamdocumentsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
