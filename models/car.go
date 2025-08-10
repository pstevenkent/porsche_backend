package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	Id                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Images                []string           `json:"images"`
	Vehicle               string             `json:"vehicle"`
	ModelYear             int                `json:"modelyear"`
	ExteriorColour        string             `json:"exteriorcolour"`
	InteriorColours       string             `json:"interiorcolours"`
	Wheels                string             `json:"wheels"`
	Seats                 string             `json:"seats"`
	RoofTransport         string             `json:"rooftransport"`
	PowertrainPerformance []string           `json:"powertrainperformance"`
	Infotainment          string             `json:"infotainment"`
}
