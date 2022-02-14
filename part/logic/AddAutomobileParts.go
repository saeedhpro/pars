package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"part/model"
	"part/repository"
	"part/requests"
	"strings"
)

type AddAutomobilePartsLogic interface {
	AddAutomobileParts(ctx *gin.Context) error
}

type addAutomobilePartsLogic struct {
}

func NewAddAutomobilePartsLogic() AddAutomobilePartsLogic {
	return &addAutomobilePartsLogic{}
}

func (u *addAutomobilePartsLogic) AddAutomobileParts(ctx *gin.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return nil
	}
	var request requests.AddAutomobilePartsRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		return err
	}
	return addAutomobilePartsByMySQL(id, request)
}

func addAutomobilePartsByMySQL(id string, request requests.AddAutomobilePartsRequest) error {
	query := "INSERT INTO `part_files`(`part_id`, `name`) VALUES "
	columns := []string{}
	var values []interface{}
	for _, r := range request.Parts {
		columns = append(columns, " (?,?) ")
		values = append(values, id, r.Name)
	}
	columnsSrt := strings.Join(columns, ",")
	query += columnsSrt
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func addAutomobilePartsByMongoDB(request requests.AddAutomobilePartsRequest) error {
	var values []interface{}
	for j := 0; j < len(request.Parts); j++ {
		values = append(values, model.Part{
			AutomobileID: 1,
			Files:        request.Parts[j].Files,
			Name:         request.Parts[j].Name,
		})
	}
	fmt.Println(request.Parts)
	_, err := repository.DBS.MongoDB.Database("parts").Collection("parts").InsertMany(*repository.DBS.Context, values)
	if err != nil {
		return err
	}
	return nil
}
