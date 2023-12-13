// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type File struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

type UploadModelFileRequest struct {
	File File `multipartForm:"file"`
}

func (o *UploadModelFileRequest) GetFile() File {
	if o == nil {
		return File{}
	}
	return o.File
}
