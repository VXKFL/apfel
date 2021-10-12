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
	"io/ioutil"
	"log"
	"net/http"
)

func ApiRegisterPost(w http.ResponseWriter, r *http.Request) {
	var input database.UserT
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unparsable JSON")
		return
	}

	code, err := database.Register(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		log.Printf("/api/register: post failed with error '%s'", err.Error())
		return
	}

	resp, err := json.Marshal(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		log.Printf("/api/register: post failed with error '%s'", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
