package validate

import (
	"context"
	//"github.com/dictyBase/gh-issue/models"
	"gh-issue/gh-issue/models"
	"io/ioutil"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
)

//JSONValidator breaks chain if JSON is not valid and passes the decodedJSON on through context
func JSONValidator(fn http.HandlerFunc) http.HandlerFunc {
	var order models.Orderinfo
	newfn := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		err = jsonapi.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		ctx := context.WithValue(r.Context(), "DecodedJson", order)
		fn(w, r.WithContext(ctx))
	}
	return newfn

}
