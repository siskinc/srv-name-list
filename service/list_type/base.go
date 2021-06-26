package list_type

import (
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *Service) makeUniqueQuery(namespace, code string) bson.D {
	return bson.D{
		bson.E{
			Key:   "namespace",
			Value: namespace,
		},
		bson.E{
			Key:   "code",
			Value: code,
		},
	}
}

func (service *Service) CheckExist(namespace, code string) (exist bool, err error) {
	filter := service.makeUniqueQuery(namespace, code)
	_, total, err := service.listTypeRepoObj.Query(filter, 1 , 10, "")
	if err != nil {
		return
	}
	if total > 0 {
		exist = true
	}
	return
}

func (service *Service) FindOne(namespace, code string) (listType *models.ListType, err error) {
	filter := service.makeUniqueQuery(namespace, code)
	result, _, err := service.listTypeRepoObj.Query(filter, 1, 10,"")
	if err != nil {
		return
	}
	if len(result) > 0 {
		listType = result[0]
	}
	return
}
