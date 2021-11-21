// Code generated by go-swagger; DO NOT EDIT.

package handler

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
    "title": "API Telegram bot",
    "version": "1.0.0"
  },
  "host": "*",
  "paths": {
    "/send": {
      "post": {
        "summary": "Send message to all subscribers",
        "parameters": [
          {
            "type": "string",
            "description": "This message will be send all subscribers",
            "name": "msg",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal Server Error"
          }
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
    "title": "API Telegram bot",
    "version": "1.0.0"
  },
  "host": "*",
  "paths": {
    "/send": {
      "post": {
        "summary": "Send message to all subscribers",
        "parameters": [
          {
            "type": "string",
            "description": "This message will be send all subscribers",
            "name": "msg",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  }
}`))
}
