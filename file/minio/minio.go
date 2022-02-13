package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

var Client *minio.Client

var BucketName = "parts"

func Run() {
	endPoint := fmt.Sprintf("%s:%d", "localhost", 9000)
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	client, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	Client = client
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", Client)
}
