package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"part/logic"
)

type PartHandler interface {
	GetAutomobileParts(c *gin.Context)
	AddFile(c *gin.Context)
	GetPartFiles(c *gin.Context)
	GetAutomobileFiles(c *gin.Context)
	AddAutomobileParts(c *gin.Context)
}

type partApi struct {
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic
	addFileLogic            logic.AddFileLogic
	getPartFilesLogic       logic.GetPartFilesLogic
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic
	addAutomobilePartsLogic logic.AddAutomobilePartsLogic
}

func NewPartApi(
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic,
	addFileLogic logic.AddFileLogic,
	getPartFilesLogic logic.GetPartFilesLogic,
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic,
	addAutomobilePartsLogic logic.AddAutomobilePartsLogic,
) *partApi {
	return &partApi{
		getAutomobilePartsLogic: getAutomobilePartsLogic,
		addFileLogic:            addFileLogic,
		getPartFilesLogic:       getPartFilesLogic,
		getAutomobileFilesLogic: getAutomobileFilesLogic,
		addAutomobilePartsLogic: addAutomobilePartsLogic,
	}
}

func (api *partApi) GetAutomobileParts(c *gin.Context) {
	url, err := api.getAutomobilePartsLogic.GetAutomobileParts(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, url)
	return
}

func (api *partApi) AddFile(c *gin.Context) {
	url, err := api.addFileLogic.AddFile(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, url)
	return
}

func (api *partApi) GetPartFiles(c *gin.Context) {
	list, err := api.getPartFilesLogic.GetPartFiles(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, list)
	return
}

func (api *partApi) GetAutomobileFiles(c *gin.Context) {
	list, err := api.getAutomobileFilesLogic.GetAutomobileFiles(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, list)
	return
}

func (api *partApi) AddAutomobileParts(c *gin.Context) {
	list := api.addAutomobilePartsLogic.AddAutomobileParts(c)
	c.JSON(200, list)
	return
}
