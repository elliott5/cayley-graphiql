package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/cayleygraph/cayley"
	_ "github.com/cayleygraph/cayley/graph/nosql/ouch"
	"github.com/gopherjs/gopherjs/js"

	"github.com/augustoroman/promise"
	"github.com/cayleygraph/cayley/query/graphql"
	"github.com/elliott5/cayley-graphiql/example/loadtestdata"
)

func init() {
	js.Global.Set("cayleyLocal", promise.Promisify(cayleyLocal))
}

func graph(paramMap map[string]interface{}) (*cayley.Handle, error) {

	gdb, found := paramMap["variables"].(map[string]interface{})["pouchdb"].(string)
	if !found || gdb == "" {
		return nil, errors.New("pouchdb name not provided")
	}

	ouch := "pouch"
	ouchGraph, err := cayley.NewGraph(ouch, gdb, nil)
	if err != nil {
		return nil, err
	}

	if ouchGraph.QuadStore.Size() == 0 && gdb == "cayley_testdata" {
		println("Browser Cayley graph is empty, inserting test data.")
		err = loadtestdata.LoadTestData(ouchGraph)
		if err != nil {
			return nil, err
		}
	}

	return ouchGraph, nil
}

func cayleyLocal(paramMap map[string]interface{}) interface{} {

	g, err := graph(paramMap)
	if err != nil {
		return wrapErr("Parse", err.Error())
	}
	defer g.Close()

	fmt.Println("Query browser Cayley graph:", paramMap)

	type Params struct {
		Query string `json:"query"`
	}

	type Result struct {
		Data   interface{} `json:"data"`
		Errors []string    `json:"errors,omitempty"`
	}

	qry := paramMap["query"].(string)
	//fmt.Println("Query", qry)

	q, err := graphql.Parse(bytes.NewReader([]byte(qry)))
	if err != nil {
		return wrapErr("Parse", err.Error())
	}
	if q == nil {
		return wrapErr("Parse", "nil pointer returned")
	}
	//println("Parsed", q)

	m, err := q.Execute(context.TODO(), g.QuadStore)
	if err != nil {
		return wrapErr("Execute", err.Error())
	}
	//fmt.Println("Output", m)

	return map[string]interface{}{
		"data": m,
	}
}

func wrapErr(where, err string) interface{} {
	msg := where + " " + err
	return map[string]interface{}{
		"data": nil,
		"errors": []map[string]interface{}{
			map[string]interface{}{
				"message":   msg,
				"locations": nil,
			},
		},
	}
}

func main() {}
