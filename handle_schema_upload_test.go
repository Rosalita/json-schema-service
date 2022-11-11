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

func Test_handleSchemaUpload(t *testing.T) {
	tests := map[string]struct {
		schemaID       string
		payload        string
		expectedResult string
		expectedCode   int
	}{
		"Can upload a valid schema": {
			schemaID: "config-schema",
			payload: `
			{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"type": "object",
				"properties": {
				  "source": {
					"type": "string"
				  },
				  "destination": {
					"type": "string"
				  },
				  "timeout": {
					"type": "integer",
					"minimum": 0,
					"maximum": 32767
				  },
				  "chunks": {
					"type": "object",
					"properties": {
					  "size": {
						"type": "integer"
					  },
					  "number": {
						"type": "integer"
					  }
					},
					"required": ["size"]
				  }
				},
				"required": ["source", "destination"]
			  }`,
			expectedResult: `{"action":"uploadSchema","id":"config-schema","status":"success"}`,
			expectedCode:   201,
		},
		"Can handle invalid json schema": {
			schemaID:       "config-schema",
			payload:        `{`,
			expectedResult: `{"action":"uploadSchema","id":"config-schema","message":"Invalid JSON","status":"error"}`,
			expectedCode:   400,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := newMockDbClient()
			router := mux.NewRouter()

			server := newServer(db, router)

			r, _ := http.NewRequest("POST", fmt.Sprintf("/schema/%s", test.schemaID), bytes.NewBuffer([]byte(test.payload)))

			// As unit testing individual handler directly, without calling .ServeHTTP on the router,
			// these tests will need to manually set URL variables on the router as part of test setup.
			r = mux.SetURLVars(r, map[string]string{
				"id": test.schemaID,
			})

			w := httptest.NewRecorder()

			handlerUnderTest := server.handleSchemaUpload()
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
