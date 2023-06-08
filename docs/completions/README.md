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
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
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
            BestOf: writer.Int64(978619),
            FrequencyPenalty: writer.Float64(4736.08),
            Logprobs: writer.Int64(799159),
            MaxTokens: writer.Int64(800911),
            MinTokens: writer.Int64(461479),
            N: writer.Int64(520478),
            PresencePenalty: writer.Float64(7805.29),
            Prompt: "dolorum",
            Stop: []string{
                "nam",
            },
            Temperature: writer.Float64(6399.21),
            TopP: writer.Float64(5820.2),
        },
        ModelID: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompletionResponse != nil {
        // handle response
    }
}
```

## CreateModelCustomizationCompletion

Create completion for LLM customization model

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
            APIKey: "",
        }),
        writer.WithOrganizationID(537373),
    )

    ctx := context.Background()
    res, err := s.Completions.CreateModelCustomizationCompletion(ctx, operations.CreateModelCustomizationCompletionRequest{
        CompletionRequest: shared.CompletionRequest{
            BestOf: writer.Int64(944669),
            FrequencyPenalty: writer.Float64(7586.16),
            Logprobs: writer.Int64(521848),
            MaxTokens: writer.Int64(105907),
            MinTokens: writer.Int64(414662),
            N: writer.Int64(473600),
            PresencePenalty: writer.Float64(2645.55),
            Prompt: "qui",
            Stop: []string{
                "cum",
                "esse",
                "ipsum",
                "excepturi",
            },
            Temperature: writer.Float64(1352.18),
            TopP: writer.Float64(187.89),
        },
        CustomizationID: "ad",
        ModelID: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompletionResponse != nil {
        // handle response
    }
}
```
