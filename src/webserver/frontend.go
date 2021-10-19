package swagger

import (
	"html/template"
	"net/http"
)

//serve main page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").ParseFiles("pages/index.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

//serve admin page
func Admin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("admin.html").ParseFiles("pages/admin.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

//serve register page
func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("register.html").ParseFiles("pages/register.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
