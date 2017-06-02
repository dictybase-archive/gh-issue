package validate

import (
	"context"
	"gh-issue/gh-issue/models"
	"io/ioutil"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
)

//JsonValidator breaks chain if JSON is not validat and passes the decodedJSON on through context
func JsonValidator(fn http.HandlerFunc) http.HandlerFunc {
	var order models.Orderinfo

	newfn := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading Body", http.StatusInternalServerError)
		}
		err = jsonapi.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, "error unmarshaling json struct", http.StatusInternalServerError)

		}
		ctx := context.WithValue(r.Context(), "DecodedJson", order)
		fn(w, r.WithContext(ctx))
	}
	return newfn

}
