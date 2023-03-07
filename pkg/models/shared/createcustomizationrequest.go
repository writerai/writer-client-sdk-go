package shared

type CreateCustomizationRequest struct {
	AdditionalHyperParameters *HyperParameters `json:"additionalHyperParameters,omitempty"`
	BatchSize                 *int64           `json:"batchSize,omitempty"`
	Description               *string          `json:"description,omitempty"`
	Epochs                    *int64           `json:"epochs,omitempty"`
	LearningRate              *float64         `json:"learningRate,omitempty"`
	Name                      string           `json:"name"`
	PromptTemplate            *string          `json:"promptTemplate,omitempty"`
	TrainingDatasetFileID     string           `json:"trainingDatasetFileId"`
	ValidationDatasetFileID   *string          `json:"validationDatasetFileId,omitempty"`
}
