package models

type ElectronicPlatform struct {
	Name string `bson:"name"`
	URL  string `bson:"url"`
}
