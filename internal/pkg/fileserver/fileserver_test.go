package fileserver

import (
	"testing"
)

func TestFileServer(t *testing.T) {

	go func() {
		Start()
	}()

}
