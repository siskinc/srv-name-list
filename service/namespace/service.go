package namespace

import (
	namespaceRepo "github.com/siskinc/srv-name-list/repository/namespace"
)

type Service struct {
	namespaceMongoRepo *namespaceRepo.MongoRepo
}

func NewService() *Service {
	return &Service{
		namespaceMongoRepo: namespaceRepo.NewNamespaceMongoRepo(),
	}
}