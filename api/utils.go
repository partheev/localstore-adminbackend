package main

import (
	"encoding/json"
	"net/http"
)

type res struct {
	Message string
}

func resMessage(s string) res {
	response := res{
		Message: s,
	}
	return response
}
func (app *application) writeJson(w http.ResponseWriter, data interface{}, status int) {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
func (app *application) writeError(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJson(w, theError, http.StatusBadRequest)

}
