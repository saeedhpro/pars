package model

type Automobile struct {
	ID          int    `bson:"id,omitempty"`
	Model       string `bson:"model,omitempty"`
	Type        string `bson:"type,omitempty"`
	Manufacture string `bson:"manufacture,omitempty"`
	Parts       []Part `bson:"parts,omitempty"`
}
