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

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(496531),
	)

	contentDetectorRequest := shared.ContentDetectorRequest{
		Input: "string",
	}

	var organizationID *int64 = 592237

	ctx := context.Background()
	res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
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


### [AIContentDetector](docs/sdks/aicontentdetector/README.md)

* [Detect](docs/sdks/aicontentdetector/README.md#detect) - Content detector api

### [Billing](docs/sdks/billing/README.md)

* [GetSubscriptionDetails](docs/sdks/billing/README.md#getsubscriptiondetails) - Get your organization subscription details

### [CoWrite](docs/sdks/cowrite/README.md)

* [GenerateContent](docs/sdks/cowrite/README.md#generatecontent) - Generate content using predefined templates
* [ListTemplates](docs/sdks/cowrite/README.md#listtemplates) - Get a list of your existing CoWrite templates

### [Completions](docs/sdks/completions/README.md)

* [Create](docs/sdks/completions/README.md#create) - Create completion for LLM model
* [CreateModelCustomizationCompletion](docs/sdks/completions/README.md#createmodelcustomizationcompletion) - Create completion for LLM customization model

### [Content](docs/sdks/content/README.md)

* [Check](docs/sdks/content/README.md#check) - Check your content against your preset styleguide.
* [Correct](docs/sdks/content/README.md#correct) - Apply the style guide suggestions directly to your content.

### [DownloadTheCustomizedModel](docs/sdks/downloadthecustomizedmodel/README.md)

* [FetchFile](docs/sdks/downloadthecustomizedmodel/README.md#fetchfile) - Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

### [Files](docs/sdks/files/README.md)

* [Delete](docs/sdks/files/README.md#delete) - Delete file
* [Get](docs/sdks/files/README.md#get) - Get file
* [List](docs/sdks/files/README.md#list) - List files
* [Upload](docs/sdks/files/README.md#upload) - Upload file

### [ModelCustomization](docs/sdks/modelcustomization/README.md)

* [Create](docs/sdks/modelcustomization/README.md#create) - Create model customization
* [Delete](docs/sdks/modelcustomization/README.md#delete) - Delete Model customization
* [Get](docs/sdks/modelcustomization/README.md#get) - Get model customization
* [List](docs/sdks/modelcustomization/README.md#list) - List model customizations

### [Models](docs/sdks/models/README.md)

* [List](docs/sdks/models/README.md#list) - List available LLM models

### [Snippet](docs/sdks/snippet/README.md)

* [Delete](docs/sdks/snippet/README.md#delete) - Delete snippets
* [Find](docs/sdks/snippet/README.md#find) - Find snippets
* [Update](docs/sdks/snippet/README.md#update) - Update snippets

### [Styleguide](docs/sdks/styleguide/README.md)

* [Get](docs/sdks/styleguide/README.md#get) - Page details
* [ListPages](docs/sdks/styleguide/README.md#listpages) - List your styleguide pages

### [Terminology](docs/sdks/terminology/README.md)

* [Add](docs/sdks/terminology/README.md#add) - Add terms
* [Delete](docs/sdks/terminology/README.md#delete) - Delete terms
* [Find](docs/sdks/terminology/README.md#find) - Find terms
* [Update](docs/sdks/terminology/README.md#update) - Update terms

### [User](docs/sdks/user/README.md)

* [List](docs/sdks/user/README.md#list) - List users

### [Document](docs/sdks/document/README.md)

* [Get](docs/sdks/document/README.md#get) - Get document details
* [List](docs/sdks/document/README.md#list) - List team documents
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Global Parameters -->
# Global Parameters

A parameter is configured globally. This parameter must be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `organizationId` to `547272` at SDK initialization and then you do not have to pass the same value on calls to operations like `Detect`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


## Available Globals

The following global parameter is available. The required parameter must be set when you initialize the SDK client.

| Name | Type | Required | Description |
| ---- | ---- |:--------:| ----------- |
| organizationId | int64 | ✔️ | The organizationId parameter. |



## Example

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(496531),
	)

	contentDetectorRequest := shared.ContentDetectorRequest{
		Input: "string",
	}

	var organizationID *int64 = 592237

	ctx := context.Background()
	res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
	if err != nil {
		log.Fatal(err)
	}

	if res.ContentDetectorResponses != nil {
		// handle response
	}
}

```
<!-- End Global Parameters -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
