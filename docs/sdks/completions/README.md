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
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(486589),
    )

    ctx := context.Background()
    res, err := s.Completions.Create(ctx, operations.CreateCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writerclientsdkgo.Int64(1),
            MaxTokens: writerclientsdkgo.Int64(1024),
            MinTokens: writerclientsdkgo.Int64(1),
            Prompt: "Configuration Money",
            Stop: []string{
                "the",
                "is",
                "and",
            },
            Temperature: writerclientsdkgo.Float64(0.7),
            TopP: writerclientsdkgo.Float64(1),
        },
        ModelID: "Cambridgeshire grey technology",
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
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(919503),
    )

    ctx := context.Background()
    res, err := s.Completions.CreateModelCustomizationCompletion(ctx, operations.CreateModelCustomizationCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writerclientsdkgo.Int64(1),
            MaxTokens: writerclientsdkgo.Int64(1024),
            MinTokens: writerclientsdkgo.Int64(1),
            Prompt: "error",
            Stop: []string{
                "the",
                "is",
                "and",
            },
            Temperature: writerclientsdkgo.Float64(0.7),
            TopP: writerclientsdkgo.Float64(1),
        },
        CustomizationID: "newton",
        ModelID: "convergence",
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

