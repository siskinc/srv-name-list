package list_type

import listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"

type Service struct {
	listTypeRepoObj *listTypeRepo.MongoRepo
}

func NewService() *Service {
	return &Service{
		listTypeRepoObj: listTypeRepo.NewMongoRepo(),
	}
}
