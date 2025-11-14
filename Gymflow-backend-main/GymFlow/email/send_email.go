package email

import (
    "fmt"
    "net/smtp"
    "gymflow/models"
)

func SendVerificationEmail(user *models.User, token string) error {

    smtpHost := "smtp.example.com"  // замените на реальный smtp
    smtpPort := "587"

    fromEmail := "noreply@example.com"
    smtpUser := "noreply@example.com"
    smtpPass := "yourpassword"

    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

    to := []string{user.Email}

    verifyURL := fmt.Sprintf("https://yourapp.com/auth/verify?token=%s", token)

    msg := []byte(
        "Subject: Подтвердите email\r\n" +
            "To: " + user.Email + "\r\n" +
            "\r\n" +
            "Перейдите по ссылке для подтверждения:\n" + verifyURL + "\n",
    )

    return smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, msg)
}
