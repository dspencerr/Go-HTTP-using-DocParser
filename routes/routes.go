package routes

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/controllers/docsearch"
    "github.com/martini-contrib/binding"
    "github.com/dspencerr/docParser/fileMgr"
    "github.com/dspencerr/ComplianceApp/batch"
)



func Routing(app *martini.ClassicMartini){
    staticFileRoutes(app)

    docSearchRoutes(app)
}

func docSearchRoutes(app *martini.ClassicMartini){
    app.Get("/docsearch/settings", docsearch.GetSettings)
    app.Get("/docsearch/source/:source/target/:target", docsearch.RunSearch)
    app.Get("/docsearch/results/:path/:start/:length", docsearch.GetDataSet)
    app.Post("/docsearch/save-revision", binding.Bind(fileMgr.DocFile{}), docsearch.UpdateRevision)
    app.Post("/docsearch/save-revision-batch", binding.Bind(batch.DocArray{}), docsearch.UpdateRevisionBatch)
    app.Post("/docsearch/export-to-csv", binding.Bind(batch.CsvExport{}), docsearch.ExportToCSV)
    app.Post("/docsearch/open-file-in-app", binding.Bind(batch.FileOpen{}), docsearch.OpenFileInApp)
}

func staticFileRoutes(app *martini.ClassicMartini){
    app.Get("/", rootServer)
    app.Get("/lib/**", setFileHeader, libsFiles)
    app.Get("/js/**", setFileHeader, jsFiles)
    app.Get("/css/**", setFileHeader, cssFiles)
    app.Get("/views/**", setFileHeader, htmlFiles)
}
