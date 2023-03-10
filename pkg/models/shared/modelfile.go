package shared

import (
	"time"
)

type ModelFile struct {
	CreatedAt       time.Time `json:"createdAt"`
	Format          string    `json:"format"`
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	NumberOfSamples int64     `json:"numberOfSamples"`
	Size            int64     `json:"size"`
}
