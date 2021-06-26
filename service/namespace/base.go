package namespace

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (service *Service) makeFindOneQuery(code string) bson.D {
	return bson.D{
		bson.E{
			Key:   "code",
			Value: code,
		},
	}
}

func (service *Service) CheckExist(code string) (exist bool, err error) {
	filter := service.makeFindOneQuery(code)
	var total int64
	_, total, err = service.namespaceMongoRepo.Query(filter, 0, 0, "")
	if err != nil {
		return
	}
	if total > 0 {
		exist = true
	}
	return
}
