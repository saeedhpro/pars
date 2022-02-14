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
	parts, err := getAutomobilePartsByMySQL(id)
	if err != nil {
		return parts, nil
	}
	return parts, nil
}

func getAutomobilePartsByMySQL(id string) ([]model.Part, error) {
	parts := []model.Part{}
	query := "SELECT id, name, automobile_id FROM `parts` where automobile_id = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return parts, nil
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return parts, nil
	}
	part := model.Part{}
	for rows.Next() {
		err = rows.Scan(
			&part.ID,
			&part.Name,
			&part.AutomobileID,
		)
		if err != nil {
			log.Println(err.Error())
			return parts, nil
		}
		parts = append(parts, part)
	}
	return parts, nil
}

func getAutomobilePartsByMongodb(ctx *gin.Context, id string) ([]model.Part, error) {
	parts := []model.Part{}
	filter := bson.M{"automobile_id": id}
	cursor, err := repository.DBS.MongoDB.Database("parts").Collection("parts").Find(*repository.DBS.Context, filter)
	if err != nil {
		log.Fatal(err)
		return parts, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item model.Part
		if err = cursor.Decode(&item); err != nil {
			log.Fatal(err)
			return parts, err
		}
		parts = append(parts, item)
	}
	fmt.Println(parts)
	return parts, nil
}
