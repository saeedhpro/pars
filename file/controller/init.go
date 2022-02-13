package controller

import "file/logic"

var (
	UploadFile logic.UploadFileLogic
	GetFile    logic.GetFileLogic
	GetFiles   logic.GetFilesLogic
)

func init() {
	UploadFile = logic.NewUploadFileLogic()
	GetFile = logic.NewGetFileLogic()
	GetFiles = logic.NewGetFilesLogic()
}
