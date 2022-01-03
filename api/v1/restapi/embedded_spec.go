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
    "description": "bpflock",
    "title": "bpflock API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1beta"
  },
  "basePath": "/v1",
  "paths": {
    "/config": {
      "get": {
        "description": "Returns the configuration of the bpflock daemon.",
        "tags": [
          "daemon"
        ],
        "summary": "Get configuration of bpflock daemon",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DaemonConfiguration"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "description": "Returns health and status information of the bpflock daemon.",
        "tags": [
          "daemon"
        ],
        "summary": "Get health of bpflock daemon",
        "parameters": [
          {
            "type": "boolean",
            "x-exportParamName": "Brief",
            "x-optionalDataType": "Bool",
            "description": "Brief will return a brief representation of the bpflock status.",
            "name": "brief",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ConfigurationMap": {
      "description": "Map of configuration key/value pairs.",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "DaemonConfiguration": {
      "description": "Response to a daemon configuration request.",
      "type": "object",
      "properties": {
        "spec": {
          "$ref": "#/definitions/DaemonConfigurationSpec"
        },
        "status": {
          "$ref": "#/definitions/DaemonConfigurationStatus"
        }
      },
      "example": {
        "spec": {
          "options": {}
        },
        "status": {
          "applied": {
            "options": {}
          },
          "daemonConfigurationMap": ""
        }
      }
    },
    "DaemonConfigurationSpec": {
      "description": "The controllable and changeable configuration of the daemon.",
      "type": "object",
      "properties": {
        "options": {
          "$ref": "#/definitions/ConfigurationMap"
        }
      },
      "example": {
        "options": {}
      }
    },
    "DaemonConfigurationStatus": {
      "description": "Response to a daemon configuration request.",
      "type": "object",
      "properties": {
        "applied": {
          "$ref": "#/definitions/DaemonConfigurationSpec"
        },
        "daemonConfigurationMap": {
          "description": "Config map which contains all the active daemon configurations"
        },
        "immutable": {
          "$ref": "#/definitions/ConfigurationMap"
        }
      },
      "example": {
        "applied": {
          "options": {}
        },
        "daemonConfigurationMap": ""
      }
    },
    "Status": {
      "description": "Status of an individual component",
      "type": "object",
      "properties": {
        "msg": {
          "description": "Human readable status/error/warning message",
          "type": "string"
        },
        "state": {
          "description": "State the component is in",
          "type": "string",
          "enum": [
            "Ok",
            "Warning",
            "Failure",
            "Disabled"
          ]
        }
      },
      "example": {
        "msg": "msg",
        "state": "Ok"
      }
    },
    "StatusResponse": {
      "description": "Health and status information of daemon",
      "type": "object",
      "properties": {
        "bpflock": {
          "$ref": "#/definitions/Status"
        },
        "stale": {
          "description": "List of stale information in the status",
          "type": "object",
          "additionalProperties": {
            "description": "Timestamp when the probe was started",
            "type": "string",
            "format": "date-time"
          }
        }
      },
      "example": {
        "bpflock": {
          "msg": "msg",
          "state": "Ok"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
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
    "description": "bpflock",
    "title": "bpflock API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1beta"
  },
  "basePath": "/v1",
  "paths": {
    "/config": {
      "get": {
        "description": "Returns the configuration of the bpflock daemon.",
        "tags": [
          "daemon"
        ],
        "summary": "Get configuration of bpflock daemon",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DaemonConfiguration"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "description": "Returns health and status information of the bpflock daemon.",
        "tags": [
          "daemon"
        ],
        "summary": "Get health of bpflock daemon",
        "parameters": [
          {
            "type": "boolean",
            "x-exportParamName": "Brief",
            "x-optionalDataType": "Bool",
            "description": "Brief will return a brief representation of the bpflock status.",
            "name": "brief",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ConfigurationMap": {
      "description": "Map of configuration key/value pairs.",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "DaemonConfiguration": {
      "description": "Response to a daemon configuration request.",
      "type": "object",
      "properties": {
        "spec": {
          "$ref": "#/definitions/DaemonConfigurationSpec"
        },
        "status": {
          "$ref": "#/definitions/DaemonConfigurationStatus"
        }
      },
      "example": {
        "spec": {
          "options": {}
        },
        "status": {
          "applied": {
            "options": {}
          },
          "daemonConfigurationMap": ""
        }
      }
    },
    "DaemonConfigurationSpec": {
      "description": "The controllable and changeable configuration of the daemon.",
      "type": "object",
      "properties": {
        "options": {
          "$ref": "#/definitions/ConfigurationMap"
        }
      },
      "example": {
        "options": {}
      }
    },
    "DaemonConfigurationStatus": {
      "description": "Response to a daemon configuration request.",
      "type": "object",
      "properties": {
        "applied": {
          "$ref": "#/definitions/DaemonConfigurationSpec"
        },
        "daemonConfigurationMap": {
          "description": "Config map which contains all the active daemon configurations"
        },
        "immutable": {
          "$ref": "#/definitions/ConfigurationMap"
        }
      },
      "example": {
        "applied": {
          "options": {}
        },
        "daemonConfigurationMap": ""
      }
    },
    "Status": {
      "description": "Status of an individual component",
      "type": "object",
      "properties": {
        "msg": {
          "description": "Human readable status/error/warning message",
          "type": "string"
        },
        "state": {
          "description": "State the component is in",
          "type": "string",
          "enum": [
            "Ok",
            "Warning",
            "Failure",
            "Disabled"
          ]
        }
      },
      "example": {
        "msg": "msg",
        "state": "Ok"
      }
    },
    "StatusResponse": {
      "description": "Health and status information of daemon",
      "type": "object",
      "properties": {
        "bpflock": {
          "$ref": "#/definitions/Status"
        },
        "stale": {
          "description": "List of stale information in the status",
          "type": "object",
          "additionalProperties": {
            "description": "Timestamp when the probe was started",
            "type": "string",
            "format": "date-time"
          }
        }
      },
      "example": {
        "bpflock": {
          "msg": "msg",
          "state": "Ok"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
