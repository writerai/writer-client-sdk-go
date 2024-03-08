# Content
(*Content*)

## Overview

Methods related to Content

### Available Operations

* [Check](#check) - Check your content against your preset styleguide.
* [Correct](#correct) - Apply the style guide suggestions directly to your content.

## Check

Check your content against your preset styleguide.

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
        writerclientsdkgo.WithOrganizationID(935464),
    )


    contentRequest := shared.ContentRequest{
        Content: "<value>",
        Settings: shared.ContentSettings{
            AgeAndFamilyStatus: false,
            Confidence: false,
            ContentSafeguards: false,
            Disability: false,
            GenderIdentitySensitivity: false,
            GenderInclusiveNouns: false,
            GenderInclusivePronouns: false,
            Grammar: false,
            HealthyCommunication: false,
            PassiveVoice: false,
            RaceEthnicityNationalitySensitivity: false,
            SexualOrientationSensitivity: false,
            Spelling: false,
            SubstanceUseSensitivity: false,
            UnclearReference: false,
            Wordiness: false,
        },
    }

    var teamID int64 = 38270

    var organizationID *int64 = writerclientsdkgo.Int64(919579)

    ctx := context.Background()
    res, err := s.Content.Check(ctx, contentRequest, teamID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcessedContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `contentRequest`                                                   | [shared.ContentRequest](../../pkg/models/shared/contentrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `teamID`                                                           | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `organizationID`                                                   | **int64*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |


### Response

**[*operations.ContentCheckResponse](../../pkg/models/operations/contentcheckresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Correct

Apply the style guide suggestions directly to your content.

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
        writerclientsdkgo.WithOrganizationID(501355),
    )


    contentRequest := shared.ContentRequest{
        Content: "<value>",
        Settings: shared.ContentSettings{
            AgeAndFamilyStatus: false,
            Confidence: false,
            ContentSafeguards: false,
            Disability: false,
            GenderIdentitySensitivity: false,
            GenderInclusiveNouns: false,
            GenderInclusivePronouns: false,
            Grammar: false,
            HealthyCommunication: false,
            PassiveVoice: false,
            RaceEthnicityNationalitySensitivity: false,
            SexualOrientationSensitivity: false,
            Spelling: false,
            SubstanceUseSensitivity: false,
            UnclearReference: false,
            Wordiness: false,
        },
    }

    var teamID int64 = 31310

    var xRequestID *string = writerclientsdkgo.String("<value>")

    var organizationID *int64 = writerclientsdkgo.Int64(383223)

    ctx := context.Background()
    res, err := s.Content.Correct(ctx, contentRequest, teamID, xRequestID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CorrectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `contentRequest`                                                   | [shared.ContentRequest](../../pkg/models/shared/contentrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `teamID`                                                           | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                |
| `xRequestID`                                                       | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `organizationID`                                                   | **int64*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |


### Response

**[*operations.ContentCorrectResponse](../../pkg/models/operations/contentcorrectresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |
