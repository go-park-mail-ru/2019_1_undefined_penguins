package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")

		if accessControlRequestHeaders := r.Header.Get("Access-Control-Request-Headers"); accessControlRequestHeaders != "" {
			responseHeader.Set("Access-Control-Allow-Headers", accessControlRequestHeaders)
		}
		responseHeader.Set("Access-Control-Allow-Origin", origin)
		w.WriteHeader(200)
		return
	} else {
		fmt.Println(r.Method)
		fmt.Println(r.RequestURI)
		cookie, err := r.Cookie("sessionid")
		if err != nil {

			fmt.Println("Я не нашел сессию")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("You are not authorized"))
			return
		}
		user, found := sessions[cookie.Value]
		if !found {

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("You are not authorized"))
			return
		}
		setupResponse(&w, r)
		bytes, err := json.Marshal(user)

		w.Write(bytes)
		w.WriteHeader(http.StatusOK)
		fmt.Print(w)

	}

}
