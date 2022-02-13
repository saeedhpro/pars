package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"part/model"
	"part/repository"
	"part/requests"
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
	var request requests.AddAutomobilePartsRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		return err
	}
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
