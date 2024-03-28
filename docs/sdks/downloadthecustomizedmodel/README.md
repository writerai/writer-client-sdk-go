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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(501762),
    )


    var customizationID string = "<value>"

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(948692)

    ctx := context.Background()
    res, err := s.DownloadTheCustomizedModel.FetchFile(ctx, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
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

**[*operations.FetchCustomizedModelFileResponse](../../pkg/models/operations/fetchcustomizedmodelfileresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
