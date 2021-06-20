package list_type

import (
	"context"
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoListTypeMgo struct {
	collection *mongo.Collection
}

func (repo *RepoListTypeMgo) makeQueryByCode(listType *models.ListType) bson.D {
	return bson.D{
		{
			Key:   "code",
			Value: listType.Code,
		},
	}
}

func (repo *RepoListTypeMgo) makeQueryById(listTypeId primitive.ObjectID) bson.D {
	return bson.D{
		{
			Key:   "_id",
			Value: listTypeId,
		},
	}
}

func (repo *RepoListTypeMgo) Create(listType *models.ListType) error {
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
	}

	result, err := repo.collection.InsertOne(context.Background(), listType)
	if err != nil {
		logrus.Errorf("insert new list type have an err: %v", err)
		return err
	}
	logrus.Infof("create list type successful, list type: %+v, insert _id: %v", listType, result.InsertedID)

	return nil
}

func (repo *RepoListTypeMgo) Delete(listTypeId primitive.ObjectID) error {
	if listTypeId == primitive.NilObjectID {
		return fmt.Errorf("list type id is nil")
	}
	queryById := repo.makeQueryById(listTypeId)
	findById := repo.collection.FindOne(context.Background(), queryById)
	err := findById.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			logrus.Errorf("find list type by code when create a list type have an err: %v", err)
			return err
		}
	}

	_, err = repo.collection.DeleteOne(context.Background(), queryById)
	if err != nil {
		logrus.Errorf("delete list type have an err: %v, object id: %s", err, listTypeId.String())
		return err
	}
	logrus.Infof("delete list type successful, object id: %s", listTypeId.String())
	return nil
}

func (repo *RepoListTypeMgo) Query(filter bson.D, pageIndex, pageSize int64, sortedField string) ([]*models.ListType, error) {
	var err error
	if filter == nil {
		filter = bson.D{}
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	opt := options.Find()
	opt.SetLimit(pageSize)
	if pageIndex > 1 {
		skip := (pageIndex - 1) * pageSize
		opt.SetSkip(skip)
	}
	if sortedField != "" {
		opt.SetSort(sortedField)
	}
	cursor, err := repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		err = errorx.NewErrorWithLog("find list type have an err: %v, filter: %v, pageIndex: %d, pageSize: %d, "+
			"sortedField: %s", err, filter, pageIndex, pageSize, sortedField)
		return nil, err
	}
	defer cursor.Close(context.Background())
	var result []*models.ListType
	for cursor.Next(context.Background()) {
		listType := &models.ListType{}
		err = cursor.Decode(listType)
		if err != nil {
			return nil, err
		}
		result = append(result, listType)
	}
	return result, nil
}

func (repo *RepoListTypeMgo) makeUpdate(isValid bool, description string) bson.M {
	return bson.M{
		"$set": bson.M{
			"is_valid": isValid,
			"description": description,
		},
	}
}

// Update 只允许修改is_valid和description字段
func (repo *RepoListTypeMgo) Update(listTypeId primitive.ObjectID, isValid bool, description string) error {
	if listTypeId == primitive.NilObjectID {
		return fmt.Errorf("list type object is nil")
	}
	queryById := repo.makeQueryById(listTypeId)
	update := repo.makeUpdate(isValid, description)
	findAndUpdateByIdResult := repo.collection.FindOneAndUpdate(context.Background(), queryById, update)
	err := findAndUpdateByIdResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("not found object in db"))
		}
	}
	return nil
}
