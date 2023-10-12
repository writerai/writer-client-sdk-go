# DownloadTheCustomizedModel
(*DownloadTheCustomizedModel*)

## Overview

Methods related to Download the customized model

### Available Operations

* [FetchFile](#fetchfile) - Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

## FetchFile

Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

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
        writerclientsdkgo.WithOrganizationID(501762),
    )


    var customizationID string = "apology"

    var modelID string = "Silver"

    var organizationID *int64 = 432823

    ctx := context.Background()
    res, err := s.DownloadTheCustomizedModel.FetchFile(ctx, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchCustomizedModelFile200ApplicationOctetStreamBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `customizationID`                                     | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `modelID`                                             | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.FetchCustomizedModelFileResponse](../../models/operations/fetchcustomizedmodelfileresponse.md), error**

