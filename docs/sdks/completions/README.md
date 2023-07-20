# Completions

## Overview

Methods related to Completions

### Available Operations

* [Create](#create) - Create completion for LLM model
* [CreateModelCustomizationCompletion](#createmodelcustomizationcompletion) - Create completion for LLM customization model

## Create

Create completion for LLM model

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
        writer.WithOrganizationID(870088),
    )

    ctx := context.Background()
    res, err := s.Completions.Create(ctx, operations.CreateCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writer.Int64(1),
            FrequencyPenalty: writer.Float64(9786.19),
            Logprobs: writer.Int64(473608),
            MaxTokens: writer.Int64(1024),
            MinTokens: writer.Int64(1),
            N: writer.Int64(799159),
            PresencePenalty: writer.Float64(8009.11),
            Prompt: "esse",
            Stop: []string{
                "porro",
                "dolorum",
                "dicta",
            },
            Temperature: writer.Float64(0.7),
            TopP: writer.Float64(1),
        },
        ModelID: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCompletionRequest](../../models/operations/createcompletionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateCompletionResponse](../../models/operations/createcompletionresponse.md), error**


## CreateModelCustomizationCompletion

Create completion for LLM customization model

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
        writer.WithOrganizationID(639921),
    )

    ctx := context.Background()
    res, err := s.Completions.CreateModelCustomizationCompletion(ctx, operations.CreateModelCustomizationCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writer.Int64(1),
            FrequencyPenalty: writer.Float64(5820.2),
            Logprobs: writer.Int64(143353),
            MaxTokens: writer.Int64(1024),
            MinTokens: writer.Int64(1),
            N: writer.Int64(537373),
            PresencePenalty: writer.Float64(9446.69),
            Prompt: "optio",
            Stop: []string{
                "beatae",
                "commodi",
                "molestiae",
            },
            Temperature: writer.Float64(0.7),
            TopP: writer.Float64(1),
        },
        CustomizationID: "modi",
        ModelID: "qui",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.CreateModelCustomizationCompletionRequest](../../models/operations/createmodelcustomizationcompletionrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.CreateModelCustomizationCompletionResponse](../../models/operations/createmodelcustomizationcompletionresponse.md), error**

