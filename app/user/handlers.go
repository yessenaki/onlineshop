package user

import (
	"log"
	"net/http"
	"onlineshop/helper"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const maxAge = 3600

// Header struct
type Header struct {
	Auth User
	Link string
}

func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		data := struct {
			Header Header
			User   *User
		}{}

		if r.Method == http.MethodGet {
			// Check if session already exists
			result, err := SessionExists(r)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			if result {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			data.User = &User{}
			helper.Render(w, "login.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			user := &User{
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
			}

			result, err := user.validate(false)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			if result == false {
				data.User = user
				helper.Render(w, "login.gohtml", data)
				return
			}

			// Generate a new uuid && create a new session in the db
			sessionID, _ := uuid.NewV4()
			err = createSession(sessionID.String(), user.ID)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			// Create a new session in the browser cookie
			cookie := &http.Cookie{
				Name:     "session_id",
				Value:    sessionID.String(),
				Path:     "/",
				HttpOnly: true,
				MaxAge:   maxAge,
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}

func Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/register/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		data := struct {
			Header Header
			User   *User
		}{}

		if r.Method == http.MethodGet {
			// Check if session already exists
			result, err := SessionExists(r)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			if result {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			data.User = &User{}
			helper.Render(w, "register.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			user := &User{
				FirstName: r.PostFormValue("first_name"),
				LastName:  r.PostFormValue("last_name"),
				Email:     r.PostFormValue("email"),
				Password:  r.PostFormValue("password"),
				Password2: r.PostFormValue("password_confirm"),
			}

			// Form validation
			result, err := user.validate(true)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			if result == false {
				data.User = user
				helper.Render(w, "register.gohtml", data)
				return
			}

			// Encrypt password with bcrypt
			passwordSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			user.Password = string(passwordSlice)

			// Create a new user
			userID, err := user.create()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			// Generate a new uuid && create a new session in the db
			sessionID, _ := uuid.NewV4()
			err = createSession(sessionID.String(), userID)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			// Create a new session in the browser cookie
			cookie := &http.Cookie{
				Name:     "session_id",
				Value:    sessionID.String(),
				Path:     "/",
				HttpOnly: true,
				MaxAge:   maxAge,
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}

func Logout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := SessionExists(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if result == false {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Delete the session from db
		cookie, _ := r.Cookie("session_id")
		err = deleteSession(cookie.Value)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		// Remove the cookie from browser
		cookie = &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
}
