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
