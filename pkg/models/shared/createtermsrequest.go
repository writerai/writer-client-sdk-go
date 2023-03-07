package shared

type CreateTermsRequestFailHandlingEnum string

const (
	CreateTermsRequestFailHandlingEnumAccumulate   CreateTermsRequestFailHandlingEnum = "accumulate"
	CreateTermsRequestFailHandlingEnumValidate     CreateTermsRequestFailHandlingEnum = "validate"
	CreateTermsRequestFailHandlingEnumSkip         CreateTermsRequestFailHandlingEnum = "skip"
	CreateTermsRequestFailHandlingEnumValidateOnly CreateTermsRequestFailHandlingEnum = "validateOnly"
)

type CreateTermsRequest struct {
	FailHandling *CreateTermsRequestFailHandlingEnum `json:"failHandling,omitempty"`
	Models       []TermCreate                        `json:"models,omitempty"`
}
