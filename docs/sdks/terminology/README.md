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
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(551477),
    )

    ctx := context.Background()
    res, err := s.Terminology.Add(ctx, operations.AddTermsRequest{
        CreateTermsRequest: shared.CreateTermsRequest{
            FailHandling: shared.CreateTermsRequestFailHandlingValidateOnly.ToPointer(),
            Models: []shared.TermCreate{
                shared.TermCreate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writerclientsdkgo.String("Optional mobile secured line"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "noon bypass Chief",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writerclientsdkgo.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writerclientsdkgo.Int64(753323),
                            Reference: writerclientsdkgo.String("Southeast"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "generating payment quantify",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writerclientsdkgo.String("Urbandale Gasoline"),
                        },
                    },
                    Pos: shared.TermCreatePosAdverb.ToPointer(),
                    Reference: writerclientsdkgo.String("methodologies"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "Northeast",
                        },
                    },
                    Term: "Pomona auxiliary",
                    Type: shared.TermCreateTypeBanned,
                },
            },
        },
        TeamID: 739517,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.AddTermsRequest](../../models/operations/addtermsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


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
    res, err := s.Terminology.Delete(ctx, operations.DeleteTermsRequest{
        XRequestID: writerclientsdkgo.String("Designer complexity gladiolus"),
        Ids: []int64{
            921415,
        },
        TeamID: 953140,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteTermsRequest](../../models/operations/deletetermsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(40141),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Limit: writerclientsdkgo.Int64(326883),
        Offset: writerclientsdkgo.Int64(488098),
        PartOfSpeech: operations.FindTermsPartOfSpeechAdjective.ToPointer(),
        SortField: operations.FindTermsSortFieldType.ToPointer(),
        SortOrder: operations.FindTermsSortOrderAsc.ToPointer(),
        Tags: []string{
            "West",
        },
        TeamID: 413686,
        Term: writerclientsdkgo.String("Bike generating"),
        Type: operations.FindTermsTypeApproved.ToPointer(),
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
    res, err := s.Terminology.Update(ctx, operations.UpdateTermsRequest{
        UpdateTermsRequest: shared.UpdateTermsRequest{
            FailHandling: shared.UpdateTermsRequestFailHandlingAccumulate.ToPointer(),
            Models: []shared.TermUpdate{
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writerclientsdkgo.String("Persistent 24/7 focus group"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "dock Quality redundant",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writerclientsdkgo.Bool(false),
                    ID: 134151,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writerclientsdkgo.Int64(54062),
                            Reference: writerclientsdkgo.String("mostly"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "dynamic white",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writerclientsdkgo.String("Forward"),
                        },
                    },
                    Pos: shared.TermUpdatePosAdjective.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "East Baht Quality",
                        },
                    },
                    Term: "Home users Sharable",
                    Type: shared.TermUpdateTypeBanned,
                },
            },
        },
        XRequestID: writerclientsdkgo.String("Northeast Wooden"),
        TeamID: 352919,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateTermsRequest](../../models/operations/updatetermsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.UpdateTermsResponse](../../models/operations/updatetermsresponse.md), error**

