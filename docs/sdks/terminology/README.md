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
        writer.WithOrganizationID(248753),
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
                    Description: writer.String("aliquid"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "necessitatibus",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "officia",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "debitis",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(447125),
                            Reference: writer.String("in"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(846409),
                            Reference: writer.String("maiores"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(699479),
                            Reference: writer.String("dicta"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "cumque",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("ea"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "aliquid",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("accusamus"),
                        },
                    },
                    Pos: shared.TermCreatePosNoun.ToPointer(),
                    Reference: writer.String("occaecati"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "accusamus",
                        },
                        shared.TermTagCreate{
                            Tag: "delectus",
                        },
                    },
                    Term: "quidem",
                    Type: shared.TermCreateTypeBanned,
                },
                shared.TermCreate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("nam"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "blanditiis",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "sapiente",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "deserunt",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(618809),
                            Reference: writer.String("omnis"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(474867),
                            Reference: writer.String("perferendis"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "magnam",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("id"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "labore",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writer.String("suscipit"),
                        },
                    },
                    Pos: shared.TermCreatePosAdverb.ToPointer(),
                    Reference: writer.String("nobis"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "vero",
                        },
                        shared.TermTagCreate{
                            Tag: "aspernatur",
                        },
                    },
                    Term: "architecto",
                    Type: shared.TermCreateTypeApproved,
                },
                shared.TermCreate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("et"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "ullam",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "quos",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                        shared.TermExampleCreate{
                            Example: "accusantium",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(652103),
                            Reference: writer.String("ad"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(431418),
                            Reference: writer.String("dolor"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(896547),
                            Reference: writer.String("odit"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(367562),
                            Reference: writer.String("quasi"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "doloribus",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("eius"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "maxime",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("facilis"),
                        },
                    },
                    Pos: shared.TermCreatePosVerb.ToPointer(),
                    Reference: writer.String("architecto"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "repudiandae",
                        },
                    },
                    Term: "ullam",
                    Type: shared.TermCreateTypePending,
                },
            },
        },
        TeamID: 469249,
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
        writer.WithOrganizationID(998848),
    )

    ctx := context.Background()
    res, err := s.Terminology.Delete(ctx, operations.DeleteTermsRequest{
        XRequestID: writer.String("quibusdam"),
        Ids: []int64{
            904648,
        },
        TeamID: 868126,
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
        writer.WithOrganizationID(37559),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Limit: writer.Int64(162493),
        Offset: writer.Int64(508315),
        PartOfSpeech: operations.FindTermsPartOfSpeechAdverb.ToPointer(),
        SortField: operations.FindTermsSortFieldTerm.ToPointer(),
        SortOrder: operations.FindTermsSortOrderAsc.ToPointer(),
        Tags: []string{
            "illum",
            "pariatur",
            "maxime",
            "ea",
        },
        TeamID: 569101,
        Term: writer.String("odit"),
        Type: operations.FindTermsTypeBanned.ToPointer(),
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
        writer.WithOrganizationID(33222),
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
                    Description: writer.String("quidem"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "voluptate",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "nam",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 866383,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(975522),
                            Reference: writer.String("perferendis"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(855804),
                            Reference: writer.String("amet"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "cumque",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writer.String("hic"),
                        },
                    },
                    Pos: shared.TermUpdatePosAdverb.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "dolores",
                        },
                        shared.TermTagCreate{
                            Tag: "quis",
                        },
                        shared.TermTagCreate{
                            Tag: "totam",
                        },
                    },
                    Term: "dignissimos",
                    Type: shared.TermUpdateTypeApproved,
                },
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("quis"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "eos",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 170986,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(463451),
                            Reference: writer.String("dolor"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(874573),
                            Reference: writer.String("nostrum"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(944120),
                            Reference: writer.String("recusandae"),
                        },
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(608253),
                            Reference: writer.String("facilis"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "voluptatem",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("consequuntur"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "blanditiis",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("eaque"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "occaecati",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("adipisci"),
                        },
                    },
                    Pos: shared.TermUpdatePosAdjective.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "modi",
                        },
                        shared.TermTagCreate{
                            Tag: "iste",
                        },
                        shared.TermTagCreate{
                            Tag: "dolorum",
                        },
                        shared.TermTagCreate{
                            Tag: "deleniti",
                        },
                    },
                    Term: "pariatur",
                    Type: shared.TermUpdateTypeBanned,
                },
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("nobis"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "delectus",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "quos",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "dolorem",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 222443,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(218749),
                            Reference: writer.String("hic"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "cum",
                            Pos: shared.TermMistakeCreatePosVerb.ToPointer(),
                            Reference: writer.String("dignissimos"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "reiciendis",
                            Pos: shared.TermMistakeCreatePosNoun.ToPointer(),
                            Reference: writer.String("dolorum"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "numquam",
                            Pos: shared.TermMistakeCreatePosNoun.ToPointer(),
                            Reference: writer.String("ipsa"),
                        },
                    },
                    Pos: shared.TermUpdatePosNoun.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "odio",
                        },
                        shared.TermTagCreate{
                            Tag: "quaerat",
                        },
                    },
                    Term: "accusamus",
                    Type: shared.TermUpdateTypePending,
                },
                shared.TermUpdate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("voluptatibus"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "natus",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                        shared.TermExampleCreate{
                            Example: "atque",
                            Type: shared.TermExampleCreateTypeGood,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 854614,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(743835),
                            Reference: writer.String("dolorum"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "voluptate",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("deleniti"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "omnis",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("distinctio"),
                        },
                    },
                    Pos: shared.TermUpdatePosAdjective.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "ipsum",
                        },
                        shared.TermTagCreate{
                            Tag: "voluptate",
                        },
                    },
                    Term: "id",
                    Type: shared.TermUpdateTypePending,
                },
            },
        },
        XRequestID: writer.String("eius"),
        TeamID: 137220,
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

