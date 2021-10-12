/*
 * apfel API
 *
 * [ap]plication [f]or [e]vent [l]ogin
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"apfel/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ApiEventEventIDGet(w http.ResponseWriter, r *http.Request) {
	attendants, err := database.GetAttendants()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		log.Printf("/api/event/{EventID}: get failed with error '%s'", err.Error())
		return
	}

	resp, err := json.Marshal(attendants)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		log.Printf("/api/event/{EventID}: get failed with error '%s'", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func ApiEventEventIDPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
