package driver

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDbDriver struct {
	URI    string `json:"uri"`
	client *mongo.Client
}

func (mongoDbDriver *MongoDbDriver) Init() {
	logrus.Infof("uri: %s", mongoDbDriver.URI)
}

func (mongoDbDriver *MongoDbDriver) Connect() error {
	var err error
	logrus.Infof("begin connect mongodb")
	defer func() {
		if err != nil {
			logrus.Errorf("connect mongodb failed, err: %v", err)
		} else {

		}
	}()
	mongoDbDriver.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDbDriver.URI))
	if err != nil {
		err = fmt.Errorf("connect mongodb have an err: %v, uri: %s", err, mongoDbDriver.URI)
		return err
	}
	err = mongoDbDriver.client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		err = fmt.Errorf("ping mongodb have an err: %v", err)
		return err
	}
	return nil
}

func (mongoDbDriver *MongoDbDriver) checkConnect() {
	if mongoDbDriver.client == nil {
		panic(fmt.Errorf("mongodb client not init"))
	}
}

func (mongoDbDriver *MongoDbDriver) DataBase(dbName string, opts ...*options.DatabaseOptions) *mongo.Database {
	mongoDbDriver.checkConnect()
	return mongoDbDriver.client.Database(dbName, opts...)
}

