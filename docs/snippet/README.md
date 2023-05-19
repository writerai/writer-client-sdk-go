# Snippet

## Overview

Methods related to Snippets

### Available Operations

* [Delete](#delete) - Delete snippets
* [Find](#find) - Find snippets
* [Update](#update) - Update snippets

## Delete

Delete snippets

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
        writer.WithOrganizationID(581850),
    )

    ctx := context.Background()
    res, err := s.Snippet.Delete(ctx, operations.DeleteSnippetsRequest{
        XRequestID: writer.String("numquam"),
        Ids: []string{
            "quam",
            "molestiae",
        },
        TeamID: 244425,
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

Find snippets

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
        writer.WithOrganizationID(623510),
    )

    ctx := context.Background()
    res, err := s.Snippet.Find(ctx, operations.FindSnippetsRequest{
        Limit: writer.Int64(158969),
        Offset: writer.Int64(338007),
        Search: writer.String("vitae"),
        Shortcuts: []string{
            "animi",
            "enim",
            "odit",
        },
        SortField: operations.FindSnippetsSortFieldModificationTime.ToPointer(),
        SortOrder: operations.FindSnippetsSortOrderAsc.ToPointer(),
        Tags: []string{
            "ipsam",
            "id",
            "possimus",
            "aut",
        },
        TeamID: 97101,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultSnippetWithUser != nil {
        // handle response
    }
}
```

## Update

Update snippets

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
        writer.WithOrganizationID(622846),
    )

    ctx := context.Background()
    res, err := s.Snippet.Update(ctx, operations.UpdateSnippetsRequest{
        RequestBody: []shared.SnippetUpdate{
            shared.SnippetUpdate{
                Description: writer.String("laborum"),
                ID: "1ffe78f0-97b0-4074-b154-71b5e6e13b99",
                Shortcut: writer.String("pariatur"),
                Snippet: "modi",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "rem",
                    },
                    shared.SnippetTagV2{
                        Tag: "voluptates",
                    },
                    shared.SnippetTagV2{
                        Tag: "quasi",
                    },
                },
            },
            shared.SnippetUpdate{
                Description: writer.String("repudiandae"),
                ID: "91e450ad-2abd-4442-a980-2d502a94bb4f",
                Shortcut: writer.String("eum"),
                Snippet: "non",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "sint",
                    },
                    shared.SnippetTagV2{
                        Tag: "aliquid",
                    },
                    shared.SnippetTagV2{
                        Tag: "provident",
                    },
                    shared.SnippetTagV2{
                        Tag: "necessitatibus",
                    },
                },
            },
            shared.SnippetUpdate{
                Description: writer.String("sint"),
                ID: "a3efa77d-fb14-4cd6-aae3-95efb9ba88f3",
                Shortcut: writer.String("deserunt"),
                Snippet: "nisi",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "natus",
                    },
                    shared.SnippetTagV2{
                        Tag: "omnis",
                    },
                },
            },
            shared.SnippetUpdate{
                Description: writer.String("molestiae"),
                ID: "074ba446-9b6e-4214-9959-890afa563e25",
                Shortcut: writer.String("quasi"),
                Snippet: "iure",
                Tags: []shared.SnippetTagV2{
                    shared.SnippetTagV2{
                        Tag: "debitis",
                    },
                    shared.SnippetTagV2{
                        Tag: "eius",
                    },
                    shared.SnippetTagV2{
                        Tag: "maxime",
                    },
                    shared.SnippetTagV2{
                        Tag: "deleniti",
                    },
                },
            },
        },
        XRequestID: writer.String("facilis"),
        TeamID: 447926,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SnippetWithUsers != nil {
        // handle response
    }
}
```
