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
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
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
