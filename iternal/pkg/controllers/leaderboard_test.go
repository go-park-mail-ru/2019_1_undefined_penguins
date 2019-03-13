package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLeaders(t *testing.T) {
	cases := []TestCase{
		TestCase{
			ID:         "42",
			Response:   `{"status": 200, "resp": {"user": 42}}`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			ID:         "500",
			Response:   `{"status": 500, "err": "db_error"}`,
			StatusCode: http.StatusInternalServerError,
		},
	}
	for caseNum, item := range cases {
		url := "http://example.com/api/user?id=" + item.ID
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		GetLeaders(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}
