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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(486589),
    )


    completionRequest := shared.CompletionRequest{
        BestOf: writerclientsdkgo.Int64(1),
        MaxTokens: writerclientsdkgo.Int64(1024),
        MinTokens: writerclientsdkgo.Int64(1),
        Prompt: "<value>",
        Stop: []string{
            "the",
            "is",
            "and",
        },
        Temperature: writerclientsdkgo.Float64(0.7),
        TopP: writerclientsdkgo.Float64(1),
    }

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(489382)

    ctx := context.Background()
    res, err := s.Completions.Create(ctx, completionRequest, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `completionRequest`                                                      | [shared.CompletionRequest](../../pkg/models/shared/completionrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `modelID`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `organizationID`                                                         | **int64*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |


### Response

**[*operations.CreateCompletionResponse](../../pkg/models/operations/createcompletionresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## CreateModelCustomizationCompletion

Create completion for LLM customization model

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(919503),
    )


    completionRequest := shared.CompletionRequest{
        BestOf: writerclientsdkgo.Int64(1),
        MaxTokens: writerclientsdkgo.Int64(1024),
        MinTokens: writerclientsdkgo.Int64(1),
        Prompt: "<value>",
        Stop: []string{
            "the",
            "is",
            "and",
        },
        Temperature: writerclientsdkgo.Float64(0.7),
        TopP: writerclientsdkgo.Float64(1),
    }

    var customizationID string = "<value>"

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(41297)

    ctx := context.Background()
    res, err := s.Completions.CreateModelCustomizationCompletion(ctx, completionRequest, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `completionRequest`                                                      | [shared.CompletionRequest](../../pkg/models/shared/completionrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `customizationID`                                                        | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `modelID`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `organizationID`                                                         | **int64*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |


### Response

**[*operations.CreateModelCustomizationCompletionResponse](../../pkg/models/operations/createmodelcustomizationcompletionresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
