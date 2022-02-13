package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"part/model"
	"part/repository"
)

type GetAutomobileFilesLogic interface {
	GetAutomobileFiles(ctx *gin.Context) ([]model.Part, error)
}

type getAutomobileFilesLogic struct {
}

func NewGetAutomobileFilesLogic() GetAutomobileFilesLogic {
	return &getAutomobileFilesLogic{}
}

func (u *getAutomobileFilesLogic) GetAutomobileFiles(ctx *gin.Context) ([]model.Part, error) {
	project := bson.M{"files": 1, "_id": 0}
	filter := bson.M{"automobile_id": 1, "$project": project}
	results, err := repository.DBS.MongoDB.Database("parts").Collection("parts").Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var parts []model.Part
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
