package helper

import (
	"context"
	"log"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
	"strconv"
	"strings"
	"time"
)

func AuthUserFromContext(ctx context.Context) user.User {
	return ctx.Value(AuthUserKey).(user.User)
}

func DefineAction(r *http.Request) string {
	p := strings.Trim(r.URL.Path, "/")
	if len(strings.Split(p, "/")) > 2 {
		return "notFound"
	}

	action := "notAllowed"
	switch r.Method {
	case http.MethodGet:
		idExists := false
		if id, ok := r.URL.Query()["id"]; ok {
			if _, err := strconv.Atoi(id[0]); err == nil {
				idExists = true
			}
		}
		queryAction, ok := r.URL.Query()["action"]

		if ok && queryAction[0] == "edit" && idExists {
			action = "edit"
		} else if ok && queryAction[0] == "create" {
			action = "create"
		} else {
			action = "index"
		}
	case http.MethodPost:
		action = "store"
	case http.MethodPut:
		action = "update"
	case http.MethodDelete:
		action = "destroy"
	}

	return action
}

func Render(w http.ResponseWriter, filename string, data interface{}) {
	err := config.Tpl.ExecuteTemplate(w, filename, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func currentTime() string {
	t := time.Now()
	datetime := t.Format("2006-01-02 15:04:05")
	return datetime
}
