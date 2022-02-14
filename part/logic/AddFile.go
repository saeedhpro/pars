package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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
	fmt.Println(request.Name)
	return addFileByMySQL(id, request)
}

func addFileByMySQL(id string, request requests.AddFileRequest) (*string, error) {
	query := "INSERT INTO `part_files`(`part_id`, `name`) VALUES (?,?)"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	_, err = stmt.Exec(id, request.Name)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &request.Name, nil
}

func addFileByMongoDB(ctx *gin.Context, id string, request requests.AddFileRequest) (*string, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$addToSet": bson.M{
			"files": request.Name,
		},
	}
	res := repository.DBS.MongoDB.Database("parts").Collection("parts").FindOneAndUpdate(
		*repository.DBS.Context,
		filter,
		update,
	)
	fmt.Println(res)
	return &request.Name, nil
}
