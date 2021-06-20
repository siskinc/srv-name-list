package list_type

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
