// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/email/authCode/": {
            "post": {
                "description": "发送验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "邮箱模块"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    }
                }
            }
        },
        "/user/login/": {
            "post": {
                "description": "已注册用户可通过登录进入app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录所需信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    }
                }
            }
        },
        "/user/register/": {
            "post": {
                "description": "用户注册入口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "注册所需信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    }
                }
            }
        },
        "/video/clip/": {
            "post": {
                "description": "用于剪辑视频",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "*/*"
                ],
                "tags": [
                    "视频模块"
                ],
                "summary": "视频剪辑",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token值",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "剪辑一段视频所需信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VideoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resp_msg.RespMsg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.RegisterDto": {
            "type": "object",
            "properties": {
                "auth_code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.VideoDto": {
            "type": "object",
            "properties": {
                "end_time": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "resp_msg.RespMsg": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "any"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "120.79.155.59:9997",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "视频剪辑项目接口文档",
	Description:      "author: 林叶润",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
