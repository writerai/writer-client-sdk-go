# CoWrite

## Overview

Methods related to CoWrite

### Available Operations

* [GenerateContent](#generatecontent) - Generate content using predefined templates
* [ListTemplates](#listtemplates) - Get a list of your existing CoWrite templates

## GenerateContent

Generate content using predefined templates

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
        writer.WithOrganizationID(857946),
    )

    ctx := context.Background()
    res, err := s.CoWrite.GenerateContent(ctx, operations.GenerateContentRequest{
        GenerateTemplateRequest: shared.GenerateTemplateRequest{
            Inputs: []shared.MagicRequestInput{
                shared.MagicRequestInput{
                    Name: "Dallas Kassulke",
                    Value: []string{
                        "suscipit",
                    },
                },
            },
            TemplateID: "iure",
        },
        TeamID: 297534,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Draft != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GenerateContentRequest](../../models/operations/generatecontentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GenerateContentResponse](../../models/operations/generatecontentresponse.md), error**


## ListTemplates

Get a list of your existing CoWrite templates

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
        writer.WithOrganizationID(891773),
    )

    ctx := context.Background()
    res, err := s.CoWrite.ListTemplates(ctx, operations.ListTemplatesRequest{
        TeamID: 56713,
        TemplateID: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TemplateDetailsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListTemplatesRequest](../../models/operations/listtemplatesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListTemplatesResponse](../../models/operations/listtemplatesresponse.md), error**

