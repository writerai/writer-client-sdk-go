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
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(998848),
    )

    ctx := context.Background()
    res, err := s.Terminology.Add(ctx, operations.AddTermsRequest{
        CreateTermsRequest: shared.CreateTermsRequest{
            FailHandling: shared.CreateTermsRequestFailHandlingEnumValidateOnly.ToPointer(),
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
                            Type: shared.TermExampleCreateTypeEnumGood,
                        },
                        shared.TermExampleCreate{
                            Example: "praesentium",
                            Type: shared.TermExampleCreateTypeEnumBad,
                        },
                        shared.TermExampleCreate{
                            Example: "magni",
                            Type: shared.TermExampleCreateTypeEnumGood,
                        },
                        shared.TermExampleCreate{
                            Example: "quo",
                            Type: shared.TermExampleCreateTypeEnumBad,
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
                            Pos: shared.TermMistakeCreatePosEnumVerb.ToPointer(),
                            Reference: writer.String("autem"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "nam",
                            Pos: shared.TermMistakeCreatePosEnumNoun.ToPointer(),
                            Reference: writer.String("pariatur"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "nemo",
                            Pos: shared.TermMistakeCreatePosEnumAdjective.ToPointer(),
                            Reference: writer.String("perferendis"),
                        },
                    },
                    Pos: shared.TermCreatePosEnumAdjective.ToPointer(),
                    Reference: writer.String("amet"),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "cumque",
                        },
                    },
                    Term: "corporis",
                    Type: shared.TermCreateTypeEnumPending,
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

## Delete

Delete terms

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
            APIKey: "YOUR_API_KEY_HERE",
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

## Find

Find terms

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
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(338985),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        Limit: writer.Int64(199996),
        Offset: writer.Int64(179490),
        PartOfSpeech: operations.FindTermsPartOfSpeechEnumNoun.ToPointer(),
        SortField: operations.FindTermsSortFieldEnumTerm.ToPointer(),
        SortOrder: operations.FindTermsSortOrderEnumDesc.ToPointer(),
        Tags: []string{
            "dolor",
            "vero",
        },
        TeamID: 345352,
        Term: writer.String("hic"),
        Type: operations.FindTermsTypeEnumPending.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultFullTermWithUser != nil {
        // handle response
    }
}
```

## Update

Update terms

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(608253),
    )

    ctx := context.Background()
    res, err := s.Terminology.Update(ctx, operations.UpdateTermsRequest{
        UpdateTermsRequest: shared.UpdateTermsRequest{
            FailHandling: shared.UpdateTermsRequestFailHandlingEnumSkip.ToPointer(),
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
                            Type: shared.TermExampleCreateTypeEnumBad,
                        },
                        shared.TermExampleCreate{
                            Example: "error",
                            Type: shared.TermExampleCreateTypeEnumGood,
                        },
                        shared.TermExampleCreate{
                            Example: "occaecati",
                            Type: shared.TermExampleCreateTypeEnumBad,
                        },
                        shared.TermExampleCreate{
                            Example: "adipisci",
                            Type: shared.TermExampleCreateTypeEnumBad,
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
                            Pos: shared.TermMistakeCreatePosEnumAdverb.ToPointer(),
                            Reference: writer.String("delectus"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "quaerat",
                            Pos: shared.TermMistakeCreatePosEnumAdverb.ToPointer(),
                            Reference: writer.String("aliquid"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "dolorem",
                            Pos: shared.TermMistakeCreatePosEnumNoun.ToPointer(),
                            Reference: writer.String("dolor"),
                        },
                    },
                    Pos: shared.TermUpdatePosEnumNoun.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "hic",
                        },
                    },
                    Term: "excepturi",
                    Type: shared.TermUpdateTypeEnumPending,
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
                            Type: shared.TermExampleCreateTypeEnumGood,
                        },
                        shared.TermExampleCreate{
                            Example: "dolorum",
                            Type: shared.TermExampleCreateTypeEnumGood,
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
                            Pos: shared.TermMistakeCreatePosEnumAdjective.ToPointer(),
                            Reference: writer.String("quidem"),
                        },
                        shared.TermMistakeCreate{
                            CaseSensitive: false,
                            Mistake: "voluptatibus",
                            Pos: shared.TermMistakeCreatePosEnumVerb.ToPointer(),
                            Reference: writer.String("natus"),
                        },
                    },
                    Pos: shared.TermUpdatePosEnumNoun.ToPointer(),
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
                    Type: shared.TermUpdateTypeEnumPending,
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
                            Type: shared.TermExampleCreateTypeEnumBad,
                        },
                        shared.TermExampleCreate{
                            Example: "omnis",
                            Type: shared.TermExampleCreateTypeEnumBad,
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
                            Pos: shared.TermMistakeCreatePosEnumAdjective.ToPointer(),
                            Reference: writer.String("ad"),
                        },
                    },
                    Pos: shared.TermUpdatePosEnumAdjective.ToPointer(),
                    Tags: []shared.TermTagCreate{
                        shared.TermTagCreate{
                            Tag: "deserunt",
                        },
                        shared.TermTagCreate{
                            Tag: "provident",
                        },
                    },
                    Term: "minima",
                    Type: shared.TermUpdateTypeEnumPending,
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
