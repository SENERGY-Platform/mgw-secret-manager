// Package swagger_docs Code generated by swaggo/swag. DO NOT EDIT
package swagger_docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache-2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/confidential/value-variant": {
            "post": {
                "description": "Get the value of a secret.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Get secret value",
                "parameters": [
                    {
                        "description": "secret request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretVariantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "secret with value",
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretValueVariant"
                        }
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "Get basic service and runtime information.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get service info",
                "responses": {
                    "200": {
                        "description": "info",
                        "schema": {
                            "$ref": "#/definitions/lib.SrvInfo"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/key": {
            "post": {
                "description": "Set key for database entry encryption.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Encryption"
                ],
                "summary": "Set encryption key",
                "parameters": [
                    {
                        "description": "encryption key",
                        "name": "key",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/path-variant/clean": {
            "post": {
                "description": "Remove all secret files with the same reference.",
                "tags": [
                    "Secrets"
                ],
                "summary": "Delete secret files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "reference",
                        "name": "reference",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/path-variant/init": {
            "post": {
                "description": "Create a placeholder file for a secret.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Init secret file",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretVariantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "secret file info",
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretPathVariant"
                        }
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/path-variant/load": {
            "post": {
                "description": "Write secret value to file. File must be initialised first.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Write secret file",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretVariantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/path-variant/unload": {
            "delete": {
                "description": "Remove a secret file.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Delete secret file",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretVariantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/secrets": {
            "get": {
                "description": "List stored secrets.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Get Secrets",
                "responses": {
                    "200": {
                        "description": "secrets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_model.Secret"
                            }
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Store a secret.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Create secret",
                "parameters": [
                    {
                        "description": "secret data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "secret ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/secrets/{id}": {
            "get": {
                "description": "Get a secret.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Get secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "secret ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "secret",
                        "schema": {
                            "$ref": "#/definitions/api_model.Secret"
                        }
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a secret.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Update secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "secret ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "secret data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_model.SecretCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a secret.",
                "tags": [
                    "Secrets"
                ],
                "summary": "Delete secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "secret ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/types": {
            "get": {
                "description": "List supported secret types.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Get secret types",
                "responses": {
                    "200": {
                        "description": "types",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "additionalProperties": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_model.Secret": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "api_model.SecretCreateRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api_model.SecretPathVariant": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "item": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "api_model.SecretValueVariant": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "item": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api_model.SecretVariantRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "item": {
                    "type": "string"
                },
                "ref": {
                    "type": "string"
                }
            }
        },
        "lib.MemStats": {
            "type": "object",
            "properties": {
                "alloc": {
                    "type": "integer"
                },
                "alloc_total": {
                    "type": "integer"
                },
                "gc_cycles": {
                    "type": "integer"
                },
                "sys_total": {
                    "type": "integer"
                }
            }
        },
        "lib.SrvInfo": {
            "type": "object",
            "properties": {
                "mem_stats": {
                    "$ref": "#/definitions/lib.MemStats"
                },
                "name": {
                    "type": "string"
                },
                "up_time": {
                    "$ref": "#/definitions/time.Duration"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                -9223372036854775808,
                9223372036854775807,
                1,
                1000,
                1000000,
                1000000000,
                60000000000,
                3600000000000
            ],
            "x-enum-varnames": [
                "minDuration",
                "maxDuration",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Secret Manager API",
	Description:      "Provides access to secret management functions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
