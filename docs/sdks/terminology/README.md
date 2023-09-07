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
        writer.WithOrganizationID(13571),
    )

    ctx := context.Background()
    res, err := s.Terminology.Add(ctx, operations.AddTermsRequest{
        CreateTermsRequest: shared.CreateTermsRequest{
            FailHandling: shared.CreateTermsRequestFailHandlingAccumulate.ToPointer(),
            Models: []shared.TermCreate{
                shared.TermCreate{
                    ApprovedTermExtension: &shared.ApprovedTermExtensionCreate{
                        Capitalize: false,
                        FixCase: false,
                        FixCommonMistakes: false,
                    },
                    CaseSensitive: false,
                    Description: writer.String("error"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "temporibus",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(96098),
                            Reference: writer.String("reiciendis"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "voluptatibus",
                            Pos: shared.TermMistakeCreatePosAdjective.ToPointer(),
                            Reference: writer.String("nihil"),
                        },
                    },
                    Pos: shared.TermCreatePosAdverb.ToPointer(),
                    Reference: writer.String("voluptatibus"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "ipsa",
                        },
                    },
                    Term: "omnis",
                    Type: shared.TermCreateTypeBanned,
                },
            },
        },
        TeamID: 739264,
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
        writer.WithOrganizationID(19987),
    )

    ctx := context.Background()
    res, err := s.Terminology.Delete(ctx, operations.DeleteTermsRequest{
        XRequestID: writer.String("doloremque"),
        Ids: []int64{
            441711,
        },
        TeamID: 282807,
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
        writer.WithOrganizationID(979587),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Limit: writer.Int64(120196),
        Offset: writer.Int64(359444),
        PartOfSpeech: operations.FindTermsPartOfSpeechVerb.ToPointer(),
        SortField: operations.FindTermsSortFieldCreationTime.ToPointer(),
        SortOrder: operations.FindTermsSortOrderAsc.ToPointer(),
        Tags: []string{
            "harum",
        },
        TeamID: 317983,
        Term: writer.String("accusamus"),
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
        writer.WithOrganizationID(918236),
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
                    Description: writer.String("ipsum"),
                    Examples: []shared.TermExampleCreate{
                        shared.TermExampleCreate{
                            Example: "quidem",
                            Type: shared.TermExampleCreateTypeBad,
                        },
                    },
                    Highlight: writer.Bool(false),
                    ID: 566602,
                    LinkedTerms: []shared.LinkedTermCreate{
                        shared.LinkedTermCreate{
                            LinkedTermID: writer.Int64(865103),
                            Reference: writer.String("modi"),
                        },
                    },
                    Mistakes: []shared.TermMistakeCreate{
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "praesentium",
                            Pos: shared.TermMistakeCreatePosAdverb.ToPointer(),
                            Reference: writer.String("voluptates"),
                        },
                    },
                    Pos: shared.TermUpdatePosNoun.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "repudiandae",
                        },
                    },
                    Term: "sint",
                    Type: shared.TermUpdateTypeApproved,
                },
            },
        },
        XRequestID: writer.String("itaque"),
        TeamID: 277718,
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

