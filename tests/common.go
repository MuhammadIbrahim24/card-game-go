package testcases

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/card-game/routes"
	"github.com/gorilla/mux"
)

const URI = "http://localhost:8000/deck"

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
