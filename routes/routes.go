package routes

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/controllers/docsearch"
)


func Routing(app *martini.ClassicMartini){
    staticFileRoutes(app)

    docSearchRoutes(app)
}

func docSearchRoutes(app *martini.ClassicMartini){
    app.Get("/docsearch/settings", docsearch.GetSettings)
    app.Get("/docsearch/source/:source/target/:target", docsearch.RunSearch)
    app.Get("/docsearch/results/:start/:length", docsearch.GetDataSet)
}

func staticFileRoutes(app *martini.ClassicMartini){
    app.Get("/", rootServer)
    app.Get("/lib/**", setFileHeader, libsFiles)
    app.Get("/js/**", setFileHeader, jsFiles)
    app.Get("/css/**", setFileHeader, cssFiles)
    app.Get("/views/**", setFileHeader, htmlFiles)
}
