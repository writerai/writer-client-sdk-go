# Billing

## Overview

Methods related to Billing

### Available Operations

* [GetSubscriptionDetails](#getsubscriptiondetails) - Get your organization subscription details

## GetSubscriptionDetails

Get your organization subscription details

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(602763),
    )

    ctx := context.Background()
    res, err := s.Billing.GetSubscriptionDetails(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SubscriptionPublicResponseAPI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSubscriptionDetailsResponse](../../models/operations/getsubscriptiondetailsresponse.md), error**

