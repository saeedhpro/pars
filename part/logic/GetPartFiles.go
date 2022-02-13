package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"part/model"
	"part/repository"
)

type GetPartFilesLogic interface {
	GetPartFiles(ctx *gin.Context) ([]model.Part, error)
}

type getPartFilesLogic struct {
}

func NewGetPartFilesLogic() GetPartFilesLogic {
	return &getPartFilesLogic{}
}

func (u *getPartFilesLogic) GetPartFiles(ctx *gin.Context) ([]model.Part, error) {
	var parts []model.Part
	id := ctx.Param("id")
	if id == "" {
		return parts, nil
	}
	project := bson.M{"files": 1, "_id": 0}
	filter := bson.M{"_id": 1, "$project": project}
	results, err := repository.DBS.MongoDB.Database("parts").Collection("parts").Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var decode interface{}
	for results.Next(ctx) {
		err := results.Decode(decode)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			parts = append(parts, decode.(model.Part))
		}
	}
	return parts, nil
}
