package list_item

import listItemRepo "github.com/siskinc/srv-name-list/repository/list_item"

type Service struct {
	listItemRepoObj *listItemRepo.MongoRepo
}

func NewService() *Service {
	return &Service{
		listItemRepoObj: listItemRepo.NewMongoRepo(),
	}
}
