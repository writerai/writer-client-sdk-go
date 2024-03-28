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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(496531),
    )


    contentDetectorRequest := shared.ContentDetectorRequest{
        Input: "<value>",
    }

    var organizationID *int64 = writerclientsdkgo.Int64(592237)

    ctx := context.Background()
    res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `contentDetectorRequest`                                                           | [shared.ContentDetectorRequest](../../pkg/models/shared/contentdetectorrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `organizationID`                                                                   | **int64*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |


### Response

**[*operations.DetectContentResponse](../../pkg/models/operations/detectcontentresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
