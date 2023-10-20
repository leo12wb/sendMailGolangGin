package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

	r := gin.Default()

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

		fmt.Println("Email enviado.")
		c.JSON(http.StatusOK, gin.H{"message": "Enviado"})
	})

	r.Run(":3000")
}
