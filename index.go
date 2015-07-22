package main

import (
	"github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/controllers/docsearch"
    "github.com/dspencerr/ComplianceApp/boltDb"
)

var bucket []byte

func main() {
    boltDb.Setup()

    app := martini.Classic()
	app.Get("/", rootFallBack)
    app.Get("/docsearch/settings", docsearch.GetSettings)
	app.Get("/docsearch/source/:source/target/:target", docsearch.RunSearch)
    app.Get("/docsearch/results/:start/:length", docsearch.GetDataSet)
	app.Run();
}

func rootFallBack() string {
    return "Well hello there friend. What happened to the public folder?"
}
