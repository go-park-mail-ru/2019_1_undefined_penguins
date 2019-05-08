package controllers

import (

	//db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

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

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer grcpConn.Close()

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()

	users, _ := authManager.GetUserArray(ctx, leaders)

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
	grcpConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer grcpConn.Close()

	authManager := models.NewAuthCheckerClient(grcpConn)
	ctx := context.Background()

	//info, err := db.UsersCount()
	info, _ := authManager.GetUserCountInfo(ctx, new(models.Nothing))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
