package shared

type GenerationModelInfoResponseTypeEnum string

const (
	GenerationModelInfoResponseTypeEnumGpt      GenerationModelInfoResponseTypeEnum = "GPT"
	GenerationModelInfoResponseTypeEnumInstruct GenerationModelInfoResponseTypeEnum = "Instruct"
)

type GenerationModelInfoResponse struct {
	ID   string                              `json:"id"`
	Name string                              `json:"name"`
	Type GenerationModelInfoResponseTypeEnum `json:"type"`
}
