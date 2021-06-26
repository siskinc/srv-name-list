package namespace

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
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewCollection() *mongo.Collection {
	driver := global.Config.MongoDbDriver
	return driver.DataBase(driver.DatabaseName).Collection(types.CollectionNamespace)
}

func NewNamespaceMongoRepo() *MongoRepo {
	return &MongoRepo{
		collection: NewCollection(),
	}
}

func (repo *MongoRepo) Create(namespaceObj *models.Namespace) (err error) {
	if namespaceObj == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("namespace object is nil"))
		return
	}
	_, err = repo.collection.InsertOne(context.Background(), namespaceObj)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoRepo) Delete(oid primitive.ObjectID) (err error) {
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("namespace oid is nil"))
		return
	}
	filter := mongox.MakeQueryByID(oid)
	_, err = repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return
	}
	return
}

func (repo *MongoRepo) Query(filter bson.D, pageIndex, pageSize int64, sortedField string) (result []*models.Namespace, total int64, err error) {
	if filter == nil {
		filter = bson.D{}
	}
	opt := mongox.MakeFindPageOpt(nil, pageIndex, pageSize)
	opt = mongox.MakeSortedFieldOpt(opt, sortedField)
	total, err = repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		err = errorx.NewErrorWithLog("count list type have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return
	}
	cursor, err := repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		err = errorx.NewErrorWithLog("find list type have an err: %v, filter: %+v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		npObj := &models.Namespace{}
		err = cursor.Decode(npObj)
		if err != nil {
			return
		}
		result = append(result, npObj)
	}
	return
}

func (repo *MongoRepo) makeUpdate(description string) bson.M {
	return bson.M{
		"$set": bson.M{
			"description": description,
		},
	}
}

// Update 只允许修改description字段
func (repo *MongoRepo) Update(oid primitive.ObjectID, description string) (npObj *models.Namespace, err error) {
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("namespace oid is nil"))
		return
	}
	filter := mongox.MakeQueryByID(oid)
	update := repo.makeUpdate(description)
	findAndUpdateByIdResult := repo.collection.FindOneAndUpdate(context.Background(), filter, update, mongox.MakeReturnAfter(nil))
	err = findAndUpdateByIdResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("not found object in db"))
		}
		return
	}
	npObj = &models.Namespace{}
	err = findAndUpdateByIdResult.Decode(npObj)
	if err != nil {
		logrus.Errorf("cannot decode to Namespace, err: %v", err)
		return
	}
	return
}
