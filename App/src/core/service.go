package core

//Service interface
type Service interface {
	CreateItem(item CreateItemReq) (string, error)
	GetItemByID(id string) (Item, error)
	GetAllItems() ([]Item, error)
	UpdateItem(id string, update UpdateItemReq) (Item, error)
	DeleteItem(id string, deletedby string) (Item, error)
	// GetRepo() Repository
}
