package models

type AttachmentGroup struct {
	Title string   `bson:"title"`
	Items []string `bson:"items"`
}
