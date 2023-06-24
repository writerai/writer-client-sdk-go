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
            APIKey: "",
        }),
        writer.WithOrganizationID(715190),
    )

    ctx := context.Background()
    res, err := s.AIContentDetector.Detect(ctx, operations.DetectContentRequest{
        ContentDetectorRequest: shared.ContentDetectorRequest{
            Input: "quibusdam",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DetectContentRequest](../../models/operations/detectcontentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DetectContentResponse](../../models/operations/detectcontentresponse.md), error**
