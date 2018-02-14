# cayley-graphiql

[Cayley graph](https://github.com/cayleygraph/cayley) adaptor for Facebook's [GraphQL IDE](https://github.com/graphql/graphiql) (including some of that project's code, so please also see their licence).

Full GraphQL is not supported, as Cayley's ["GraphQL-inspired" language](https://github.com/cayleygraph/cayley/blob/master/docs/GraphQL.md) is used.
This means that the "Documentation Explorer" section on the right of the IDE and "QUERY VARIABLES" section at the bottom of the IDE are redundant for their original use. It also means that the GraphQL validation code often marks valid Cayley-GraphQL-inspired-language as invalid ... which can be confusing initially.

Not everything works (there will probably be an error shown when you first start the page), but I've found it a useful basic tool for graph debugging when using [Cayley as a library](https://github.com/cayleygraph/cayley/blob/master/docs/Quickstart-As-Lib.md). It also demonstrates Cayley running in the browser and allows graph debuging in that context too.

The following query ([from the Cayley docs](https://github.com/cayleygraph/cayley/blob/master/docs/GraphQL.md)) works well:
```
{
  nodes {
    id
    follows {
      id
    }
  }
}
```

To query the same testdata cayley graph in the browser, add the following JSON to the "QUERY VARIABLES" section of the ide:
```
{
  "pouchdb":"cayley_testdata"
}
```

See this running from a static web page [here](https://elliott5.github.io/cayleygraphiql/index.html?query=%7B%0A%20%20nodes%20%7B%0A%20%20%20%20id%0A%20%20%20%20follows%20%7B%0A%20%20%20%20%20%20id%0A%20%20%20%20%7D%0A%20%20%7D%0A%7D&variables=%7B%0A%20%20%22pouchdb%22%3A%22cayley_testdata%22%0A%7D).

Whenever you provide a `"pouchdb"` database name, that database will be created in the browser if it does not exist. If that database has the name `"cayley_testdata"` then the [testdata used in the Cayley documentation](https://github.com/cayleygraph/cayley/blob/master/data/testdata.nq) will be inserted (if not already present).

To see it in action on your own machine `go run example/main.go` then open your browser at the url reported.

If you want to hack, `build.sh` works on OSX. You will need [GopherJS](https://gopherjs.github.io) installed and will probably need to `go get` some missing libraries.

Enjoy! PRs welcome.
