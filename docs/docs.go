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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/post": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子模块"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少条数据",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/post/{post_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子模块"
                ],
                "summary": "获取帖子详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "帖子id",
                        "name": "post_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {}
            }
        },
        "dto.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterDTO": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 0
                },
                "email": {
                    "description": "Email      string ` + "`" + `json:\"email\" binding:\"required,email\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "re_password": {
                    "description": "跨字段验证",
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "minLength": 3
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
