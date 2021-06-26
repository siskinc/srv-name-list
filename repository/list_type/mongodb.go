package list_type

import (
	"context"
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/contants/types"
	"github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCollection() *mongo.Collection {
	driver := global.Config.MongoDbDriver
	return driver.DataBase(driver.DatabaseName).Collection(types.CollectionNameListType)
}

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo() *MongoRepo {
	return &MongoRepo{collection: NewCollection()}
}

func (repo *MongoRepo) makeQueryByCode(listType *models.ListType) bson.D {
	return bson.D{
		{
			Key:   "code",
			Value: listType.Code,
		},
		{
			Key:   "namespace",
			Value: listType.Namespace,
		},
	}
}

func (repo *MongoRepo) makeQueryById(listTypeId primitive.ObjectID) bson.D {
	return bson.D{
		{
			Key:   "_id",
			Value: listTypeId,
		},
	}
}

func (repo *MongoRepo) Create(listType *models.ListType) error {
	if listType == nil {
		return fmt.Errorf("list type object is nil")
	}

	queryByCode := repo.makeQueryByCode(listType)

	findByCode := repo.collection.FindOne(context.Background(), queryByCode)
	err := findByCode.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		} else {
			logrus.Errorf("find list type by code when create a list type have an err: %v", err)
			return err
		}
	} else {
		err = errorx.NewError(error_code.CustomForbiddenConflictListType, fmt.Errorf("list type exist"))
		return err
	}

	result, err := repo.collection.InsertOne(context.Background(), listType)
	if err != nil {
		logrus.Errorf("insert new list type have an err: %v", err)
		return err
	}
	logrus.Infof("create list type successful, list type: %+v, insert _id: %v", listType, result.InsertedID)

	return nil
}

func (repo *MongoRepo) Delete(listTypeId primitive.ObjectID) error {
	filter := mongox.MakeQueryByID(listTypeId)
	findById := repo.collection.FindOne(context.Background(), filter)
	err := findById.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			logrus.Errorf("find list type by code when create a list type have an err: %v", err)
			return err
		}
	}

	_, err = repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logrus.Errorf("delete list type have an err: %v, object id: %s", err, listTypeId.String())
		return err
	}
	logrus.Infof("delete list type successful, object id: %s", listTypeId.String())
	return nil
}

func (repo *MongoRepo) Query(filter bson.D, pageIndex, pageSize int64, sortedField string) ([]*models.ListType, int64, error) {
	var err error
	if filter == nil {
		filter = bson.D{}
	}
	var opt *options.FindOptions
	if pageSize != 0 {
		opt = mongox.MakeFindPageOpt(nil, pageIndex, pageSize)
	}
	opt = mongox.MakeSortedFieldOpt(opt, sortedField)
	total, err := repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		err = errorx.NewErrorWithLog("count list type have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return nil, 0, err
	}
	cursor, err := repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		err = errorx.NewErrorWithLog("find list type have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	var result []*models.ListType
	for cursor.Next(context.Background()) {
		listType := &models.ListType{}
		err = cursor.Decode(listType)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, listType)
	}
	return result, total, nil
}

func (repo *MongoRepo) makeUpdate(isValid bool, description string) bson.M {
	return bson.M{
		"$set": bson.M{
			"is_valid":    isValid,
			"description": description,
		},
	}
}

// Update 只允许修改is_valid和description字段
func (repo *MongoRepo) Update(listTypeId primitive.ObjectID, isValid bool, description string) (*models.ListType, error) {
	filter := mongox.MakeQueryByID(listTypeId)
	update := repo.makeUpdate(isValid, description)
	findAndUpdateByIdResult := repo.collection.FindOneAndUpdate(context.Background(), filter, update, mongox.MakeReturnAfter(nil))
	err := findAndUpdateByIdResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("not found object in db"))
		}
		return nil, err
	}
	listType := &models.ListType{}
	err = findAndUpdateByIdResult.Decode(listType)
	if err != nil {
		logrus.Errorf("cannot decode to ListType, err: %v", err)
		return nil, err
	}
	return listType, nil
}
