// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContentRequest struct {
	Content  string          `json:"content"`
	Settings ContentSettings `json:"settings"`
}

func (o *ContentRequest) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *ContentRequest) GetSettings() ContentSettings {
	if o == nil {
		return ContentSettings{}
	}
	return o.Settings
}
