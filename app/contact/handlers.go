package contact

import (
	"fmt"
	"net/http"
	"net/smtp"
	"onlineshop/helper"
)

// Header struct
type Header struct {
	Context helper.ContextData
	Link    string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/contact/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			data := struct {
				Header Header
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "contact",
				},
			}

			helper.Render(w, "contact.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			// Sender data.
			from := "email@gmail.com"
			password := "password"

			// Receiver email address.
			to := []string{
				"email@gmail.com",
			}

			// smtp server configuration.
			smtpHost := "smtp.gmail.com"
			smtpPort := "587"

			// Message.
			message := []byte("This is a test email message.")

			// Authentication.
			auth := smtp.PlainAuth("", from, password, smtpHost)

			// Sending email.
			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Email Sent Successfully!")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
