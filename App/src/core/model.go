package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item struct
type Item struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ItemName       string              `bson:"ItemName" json:"item_name"`
	ItemDesc       string              `bson:"ItemDesc" json:"item_desc"`
	AdditionalInfo []string            `bson:"AdditionalInfo" json:"additional_info"`
	Status         int                 `bson:"Status" json:"status"` //0, 1, 9 => notactive, activate, ban
	IsDeleted      bool                `bson:"IsDeleted" json:"is_deleted"`
	CreatedBy      *primitive.ObjectID `bson:"CreatedBy,omitempty" json:"created_by,omitempty"`
	UpdatedBy      *primitive.ObjectID `bson:"UpdatedBy,omitempty" json:"updated_by,omitempty"`
	DeletedBy      *primitive.ObjectID `bson:"DeletedBy,omitempty" json:"deleted_by,omitempty"`
	CreatedAt      *time.Time          `bson:"CreatedAt,omitempty" json:"created_at,omitempty"`
	UpdatedAt      *time.Time          `bson:"UpdatedAt,omitempty" json:"updated_at,omitempty"`
	DeletedAt      *time.Time          `bson:"DeletedAt,omitempty" json:"deleted_at,omitempty"`
}

//Repository interface
type Repository interface {
	CreateItem(item Item) (string, error)
	GetItemByID(id string) (Item, error)
	GetAllItems() ([]Item, error)
	UpdateItem(id string, update interface{}) (Item, error)
}
