package list_item

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/types"
	"github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/models"
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
