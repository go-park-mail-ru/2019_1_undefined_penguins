package controllers

import (
	"net/http"
)

func SetupCORS(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	responseHeader := (*w).Header()
	responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	responseHeader.Set("Access-Control-Allow-Credentials", "true")

	responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")

	responseHeader.Set("Access-Control-Allow-Origin", origin)
}
