package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"part/repository"
	"part/requests"
)

type AddFileLogic interface {
	AddFile(ctx *gin.Context) (*string, error)
}

type addFileLogic struct {
}

func NewAddFileLogic() AddFileLogic {
	return &addFileLogic{}
}

func (u *addFileLogic) AddFile(ctx *gin.Context) (*string, error) {
	id := ctx.Param("id")
	if id == "" {
		return nil, nil
	}
	var request requests.AddFileRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		return nil, err
	}
	fmt.Println(request.File)
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$addToSet": bson.M{
			"files": request.File,
		},
	}
	res := repository.DBS.MongoDB.Database("parts").Collection("parts").FindOneAndUpdate(
		*repository.DBS.Context,
		filter,
		update,
	)
	fmt.Println(res)
	return &request.File, nil
}
