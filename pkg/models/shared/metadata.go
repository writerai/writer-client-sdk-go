package shared

type MetaDataTierEnum string

const (
	MetaDataTierEnumEnterprise1 MetaDataTierEnum = "enterprise-1"
	MetaDataTierEnumEnterprise2 MetaDataTierEnum = "enterprise-2"
	MetaDataTierEnumEnterprise3 MetaDataTierEnum = "enterprise-3"
	MetaDataTierEnumEnterprise4 MetaDataTierEnum = "enterprise-4"
)

type MetaData struct {
	Portal        map[string]string `json:"portal"`
	Reporting     map[string]string `json:"reporting"`
	SnippetsCount int64             `json:"snippetsCount"`
	SsoAccess     bool              `json:"ssoAccess"`
	Styleguide    map[string]string `json:"styleguide"`
	TeamCount     int64             `json:"teamCount"`
	TermsCount    int64             `json:"termsCount"`
	Tier          *MetaDataTierEnum `json:"tier,omitempty"`
}
