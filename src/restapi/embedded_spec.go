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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test Task",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "the app is healthy"
          }
        }
      }
    },
    "/state": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "postState",
        "parameters": [
          {
            "name": "stateObj",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/stateObj"
            }
          },
          {
            "enum": [
              "game",
              "server",
              "payment"
            ],
            "type": "string",
            "name": "Source-Type",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "state updates",
            "schema": {
              "$ref": "#/definitions/stateObj"
            }
          },
          "default": {
            "description": "error message",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "reason": {
          "type": "string"
        }
      }
    },
    "stateObj": {
      "type": "object",
      "required": [
        "state",
        "transactionId",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "string"
        },
        "state": {
          "type": "string",
          "enum": [
            "win",
            "lost"
          ]
        },
        "transactionId": {
          "type": "string",
          "format": "uuid"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test Task",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "the app is healthy"
          }
        }
      }
    },
    "/state": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "postState",
        "parameters": [
          {
            "name": "stateObj",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/stateObj"
            }
          },
          {
            "enum": [
              "game",
              "server",
              "payment"
            ],
            "type": "string",
            "name": "Source-Type",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "state updates",
            "schema": {
              "$ref": "#/definitions/stateObj"
            }
          },
          "default": {
            "description": "error message",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "reason": {
          "type": "string"
        }
      }
    },
    "stateObj": {
      "type": "object",
      "required": [
        "state",
        "transactionId",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "string"
        },
        "state": {
          "type": "string",
          "enum": [
            "win",
            "lost"
          ]
        },
        "transactionId": {
          "type": "string",
          "format": "uuid"
        }
      }
    }
  }
}`))
}
