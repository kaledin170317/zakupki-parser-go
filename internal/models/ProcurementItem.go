package models

type ProcurementItem struct {
	Name            string               `bson:"name"`
	Identifier      string               `bson:"identifier"`
	Code            string               `bson:"code"`
	Customer        string               `bson:"customer"`
	ItemType        string               `bson:"item_type"`
	Unit            string               `bson:"unit"`
	UnitPrice       float64              `bson:"unit_price"`
	Quantity        float64              `bson:"quantity"`
	TotalPrice      float64              `bson:"total_price"`
	Characteristics []ItemCharacteristic `bson:"characteristics"`
}
