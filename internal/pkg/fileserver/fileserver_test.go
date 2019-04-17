package fileserver

import (
	"testing"
)

func TestFileServer(t *testing.T) {

	err := Start()
	if err == nil {
		t.Error(err)
	}

}
