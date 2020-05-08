package config

type (

	//StatusUnauthorizedResponse struct
	StatusUnauthorizedResponse struct {
		Message string `json:"message" example:"Unauthorized/No API key found in request"`
	}

	//StatusNotFoundResponse struct
	StatusNotFoundResponse struct {
		Code    int    `json:"code" example:"404"`
		Message string `json:"message" example:"Not Found"`
		Error   string `json:"error,omitempty"`
	}

	//StatusBadRequestResponse struct
	StatusBadRequestResponse struct {
		Code    int    `json:"code" example:"400"`
		Message string `json:"message" example:"Bad Request"`
		Error   string `json:"error,omitempty"`
	}

	//StatusInternalServerErrorResponse struct
	StatusInternalServerErrorResponse struct {
		Code    int    `json:"code" example:"500"`
		Message string `json:"message" example:"error"`
		Error   string `json:"error,omitempty"`
	}

	//StatusNotImplementedResponse struct
	StatusNotImplementedResponse struct {
		Code    int    `json:"code" example:"501"`
		Message string `json:"message" example:"error"`
		Error   string `json:"error,omitempty"`
	}

	//StatusOKResponse struct
	StatusOKResponse struct {
		Code    int    `json:"code" example:"200"`
		Message string `json:"message" example:"Ok"`
	}

	//StatusCreatedResponse struct
	StatusCreatedResponse struct {
		Code    int    `json:"code" example:"201"`
		Message string `json:"message" example:"Created"`
	}
)
