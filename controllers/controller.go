package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func StudentHandler(rw http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	name := queries.Get("name")
	studentResponse := &StudentResponse {
		Name: name,
	}

	responseData, err := json.Marshal(studentResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Oops, something went wrong"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(responseData)
}


func UsersController(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "users")
}
func ListsController(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "list")
}

func StudentsController(rw http.ResponseWriter, r *http.Request) {
	StudentHandler(rw, r)
	fmt.Fprint(rw, "students")
}


