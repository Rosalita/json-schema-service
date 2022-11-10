package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_handleSchemaDownload(t *testing.T) {
	tests := map[string]struct {
		schemaID       string
		expectedResult string
		expectedCode   int
	}{
		"can download a valid schema": {
			schemaID:       "config-schema",
			expectedResult: `{"mock":"schema"}`,
			expectedCode:   200,
		},
		"can handle schema not found error": {
			schemaID:       "not-found",
			expectedResult: "not found\n",
			expectedCode:   404,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := newMockDbClient()
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("GET", fmt.Sprintf("/schema/%s", test.schemaID), nil)
			w := httptest.NewRecorder()

			// As unit testing individual handler directly, without calling .ServeHTTP on the router,
			// these tests will need to manually set URL variables on the router as part of test setup.
			r = mux.SetURLVars(r, map[string]string{
				"id": test.schemaID,
			})

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
