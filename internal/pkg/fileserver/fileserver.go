package fileserver

import (
	"fmt"
	"net/http"
)

func Start() error {
	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)
	http.Handle("/data/", staticHandler)

	fmt.Println("FileServer started at 8081")
	return http.ListenAndServe(":8081", nil)

}
