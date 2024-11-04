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
        "/api/dashboard/charts": {
            "post": {
                "description": "Получение графиков",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Получение графиков",
                "parameters": [
                    {
                        "description": "Фильтр",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Filter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ChartReport"
                        }
                    }
                }
            }
        },
        "/api/session": {
            "post": {
                "description": "Добавление сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Добавление сессии",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Session"
                        }
                    }
                }
            }
        },
        "/api/session/list": {
            "get": {
                "description": "Получение списка сессий",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Получение списка сессий",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Session"
                            }
                        }
                    }
                }
            }
        },
        "/api/session/{session_id}": {
            "get": {
                "description": "Получение контента сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Получение контента сессии",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID сессии в формате UUID",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.SessionContentDto"
                            }
                        }
                    }
                }
            }
        },
        "/api/session/{session_id}/{title}": {
            "put": {
                "description": "Обновление названия сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Обновление названия сессии",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID сессии в формате UUID",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Название сессии",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Логин пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Запрос на логин",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TokenResp"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Регистрация пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "Запрос на регистрацию",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignupReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.TokenResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Chart": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "records": {
                    "description": "для отрисовки графа",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Record"
                    }
                },
                "type": {
                    "description": "пока что bar chart",
                    "type": "string"
                }
            }
        },
        "model.ChartReport": {
            "type": "object",
            "properties": {
                "charts": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/model.Chart"
                        }
                    }
                },
                "endDate": {
                    "type": "integer"
                },
                "legend": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "multipliers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Multiplier"
                    }
                },
                "startDate": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "model.Filter": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "model.LoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Multiplier": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "model.QueryDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "prompt": {
                    "type": "string"
                }
            }
        },
        "model.Record": {
            "type": "object",
            "properties": {
                "x": {
                    "description": "формат: квартал - год",
                    "type": "string"
                },
                "y": {
                    "type": "number"
                }
            }
        },
        "model.ResponseDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "description": "llm response",
                    "type": "string"
                },
                "error": {},
                "filenames": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isLast": {
                    "type": "boolean"
                },
                "queryId": {
                    "type": "integer"
                },
                "sources": {
                    "description": "s3 link",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Session": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.SessionContentDto": {
            "type": "object",
            "properties": {
                "query": {
                    "$ref": "#/definitions/model.QueryDto"
                },
                "response": {
                    "$ref": "#/definitions/model.ResponseDto"
                }
            }
        },
        "model.SignupReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.TokenResp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9999",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MTS AI Docs",
	Description:      "Документация для сервиса решения команды MISIS Banach Space к задаче MTS AI Docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}