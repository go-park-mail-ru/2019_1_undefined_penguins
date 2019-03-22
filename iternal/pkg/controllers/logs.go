package controllers

import (
	"fmt"
	"net/http"
)

func logMethodAndURL(r *http.Request) {
	fmt.Println(r.Method + r.RequestURI)
}
