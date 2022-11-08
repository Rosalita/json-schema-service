package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_ServeHTTP(t *testing.T) {
	tests := map[string]struct {
		route        string
		expectedCode int
	}{
		"can serve a known route": {
			route:        `/schema/123`,
			expectedCode: 200,
		},
		"can handle an unknown route": {
			route:        `/unknown`,
			expectedCode: 404,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := "todo"
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("GET", test.route, nil)
			w := httptest.NewRecorder()

			server.ServeHTTP(w, r)
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
