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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Implementing a User API with a Cassandra backend",
    "title": "A User API with Cassandra backend",
    "version": "1.0.0"
  },
  "paths": {
    "/user": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "getAll",
        "responses": {
          "200": {
            "description": "List the users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "An error occured (GET /user)",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "createOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured (POST user)",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "getOne",
        "responses": {
          "200": {
            "description": "User with specific id",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured(GET /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "deleteOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "An error occured (DELETE /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "users"
        ],
        "operationId": "patchOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/patchDocument"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Patch user with specific id",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured (PATCH /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "string"
        },
        "status_code": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "patchDocument": {
      "type": "object",
      "properties": {
        "first_name": {
          "type": "string",
          "minLength": 3
        },
        "last_name": {
          "type": "string",
          "minLength": 3
        },
        "username": {
          "type": "string",
          "minLength": 6
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "username",
        "first_name",
        "last_name"
      ],
      "properties": {
        "first_name": {
          "type": "string",
          "minLength": 3
        },
        "id": {
          "type": "string",
          "readOnly": true
        },
        "last_name": {
          "type": "string",
          "minLength": 3
        },
        "username": {
          "type": "string",
          "minLength": 6
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Implementing a User API with a Cassandra backend",
    "title": "A User API with Cassandra backend",
    "version": "1.0.0"
  },
  "paths": {
    "/user": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "getAll",
        "responses": {
          "200": {
            "description": "List the users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "An error occured (GET /user)",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "createOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured (POST user)",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "getOne",
        "responses": {
          "200": {
            "description": "User with specific id",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured(GET /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "deleteOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "An error occured (DELETE /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "users"
        ],
        "operationId": "patchOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/patchDocument"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Patch user with specific id",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "An error occured (PATCH /user/{id})",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "string"
        },
        "status_code": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "patchDocument": {
      "type": "object",
      "properties": {
        "first_name": {
          "type": "string",
          "minLength": 3
        },
        "last_name": {
          "type": "string",
          "minLength": 3
        },
        "username": {
          "type": "string",
          "minLength": 6
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "username",
        "first_name",
        "last_name"
      ],
      "properties": {
        "first_name": {
          "type": "string",
          "minLength": 3
        },
        "id": {
          "type": "string",
          "readOnly": true
        },
        "last_name": {
          "type": "string",
          "minLength": 3
        },
        "username": {
          "type": "string",
          "minLength": 6
        }
      }
    }
  }
}`))
}
