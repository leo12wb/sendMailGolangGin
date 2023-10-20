package docs

import (
	"github.com/swaggo/swag"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
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
                        "name": "to",
                        "in": "body",
                        "required": true,
                        "type": "string",
                        "description": "Informe o email de envio",
                        "schema": {}
                    },
                    {
                        "name": "subject",
                        "in": "body",
                        "required": true,
                        "type": "string",
                        "description": "Informe o Título da Mensagem",
                        "schema": {}
                    },
                    {
                        "name": "text",
                        "in": "body",
                        "required": true,
                        "type": "string",
                        "description": "Informe o Corpo da Mensagem",
                        "schema": {}
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
    SwaggerTemplate:  docTemplate,
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
