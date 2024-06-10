// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
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
        "/api/users/{userid}": {
            "get": {
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
                "summary": "get user by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
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
            }
        }
    },
    "definitions": {
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "cause": {
                    "type": "string",
                    "default": "Something went wrong"
                },
                "code": {
                    "type": "integer",
                    "default": 500
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
        "ApiKeyAuth": {
            "description": "Jwt Bearer Token",
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
	Schemes:          []string{},
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
