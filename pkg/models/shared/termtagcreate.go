// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TermTagCreate struct {
	Tag string `json:"tag"`
}

func (o *TermTagCreate) GetTag() string {
	if o == nil {
		return ""
	}
	return o.Tag
}
