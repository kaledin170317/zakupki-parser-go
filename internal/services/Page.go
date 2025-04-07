package services

import (
	"ZakupkiParser/internal/database/MongoDB"
	"ZakupkiParser/internal/models"
	"log"
)

func download44FZ(id string) error {

	url := "https://zakupki.gov.ru/epz/order/notice/printForm/view.html?regNumber=" + id
	body, err := FetchHTML(url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	page := models.SavedPage{
		NoticeNumber: id,
		URL:          url,
		HTML:         body,
	}

	return MongoDB.SavePage(page)
}

func Get44FZ(id string) (*models.SavedPage, error) {

	page, _ := MongoDB.GetPageByNoticeNumber(id)

	if page != nil {
		return page, nil
	}

	err := download44FZ(id)

	if err != nil {
		return nil, err
	}

	return MongoDB.GetPageByNoticeNumber(id)
}
