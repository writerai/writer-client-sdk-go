# Completions
(*Completions*)

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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(486589),
    )

    ctx := context.Background()
    res, err := s.Completions.Create(ctx, operations.CreateCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writerclientsdkgo.Int64(1),
            FrequencyPenalty: writerclientsdkgo.Float64(4893.82),
            Logprobs: writerclientsdkgo.Int64(638424),
            MaxTokens: writerclientsdkgo.Int64(1024),
            MinTokens: writerclientsdkgo.Int64(1),
            N: writerclientsdkgo.Int64(859213),
            PresencePenalty: writerclientsdkgo.Float64(4174.58),
            Prompt: "South",
            Stop: []string{
                "shred",
            },
            Temperature: writerclientsdkgo.Float64(0.7),
            TopP: writerclientsdkgo.Float64(1),
        },
        ModelID: "technology East",
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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writerclientsdkgo.WithOrganizationID(919503),
    )

    ctx := context.Background()
    res, err := s.Completions.CreateModelCustomizationCompletion(ctx, operations.CreateModelCustomizationCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writerclientsdkgo.Int64(1),
            FrequencyPenalty: writerclientsdkgo.Float64(412.97),
            Logprobs: writerclientsdkgo.Int64(951257),
            MaxTokens: writerclientsdkgo.Int64(1024),
            MinTokens: writerclientsdkgo.Int64(1),
            N: writerclientsdkgo.Int64(314952),
            PresencePenalty: writerclientsdkgo.Float64(657.2),
            Prompt: "platforms convergence Bicycle",
            Stop: []string{
                "heavily",
            },
            Temperature: writerclientsdkgo.Float64(0.7),
            TopP: writerclientsdkgo.Float64(1),
        },
        CustomizationID: "Beauty cum Southeast",
        ModelID: "Tandem BMX Lead",
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

