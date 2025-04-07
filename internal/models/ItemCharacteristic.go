package models

type ItemCharacteristic struct {
	Name        string `bson:"name"`
	Value       string `bson:"value"`
	Unit        string `bson:"unit"`
	Instruction string `bson:"instruction"`
}
