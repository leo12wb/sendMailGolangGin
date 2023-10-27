package mainN

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "gosite/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/gomail.v2"
)

// swagger info
//	@title			SendMail API
//	@version		1.0
//	@description	Aplicativo para envio de emails.

//	@Summary	Enviar um email
//	@Schemes
//	@Description	Envie um email com assunto e corpo personalizado
//	@Tags			Enviar um email
//	@Accept			json
//	@Produce		json
//	@Param			to		body		string	true	"Informe o email de envio"
//	@Param			subject	body		string	true	"Informe o Titulo da Mensagem"
//	@Param			text	body		string	true	"Informe o Corpo da Mensagem"
//	@Success		200		{string}	true	"E-mail enviado com sucesso"
//	@Router			/notification/sendmail [post]

func mainn() {
	// Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configuração do middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} 
	r.Use(cors.New(config))

	r.POST("/notification/sendmail", func(c *gin.Context) {
		var input struct {
			To      string `json:"to"`
			Subject string `json:"subject"`
			Text    string `json:"text"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.To == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Informe o email de envio"})
			return
		}
		if input.Subject == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Informe o Titulo da Mensagem"})
			return
		}
		if input.Text == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Informe o Corpo da Mensagem"})
			return
		}

		emailNaoResponda := os.Getenv("EMAIL_NAO_RESPONDA")
		senhaNaoResponda := os.Getenv("SENHA_NAO_RESPONDA")
		servidorSmtp := os.Getenv("SERVIDOR_SMTP")
		portString := os.Getenv("PORT")
		port, err := strconv.Atoi(portString)
		if err != nil {
			panic(err)
		}

		message := gomail.NewMessage()
		message.SetHeader("From", emailNaoResponda)
		message.SetHeader("To", input.To)
		message.SetHeader("Subject", input.Subject)
		message.SetBody("text/plain", input.Text)

		d := gomail.NewDialer(servidorSmtp, port, emailNaoResponda, senhaNaoResponda)

		if err := d.DialAndSend(message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		fmt.Println("Email enviado com sucesso!")
		c.JSON(http.StatusOK, gin.H{"message": "Enviado"})
	})

	r.Run(":3000")
}
