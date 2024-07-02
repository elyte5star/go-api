// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Elyte Fiber Application.",
            "url": "https://github.com/elyte5star.",
            "email": "elyte5star@gmail.com"
        },
        "license": {
            "name": "Proprietary",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login": {
            "post": {
                "description": "Create a new bearer token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create a new bearer token",
                "parameters": [
                    {
                        "description": "Login data",
                        "name": "credential",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        },
        "/api/products": {
            "get": {
                "description": "Get all existing products.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all existing products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.RequestResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/products/create": {
            "post": {
                "description": "Create a new product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Create product",
                        "name": "create_product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        },
        "/api/products/create/review": {
            "post": {
                "description": "Create a new product review.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product review",
                "parameters": [
                    {
                        "description": "Create a product review",
                        "name": "product_review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateProductReviewRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        },
        "/api/products/{pid}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Product by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get product by given pid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pid",
                        "name": "pid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete product by a given pid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete Product by given pid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pid",
                        "name": "pid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        },
        "/api/status": {
            "get": {
                "description": "API status check",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all existing users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all existing users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.RequestResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/signup": {
            "post": {
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "create_user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{userid}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get User by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by given userid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update User.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Modify User",
                        "name": "modify_user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ModifyUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete user by a given userid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user by given userid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RequestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateAddressReq": {
            "type": "object",
            "required": [
                "country",
                "fullName",
                "state",
                "streetAddress",
                "zip"
            ],
            "properties": {
                "country": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "streetAddress": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "request.CreateProductRequest": {
            "type": "object",
            "required": [
                "category",
                "description",
                "details",
                "image",
                "name",
                "price",
                "productDiscount",
                "stockQuantity"
            ],
            "properties": {
                "category": {
                    "type": "string",
                    "maxLength": 255
                },
                "description": {
                    "type": "string",
                    "maxLength": 555
                },
                "details": {
                    "type": "string",
                    "maxLength": 555
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "productDiscount": {
                    "type": "number"
                },
                "stockQuantity": {
                    "type": "integer"
                }
            }
        },
        "request.CreateProductReviewRequest": {
            "type": "object",
            "required": [
                "comment",
                "email",
                "pid",
                "reviewerName"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "pid": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "reviewerName": {
                    "type": "string"
                }
            }
        },
        "request.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "telephone"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "telephone": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 5
                }
            }
        },
        "request.ModifyUser": {
            "type": "object",
            "required": [
                "address"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/request.CreateAddressReq"
                },
                "telephone": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "default": 500
                },
                "message": {
                    "type": "string",
                    "default": "Something went wrong"
                },
                "success": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "response.RequestResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "default": 200
                },
                "message": {
                    "type": "string",
                    "default": "Operation Successful!"
                },
                "path": {
                    "type": "string",
                    "default": "0"
                },
                "result": {},
                "success": {
                    "type": "boolean",
                    "default": true
                },
                "timeStamp": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Bearer Token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Elyte Realm API",
	Description:      "Interactive Documentation for Elyte-Realm API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
