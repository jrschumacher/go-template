// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Delete an item from the store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Delete an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of data to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.StoreDeleteResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fetch/{id}": {
            "get": {
                "description": "Fetch an item from the store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Fetch an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of data to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.StoreSingleResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get health of service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "utils"
                ],
                "summary": "Health endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.HealthResult"
                            }
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Search for data in store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Search for data in store",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.StoreMultiResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Write data to store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Update data in store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of data to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Data to store",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.StoreUpdateResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/write": {
            "post": {
                "description": "Write data to store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Write to store",
                "parameters": [
                    {
                        "description": "Data to store",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.StoreSingleResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.HealthResult": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "api.VersionStat": {
            "type": "object",
            "properties": {
                "buildTime": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                },
                "versionLong": {
                    "type": "string"
                }
            }
        },
        "store.StoreDeleteResult": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "prevData": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "store.StoreMultiResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/store.StoreSingleResult"
                    }
                }
            }
        },
        "store.StoreSingleResult": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "store.StoreUpdateResult": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "prevData": {
                    "type": "object",
                    "additionalProperties": true
                },
                "replaced": {
                    "type": "boolean"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go Template API",
	Description: "A go template project",
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
