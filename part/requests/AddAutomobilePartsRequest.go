package requests

import "part/model"

type AddAutomobilePartsRequest struct {
	Parts []model.Part `json:"parts"`
}
