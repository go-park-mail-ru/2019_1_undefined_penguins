package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

func CheckWsSingle(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	cookie, err := r.Cookie("sessionid")

	ctx := context.Background()
	//if err != nil || cookie.Value == "" {
	if err != nil || cookie.Value == "" {
		user.Login = "Anonumys"
	} else {
		user, _ = models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})
	}

	_, _ = models.AuthManager.AddUserToGame(ctx, user)
	w.WriteHeader(http.StatusOK)
}

func CheckWsMulti(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)

	cookie, err := r.Cookie("sessionid")
	fmt.Println(cookie)
	ctx := context.Background()

	if err != nil || cookie.Value == "" {
		helpers.LogMsg("No Cookie in Multi")
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {

		user, err = models.AuthManager.GetUser(ctx, &models.JWT{Token: cookie.Value})
		if err != nil {
			helpers.LogMsg("Invalid Cookie in Multi")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	user.IsPlaying = true
	_, _ = models.AuthManager.AddUserToGame(ctx, user)
	w.WriteHeader(http.StatusOK)
}
