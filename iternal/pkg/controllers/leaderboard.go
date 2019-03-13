package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	bytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"=("}`))
	}
	w.Write(bytes)
}
