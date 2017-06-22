package validate

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gh-issue/gh-issue/models"

	"github.com/manyminds/api2go/jsonapi"
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
		t.Errorf("JSON unmarshal unsuccessful. handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestSuccess(t *testing.T) {

	testHandlerFn := http.HandlerFunc(temp)
	w := httptest.NewRecorder()

	postBody := models.Orderinfo{"1223", "Date1", "Date2", "Fedex", "FedexAccount", "No comment", "Fake payment", "Num 3", "OK status"}
	body, _ := jsonapi.Marshal(postBody)
	fmt.Printf("whats happening")
	fmt.Print(body)
	fmt.Print(postBody)
	b := bytes.NewBuffer(body)
	fmt.Printf("test")
	fmt.Print(b)
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
