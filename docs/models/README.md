# Models

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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(653108),
    )

    ctx := context.Background()    
    req := operations.ListModelsRequest{}

    res, err := s.Models.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerationModelsResponse != nil {
        // handle response
    }
}
```
