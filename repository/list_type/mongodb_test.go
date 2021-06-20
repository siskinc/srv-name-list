package list_type

import (
	"fmt"
	"github.com/siskinc/srv-name-list/contants/types"
	"github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func createCollection() *mongo.Collection {
	database := global.Config.MongoDbDriver.DataBase(global.Config.MongoDbDriver.DatabaseName)
	collection := database.Collection(types.CollectionNameListType)
	return collection
}

func TestRepoListTypeMgo_Create(t *testing.T) {
	collection := createCollection()
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		listType *models.ListType
	}
	now := time.Now().Unix()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "no conflict",
			fields: fields{
				collection: collection,
			},
			args: args{
				listType: &models.ListType{
					Id:          primitive.NewObjectID(),
					Code:        fmt.Sprintf("new_code_%d", now),
					Fields:      nil,
					IsValid:     false,
					Description: "no conflict",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepoListTypeMgo{
				collection: tt.fields.collection,
			}
			if err := repo.Create(tt.args.listType); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepoListTypeMgo_Delete(t *testing.T) {
	collection := createCollection()
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		listTypeId primitive.ObjectID
	}
	oid, _ := primitive.ObjectIDFromHex("60cf3703685b429d5bf67155")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete",
			fields: fields{
				collection: collection,
			},
			args: args{listTypeId: oid},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepoListTypeMgo{
				collection: tt.fields.collection,
			}
			if err := repo.Delete(tt.args.listTypeId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
