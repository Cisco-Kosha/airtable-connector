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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "eti@cisco.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v0/meta/bases": {
            "get": {
                "description": "Get Records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get airtable",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BasesStruct"
                        }
                    }
                }
            }
        },
        "/api/v0/specification/list": {
            "get": {
                "description": "Get all environment variables that need to be supplied",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "specification"
                ],
                "summary": "Get details of the connector specification",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "permission denied",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v0/{baseId}/{tableIdOrName}/{recordId}": {
            "get": {
                "description": "Get Records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get airtable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Base Id",
                        "name": "baseId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Table Id",
                        "name": "tableIdOrName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Base Id",
                        "name": "recordId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    }
                }
            }
        },
        "/api/v0/{baseId}/{tableIdOrName}/{recordId}/comments": {
            "get": {
                "description": "Get Comments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get airtable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Base Id",
                        "name": "baseId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Table Id",
                        "name": "tableIdOrName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Base Id",
                        "name": "recordId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommentStruct"
                        }
                    }
                }
            }
        },
        "/api/v0/{baseId}/{tableId}": {
            "get": {
                "description": "Get Records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get airtable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Base Id",
                        "name": "baseId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Table Id",
                        "name": "tableId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RecordsStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BasesStruct": {
            "type": "object",
            "properties": {
                "bases": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "string"
                            },
                            "name": {
                                "type": "string"
                            },
                            "permissionLevel": {
                                "type": "string"
                            }
                        }
                    }
                },
                "offset": {
                    "type": "string"
                }
            }
        },
        "models.CommentStruct": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "author": {
                                "type": "object",
                                "properties": {
                                    "email": {
                                        "type": "string"
                                    },
                                    "id": {
                                        "type": "string"
                                    },
                                    "name": {
                                        "type": "string"
                                    }
                                }
                            },
                            "createdTime": {
                                "type": "string"
                            },
                            "id": {
                                "type": "string"
                            },
                            "lastUpdatedTime": {},
                            "mentioned": {
                                "type": "object",
                                "properties": {
                                    "usrL2PNC5o3H4lBEi": {
                                        "type": "object",
                                        "properties": {
                                            "displayName": {
                                                "type": "string"
                                            },
                                            "email": {
                                                "type": "string"
                                            },
                                            "id": {
                                                "type": "string"
                                            },
                                            "type": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            },
                            "text": {
                                "type": "string"
                            }
                        }
                    }
                },
                "offset": {}
            }
        },
        "models.Record": {
            "type": "object",
            "properties": {
                "createdTime": {
                    "type": "string"
                },
                "fields": {
                    "type": "object",
                    "properties": {
                        "End date": {
                            "type": "string"
                        },
                        "Field 3": {
                            "type": "string"
                        },
                        "Field 4": {
                            "type": "string"
                        },
                        "Projects": {
                            "type": "string"
                        }
                    }
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.RecordsStruct": {
            "type": "object",
            "properties": {
                "records": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "createdTime": {
                                "type": "string"
                            },
                            "fields": {
                                "type": "object",
                                "properties": {
                                    "End date": {
                                        "type": "string"
                                    },
                                    "Field 3": {
                                        "type": "string"
                                    },
                                    "Field 4": {
                                        "type": "string"
                                    },
                                    "Projects": {
                                        "type": "string"
                                    }
                                }
                            },
                            "id": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8012",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Airtable Connector API",
	Description:      "This is a Kosha REST serice for exposing many freshservice features as REST APIs with better consistency, observability etc",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
