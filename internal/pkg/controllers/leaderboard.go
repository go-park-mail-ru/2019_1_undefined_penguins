package controllers

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	//"google.golang.org/grpc/status"
)

func GetLeaderboardPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	leaders := new(models.LeadersInfo)

	if err != nil {
		id = 1
		leaders.ID = uint64(id)
	}
	leaders.ID = uint64(id)


	ctx := context.Background()
	users, _ := models.AuthManager.GetUserArray(ctx, leaders)

	fmt.Println("led: ", users)

	u := users.Users
	if respBody, err := json.Marshal(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}
}

func GetLeaderboardInfo(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	info, _ := models.AuthManager.GetUserCountInfo(ctx, new(models.Nothing))

	var i models.LeadersInfo1
	i.UsersOnPage = uint(info.UsersOnPage)
	i.Count = uint(info.Count)
	if respBody, err := json.Marshal(i); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}

}
