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
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(935464),
    )


    contentRequest := shared.ContentRequest{
        Content: "Tustin",
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

    var teamID int64 = 488169

    var organizationID *int64 = 740907

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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `contentRequest`                                               | [shared.ContentRequest](../../models/shared/contentrequest.md) | :heavy_check_mark:                                             | N/A                                                            |
| `teamID`                                                       | *int64*                                                        | :heavy_check_mark:                                             | N/A                                                            |
| `organizationID`                                               | **int64*                                                       | :heavy_minus_sign:                                             | N/A                                                            |


### Response

**[*operations.ContentCheckResponse](../../models/operations/contentcheckresponse.md), error**


## Correct

Apply the style guide suggestions directly to your content.

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
        writerclientsdkgo.WithOrganizationID(501355),
    )


    contentRequest := shared.ContentRequest{
        Content: "Hattiesburg",
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

    var teamID int64 = 754764

    var xRequestID *string = "Folk"

    var organizationID *int64 = 874845

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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `contentRequest`                                               | [shared.ContentRequest](../../models/shared/contentrequest.md) | :heavy_check_mark:                                             | N/A                                                            |
| `teamID`                                                       | *int64*                                                        | :heavy_check_mark:                                             | N/A                                                            |
| `xRequestID`                                                   | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `organizationID`                                               | **int64*                                                       | :heavy_minus_sign:                                             | N/A                                                            |


### Response

**[*operations.ContentCorrectResponse](../../models/operations/contentcorrectresponse.md), error**

