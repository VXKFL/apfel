package swagger

import (
	"net/http"
	"html/template"
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
