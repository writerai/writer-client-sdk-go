# Terminology
(*Terminology*)

## Overview

Methods related to Terminology

### Available Operations

* [Add](#add) - Add terms
* [Delete](#delete) - Delete terms
* [Find](#find) - Find terms
* [Update](#update) - Update terms

## Add

Add terms

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
        writerclientsdkgo.WithOrganizationID(551477),
    )


    createTermsRequest := shared.CreateTermsRequest{
        Models: []shared.TermCreate{
            shared.TermCreate{
                ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                    Capitalize: false,
                    FixCase: false,
                    FixCommonMistakes: false,
                },
                CaseSensitive: false,
                Examples: []shared.TermExampleCreate{
                    shared.TermExampleCreate{
                        Example: "string",
                        Type: shared.TermExampleCreateTypeBad,
                    },
                },
                LinkedTerms: []shared.LinkedTermCreate{
                    shared.LinkedTermCreate{},
                },
                Mistakes: []shared.TermMistakeCreate{
                    shared.TermMistakeCreate{
                        CaseSensitive: false,
                        Mistake: "string",
                    },
                },
                Tags: []shared.TermTagCreate{
                    shared.TermTagCreate{
                        Tag: "string",
                    },
                },
                Term: "string",
                Type: shared.TermCreateTypeBanned,
            },
        },
    }

    var teamID int64 = 623445

    var organizationID *int64 = 822001

    ctx := context.Background()
    res, err := s.Terminology.Add(ctx, createTermsRequest, teamID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `createTermsRequest`                                                          | [shared.CreateTermsRequest](../../../pkg/models/shared/createtermsrequest.md) | :heavy_check_mark:                                                            | N/A                                                                           |
| `teamID`                                                                      | *int64*                                                                       | :heavy_check_mark:                                                            | N/A                                                                           |
| `organizationID`                                                              | **int64*                                                                      | :heavy_minus_sign:                                                            | N/A                                                                           |


### Response

**[*operations.AddTermsResponse](../../pkg/models/operations/addtermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |

## Delete

Delete terms

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

    ids := []int64{
        698486,
    }

    var organizationID *int64 = 557937

    ctx := context.Background()
    res, err := s.Terminology.Delete(ctx, teamID, xRequestID, ids, organizationID)
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
| `ids`                                                 | []*int64*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.DeleteTermsResponse](../../pkg/models/operations/deletetermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |

## Find

Find terms

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
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Tags: []string{
            "string",
        },
        TeamID: 326883,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultFullTermWithUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.FindTermsRequest](../../pkg/models/operations/findtermsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.FindTermsResponse](../../pkg/models/operations/findtermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |

## Update

Update terms

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


    updateTermsRequest := shared.UpdateTermsRequest{
        Models: []shared.TermUpdate{
            shared.TermUpdate{
                ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                    Capitalize: false,
                    FixCase: false,
                    FixCommonMistakes: false,
                },
                CaseSensitive: false,
                Examples: []shared.TermExampleCreate{
                    shared.TermExampleCreate{
                        Example: "string",
                        Type: shared.TermExampleCreateTypeGood,
                    },
                },
                ID: 597129,
                LinkedTerms: []shared.LinkedTermCreate{
                    shared.LinkedTermCreate{},
                },
                Mistakes: []shared.TermMistakeCreate{
                    shared.TermMistakeCreate{
                        CaseSensitive: false,
                        Mistake: "string",
                    },
                },
                Tags: []shared.TermTagCreate{
                    shared.TermTagCreate{
                        Tag: "string",
                    },
                },
                Term: "string",
                Type: shared.TermUpdateTypeApproved,
            },
        },
    }

    var teamID int64 = 344620

    var xRequestID *string = "string"

    var organizationID *int64 = 708455

    ctx := context.Background()
    res, err := s.Terminology.Update(ctx, updateTermsRequest, teamID, xRequestID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `updateTermsRequest`                                                          | [shared.UpdateTermsRequest](../../../pkg/models/shared/updatetermsrequest.md) | :heavy_check_mark:                                                            | N/A                                                                           |
| `teamID`                                                                      | *int64*                                                                       | :heavy_check_mark:                                                            | N/A                                                                           |
| `xRequestID`                                                                  | **string*                                                                     | :heavy_minus_sign:                                                            | N/A                                                                           |
| `organizationID`                                                              | **int64*                                                                      | :heavy_minus_sign:                                                            | N/A                                                                           |


### Response

**[*operations.UpdateTermsResponse](../../pkg/models/operations/updatetermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |
