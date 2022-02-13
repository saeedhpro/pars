package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"part/model"
	"part/repository"
)

type GetAutomobilePartsLogic interface {
	GetAutomobileParts(ctx *gin.Context) ([]model.Part, error)
}

type getAutomobilePartsLogic struct {
}

func NewGetAutomobilePartsLogic() GetAutomobilePartsLogic {
	return &getAutomobilePartsLogic{}
}

func (u *getAutomobilePartsLogic) GetAutomobileParts(ctx *gin.Context) ([]model.Part, error) {
	id := ctx.Param("id")
	parts := []model.Part{}
	if id == "" {
		return parts, nil
	}
	filter := bson.M{"automobile_id": id}
	cursor, err := repository.DBS.MongoDB.Database("parts").Collection("parts").Find(*repository.DBS.Context, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item model.Part
		if err = cursor.Decode(&item); err != nil {
			log.Fatal(err)
		}
		parts = append(parts, item)
		fmt.Println(item)
	}
	fmt.Println(parts)
	return parts, nil
}
