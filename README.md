<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/223574357-9a053550-02f9-49f1-b453-1b11db148d7b.svg" media="(prefers-color-scheme: dark)" width="500">
        <img src="https://user-images.githubusercontent.com/6267663/223574369-77805bfe-6d95-44e8-ac48-f9494101e1dc.svg" width="500">
    </picture>
    <h1>Go SDK</h1>
   <p>AI for everyone.</p>
   <a href="https://dev.writer.com/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/writerai/writer-client-sdk-go/releases"><img src="https://img.shields.io/github/v/release/writerai/writer-client-sdk-go?sort=semver&style=for-the-badge" /></a>
</div>


<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/writerai/writer-client-sdk-go
```
<!-- End SDK Installation -->

## Authentication

Writer authenticates your API requests using your account’s API keys. If you do not include your key when making an API request, or use one that is incorrect or outdated, Writer returns an error.

Your API keys are available in the account dashboard. We include randomly generated API keys in our code examples if you are not logged in. Replace these with your own or log in to see code examples populated with your own API keys.

<img width="1070" alt="writer-auth" src="https://user-images.githubusercontent.com/6267663/223578295-89087c24-c55a-48bf-b74a-5f057e21e14f.png">

If you cannot see your secret API keys in the Dashboard, this means you do not have access to them. Contact your Writer account owner and ask to be added to their team as a developer.

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
            APIKey: "",
        }),
        writer.WithOrganizationID(548814),
    )

    ctx := context.Background()
    res, err := s.AIContentDetector.Detect(ctx, operations.DetectContentRequest{
        ContentDetectorRequest: shared.ContentDetectorRequest{
            Input: "provident",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AIContentDetector](docs/aicontentdetector/README.md)

* [Detect](docs/aicontentdetector/README.md#detect) - Content detector api

### [Billing](docs/billing/README.md)

* [GetSubscriptionDetails](docs/billing/README.md#getsubscriptiondetails) - Get your organization subscription details

### [CoWrite](docs/cowrite/README.md)

* [GenerateContent](docs/cowrite/README.md#generatecontent) - Generate content using predefined templates
* [ListTemplates](docs/cowrite/README.md#listtemplates) - Get a list of your existing CoWrite templates

### [Completions](docs/completions/README.md)

* [Create](docs/completions/README.md#create) - Create completion for LLM model
* [CreateModelCustomizationCompletion](docs/completions/README.md#createmodelcustomizationcompletion) - Create completion for LLM customization model

### [Content](docs/content/README.md)

* [Check](docs/content/README.md#check) - Check your content against your preset styleguide.
* [Correct](docs/content/README.md#correct) - Apply the style guide suggestions directly to your content.

### [DownloadTheCustomizedModel](docs/downloadthecustomizedmodel/README.md)

* [FetchFile](docs/downloadthecustomizedmodel/README.md#fetchfile) - Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

### [Files](docs/files/README.md)

* [Delete](docs/files/README.md#delete) - Delete file
* [Get](docs/files/README.md#get) - Get file
* [List](docs/files/README.md#list) - List files
* [Upload](docs/files/README.md#upload) - Upload file

### [ModelCustomization](docs/modelcustomization/README.md)

* [Create](docs/modelcustomization/README.md#create) - Create model customization
* [Delete](docs/modelcustomization/README.md#delete) - Delete Model customization
* [Get](docs/modelcustomization/README.md#get) - Get model customization
* [List](docs/modelcustomization/README.md#list) - List model customizations

### [Models](docs/models/README.md)

* [List](docs/models/README.md#list) - List available LLM models

### [Snippet](docs/snippet/README.md)

* [Delete](docs/snippet/README.md#delete) - Delete snippets
* [Find](docs/snippet/README.md#find) - Find snippets
* [Update](docs/snippet/README.md#update) - Update snippets

### [Styleguide](docs/styleguide/README.md)

* [Get](docs/styleguide/README.md#get) - Page details
* [ListPages](docs/styleguide/README.md#listpages) - List your styleguide pages

### [Terminology](docs/terminology/README.md)

* [Add](docs/terminology/README.md#add) - Add terms
* [Delete](docs/terminology/README.md#delete) - Delete terms
* [Find](docs/terminology/README.md#find) - Find terms
* [Update](docs/terminology/README.md#update) - Update terms

### [User](docs/user/README.md)

* [List](docs/user/README.md#list) - List users
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
