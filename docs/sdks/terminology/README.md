# Terminology

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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(998848),
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
                    Description: writer.String("saepe"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "accusantium",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "praesentium",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "magni",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "quo",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(807319),
                            Reference: writer.String("ea"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(569101),
                            Reference: writer.String("odit"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(407183),
                            Reference: writer.String("accusantium"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(69167),
                            Reference: writer.String("maiores"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "ipsam",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writer.String("autem"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "nam",
                            Pos: shared.TermMistakeCreatePosNoun.ToPointer(),
                            Reference: writer.String("pariatur"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "nemo",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("perferendis"),
                        },
                    },
                    Pos: shared.TermCreatePosAdjective.ToPointer(),
                    Reference: writer.String("amet"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "cumque",
                        },
                    },
                    Term: "corporis",
                    Type: shared.TermCreateTypePending,
                },
            },
        },
        TeamID: 729991,
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(749999),
    )

    ctx := context.Background()
    res, err := s.Terminology.Delete(ctx, operations.DeleteTermsRequest{
        XRequestID: writer.String("dolores"),
        Ids: []int64{
            521037,
            489549,
        },
        TeamID: 54338,
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(338985),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Limit: writer.Int64(199996),
        Offset: writer.Int64(179490),
        PartOfSpeech: operations.FindTermsPartOfSpeechNoun.ToPointer(),
        SortField: operations.FindTermsSortFieldTerm.ToPointer(),
        SortOrder: operations.FindTermsSortOrderDesc.ToPointer(),
        Tags: []string{
            "dolor",
            "vero",
        },
        TeamID: 345352,
        Term: writer.String("hic"),
        Type: operations.FindTermsTypePending.ToPointer(),
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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(608253),
    )

    ctx := context.Background()
    res, err := s.Terminology.Update(ctx, operations.UpdateTermsRequest{
        UpdateTermsRequest: shared.UpdateTermsRequest{
            FailHandling: shared.UpdateTermsRequestFailHandlingSkip.ToPointer(),
            Models: []shared.TermUpdate{
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("voluptatem"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "consequuntur",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "error",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "occaecati",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "adipisci",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 934214,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(613966),
                            Reference: writer.String("dolorum"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(535633),
                            Reference: writer.String("pariatur"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "nobis",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("delectus"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "quaerat",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("aliquid"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "dolorem",
                            Pos: shared.TermMistakeCreatePosNoun.ToPointer(),
                            Reference: writer.String("dolor"),
                        },
                    },
                    Pos: shared.TermUpdatePosNoun.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "hic",
                        },
                    },
                    Term: "excepturi",
                    Type: shared.TermUpdateTypePending,
                },
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("voluptate"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "reiciendis",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "dolorum",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 85295,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(56418),
                            Reference: writer.String("iure"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "quaerat",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("quidem"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "voluptatibus",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writer.String("natus"),
                        },
                    },
                    Pos: shared.TermUpdatePosNoun.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "sit",
                        },
                        shared.TermTagCreate{
                            Tag: "fugiat",
                        },
                        shared.TermTagCreate{
                            Tag: "ab",
                        },
                    },
                    Term: "soluta",
                    Type: shared.TermUpdateTypePending,
                },
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("iusto"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "dolorum",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "omnis",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 714697,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(469497),
                            Reference: writer.String("ipsum"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(456015),
                            Reference: writer.String("id"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(906418),
                            Reference: writer.String("eius"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(137220),
                            Reference: writer.String("perferendis"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "optio",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("ad"),
                        },
                    },
                    Pos: shared.TermUpdatePosAdjective.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "deserunt",
                        },
                        shared.TermTagCreate{
                            Tag: "provident",
                        },
                    },
                    Term: "minima",
                    Type: shared.TermUpdateTypePending,
                },
            },
        },
        XRequestID: writer.String("totam"),
        TeamID: 628982,
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

