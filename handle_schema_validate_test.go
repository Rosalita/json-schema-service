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

func Test_handleSchemaValidate(t *testing.T) {
	tests := map[string]struct {
		schemaID       string
		payload        string
		expectedResult string
		expectedCode   int
	}{
		"can validate valid data against a schema": {
			schemaID: "config-schema",
			payload: `
			{
				"source": "/home/alice/image.iso",
				"destination": "/mnt/storage",
				"timeout": null,
				"chunks": {
				  "size": 1024,
				  "number": null
				}
			  }`,
			expectedResult: `{"action":"validateDocument","id":"config-schema","status":"success"}`,
			expectedCode:   200,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := newMockDbClient()
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("POST", fmt.Sprintf("/validate/%s", test.schemaID), bytes.NewBuffer([]byte(test.payload)))
			w := httptest.NewRecorder()

			// As unit testing individual handler directly, without calling .ServeHTTP on the router,
			// these tests will need to manually set URL variables on the router as part of test setup.
			r = mux.SetURLVars(r, map[string]string{
				"id": test.schemaID,
			})

			handlerUnderTest := server.handleSchemaValidate()
			handlerUnderTest(w, r)

			var respBody bytes.Buffer
			if _, err := respBody.ReadFrom(w.Body); err != nil {
				t.Errorf("Could not read body: %v", err)
			}

			assert.Equal(t, test.expectedCode, w.Code)
			assert.JSONEq(t, test.expectedResult, respBody.String())
		})
	}
}
