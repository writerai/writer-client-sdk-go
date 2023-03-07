package shared

type Usage struct {
	CoWriteWords UsageItem `json:"coWriteWords"`
	Team         UsageItem `json:"team"`
	User         UsageItem `json:"user"`
	Words        UsageItem `json:"words"`
}
