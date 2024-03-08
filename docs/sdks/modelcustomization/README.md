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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(486589),
    )


    createCustomizationRequest := shared.CreateCustomizationRequest{
        Name: "<value>",
        TrainingDatasetFileID: "<value>",
    }

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(489382)

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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `createCustomizationRequest`                                                               | [shared.CreateCustomizationRequest](../../pkg/models/shared/createcustomizationrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `modelID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `organizationID`                                                                           | **int64*                                                                                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |


### Response

**[*operations.CreateModelCustomizationResponse](../../pkg/models/operations/createmodelcustomizationresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Delete

Delete Model customization

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
        writerclientsdkgo.WithOrganizationID(545907),
    )


    var customizationID string = "<value>"

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(841399)

    ctx := context.Background()
    res, err := s.ModelCustomization.Delete(ctx, customizationID, modelID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.DeleteModelCustomizationResponse](../../pkg/models/operations/deletemodelcustomizationresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Get

Get model customization

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
        writerclientsdkgo.WithOrganizationID(700347),
    )


    var customizationID string = "<value>"

    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(90065)

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

**[*operations.GetModelCustomizationResponse](../../pkg/models/operations/getmodelcustomizationresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## List

List model customizations

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
        writerclientsdkgo.WithOrganizationID(768578),
    )


    var modelID string = "<value>"

    var organizationID *int64 = writerclientsdkgo.Int64(99895)

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

**[*operations.ListModelCustomizationsResponse](../../pkg/models/operations/listmodelcustomizationsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
