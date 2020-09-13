package helper

import (
	"context"
	"net/http"
	"onlineshop/app/user"
	"text/template"
	"time"
)

func AuthUserFromContext(ctx context.Context) user.User {
	return ctx.Value(AuthUserKey).(user.User)
}

func RenderTemplate(w http.ResponseWriter, path map[string]string, ctx interface{}) {
	t := template.Must(template.ParseGlob("templates/layouts/*.gohtml"))
	t = template.Must(t.ParseGlob("templates/" + path["folder"] + "/*.gohtml"))
	err := t.ExecuteTemplate(w, path["file"], ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func currentTime() string {
	t := time.Now()
	datetime := t.Format("2006-01-02 15:04:05")
	return datetime
}
