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
	pageCollection     *mongo.Collection
	oncePageCollection sync.Once
)

func GetPageCollection() *mongo.Collection {
	oncePageCollection.Do(func() {
		pageCollection = GetMongoClient().Database("Tenders").Collection("html")
	})
	return pageCollection
}

func SavePage(page models.SavedPage) error {
	filter := bson.M{"notice_number": page.NoticeNumber}
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": page}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := GetPageCollection().UpdateOne(ctx, filter, update, opts)
	return err
}

func GetPageByNoticeNumber(noticeNumber string) (*models.SavedPage, error) {
	var page models.SavedPage
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := GetPageCollection().FindOne(ctx, bson.M{"notice_number": noticeNumber}).Decode(&page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
