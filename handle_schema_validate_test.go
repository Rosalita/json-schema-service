package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_handleSchemaValidate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedResult string
		expectedCode   int
	}{
		"can upload": {
			input: `
			{
				"source": "/home/alice/image.iso",
				"destination": "/mnt/storage",
				"timeout": null,
				"chunks": {
				  "size": 1024,
				  "number": null
				}
			  }`,
			expectedResult: `validate handler`,
			expectedCode:   200,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db := "todo"
			router := mux.NewRouter()
			server := newServer(db, router)

			r, _ := http.NewRequest("POST", "/validate/123", bytes.NewBuffer([]byte(test.input)))
			w := httptest.NewRecorder()

			handlerUnderTest := server.handleSchemaValidate()
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
