package [[ with .ModuleToParse ]][[.Name]]

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
[[ with .Model ]]
//[[.Name | ToCamel]] struct
type [[.Name | ToCamel]] struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`[[ with .Structures ]][[ range . ]]
	[[.Name | ToCamel]]       [[.Type]]              `bson:"[[.Name | ToCamel]]" json:"[[.Name | ToCamel | ToSnake]]"`[[ end ]][[ end ]]
	IsDeleted      bool                `bson:"IsDeleted" json:"is_deleted"`
	CreatedBy      *primitive.ObjectID `bson:"CreatedBy,omitempty" json:"created_by,omitempty"`
	UpdatedBy      *primitive.ObjectID `bson:"UpdatedBy,omitempty" json:"updated_by,omitempty"`
	DeletedBy      *primitive.ObjectID `bson:"DeletedBy,omitempty" json:"deleted_by,omitempty"`
	CreatedAt      *time.Time          `bson:"CreatedAt,omitempty" json:"created_at,omitempty"`
	UpdatedAt      *time.Time          `bson:"UpdatedAt,omitempty" json:"updated_at,omitempty"`
	DeletedAt      *time.Time          `bson:"DeletedAt,omitempty" json:"deleted_at,omitempty"`
}
[[ end ]]
[[ with .AddStructs ]] [[ range . ]]
//[[.Name | ToCamel]] struct
type [[.Name | ToCamel]] struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`[[ with .Structures ]][[ range . ]]
	[[.Name | ToCamel]]       [[.Type]]              `bson:"[[.Name | ToCamel]]" json:"[[.Name | ToSnake]]"`[[ end ]][[ end ]]
	IsDeleted      bool                `bson:"IsDeleted" json:"is_deleted"`
	CreatedBy      *primitive.ObjectID `bson:"CreatedBy,omitempty" json:"created_by,omitempty"`
	UpdatedBy      *primitive.ObjectID `bson:"UpdatedBy,omitempty" json:"updated_by,omitempty"`
	DeletedBy      *primitive.ObjectID `bson:"DeletedBy,omitempty" json:"deleted_by,omitempty"`
	CreatedAt      *time.Time          `bson:"CreatedAt,omitempty" json:"created_at,omitempty"`
	UpdatedAt      *time.Time          `bson:"UpdatedAt,omitempty" json:"updated_at,omitempty"`
	DeletedAt      *time.Time          `bson:"DeletedAt,omitempty" json:"deleted_at,omitempty"`
}
[[ end ]][[ end ]]

//Repository interface
type Repository interface { [[ with .Model ]]
	Create[[.Name | ToCamel]]([[.Name | ToLower]] [[.Name | ToCamel]]) (string, error)
	Get[[.Name | ToCamel]]ByID(id string) ([[.Name | ToCamel]], error)
	GetAll[[.Name | ToCamel | ToPlural]]() ([][[.Name | ToCamel]], error)
	Update[[.Name | ToCamel]](id string, update interface{}) ([[.Name | ToCamel]], error) [[ end ]]
}[[ end ]]