package helpers

import (
	"net/http"
	"time"

	"2019_1_undefined_penguins/internal/pkg/models"
)

func DeleteCookie(w *http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie((*w), &http.Cookie{
		Name:     "sessionid",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		// Secure:   true,
	})
	delete(models.Sessions, cookie.Value)
}
