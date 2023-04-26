# ModelCustomization

## Overview

Methods related to Model Customization

### Available Operations

* [Create](#create) - Create model customization
* [Delete](#delete) - Delete Model customization
* [Get](#get) - Get model customization
* [List](#list) - List model customizations

## Create

Create model customization

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
        writer.WithOrganizationID(653140),
    )

    ctx := context.Background()    
    req := operations.CreateModelCustomizationRequest{
        CreateCustomizationRequest: shared.CreateCustomizationRequest{
            AdditionalHyperParameters: &shared.HyperParameters{
                NumVirtualTokens: 670638,
            },
            BatchSize: writer.Int64(170909),
            Description: writer.String("dolorem"),
            Epochs: writer.Int64(358152),
            LearningRate: writer.Float64(1289.26),
            Name: "Ronnie Mohr",
            PromptTemplate: writer.String("excepturi"),
            TrainingDatasetFileID: "accusantium",
            ValidationDatasetFileID: writer.String("iure"),
        },
        ModelID: "culpa",
    }

    res, err := s.ModelCustomization.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
        // handle response
    }
}
```

## Delete

Delete Model customization

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
        writer.WithOrganizationID(988374),
    )

    ctx := context.Background()    
    req := operations.DeleteModelCustomizationRequest{
        CustomizationID: "sapiente",
        ModelID: "architecto",
    }

    res, err := s.ModelCustomization.Delete(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteModelCustomization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Get

Get model customization

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
        writer.WithOrganizationID(652790),
    )

    ctx := context.Background()    
    req := operations.GetModelCustomizationRequest{
        CustomizationID: "dolorem",
        ModelID: "culpa",
    }

    res, err := s.ModelCustomization.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
        // handle response
    }
}
```

## List

List model customizations

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
        writer.WithOrganizationID(161309),
    )

    ctx := context.Background()    
    req := operations.ListModelCustomizationsRequest{
        ModelID: "repellat",
    }

    res, err := s.ModelCustomization.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomizationsResponse != nil {
        // handle response
    }
}
```
