package main

import (
	"log"
	"net/http"

	"github.com/cayleygraph/cayley"

	"github.com/elliott5/cayley-graphiql"
	"github.com/elliott5/cayley-graphiql/example/loadtestdata"
)

func main() {
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}

	err = loadtestdata.LoadTestData(store)
	if err != nil {
		log.Fatalln(err)
	}

	graphiql.AddHandlers(store) // call the GraphQL IDE package

	port := "4080"
	httpURL := "http://localhost:" + port + "/graphiql/"
	log.Println("Serving on " + httpURL)
	log.Println(http.ListenAndServe(":"+port, nil))
}
