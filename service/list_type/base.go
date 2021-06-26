package list_type

import "go.mongodb.org/mongo-driver/bson"

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
	_, total, err := service.listTypeRepoObj.Query(filter, 0 , 0, "")
	if err != nil {
		return
	}
	if total > 0 {
		exist = true
	}
	return
}
