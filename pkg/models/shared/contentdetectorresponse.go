package shared

type ContentDetectorResponseLabelEnum string

const (
	ContentDetectorResponseLabelEnumFake ContentDetectorResponseLabelEnum = "fake"
	ContentDetectorResponseLabelEnumReal ContentDetectorResponseLabelEnum = "real"
)

type ContentDetectorResponse struct {
	Label ContentDetectorResponseLabelEnum `json:"label"`
	Score float64                          `json:"score"`
}
