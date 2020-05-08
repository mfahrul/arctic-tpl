// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-16 12:17:40.523685 +0700 WIB m=+0.080358171

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Giftano API Support",
            "url": "http://giftano.io",
            "email": "fahrul@giftano.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/item": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create item for access",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Create a item",
                "parameters": [
                    {
                        "description": "Item Param",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.CreateItemReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/core.CreateItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.StatusBadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.StatusUnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.StatusInternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/item/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get a item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Get a item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.GetItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.StatusBadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.StatusUnauthorizedResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.StatusNotFoundResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.StatusInternalServerErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Update a item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item Param",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.GetItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.StatusBadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.StatusUnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.StatusInternalServerErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Delete a item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Delete by",
                        "name": "user_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.DeleteItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.StatusBadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.StatusUnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.StatusInternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/items": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get all items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Get all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.GetAllItemsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.StatusUnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.StatusInternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.StatusBadRequestResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "config.StatusInternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "config.StatusNotFoundResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Not Found"
                }
            }
        },
        "config.StatusUnauthorizedResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Unauthorized/No API key found in request"
                }
            }
        },
        "core.CreateItemReq": {
            "type": "object",
            "required": [
                "item_name"
            ],
            "properties": {
                "additional_info": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_by": {
                    "type": "string"
                },
                "item_desc": {
                    "type": "string"
                },
                "item_name": {
                    "type": "string"
                }
            }
        },
        "core.CreateItemResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 201
                },
                "message": {
                    "type": "string",
                    "example": "Created"
                }
            }
        },
        "core.DeleteItemResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/core.DeletedAtItem"
                },
                "message": {
                    "type": "string",
                    "example": "Ok"
                }
            }
        },
        "core.DeletedAtItem": {
            "type": "object",
            "properties": {
                "deleted_at": {
                    "type": "string"
                }
            }
        },
        "core.GetAllItemsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.Item"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Ok"
                }
            }
        },
        "core.GetItemResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/core.Item"
                },
                "message": {
                    "type": "string",
                    "example": "Ok"
                }
            }
        },
        "core.Item": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "item_desc": {
                    "type": "string"
                },
                "item_name": {
                    "type": "string"
                },
                "status": {
                    "description": "0, 1, 9 =\u003e notactive, activate, ban",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "apikey",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.1",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Giftano Core API Docs",
	Description: "Dashboard users management service.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
