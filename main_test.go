package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/stretchr/testify/assert"
)


func Test_upload(t *testing.T){
	tests := map[string]struct {
		input           string
		expectedResult  string
		expectedCode    int
	}{
		"can upload": {
			input: `{"data": "some json payload"}`,
			expectedResult: `upload handler`,
			expectedCode: 200,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/upload", bytes.NewBuffer([]byte(test.input)))
			w := httptest.NewRecorder()

			uploadHandler(w, r)

			var respBody bytes.Buffer
			if _, err := respBody.ReadFrom(w.Body); err != nil {
				t.Errorf("Could not read body: %v", err)
			}

			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedResult, respBody.String())
		})
	}
}