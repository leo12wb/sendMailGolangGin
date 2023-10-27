package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type Request struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

type Response struct {
	Message string `json:"message"`
}

func SendMail(ctx context.Context, request Request) (Response, error) {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		return Response{Message: fmt.Sprintf("Erro ao carregar o arquivo .env: %v", err)}, err
	}

	emailNaoResponda := os.Getenv("EMAIL_NAO_RESPONDA")
	senhaNaoResponda := os.Getenv("SENHA_NAO_RESPONDA")
	servidorSmtp := os.Getenv("SERVIDOR_SMTP")
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		return Response{Message: fmt.Sprintf("Erro ao converter a porta: %v", err)}, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", emailNaoResponda)
	message.SetHeader("To", request.To)
	message.SetHeader("Subject", request.Subject)
	message.SetBody("text/plain", request.Text)

	d := gomail.NewDialer(servidorSmtp, port, emailNaoResponda, senhaNaoResponda)

	if err := d.DialAndSend(message); err != nil {
		return Response{Message: fmt.Sprintf("Erro ao enviar o email: %v", err)}, err
	}

	return Response{Message: "Email enviado com sucesso!"}, nil
}

func main() {
	lambda.Start(SendMail)
}
