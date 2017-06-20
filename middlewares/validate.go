package validate

import (
	"context"
	"io"
	//"github.com/dictyBase/gh-issue/models"
	"gh-issue/gh-issue/models"
	"io/ioutil"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
)

//TEMP
func Temp(w http.ResponseWriter, r *http.Request) {
	//var order models.Orderinfo
	io.WriteString(w, `{"alive": true}`)
	body, err := ioutil.ReadAll(r.Body)
	if body != nil {
		http.Error(w, "body not nil", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = jsonapi.Unmarshal(body, &order)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

//JSONValidator breaks chain if JSON is not valid and passes the decodedJSON on through context
func JSONValidator(fn http.HandlerFunc) http.HandlerFunc {
	var order models.Orderinfo
	newfn := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = jsonapi.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "DecodedJson", order)
		fn(w, r.WithContext(ctx))
	}
	return newfn

}
