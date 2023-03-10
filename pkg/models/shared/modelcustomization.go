package shared

import (
	"time"
)

type ModelCustomization struct {
	AdditionalHyperParameters *HyperParameters `json:"additionalHyperParameters,omitempty"`
	BaseModelID               string           `json:"baseModelId"`
	BatchSize                 *int64           `json:"batchSize,omitempty"`
	CreatedAt                 time.Time        `json:"createdAt"`
	Description               *string          `json:"description,omitempty"`
	Epochs                    *int64           `json:"epochs,omitempty"`
	ID                        string           `json:"id"`
	LearningRate              *float64         `json:"learningRate,omitempty"`
	Name                      string           `json:"name"`
	PromptTemplate            *string          `json:"promptTemplate,omitempty"`
	Status                    string           `json:"status"`
	TrainingDatasetFileID     string           `json:"trainingDatasetFileId"`
	UpdatedAt                 time.Time        `json:"updatedAt"`
	ValidationDatasetFileID   *string          `json:"validationDatasetFileId,omitempty"`
}
