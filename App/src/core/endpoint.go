package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Projectpath}}/config"
)

var responseMessage interface{}

//Endpoints struct
type Endpoints struct {
	CreateItem  gin.HandlerFunc
	GetItemByID gin.HandlerFunc
	GetAllItems gin.HandlerFunc
	UpdateItem  gin.HandlerFunc
	DeleteItem  gin.HandlerFunc
}

//MakeEndpoints functions
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateItem:  makeCreateItemEndpoint(s),
		GetItemByID: makeGetItemByIDEndpoint(s),
		GetAllItems: makeGetAllItemEndpoint(s),
		UpdateItem:  makeUpdateItemEndpoint(s),
		DeleteItem:  makeDeleteItemEndpoint(s),
	}
}

// @Summary Create a item
// @Description Create item for access
// @Accept  json
// @Produce  json
// @Param item body CreateItemReq true "Item Param"
// @Success 201 {object} CreateItemResponse
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags item
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /item [post]
func makeCreateItemEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item CreateItemReq
		errBind := c.ShouldBind(&item)
		if errBind != nil {
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   errBind.Error(),
			}
		} else {

			if _, err := s.CreateItem(item); err != nil {
				responseMessage = config.StatusInternalServerErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				}
			} else {
				responseMessage = CreateItemResponse{
					StatusCreatedResponse: config.StatusCreatedResponse{
						Code:    http.StatusCreated,
						Message: http.StatusText(http.StatusCreated),
					},
				}
			}
		}
		c.JSON(200, responseMessage)
	}
}

// @Summary Get all items
// @Description Get all items
// @Produce  json
// @Success 200 {object} GetAllItemsResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags item
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /items [get]
func makeGetAllItemEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := s.GetAllItems()
		if err != nil {
			responseMessage = config.StatusInternalServerErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			}
		} else {
			// if len(items) > 0 {
			responseMessage = GetAllItemsResponse{
				StatusOKResponse: config.StatusOKResponse{
					Code:    http.StatusOK,
					Message: http.StatusText(http.StatusOK),
				},
				Data: items,
			}
			// } else {
			// 	responseMessage = config.StatusInternalServerErrorResponse{
			// 		Code:    http.StatusInternalServerError,
			// 		Message: http.StatusText(http.StatusInternalServerError),
			// 		Error:   err.Error(),
			// 	}
			// }
		}
		c.JSON(200, responseMessage)
	}
}

// @Summary Get a item
// @Description Get a item by ID
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} GetItemResponse
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 404 {object} config.StatusNotFoundResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags item
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /item/{id} [get]
func makeGetItemByIDEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItemRequest
		errBind := c.ShouldBindUri(&req)
		if errBind != nil {
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   errBind.Error(),
			}
		} else {
			item, err := s.GetItemByID(req.ID)
			if err != nil {
				responseMessage = config.StatusNotFoundResponse{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
					Error:   err.Error(),
				}
			} else {
				if item.ID != nil {
					responseMessage = GetItemResponse{
						StatusOKResponse: config.StatusOKResponse{
							Code:    http.StatusOK,
							Message: http.StatusText(http.StatusOK),
						},
						Data: item,
					}
				} else {
					responseMessage = config.StatusInternalServerErrorResponse{
						Code:    http.StatusInternalServerError,
						Message: http.StatusText(http.StatusInternalServerError),
						Error:   err.Error(),
					}
				}
			}

		}
		c.JSON(200, responseMessage)
	}
}

// @Summary Update a item
// @Description Update item by ID
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Param item body Item true "Item Param"
// @Success 200 {object} GetItemResponse
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags item
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /item/{id} [put]
func makeUpdateItemEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItemRequest
		var item UpdateItemReq
		errBind := c.ShouldBindUri(&req)
		errBinditem := c.ShouldBind(&item)
		if errBind != nil || errBinditem != nil {
			var err string
			if errBind != nil {
				err = errBind.Error()
			} else {
				err = errBinditem.Error()
			}
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			}
		} else {
			item, err := s.UpdateItem(req.ID, item)
			if err != nil {
				responseMessage = config.StatusInternalServerErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				}
			} else {
				responseMessage = GetItemResponse{
					StatusOKResponse: config.StatusOKResponse{
						Code:    http.StatusOK,
						Message: http.StatusText(http.StatusOK),
					},
					Data: item,
				}
			}
		}
		c.JSON(200, responseMessage)
	}
}

// @Summary Delete a item
// @Description Delete item by ID
// @Produce  json
// @Param id path string true "Item ID"
// @Param UserID header string true "Delete by"
// @Success 200 {object} DeleteItemResponse
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 404 {object} config.StatusNotFoundResponse
// @Tags item
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /item/{id} [delete]
func makeDeleteItemEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItemRequest
		var header DeleteItemRequest
		errBind := c.ShouldBindUri(&req)
		errHederBind := c.ShouldBindHeader(&header)

		if errBind != nil || errHederBind != nil {
			var err string
			if errBind != nil {
				err = errBind.Error()
			} else {
				err = errHederBind.Error()
			}
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			}
		} else {
			item, err := s.DeleteItem(req.ID, header.UserID)
			if err != nil {
				responseMessage = config.StatusNotFoundResponse{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
					Error:   err.Error(),
				}
			} else {
				responseMessage = DeleteItemResponse{
					StatusOKResponse: config.StatusOKResponse{
						Code:    http.StatusOK,
						Message: http.StatusText(http.StatusOK),
					},
					Data: DeletedAtItem{
						DeletedAt: item.DeletedAt,
					},
				}
			}
		}
		c.JSON(200, responseMessage)
	}
}
