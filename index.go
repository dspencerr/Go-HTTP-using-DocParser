package main

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/boltDb"
    "github.com/dspencerr/ComplianceApp/routes"
    "github.com/skratchdot/open-golang/open"
//    "time"
//    "fmt"
)

var bucket []byte

func main() {
    boltDb.Setup()

    app := martini.Classic()

    routes.Routing(app)
    open.Start("http://localhost:3000")
    app.Run()
    //
}

