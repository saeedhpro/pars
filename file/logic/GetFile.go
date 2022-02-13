package logic

import (
	minio2 "file/minio"
	"file/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"time"
)

type GetFileLogic interface {
	GetFile(ctx *gin.Context) (*string, error)
}

type getFileLogic struct {
}

func NewGetFileLogic() GetFileLogic {
	return &getFileLogic{}
}

func (u *getFileLogic) GetFile(ctx *gin.Context) (*string, error) {
	var request requests.GetFileRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	reqParams := make(url.Values)
	preSignedURL, err := minio2.Client.PresignedGetObject(ctx, minio2.BucketName, request.Name, time.Second*24*60*60, reqParams)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &preSignedURL.RawPath, nil
}
