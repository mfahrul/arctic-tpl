package [[ with .ModuleToParse ]][[.Name]][[ end ]]

import (
	"time"

	"[[.Projectpath]]/config"
)
[[ with .ModuleToParse.Model ]]
type (

	//Create[[.Name | ToCamel]]Req : struct for create [[.Name | ToLower]] request
	Create[[.Name | ToCamel]]Req struct { [[ with .Structures ]][[ range . ]]
		[[.Name | ToCamel]]       [[.Type]]   `form:"[[.Name | ToCamel | ToSnake]],omitempty" json:"[[.Name | ToCamel | ToSnake]],omitempty"` [[ end ]][[ end ]]
		CreatedBy      string   `form:"created_by,omitempty" json:"created_by,omitempty"`
	}

	//Get[[.Name | ToCamel]]Request : struct for get [[.Name | ToLower]] request
	Get[[.Name | ToCamel]]Request struct {
		ID string `json:"id"`
	}

	//Update[[.Name | ToCamel]]Req : struct for update [[.Name | ToLower]] request
	Update[[.Name | ToCamel]]Req struct { [[ with .Structures ]][[ range . ]]
		[[.Name | ToCamel]]       *[[.Type]]   `bson:"[[.Name | ToCamel]]" form:"[[.Name | ToCamel | ToSnake]],omitempty" json:"[[.Name | ToCamel | ToSnake]],omitempty"`[[ end ]][[ end ]]
		Status         *int      `bson:"Status" form:"status,omitempty" json:"status,omitempty"` //0, 1, 9 => notactive, activate, ban
		UpdatedBy      *string   `bson:"id:UpdatedBy" form:"id:updated_by,omitempty" json:"id:updated_by,omitempty"`
	}

	//Delete[[.Name | ToCamel]]Request : struct for delete [[.Name | ToLower]] request
	Delete[[.Name | ToCamel]]Request struct {
		UserID string `binding:"required"`
	}

	//Create[[.Name | ToCamel]]Response : struct for create [[.Name | ToLower]] response
	Create[[.Name | ToCamel]]Response struct {
		config.StatusCreatedResponse
	}

	//GetAll[[.Name | ToCamel | ToPlural]]Response : struct for get all [[.Name | ToLower]] response
	GetAll[[.Name | ToCamel | ToPlural]]Response struct {
		config.StatusOKResponse
		Data [][[.Name | ToCamel]] `json:"data"`
	}

	//Get[[.Name | ToCamel]]Response : struct for get a [[.Name | ToLower]] response
	Get[[.Name | ToCamel]]Response struct {
		config.StatusOKResponse
		Data [[.Name | ToCamel]] `json:"data"`
	}

	//Delete[[.Name | ToCamel]]Response : struct for delete a [[.Name | ToLower]] response
	Delete[[.Name | ToCamel]]Response struct {
		config.StatusOKResponse
		Data DeletedAt[[.Name | ToCamel]] `json:"data"`
	}

	//DeletedAt[[.Name | ToCamel]] : struct for delete response
	DeletedAt[[.Name | ToCamel]] struct {
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}
)[[ end ]]
