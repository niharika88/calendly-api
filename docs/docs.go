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
        "/availability": {
            "get": {
                "description": "handles the retrieval of overall user availability across a range of dates, takes both day/date into account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Get availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-12-15",
                        "description": "Start Date",
                        "name": "startDate",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-12-15",
                        "description": "End Date",
                        "name": "endDate",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/UserDateAvailability"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/availability/date": {
            "post": {
                "description": "handles the creation of date-specific availability\nevery request overrides the existing availability for that date\ndate availability ALWAYS overrides the day availability",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Create date availability",
                "parameters": [
                    {
                        "description": "DateAvailabilityRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateDateAvailabilityRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/DateAvailability"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "handles the deletion of date-based availability",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Delete date availability",
                "parameters": [
                    {
                        "description": "DeleteUserAvailabilityRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DeleteUserAvailabilityRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/availability/day": {
            "post": {
                "description": "handles the creation of day-based availability\nevery request overrides the existing availability for all days\nif day is not provided, no availability is created for that day",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Create day availability",
                "parameters": [
                    {
                        "description": "DayAvailabilityRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateDayAvailabilityRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/DayAvailability"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "handles the deletion of day-based availability (` + "`" + `date` + "`" + ` param is ignored)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Delete day availability",
                "parameters": [
                    {
                        "description": "DeleteUserAvailabilityRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DeleteUserAvailabilityRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/availability/overlap": {
            "get": {
                "description": "handles the retrieval of schedule overlap between two users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "availability"
                ],
                "summary": "Get schedule overlap",
                "parameters": [
                    {
                        "type": "string",
                        "description": "First Username",
                        "name": "firstUsername",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Second Username",
                        "name": "secondUsername",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-12-15",
                        "description": "Start Date",
                        "name": "startDate",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-12-15",
                        "description": "End Date",
                        "name": "endDate",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/UserDateAvailability"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "healthcheck",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "handles the retrieval of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "post": {
                "description": "handles the creation of a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "handles the retrieval of a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "handles the update of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateUserRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "handles the deletion of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateDateAvailabilityRequest": {
            "type": "object",
            "required": [
                "date",
                "slots",
                "username"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2024-12-15T00:00:00Z"
                },
                "slots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Slot"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "CreateDayAvailabilityRequest": {
            "type": "object",
            "required": [
                "availability",
                "username"
            ],
            "properties": {
                "availability": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/UserDayAvailability"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "DateAvailability": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "slots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Slot"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "DayAvailability": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "day": {
                    "$ref": "#/definitions/github_com_niharika88_calendly-api_internal_db_models.Day"
                },
                "id": {
                    "type": "string"
                },
                "slots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Slot"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "DeleteUserAvailabilityRequest": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2024-12-15T00:00:00Z"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "$ref": "#/definitions/echo.HTTPError"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "Slot": {
            "type": "object",
            "properties": {
                "end": {
                    "description": "End time in minutes since midnight",
                    "type": "integer"
                },
                "start": {
                    "description": "Start time in minutes since midnight",
                    "type": "integer"
                }
            }
        },
        "UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                }
            }
        },
        "User": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "timezone": {
                    "description": "timezone for future use",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "UserDateAvailability": {
            "type": "object",
            "required": [
                "availability"
            ],
            "properties": {
                "availability": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/Slot"
                        }
                    }
                }
            }
        },
        "UserDayAvailability": {
            "type": "object",
            "required": [
                "day",
                "slots"
            ],
            "properties": {
                "day": {
                    "$ref": "#/definitions/github_com_niharika88_calendly-api_internal_db_models.Day"
                },
                "slots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Slot"
                    }
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "github_com_niharika88_calendly-api_internal_db_models.Day": {
            "type": "string",
            "enum": [
                "monday",
                "tuesday",
                "wednesday",
                "thursday",
                "friday",
                "saturday",
                "sunday"
            ],
            "x-enum-varnames": [
                "DayMonday",
                "DayTuesday",
                "DayWednesday",
                "DayThursday",
                "DayFriday",
                "DaySaturday",
                "DaySunday"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Calendly API",
	Description:      "Calendly clone",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
