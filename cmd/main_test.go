package main

import "testing"

func TestServer(t *testing.T) {

	go func() {
		main()
	}()

}
