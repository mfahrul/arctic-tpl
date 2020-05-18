package [[ with .ModuleToParse ]][[.Name]][[ end ]]

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"[[.Projectpath]]/config"
)

var responseMessage interface{} [[ with .ModuleToParse.Model ]]

//Endpoints struct 
type Endpoints struct {
	Create[[.Name | ToCamel]]  gin.HandlerFunc
	Get[[.Name | ToCamel]]ByID gin.HandlerFunc
	GetAll[[.Name | ToCamel | ToPlural]] gin.HandlerFunc
	Update[[.Name | ToCamel]]  gin.HandlerFunc
	Delete[[.Name | ToCamel]]  gin.HandlerFunc
}

//MakeEndpoints functions
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create[[.Name | ToCamel]]: makeCreate[[.Name | ToCamel]]Endpoint(s),
		Get[[.Name | ToCamel]]ByID: makeGet[[.Name | ToCamel]]ByIDEndpoint(s),
		GetAll[[.Name | ToCamel | ToPlural]]: makeGetAll[[.Name | ToCamel]]Endpoint(s),
		Update[[.Name | ToCamel]]: makeUpdate[[.Name | ToCamel]]Endpoint(s),
		Delete[[.Name | ToCamel]]: makeDelete[[.Name | ToCamel]]Endpoint(s),
	}
}

// @Summary Create a [[.Name | ToLower]]
// @Description Create [[.Name | ToLower]] for access
// @Accept  json
// @Produce  json
// @Param [[.Name | ToLower]] body Create[[.Name | ToCamel]]Req true "[[.Name | ToCamel]] Param"
// @Success 201 {object} Create[[.Name | ToCamel]]Response
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags [[.Name | ToLower]]
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /[[.Name | ToLower]] [post]
func makeCreate[[.Name | ToCamel]]Endpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var [[.Name | ToLower]] Create[[.Name | ToCamel]]Req
		errBind := c.ShouldBind(&[[.Name | ToLower]])
		if errBind != nil {
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   errBind.Error(),
			}
		} else {

			if _, err := s.Create[[.Name | ToCamel]]([[.Name | ToLower]]); err != nil {
				responseMessage = config.StatusInternalServerErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				}
			} else {
				responseMessage = Create[[.Name | ToCamel]]Response{
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

// @Summary Get all [[.Name | ToCamel | ToPlural | ToLower]]
// @Description Get all [[.Name | ToCamel | ToPlural | ToLower]]
// @Produce  json
// @Success 200 {object} GetAll[[.Name | ToCamel | ToPlural]]Response
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags [[.Name | ToLower]]
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /[[.Name | ToCamel | ToPlural | ToLower]] [get]
func makeGetAll[[.Name | ToCamel]]Endpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		[[.Name | ToCamel | ToPlural | ToLower]], err := s.GetAll[[.Name | ToCamel | ToPlural]]()
		if err != nil {
			responseMessage = config.StatusInternalServerErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			}
		} else {
			// if len([[.Name | ToCamel | ToPlural | ToLower]]) > 0 {
			responseMessage = GetAll[[.Name | ToCamel | ToPlural]]Response{
				StatusOKResponse: config.StatusOKResponse{
					Code:    http.StatusOK,
					Message: http.StatusText(http.StatusOK),
				},
				Data: [[.Name | ToCamel | ToPlural | ToLower]],
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

// @Summary Get a [[.Name | ToLower]]
// @Description Get a [[.Name | ToLower]] by ID
// @Produce  json
// @Param id path string true "[[.Name | ToCamel]] ID"
// @Success 200 {object} Get[[.Name | ToCamel]]Response
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 404 {object} config.StatusNotFoundResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags [[.Name | ToLower]]
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /[[.Name | ToLower]]/{id} [get]
func makeGet[[.Name | ToCamel]]ByIDEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Get[[.Name | ToCamel]]Request
		errBind := c.ShouldBindUri(&req)
		if errBind != nil {
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   errBind.Error(),
			}
		} else {
			[[.Name | ToLower]], err := s.Get[[.Name | ToCamel]]ByID(req.ID)
			if err != nil {
				responseMessage = config.StatusNotFoundResponse{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
					Error:   err.Error(),
				}
			} else {
				if [[.Name | ToLower]].ID != nil {
					responseMessage = Get[[.Name | ToCamel]]Response{
						StatusOKResponse: config.StatusOKResponse{
							Code:    http.StatusOK,
							Message: http.StatusText(http.StatusOK),
						},
						Data: [[.Name | ToLower]],
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

// @Summary Update a [[.Name | ToLower]]
// @Description Update [[.Name | ToLower]] by ID
// @Accept  json
// @Produce  json
// @Param id path string true "[[.Name | ToCamel]] ID"
// @Param [[.Name | ToLower]] body [[.Name | ToCamel]] true "[[.Name | ToCamel]] Param"
// @Success 200 {object} Get[[.Name | ToCamel]]Response
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 500 {object} config.StatusInternalServerErrorResponse
// @Tags [[.Name | ToLower]]
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /[[.Name | ToLower]]/{id} [put]
func makeUpdate[[.Name | ToCamel]]Endpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Get[[.Name | ToCamel]]Request
		var [[.Name | ToLower]] Update[[.Name | ToCamel]]Req
		errBind := c.ShouldBindUri(&req)
		errBind[[.Name | ToCamel]] := c.ShouldBind(&[[.Name | ToLower]])
		if errBind != nil || errBind[[.Name | ToCamel]] != nil {
			var err string
			if errBind != nil {
				err = errBind.Error()
			} else {
				err = errBind[[.Name | ToCamel]].Error()
			}
			responseMessage = config.StatusBadRequestResponse{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			}
		} else {
			[[.Name | ToCamel]], err := s.Update[[.Name | ToCamel]](req.ID, [[.Name | ToLower]])
			if err != nil {
				responseMessage = config.StatusInternalServerErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Error:   err.Error(),
				}
			} else {
				responseMessage = Get[[.Name | ToCamel]]Response{
					StatusOKResponse: config.StatusOKResponse{
						Code:    http.StatusOK,
						Message: http.StatusText(http.StatusOK),
					},
					Data: [[.Name | ToCamel]],
				}
			}
		}
		c.JSON(200, responseMessage)
	}
}

// @Summary Delete a [[.Name | ToLower]]
// @Description Delete [[.Name | ToLower]] by ID
// @Produce  json
// @Param id path string true "[[.Name | ToCamel]] ID"
// @Param UserID header string true "Delete by"
// @Success 200 {object} Delete[[.Name | ToCamel]]Response
// @Failure 400 {object} config.StatusBadRequestResponse
// @Failure 401 {object} config.StatusUnauthorizedResponse
// @Failure 404 {object} config.StatusNotFoundResponse
// @Tags [[.Name | ToLower]]
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /[[.Name | ToLower]]/{id} [delete]
func makeDelete[[.Name | ToCamel]]Endpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Get[[.Name | ToCamel]]Request
		var header Delete[[.Name | ToCamel]]Request
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
			[[.Name | ToLower]], err := s.Delete[[.Name | ToCamel]](req.ID, header.UserID)
			if err != nil {
				responseMessage = config.StatusNotFoundResponse{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
					Error:   err.Error(),
				}
			} else {
				responseMessage = Delete[[.Name | ToCamel]]Response{
					StatusOKResponse: config.StatusOKResponse{
						Code:    http.StatusOK,
						Message: http.StatusText(http.StatusOK),
					},
					Data: DeletedAt[[.Name | ToCamel]]{
						DeletedAt: [[.Name | ToLower]].DeletedAt,
					},
				}
			}
		}
		c.JSON(200, responseMessage)
	}
} [[ end ]]
