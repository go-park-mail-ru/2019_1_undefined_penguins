package server

import (
	"os"
	"testing"
)

func TestServerFail(t *testing.T) {

	params := Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "8080"
	}
	go func() {
		err := StartApp(params)
		if err != nil {
			t.Error(err)
		}
	}()
	

}
