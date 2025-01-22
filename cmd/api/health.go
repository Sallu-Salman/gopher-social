package main

import (
	"net/http"
)

func (a *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	data["status"] = "ok"
	
	err := WriteJson(w, http.StatusOK, data)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
	}
}