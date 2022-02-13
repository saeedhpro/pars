package logic

import (
	"file/helper"
	minio2 "file/minio"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
)

type UploadFileLogic interface {
	UploadFile(ctx *gin.Context) (*string, error)
}

type uploadFileLogic struct {
}

func NewUploadFileLogic() UploadFileLogic {
	return &uploadFileLogic{}
}

func (u *uploadFileLogic) UploadFile(ctx *gin.Context) (*string, error) {
	file, err := ctx.FormFile("file")

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	buffer, err := file.Open()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer buffer.Close()

	objectName := helper.RandomString(22)
	contentType := file.Header["Content-Type"][0]
	fileBuffer := buffer
	fileSize := file.Size

	info, err := minio2.Client.PutObject(ctx, minio2.BucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return &objectName, nil
}
