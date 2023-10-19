# SendMail API with Golang and Gin

## Overview

This is  an email sending API using the Go (Golang) programming language and the Gin web framework. The API allows you to send emails with custom subject and body to specified recipients.

## Getting Started

### Prerequisites

- Make sure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone this repository:

git clone https://github.com/leo12wb/sendMailGolangGin.git
cd sendmail-api-golang-gin
go get -u github.com/gin-gonic/gin
go get gopkg.in/gomail.v2

### Usage

1. Start the API server:

    ```bash
    go run main.go

    ```
2. Make a POST request to the /notification/sendmail route with a JSON body containing the following fields:

    to: The recipient's email address.
    subject: The email subject.
    text: The email body in plain text.
    Example request using cURL:
```bash
    curl -X POST -H "Content-Type: application/json" -d '{
    "to": "recipient@example.com",
    "subject": "Email Subject",
    "text": "This is the plain text email body."
    }' http://localhost:3000/notification/sendmail
```