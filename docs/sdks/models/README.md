# Models
(*Models*)

## Overview

Methods related to Model

### Available Operations

* [List](#list) - List available LLM models

## List

List available LLM models

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(768578),
    )


    var organizationID *int64 = 99895

    ctx := context.Background()
    res, err := s.Models.List(ctx, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerationModelsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.ListModelsResponse](../../models/operations/listmodelsresponse.md), error**

