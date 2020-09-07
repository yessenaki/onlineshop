package user

import (
	"net/http"
	"onlineshop/helper"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var ctx interface{}
	var path = map[string]string{
		"folder": "auth",
		"file":   "login.gohtml",
	}

	if r.Method == http.MethodGet {
		// Check if session already exists
		result, err := sessionExists(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if result {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		u, err := userExists(email)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if u == (User{}) {
			http.Error(w, "Email and/or password do not match", http.StatusForbidden)
			return
		}

		// Does the entered password match the stored password?
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err != nil {
			http.Error(w, "Email and/or password do not match", http.StatusForbidden)
			return
		}

		// Generate a new uuid && create a new session in the db
		sessionID, _ := uuid.NewV4()
		err = createSession(sessionID.String(), u.ID)
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

	helper.RenderTemplate(w, path, ctx)
}

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {
	var ctx interface{}
	var tplPath = map[string]string{
		"folder": "auth",
		"file":   "register.gohtml",
	}

	if r.Method == http.MethodGet {
		// Check if session already exists
		result, err := sessionExists(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if result {
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
		u, err := userExists(user.Email)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if u != (User{}) {
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
