// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

type FailMessage struct {
	Description string      `json:"description"`
	Extras      interface{} `json:"extras"`
	Key         string      `json:"key"`
}

func (o *FailMessage) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *FailMessage) GetExtras() interface{} {
	if o == nil {
		return nil
	}
	return o.Extras
}

func (o *FailMessage) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}
