///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Function Manager\n",
    "title": "Function Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/function": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "List all existing functions",
        "operationId": "getFunctions",
        "parameters": [
          {
            "type": "string",
            "description": "Function state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter based on tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "./models.json#/definitions/Function"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "default": {
            "description": "Custom error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Add a new function",
        "operationId": "addFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Function created",
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized Request",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/orgIDParam"
        }
      ]
    },
    "/function/{functionName}": {
      "get": {
        "description": "Returns a single function",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Find function by Name",
        "operationId": "getFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Update a function",
        "operationId": "updateFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Deletes a function",
        "operationId": "deleteFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/orgIDParam"
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to work on",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/runs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function runs that are being executed",
        "operationId": "getRuns",
        "responses": {
          "200": {
            "description": "List of function runs",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "./models.json#/definitions/Run"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Run a function",
        "operationId": "runFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "./models.json#/definitions/Run"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful execution (blocking call)",
            "schema": {
              "$ref": "./models.json#/definitions/Run"
            }
          },
          "202": {
            "description": "Execution started (non-blocking call)",
            "schema": {
              "$ref": "./models.json#/definitions/Run"
            }
          },
          "400": {
            "description": "User error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "422": {
            "description": "Input object validation failed",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "502": {
            "description": "Function error occurred (blocking call)",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/orgIDParam"
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to run or retreive runs for",
          "name": "functionName",
          "in": "query"
        }
      ]
    },
    "/runs/{runName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function run by its name",
        "operationId": "getRun",
        "responses": {
          "200": {
            "description": "Function Run",
            "schema": {
              "$ref": "./models.json#/definitions/Run"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Function or Run not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/orgIDParam"
        },
        {
          "type": "string",
          "format": "uuid",
          "description": "name of run to retrieve",
          "name": "runName",
          "in": "path",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to retreive a run for",
          "name": "functionName",
          "in": "query"
        }
      ]
    }
  },
  "parameters": {
    "orgIDParam": {
      "type": "string",
      "name": "X-Dispatch-Org",
      "in": "header",
      "required": true
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Crud operations on functions",
      "name": "Store"
    },
    {
      "description": "Execution operations on functions",
      "name": "Runner"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Function Manager\n",
    "title": "Function Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/function": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "List all existing functions",
        "operationId": "getFunctions",
        "parameters": [
          {
            "type": "string",
            "description": "Function state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter based on tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/function"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Custom error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Add a new function",
        "operationId": "addFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/function"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Function created",
            "schema": {
              "$ref": "#/definitions/function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "X-Dispatch-Org",
          "in": "header",
          "required": true
        }
      ]
    },
    "/function/{functionName}": {
      "get": {
        "description": "Returns a single function",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Find function by Name",
        "operationId": "getFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Update a function",
        "operationId": "updateFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Deletes a function",
        "operationId": "deleteFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "X-Dispatch-Org",
          "in": "header",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to work on",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/runs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function runs that are being executed",
        "operationId": "getRuns",
        "responses": {
          "200": {
            "description": "List of function runs",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/run"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Run a function",
        "operationId": "runFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/run"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful execution (blocking call)",
            "schema": {
              "$ref": "#/definitions/run"
            }
          },
          "202": {
            "description": "Execution started (non-blocking call)",
            "schema": {
              "$ref": "#/definitions/run"
            }
          },
          "400": {
            "description": "User error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "422": {
            "description": "Input object validation failed",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "502": {
            "description": "Function error occurred (blocking call)",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "X-Dispatch-Org",
          "in": "header",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to run or retreive runs for",
          "name": "functionName",
          "in": "query"
        }
      ]
    },
    "/runs/{runName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function run by its name",
        "operationId": "getRun",
        "responses": {
          "200": {
            "description": "Function Run",
            "schema": {
              "$ref": "#/definitions/run"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Function or Run not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "X-Dispatch-Org",
          "in": "header",
          "required": true
        },
        {
          "type": "string",
          "format": "uuid",
          "description": "name of run to retrieve",
          "name": "runName",
          "in": "path",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to retreive a run for",
          "name": "functionName",
          "in": "query"
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "description": "Error error",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "code",
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "function": {
      "description": "Function function",
      "type": "object",
      "required": [
        "source",
        "image",
        "handler",
        "name"
      ],
      "properties": {
        "createdTime": {
          "description": "created time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "CreatedTime"
        },
        "faasId": {
          "description": "faas Id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "FaasID"
        },
        "functionImageURL": {
          "description": "functionImageURL",
          "type": "string",
          "x-go-name": "FunctionImageURL"
        },
        "handler": {
          "description": "handler",
          "type": "string",
          "x-go-name": "Handler"
        },
        "id": {
          "description": "id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "ID"
        },
        "image": {
          "description": "image",
          "type": "string",
          "x-go-name": "Image"
        },
        "kind": {
          "description": "kind",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Kind",
          "readOnly": true
        },
        "modifiedTime": {
          "description": "modified time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "ModifiedTime"
        },
        "name": {
          "description": "name",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Name"
        },
        "schema": {
          "$ref": "#/definitions/functionSchema"
        },
        "secrets": {
          "description": "secrets",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Secrets"
        },
        "services": {
          "description": "services",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Services"
        },
        "source": {
          "description": "source",
          "type": "string",
          "format": "byte",
          "x-go-name": "Source"
        },
        "sourcePath": {
          "description": "only used in seed.yaml",
          "type": "string",
          "x-go-name": "SourcePath"
        },
        "status": {
          "description": "Status status",
          "type": "string",
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        },
        "tags": {
          "description": "tags",
          "type": "array",
          "items": {
            "$ref": "#/definitions/functionTagsItems"
          },
          "x-go-name": "Tags"
        },
        "timeout": {
          "description": "timeout",
          "type": "integer",
          "format": "int64",
          "x-go-name": "Timeout"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "functionSchema": {
      "description": "Schema schema",
      "type": "object",
      "properties": {
        "in": {
          "description": "in",
          "type": "object",
          "x-go-name": "In"
        },
        "out": {
          "description": "out",
          "type": "object",
          "x-go-name": "Out"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "functionTagsItems": {
      "description": "Tag tag",
      "type": "object",
      "properties": {
        "key": {
          "description": "key",
          "type": "string",
          "x-go-name": "Key"
        },
        "value": {
          "description": "value",
          "type": "string",
          "x-go-name": "Value"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "run": {
      "description": "Run run",
      "type": "object",
      "properties": {
        "blocking": {
          "description": "blocking",
          "type": "boolean",
          "x-go-name": "Blocking"
        },
        "error": {
          "$ref": "#/definitions/runError"
        },
        "event": {
          "$ref": "#/definitions/runEvent"
        },
        "executedTime": {
          "description": "executed time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "ExecutedTime",
          "readOnly": true
        },
        "faasId": {
          "description": "faas Id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "FaasID"
        },
        "finishedTime": {
          "description": "finished time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "FinishedTime",
          "readOnly": true
        },
        "functionId": {
          "description": "function Id",
          "type": "string",
          "x-go-name": "FunctionID",
          "readOnly": true
        },
        "functionName": {
          "description": "function name",
          "type": "string",
          "x-go-name": "FunctionName",
          "readOnly": true
        },
        "httpContext": {
          "description": "http context",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          },
          "x-go-name": "HTTPContext",
          "readOnly": true
        },
        "input": {
          "description": "input",
          "type": "object",
          "x-go-name": "Input"
        },
        "logs": {
          "$ref": "#/definitions/runLogs"
        },
        "name": {
          "description": "name",
          "type": "string",
          "format": "uuid",
          "x-go-name": "Name",
          "readOnly": true
        },
        "output": {
          "description": "output",
          "type": "object",
          "x-go-name": "Output",
          "readOnly": true
        },
        "reason": {
          "description": "reason",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Reason"
        },
        "secrets": {
          "description": "secrets",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Secrets"
        },
        "services": {
          "description": "services",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Services"
        },
        "status": {
          "description": "Status status",
          "type": "string",
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        },
        "tags": {
          "description": "tags",
          "type": "array",
          "items": {
            "$ref": "#/definitions/runTagsItems"
          },
          "x-go-name": "Tags"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "runError": {
      "description": "InvocationError invocation error",
      "type": "object",
      "required": [
        "message",
        "type"
      ],
      "properties": {
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        },
        "stacktrace": {
          "description": "stacktrace",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Stacktrace"
        },
        "type": {
          "description": "ErrorType error type",
          "type": "string",
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "runEvent": {
      "description": "CloudEvent cloud event",
      "type": "object",
      "required": [
        "cloud-events-version",
        "event-id",
        "event-type",
        "namespace",
        "source-id",
        "source-type"
      ],
      "properties": {
        "cloud-events-version": {
          "description": "cloud events version",
          "type": "string",
          "x-go-name": "CloudEventsVersion"
        },
        "content-type": {
          "description": "content type",
          "type": "string",
          "x-go-name": "ContentType"
        },
        "data": {
          "description": "data",
          "type": "string",
          "x-go-name": "Data"
        },
        "event-id": {
          "description": "event id",
          "type": "string",
          "x-go-name": "EventID"
        },
        "event-time": {
          "description": "event time",
          "type": "string",
          "format": "date-time",
          "x-go-name": "EventTime"
        },
        "event-type": {
          "description": "event type",
          "type": "string",
          "maxLength": 128,
          "pattern": "^[\\w\\d\\-\\.]+$",
          "x-go-name": "EventType"
        },
        "event-type-version": {
          "description": "event type version",
          "type": "string",
          "x-go-name": "EventTypeVersion"
        },
        "extensions": {
          "description": "extensions",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          },
          "x-go-name": "Extensions"
        },
        "namespace": {
          "description": "namespace",
          "type": "string",
          "x-go-name": "Namespace"
        },
        "schema-url": {
          "description": "schema url",
          "type": "string",
          "x-go-name": "SchemaURL"
        },
        "source-id": {
          "description": "source id",
          "type": "string",
          "x-go-name": "SourceID"
        },
        "source-type": {
          "description": "source type",
          "type": "string",
          "x-go-name": "SourceType"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "runLogs": {
      "description": "Logs logs",
      "type": "object",
      "required": [
        "stderr",
        "stdout"
      ],
      "properties": {
        "stderr": {
          "description": "stderr",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Stderr"
        },
        "stdout": {
          "description": "stdout",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Stdout"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "runTagsItems": {
      "description": "Tag tag",
      "type": "object",
      "properties": {
        "key": {
          "description": "key",
          "type": "string",
          "x-go-name": "Key"
        },
        "value": {
          "description": "value",
          "type": "string",
          "x-go-name": "Value"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    }
  },
  "parameters": {
    "orgIDParam": {
      "type": "string",
      "name": "X-Dispatch-Org",
      "in": "header",
      "required": true
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Crud operations on functions",
      "name": "Store"
    },
    {
      "description": "Execution operations on functions",
      "name": "Runner"
    }
  ]
}`))
}
