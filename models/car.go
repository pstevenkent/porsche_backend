package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Images          []string           `json:"images"`
	Preview         string             `json:"preview"`
	Vehicle         string             `json:"vehicle"`
	ModelYear       int                `json:"modelyear"`
	ExteriorColour  string             `json:"exteriorcolour"`
	InteriorColours string             `json:"interiorcolours"`
	Wheels          string             `json:"wheels"`
	Seats           string             `json:"seats"`

	PaintedWheels       string   `json:"paintedwheels"`
	LetteringDecals     string   `json:"letteringdecals"`
	SeatbeltsSeatdesign string   `json:"seatbeltsseatdesign"`
	ExteriorDesign      []string `json:"exteriordesign"`
	InteriorDesign      []string `json:"interiordesign"`
	AssistanceSystems   []string `json:"assistancesystems"`
	ComfortNUsability   []string `json:"comfortnusability"`
	LightsVision        []string `json:"lightsvision"`
	EquipmentPackages   string   `json:"equipmentpackages"`

	RoofTransport         string   `json:"rooftransport"`
	PowertrainPerformance []string `json:"powertrainperformance"`
	Infotainment          string   `json:"infotainment"`
	CommNr                string   `json:"commnr"`
	Price                 int      `json:"price"`
}
