package controller

import "part/logic"

var (
	GetAutomobileParts logic.GetAutomobilePartsLogic
	AddFile            logic.AddFileLogic
	GetPartFiles       logic.GetPartFilesLogic
	GetAutomobileFiles logic.GetAutomobileFilesLogic
	AddAutomobileParts logic.AddAutomobilePartsLogic
)

func init() {
	GetAutomobileParts = logic.NewGetAutomobilePartsLogic()
	AddFile = logic.NewAddFileLogic()
	GetPartFiles = logic.NewGetPartFilesLogic()
	GetAutomobileFiles = logic.NewGetAutomobileFilesLogic()
	AddAutomobileParts = logic.NewAddAutomobilePartsLogic()
}
