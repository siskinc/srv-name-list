package list_item

import (
	"context"
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/types"
	"github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCollection() *mongo.Collection {
	driver := global.Config.MongoDbDriver
	return driver.DataBase(driver.DatabaseName).Collection(types.CollectionNameListItem)
}

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo() *MongoRepo {
	return &MongoRepo{
		collection: NewCollection(),
	}
}

func (repo *MongoRepo) Create(listItem *models.ListItem) error {
	if listItem == nil {
		return fmt.Errorf("list item object is nil")
	}

	_, err := repo.collection.InsertOne(context.Background(), listItem)
	if err != nil {
		logrus.Errorf("insert list item: %+v have an err: %v", listItem, err)
		return err
	}
	return nil
}

func (repo *MongoRepo) UpdateById(oid primitive.ObjectID, updater bson.M) (listItem *models.ListItem,err error) {
	filter := mongox.MakeQueryByID(oid)
	result := repo.collection.FindOneAndUpdate(context.Background(), filter, updater, mongox.MakeReturnAfter(nil))
	err = result.Err()
	if err != nil {
		return
	}
	listItem = &models.ListItem{}
	err = result.Decode(listItem)
	return
}

func (repo *MongoRepo) DeleteById(oid primitive.ObjectID) (err error) {
	query := mongox.MakeQueryByID(oid)
	_, err = repo.collection.DeleteOne(context.Background(), query)
	return
}

func (repo *MongoRepo) FindById(oid primitive.ObjectID) (listItem *models.ListItem, err error) {
	query := mongox.MakeQueryByID(oid)
	result := repo.collection.FindOne(context.Background(), query)
	listItem = &models.ListItem{}
	err = result.Decode(listItem)
	if err != nil {
		logrus.Errorf("find list item by id decode have an err: %v", err)
		return
	}
	return
}

func (repo *MongoRepo) Query(filter bson.D, pageIndex, pageSize int64, sortedField string) ([]*models.ListItem, int64, error) {
	var err error
	if filter == nil {
		filter = bson.D{}
	}

	opt := mongox.MakeFindPageOpt(nil, pageIndex, pageSize)
	opt = mongox.MakeSortedFieldOpt(opt, sortedField)
	total, err := repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		err = errorx.NewErrorWithLog("count list item have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return nil, 0, err
	}
	cursor, err := repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		err = errorx.NewErrorWithLog("find list item have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	var result []*models.ListItem
	for cursor.Next(context.Background()) {
		listItem := &models.ListItem{}
		err = cursor.Decode(listItem)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, listItem)
	}
	return result, total, nil
}
