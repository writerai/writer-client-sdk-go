package shared

type InputTypeEnum string

const (
	InputTypeEnumTextbox  InputTypeEnum = "textbox"
	InputTypeEnumTextarea InputTypeEnum = "textarea"
	InputTypeEnumDropdown InputTypeEnum = "dropdown"
)

type Input struct {
	Dynamic   bool          `json:"dynamic"`
	Help      *string       `json:"help,omitempty"`
	MaxFields *int64        `json:"maxFields,omitempty"`
	MinFields *int64        `json:"minFields,omitempty"`
	Name      string        `json:"name"`
	Options   []string      `json:"options,omitempty"`
	Required  bool          `json:"required"`
	Subtitle  *string       `json:"subtitle,omitempty"`
	Type      InputTypeEnum `json:"type"`
	UnitCopy  *string       `json:"unitCopy,omitempty"`
}
