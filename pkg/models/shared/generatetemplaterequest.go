package shared

type GenerateTemplateRequest struct {
	Inputs     []MagicRequestInput `json:"inputs,omitempty"`
	TemplateID string              `json:"templateId"`
}
