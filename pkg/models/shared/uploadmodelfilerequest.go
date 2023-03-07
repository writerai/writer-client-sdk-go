package shared

type UploadModelFileRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type UploadModelFileRequest struct {
	File UploadModelFileRequestFile `multipartForm:"file"`
}
