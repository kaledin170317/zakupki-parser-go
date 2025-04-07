package MongoDB

import (
	"ZakupkiParser/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	tenderCollection     *mongo.Collection
	onceTenderCollection sync.Once
)

// Получение singleton-коллекции
func GetTenderCollection() *mongo.Collection {
	onceTenderCollection.Do(func() {
		tenderCollection = GetMongoClient().Database("Tenders").Collection("tender")
	})
	return tenderCollection
}

// Сохранение объекта Tender по notice_number (upsert)
func SaveTender(tender models.Tender) error {
	filter := bson.M{"notice_number": tender.NoticeNumber}
	update := bson.M{"$set": tender}
	opts := options.Update().SetUpsert(true)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := GetTenderCollection().UpdateOne(ctx, filter, update, opts)
	return err
}

// Получение Tender по номеру извещения
func GetTenderByNoticeNumber(noticeNumber string) (*models.Tender, error) {
	var tender models.Tender
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := GetTenderCollection().FindOne(ctx, bson.M{"notice_number": noticeNumber}).Decode(&tender)
	if err != nil {
		return nil, err
	}
	return &tender, nil
}
