package docs

import (
	"github.com/swaggo/swag"
)

const docJSON = `{
    "swagger": "2.0",
    "info": {
        "description": "Aplicativo para envio de emails.",
        "title": "SendMail API",
        "version": "1.0"
    },
    "paths": {
        "/notification/sendmail": {
            "post": {
                "summary": "Enviar um email",
                "description": "Envie um email com assunto e corpo personalizado",
                "tags": ["Enviar um email"],
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "parameters": [
                    {
                        "in": "body",
                        "name": "email",
                        "description": "Email details",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "properties": {
                                "to": {
                                    "type": "string",
                                    "example": "recipient@example.com"
                                },
                                "subject": {
                                    "type": "string",
                                    "example": "Email Subject"
                                },
                                "text": {
                                    "type": "string",
                                    "example": "This is the plain text email body."
                                }
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "E-mail enviado com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
    Version:          "1.0",
    Host:             "",
    BasePath:         "",
    Schemes:          []string{},
    Title:            "SendMail API",
    Description:      "Aplicativo para envio de emails.",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  docJSON, // Substitua docTemplate pelo JSON
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
