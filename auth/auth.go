package main

import (
	"html/template"
	"net/http"
	"os/user"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandlerFunc)
	http.HandleFunc("/login/", loginHandlerFunc)
	http.ListenAndServe(":3000", nil)
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("auth/first.html")
	t.Execute(w, nil)
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")

	if user == HostName() {
		w.Write([]byte("Привет " + user))
	} else {
		t, _ := template.ParseFiles("auth/nd.html")
		t.Execute(w, nil)
	}
}

func HostName() string {
	cur, _ := user.Current()
	curUser := cur.HomeDir
	contrib := strings.Trim(curUser, `C:\Users\`)
	return contrib
}
