package model

type Part struct {
	AutomobileID int      `bson:"automobile_id,omitempty"`
	Name         string   `bson:"name,omitempty"`
	Files        []string `bson:"files,omitempty"`
}
