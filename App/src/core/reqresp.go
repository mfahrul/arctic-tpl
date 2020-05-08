package core

import (
	"time"

	"io.giftano.api/go_core/config"
)

type (

	//CreateItemReq : struct for create item request
	CreateItemReq struct {
		ItemName       string   `form:"item_name,omitempty" json:"item_name,omitempty" binding:"required"`
		ItemDesc       string   `form:"item_desc,omitempty" json:"item_desc,omitempty"`
		AdditionalInfo []string `form:"additional_info,omitempty" json:"additional_info,omitempty"`
		CreatedBy      string   `form:"created_by,omitempty" json:"created_by,omitempty"`
	}

	//GetItemRequest : struct for get item request
	GetItemRequest struct {
		ID string `json:"id"`
	}

	//UpdateItemReq : struct for update item request
	UpdateItemReq struct {
		ItemName       *string   `bson:"ItemName" form:"item_name,omitempty" json:"item_name,omitempty"`
		ItemDesc       *string   `bson:"ItemDesc" form:"item_desc,omitempty" json:"item_desc,omitempty"`
		AdditionalInfo *[]string `bson:"AdditionalInfo" form:"additional_info,omitempty" json:"additional_info,omitempty"`
		Status         *int      `bson:"Status" form:"status,omitempty" json:"status,omitempty"` //0, 1, 9 => notactive, activate, ban
		UpdatedBy      *string   `bson:"id:UpdatedBy" form:"id:updated_by,omitempty" json:"id:updated_by,omitempty"`
	}

	//DeleteItemRequest : struct for delete item request
	DeleteItemRequest struct {
		UserID string `binding:"required"`
	}

	//CreateItemResponse : struct for create item response
	CreateItemResponse struct {
		config.StatusCreatedResponse
	}

	//GetAllItemsResponse : struct for get all item response
	GetAllItemsResponse struct {
		config.StatusOKResponse
		Data []Item `json:"data"`
	}

	//GetItemResponse : struct for get a item response
	GetItemResponse struct {
		config.StatusOKResponse
		Data Item `json:"data"`
	}

	//DeleteItemResponse : struct for delete a item response
	DeleteItemResponse struct {
		config.StatusOKResponse
		Data DeletedAtItem `json:"data"`
	}

	//DeletedAtItem : struct for delete response
	DeletedAtItem struct {
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}
)
