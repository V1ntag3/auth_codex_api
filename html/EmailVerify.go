package html

import (
	"auth_codex_api/models"
	"auth_codex_api/utilities"
	"bytes"
	"fmt"
	"net/smtp"
)

func HtmlEmailVerify(recipient string, user models.User) (bool, error) {
	// Dados do email
	from := utilities.GoDotEnvVariable("EMAIL_HOST")
	password := utilities.GoDotEnvVariable("PASSWORD_HOST")
	server_smtp := utilities.GoDotEnvVariable("SERVER_SMTP")
	to := recipient

	// Autenticação no servidor SMTP
	auth := smtp.PlainAuth("", from, password, server_smtp)

	htmlBody := `
	<!doctype html>
<html>
  <head>
    <title>Verificação de conta</title>
  </head>
  <body>
    <h1>Olá ` + user.Name + `</h1>
    <p>
      Para sua segurança precisamos saber se o email é seu mesmo, para isso
      basta clicar no botão abaixo
    </p>
    <a href="` + utilities.GoDotEnvVariable("HOST_FRONT") + `verify/` + user.Id + `">Verificar</a>
  </body>
  <style>
    * {
      font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
        "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
      color: white;
    }
    body {
      background-color: #2d7cd6;
    }
    p {
      font-size: 18px;
    }
    a {
      padding: 10px 20px;
      font-size: 2rem;
      border: none;
      background-color: blue;
      color: white;
      border-radius: 1rem;
      font-weight: 800;
      cursor: pointer;
    }
  </style>
</html>

	`

	// Cabeçalhos do email
	headers := map[string]string{
		"From":         from,
		"To":           to,
		"Subject":      "Olá, para acessar sua conta é necessário verificar sua conta",
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	// Construindo o corpo do email
	var emailBody bytes.Buffer
	for key, value := range headers {
		_, _ = emailBody.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	_, _ = emailBody.WriteString("\r\n" + htmlBody)

	// Envio do email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, emailBody.Bytes())
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
