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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(359508),
    )

    ctx := context.Background()
    res, err := s.ModelCustomization.Create(ctx, operations.CreateModelCustomizationRequest{
        CreateCustomizationRequest: shared.CreateCustomizationRequest{
            AdditionalHyperParameters: &shared.HyperParameters{
                NumVirtualTokens: 613064,
            },
            BatchSize: writer.Int64(437032),
            Description: writer.String("saepe"),
            Epochs: writer.Int64(697631),
            LearningRate: writer.Float64(992.8),
            Name: "Lela Orn",
            PromptTemplate: writer.String("dolores"),
            TrainingDatasetFileID: "dolorem",
            ValidationDatasetFileID: writer.String("corporis"),
        },
        ModelID: "explicabo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateModelCustomizationRequest](../../models/operations/createmodelcustomizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.CreateModelCustomizationResponse](../../models/operations/createmodelcustomizationresponse.md), error**


## Delete

Delete Model customization

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(750686),
    )

    ctx := context.Background()
    res, err := s.ModelCustomization.Delete(ctx, operations.DeleteModelCustomizationRequest{
        CustomizationID: "enim",
        ModelID: "omnis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteModelCustomization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteModelCustomizationRequest](../../models/operations/deletemodelcustomizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.DeleteModelCustomizationResponse](../../models/operations/deletemodelcustomizationresponse.md), error**


## Get

Get model customization

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(363711),
    )

    ctx := context.Background()
    res, err := s.ModelCustomization.Get(ctx, operations.GetModelCustomizationRequest{
        CustomizationID: "minima",
        ModelID: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetModelCustomizationRequest](../../models/operations/getmodelcustomizationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetModelCustomizationResponse](../../models/operations/getmodelcustomizationresponse.md), error**


## List

List model customizations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(38425),
    )

    ctx := context.Background()
    res, err := s.ModelCustomization.List(ctx, operations.ListModelCustomizationsRequest{
        ModelID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomizationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListModelCustomizationsRequest](../../models/operations/listmodelcustomizationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListModelCustomizationsResponse](../../models/operations/listmodelcustomizationsresponse.md), error**

