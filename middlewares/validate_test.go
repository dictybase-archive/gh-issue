package validate

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func temp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("body string"))
	w.WriteHeader(200)
}

//Fills body with empty string isntead of expected JSON format so unmarshal should fail
func TestUnmarshalFailure(t *testing.T) {

	testHandlerFn := http.HandlerFunc(temp)
	w := httptest.NewRecorder()

	b := bytes.NewBufferString("")
	req, err := http.NewRequest("POST", "/json-test", b)
	if err != nil {
		t.Fatal(err)
	}

	JSONValidator(testHandlerFn).ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestSuccess(t *testing.T) {

	testHandlerFn := http.HandlerFunc(temp)
	w := httptest.NewRecorder()

	b := bytes.NewBufferString("")
	req, err := http.NewRequest("POST", "/json-test", b)
	if err != nil {
		t.Fatal(err)
	}

	JSONValidator(testHandlerFn).ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
