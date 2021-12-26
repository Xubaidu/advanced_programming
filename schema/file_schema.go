package schema

import "mime/multipart"

type UploadFileRequest struct {
	File     multipart.File
	FileName string
	UserID   int64
}

type UploadFileResponse struct {
	DstPath string `json:"dst_path"`
}

type DownloadFileRequest struct {
	Filter map[string]interface{}
}

type DownloadFileResponse struct {
	FileName string `json:"file_name"`
}
