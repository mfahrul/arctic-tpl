package core

import (
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"{{.Projectpath}}/utils"
)

type service struct {
	repo   Repository
	logger log.Logger
}

//NewService function
func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// func (s service) GetRepo() Repository {
// 	return s.repo
// }

func (s service) CreateItem(itm CreateItemReq) (itemID string, err error) {
	logger := log.With(s.logger, "method", "Register Item")

	tn := time.Now()
	createdBy, _ := primitive.ObjectIDFromHex(itm.CreatedBy)
	newItem := Item{
		ItemName:       itm.ItemName,
		ItemDesc:       itm.ItemDesc,
		AdditionalInfo: itm.AdditionalInfo,
		CreatedBy:      &createdBy,
		CreatedAt:      &tn,
	}

	itemID, err = s.repo.CreateItem(newItem)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		level.Info(logger).Log("Create item success with id ", itemID)
	}

	return
}

func (s service) GetAllItems() (items []Item, err error) {
	logger := log.With(s.logger, "method", "Get All Item")

	items, err = s.repo.GetAllItems()
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		// i, _ := json.Marshal(items)
		level.Info(logger).Log("Get items ", len(items))
	}

	return
}

func (s service) GetItemByID(id string) (item Item, err error) {
	logger := log.With(s.logger, "method", "GetItem")

	item, err = s.repo.GetItemByID(id)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal(item)
		level.Info(logger).Log("Get item ", string(i))
	}

	return
}

func (s service) UpdateItem(uid string, updt UpdateItemReq) (item Item, err error) {
	logger := log.With(s.logger, "method", "Update Item")

	tn := time.Now()
	// updatedBy, _ := primitive.ObjectIDFromHex(updt.UpdatedBy)
	// update := Item{
	// 	ItemName:       updt.ItemName,
	// 	ItemDesc:       updt.ItemDesc,
	// 	AdditionalInfo: updt.AdditionalInfo,
	// 	Status:         updt.Status,
	// 	UpdatedBy:      &updatedBy,
	// 	UpdatedAt:      &tn,
	// }

	update, _ := utils.ToDoc(updt)

	update = append(update, bson.E{Key: "UpdatedAt", Value: &tn})

	item, err = s.repo.UpdateItem(uid, update)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal(item)
		level.Info(logger).Log("Updated item ", string(i))
	}

	return
}

func (s service) DeleteItem(uid string, deletedby string) (item Item, err error) {
	logger := log.With(s.logger, "method", "Delete Item")

	type deleteStruct struct {
		DeletedBy string     `bson:"id:DeletedBy"`
		IsDeleted bool       `bson:"IsDeleted"`
		DeletedAt *time.Time `bson:"DeletedAt,omitempty"`
	}

	tn := time.Now()
	itm := deleteStruct{
		DeletedBy: deletedby,
		IsDeleted: true,
		DeletedAt: &tn,
	}

	update, _ := utils.ToDoc(itm)

	item, err = s.repo.UpdateItem(uid, update)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal(item)
		level.Info(logger).Log("Deleted item ", string(i))
	}

	return
}
