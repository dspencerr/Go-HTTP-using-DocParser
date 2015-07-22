package main

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/boltDb"
    "github.com/dspencerr/ComplianceApp/routes"
)

var bucket []byte

func main() {
    boltDb.Setup()

    app := martini.Classic()

    routes.Routing(app)

    app.Run();
}

