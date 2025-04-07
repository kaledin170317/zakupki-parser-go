package models

type SavedPage struct {
	NoticeNumber string `bson:"notice_number"` // номер извещения
	URL          string `bson:"url"`           // откуда скачано
	HTML         string `bson:"html"`          // содержимое страницы
}
