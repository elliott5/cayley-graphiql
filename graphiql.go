//go:generate go-bindata-assetfs -pkg graphiql static

package graphiql

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/query"
	_ "github.com/cayleygraph/cayley/query/graphql"
)

func AddHandlers(db graph.QuadStore) error {

	gql := query.GetLanguage("graphql")
	if gql == nil {
		return errors.New("could not find cayley language graphql")
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		type IdeInput struct {
			Query string "json:'query'"
		}

		var ideInput IdeInput
		byt, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//fmt.Println("graphql query raw:", string(byt))
		if err := json.Unmarshal(byt, &ideInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(ideInput.Query) == 0 {
			http.Error(w, "no query to process", http.StatusBadRequest)
			return
		}
		//fmt.Println("graphql query unmarshalled:", ideInput.Query)
		gql.HTTPQuery(context.TODO(), db, w, bytes.NewReader([]byte(ideInput.Query)))
	})

	http.HandleFunc("/graphiql/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "bad method", 405)
			return
		}

		file := strings.TrimPrefix(r.URL.Path, "/graphiql")
		if file == "/" || file == "" {
			file = "/index.html"
		}
		file = "static" + file

		data, err := Asset(file)
		if err != nil {
			// Asset was not found.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if strings.HasSuffix(file, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}

		w.Write(data)
	})

	return nil
}
