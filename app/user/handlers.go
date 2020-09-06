package user

import (
	"net/http"
	"onlineshop/helper"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var ctx interface{}
	var tplPath = map[string]string{
		"folder": "auth",
		"file":   "register.gohtml",
	}

	if r.Method == http.MethodGet {
		if sessionExists(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Method == http.MethodPost {
		user := User{}
		user.FirstName = r.FormValue("first_name")
		user.LastName = r.FormValue("last_name")
		user.Email = r.FormValue("email")
		password := r.FormValue("password")

		// Check if user already exists
		result, err := userExists(user.Email)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if result {
			http.Error(w, "Email already taken", http.StatusForbidden)
			return
		}

		// Encrypt password with bcrypt
		ps, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		user.Password = string(ps)

		// Create a new user
		userID, err := createUser(user)
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
			Name:  "session_id",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	helper.RenderTemplate(w, tplPath, ctx)
}

func sessionExists(r *http.Request) bool {
	// c, err := r.Cookie("session_id")
	// if err != nil {
	// 	return false
	// }

	// row, err := config.DB.Query("SELECT sessions.session_id, users.id FROM sessions INNER JOIN users ON sessions.user_id = users.id LIMIT 1")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	return false
}
