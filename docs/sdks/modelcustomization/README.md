# ModelCustomization
(*ModelCustomization*)

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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(486589),
    )


    createCustomizationRequest := shared.CreateCustomizationRequest{
        AdditionalHyperParameters: &shared.HyperParameters{
            NumVirtualTokens: 489382,
        },
        Name: "Extended South",
        TrainingDatasetFileID: "grey technology East",
    }

    var modelID string = "orange"

    var organizationID *int64 = 89964

    ctx := context.Background()
    res, err := s.ModelCustomization.Create(ctx, createCustomizationRequest, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `createCustomizationRequest`                                                           | [shared.CreateCustomizationRequest](../../models/shared/createcustomizationrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `modelID`                                                                              | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `organizationID`                                                                       | **int64*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |


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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(545907),
    )


    var customizationID string = "Van"

    var modelID string = "complexity"

    var organizationID *int64 = 952479

    ctx := context.Background()
    res, err := s.ModelCustomization.Delete(ctx, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteModelCustomization200ApplicationJSONObject != nil {
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

**[*operations.DeleteModelCustomizationResponse](../../models/operations/deletemodelcustomizationresponse.md), error**


## Get

Get model customization

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
        writerclientsdkgo.WithOrganizationID(700347),
    )


    var customizationID string = "Northeast"

    var modelID string = "Hatchback"

    var organizationID *int64 = 830636

    ctx := context.Background()
    res, err := s.ModelCustomization.Get(ctx, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelCustomization != nil {
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

**[*operations.GetModelCustomizationResponse](../../models/operations/getmodelcustomizationresponse.md), error**


## List

List model customizations

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
        writerclientsdkgo.WithOrganizationID(768578),
    )


    var modelID string = "Northeast"

    var organizationID *int64 = 257649

    ctx := context.Background()
    res, err := s.ModelCustomization.List(ctx, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomizationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `modelID`                                             | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.ListModelCustomizationsResponse](../../models/operations/listmodelcustomizationsresponse.md), error**

