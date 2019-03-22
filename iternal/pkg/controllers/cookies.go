package controllers

import (
	"net/http"
	"time"
)

func CreateCookie(w *http.ResponseWriter, sessionID string) {
	http.SetCookie(*w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessionID,
		Expires:  time.Now().Add(60 * time.Hour),
		HttpOnly: true,
		// Secure:   true,
	})
}

func DeleteCookie(w *http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie((*w), &http.Cookie{
		Name:     "sessionid",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		// Secure:   true,
	})
	delete(sessions, cookie.Value)
}
