package list_item

import (
	"context"
	"fmt"
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

type RepoListItemMgo struct {
	collection *mongo.Collection
}

func NewRepoListItemMgo() *RepoListItemMgo {
	return &RepoListItemMgo{
		collection: NewCollection(),
	}
}

func (repo *RepoListItemMgo) Create(listItem *models.ListItem) error {
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

func (repo *RepoListItemMgo) UpdateById(oid primitive.ObjectID, updater bson.E) (err error) {
	_, err = repo.collection.UpdateByID(context.Background(), oid, updater)
	return
}

func (repo *RepoListItemMgo) FindById(oid primitive.ObjectID) (listItem *models.ListItem, err error) {
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
