package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_handleSchemaDownload(t *testing.T) {
	tests := map[string]struct {
		expectedResult string
		expectedCode   int
	}{
		"can download a valid schema": {
			expectedResult: `download handler`,
			expectedCode:   200,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := newMockDbClient()
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("GET", "/schema/123", nil)
			w := httptest.NewRecorder()

			handlerUnderTest := server.handleSchemaDownload()
			handlerUnderTest(w, r)

			var respBody bytes.Buffer
			if _, err := respBody.ReadFrom(w.Body); err != nil {
				t.Errorf("Could not read body: %v", err)
			}

			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedResult, respBody.String())
		})
	}
}
