package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"part/repository"
)

type GetPartFilesLogic interface {
	GetPartFiles(ctx *gin.Context) ([]string, error)
}

type getPartFilesLogic struct {
}

func NewGetPartFilesLogic() GetPartFilesLogic {
	return &getPartFilesLogic{}
}

func (u *getPartFilesLogic) GetPartFiles(ctx *gin.Context) ([]string, error) {
	files := []string{}
	id := ctx.Param("id")
	if id == "" {
		return files, nil
	}
	return getPartFilesByMySQL(id)
}

func getPartFilesByMySQL(id string) ([]string, error) {
	files := []string{}
	query := "SELECT `part_files`.`name` FROM `part_files` WHERE `part_files`.`part_id` = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return files, nil
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return files, nil
	}
	var partFile string
	for rows.Next() {
		err = rows.Scan(
			&partFile,
		)
		if err != nil {
			log.Println(err.Error())
			return files, nil
		}
		files = append(files, partFile)
	}
	return files, nil
}
