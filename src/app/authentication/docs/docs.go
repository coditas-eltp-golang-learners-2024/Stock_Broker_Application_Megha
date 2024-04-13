// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/signIn": {
            "post": {
                "description": "Sign in with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "User email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful sign-in"
                    },
                    "400": {
                        "description": "Invalid request payload"
                    },
                    "401": {
                        "description": "Invalid credentials"
                    }
                }
            }
        },
        "/signUp": {
            "post": {
                "description": "Sign up with user details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Sign-up request payload",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful sign-up",
                        "schema": {
                            "$ref": "#/definitions/models.SignUpRequest"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "pancard_number",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pancard_number": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "phone_number": {
                    "type": "string",
                    "maxLength": 9999999999,
                    "minLength": 0
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}