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
                        Example: "calculate Toyota noon",
                        Type: shared.TermExampleCreateTypeBad,
                    },
                },
                LinkedTerms: []shared.LinkedTermCreate{
                    shared.LinkedTermCreate{},
                },
                Mistakes: []shared.TermMistakeCreate{
                    shared.TermMistakeCreate{
                        CaseSensitive: false,
                        Mistake: "Chief",
                    },
                },
                Tags: []shared.TermTagCreate{
                    shared.TermTagCreate{
                        Tag: "kelvin",
                    },
                },
                Term: "lime",
                Type: shared.TermCreateTypeBanned,
            },
        },
    }

    var teamID int64 = 623862

    var organizationID *int64 = 445859

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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `createTermsRequest`                                                   | [shared.CreateTermsRequest](../../models/shared/createtermsrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `teamID`                                                               | *int64*                                                                | :heavy_check_mark:                                                     | N/A                                                                    |
| `organizationID`                                                       | **int64*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |


### Response

**[*operations.AddTermsResponse](../../models/operations/addtermsresponse.md), error**


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

    var xRequestID *string = "Designer"

    ids := []int64{
        386564,
    }

    var organizationID *int64 = 201447

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

**[*operations.DeleteTermsResponse](../../models/operations/deletetermsresponse.md), error**


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
            "underestimate",
        },
        TeamID: 111247,
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.FindTermsRequest](../../models/operations/findtermsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.FindTermsResponse](../../models/operations/findtermsresponse.md), error**


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
                        Example: "Rock",
                        Type: shared.TermExampleCreateTypeGood,
                    },
                },
                ID: 708455,
                LinkedTerms: []shared.LinkedTermCreate{
                    shared.LinkedTermCreate{},
                },
                Mistakes: []shared.TermMistakeCreate{
                    shared.TermMistakeCreate{
                        CaseSensitive: false,
                        Mistake: "Metal cheater Islands",
                    },
                },
                Tags: []shared.TermTagCreate{
                    shared.TermTagCreate{
                        Tag: "withdrawal extend",
                    },
                },
                Term: "bifurcated",
                Type: shared.TermUpdateTypeBanned,
            },
        },
    }

    var teamID int64 = 789275

    var xRequestID *string = "syndicate"

    var organizationID *int64 = 345187

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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `updateTermsRequest`                                                   | [shared.UpdateTermsRequest](../../models/shared/updatetermsrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `teamID`                                                               | *int64*                                                                | :heavy_check_mark:                                                     | N/A                                                                    |
| `xRequestID`                                                           | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `organizationID`                                                       | **int64*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |


### Response

**[*operations.UpdateTermsResponse](../../models/operations/updatetermsresponse.md), error**

