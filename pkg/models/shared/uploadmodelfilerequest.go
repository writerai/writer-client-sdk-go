// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UploadModelFileRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

func (o *UploadModelFileRequestFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadModelFileRequestFile) GetFile() string {
	if o == nil {
		return ""
	}
	return o.File
}

type UploadModelFileRequest struct {
	File UploadModelFileRequestFile `multipartForm:"file"`
}

func (o *UploadModelFileRequest) GetFile() UploadModelFileRequestFile {
	if o == nil {
		return UploadModelFileRequestFile{}
	}
	return o.File
}
