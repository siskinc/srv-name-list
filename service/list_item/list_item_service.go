package list_item

import listItemRepo "github.com/siskinc/srv-name-list/repository/list_item"

type ListItemService struct {
	listItemRepoObj *listItemRepo.RepoListItemMgo
}

func NewListItemService() *ListItemService {
	return &ListItemService{
		listItemRepoObj: listItemRepo.NewRepoListItemMgo(),
	}
}
