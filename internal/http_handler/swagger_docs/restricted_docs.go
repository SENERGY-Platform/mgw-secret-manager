// Package swagger_docs Code generated by swaggo/swag. DO NOT EDIT
package swagger_docs

import "github.com/swaggo/swag"

const docTemplaterestricted = `{
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
                3600000000000,
                1,
                1000,
                1000000,
                1000000000
            ],
            "x-enum-varnames": [
                "minDuration",
                "maxDuration",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second"
            ]
        }
    }
}`

// SwaggerInforestricted holds exported Swagger Info so clients can modify it
var SwaggerInforestricted = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/restricted",
	Schemes:          []string{},
	Title:            "Secret Manager Restricted API",
	Description:      "Provides access to secret management functions.",
	InfoInstanceName: "restricted",
	SwaggerTemplate:  docTemplaterestricted,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInforestricted.InstanceName(), SwaggerInforestricted)
}
