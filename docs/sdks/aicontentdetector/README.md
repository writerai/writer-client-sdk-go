# AIContentDetector
(*AIContentDetector*)

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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(496531),
    )


    contentDetectorRequest := shared.ContentDetectorRequest{
        Input: "Bronze Indian",
    }

    var organizationID *int64 = 558689

    ctx := context.Background()
    res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `contentDetectorRequest`                                                       | [shared.ContentDetectorRequest](../../models/shared/contentdetectorrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `organizationID`                                                               | **int64*                                                                       | :heavy_minus_sign:                                                             | N/A                                                                            |


### Response

**[*operations.DetectContentResponse](../../models/operations/detectcontentresponse.md), error**

