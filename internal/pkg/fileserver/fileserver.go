package fileserver

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"net/http"
)

func Start() error {
	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)
	http.Handle("/data/", staticHandler)

	helpers.LogMsg("FileServer started at 8081")
	return http.ListenAndServe(":8081", nil)

}
