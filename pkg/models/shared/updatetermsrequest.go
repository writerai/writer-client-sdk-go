package shared

type UpdateTermsRequestFailHandlingEnum string

const (
	UpdateTermsRequestFailHandlingEnumAccumulate   UpdateTermsRequestFailHandlingEnum = "accumulate"
	UpdateTermsRequestFailHandlingEnumValidate     UpdateTermsRequestFailHandlingEnum = "validate"
	UpdateTermsRequestFailHandlingEnumSkip         UpdateTermsRequestFailHandlingEnum = "skip"
	UpdateTermsRequestFailHandlingEnumValidateOnly UpdateTermsRequestFailHandlingEnum = "validateOnly"
)

type UpdateTermsRequest struct {
	FailHandling *UpdateTermsRequestFailHandlingEnum `json:"failHandling,omitempty"`
	Models       []TermUpdate                        `json:"models,omitempty"`
}
