package main


import (


    "os/exec"
    "fmt"
    "log"
)


func main() {

    //app := "go-bindata.exe"
    out, err := exec.Command(
        "build/go-bindata.exe", "-o",
        "static/static.go",
        "assets/index.html",
        "assets/css/app.css",

        "assets/js/docsearch/docsearch-controller.js",
        "assets/js/home/home-controller.js",
        "assets/js/app.js",

        "assets/lib/ag-grid/dist/angular-grid.css",
        "assets/lib/ag-grid/dist/angular-grid.js",

        "assets/lib/bootstrap-css/css/bootstrap.css",

        "assets/lib/jquery-ui/themes/smoothness/jquery-ui.css",
        "assets/lib/ag-grid/dist/theme-fresh.css",
        "assets/lib/jquery/dist/jquery.js",
        "assets/lib/jquery-ui/jquery-ui.js",
        "assets/lib/lodash/dist/lodash.js",
        "assets/lib/angular/angular.js",
        "assets/lib/angular-resource/angular-resource.js",
        "assets/lib/angular-route/angular-route.js",
        "assets/lib/angular-bootstrap/ui-bootstrap-tpls.js",
        "assets/lib/angular-ui-date/src/date.js",

        "assets/views/home/home.html",
        "assets/views/docsearch/docsearch.html",

    ).Output()

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(out)
    }


}
