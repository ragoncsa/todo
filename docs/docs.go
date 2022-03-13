// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks/": {
            "get": {
                "description": "Reads and returns all the tasks.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "read"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Task"
                            }
                        }
                    },
                    "default": {
                        "description": "unexpected error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new task.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "write"
                ],
                "summary": "Creates task",
                "parameters": [
                    {
                        "description": "New task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CreateTaskRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "the id of the caller",
                        "name": "CallerId",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Deletes all the tasks.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "write"
                ],
                "summary": "Delete all tasks",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/tasks/{taskid}": {
            "get": {
                "description": "Reads a single task and returns it.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "read"
                ],
                "summary": "Get task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Task"
                        }
                    },
                    "401": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "unexpected error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a single task.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "write"
                ],
                "summary": "Deletes task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Task": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "http.CreateTaskRequest": {
            "type": "object",
            "required": [
                "task"
            ],
            "properties": {
                "task": {
                    "$ref": "#/definitions/http.Task"
                }
            }
        },
        "http.Task": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "description of my-task-1"
                },
                "dueAt": {
                    "type": "string",
                    "example": "2019-10-12T07:20:50.52Z"
                },
                "name": {
                    "type": "string",
                    "example": "my-task-1"
                },
                "priority": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "userId": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tasks service",
	Description:      "This is a sample server that manages tasks.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
