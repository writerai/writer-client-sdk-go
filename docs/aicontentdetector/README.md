# AIContentDetector

## Overview

Methods related to AI Content Detector

### Available Operations

* [Detect](#detect) - Content detector api

## Detect

Content detector api

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
        writer.WithOrganizationID(715190),
    )

    ctx := context.Background()    
    req := operations.DetectContentRequest{
        ContentDetectorRequest: shared.ContentDetectorRequest{
            Input: "quibusdam",
        },
    }

    res, err := s.AIContentDetector.Detect(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```
