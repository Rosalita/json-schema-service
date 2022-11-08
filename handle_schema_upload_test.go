package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_handleSchemaUpload(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedResult string
		expectedCode   int
	}{
		"Can upload a valid schema": {
			input: `
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
			expectedResult: `upload handler`,
			expectedCode:   200,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := "todo"
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("POST", "/schema/123", bytes.NewBuffer([]byte(test.input)))
			w := httptest.NewRecorder()

			handlerUnderTest := server.handleSchemaUpload()
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
