package routes


import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/ComplianceApp/static"
    "fmt"
    "net/http"
    "strings"
)


func rootServer() string {
    data, err := static.Asset("assets/index.html")
    if err != nil {
        fmt.Println(err)
    }
    return string(data)
}

func setFileHeader(res http.ResponseWriter, req *http.Request) {
    if strings.Contains(req.RequestURI, ".css") {
        res.Header().Set("Content-Type", "text/css")
    } else if strings.Contains(req.RequestURI, ".js"){
        res.Header().Set("Content-Type", "application/javascript")
    } else if strings.Contains(req.RequestURI, ".html"){
        res.Header().Set("Content-Type", "text/html")
    }
}

func jsFiles(params martini.Params) string {
    data, err := static.Asset("assets/js/"+params["_1"])
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(data)
    return string(data)
}

func cssFiles(params martini.Params) string {
    data, err := static.Asset("assets/css/"+params["_1"])
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(data)
    return string(data)
}

func libsFiles(params martini.Params) string {
    fmt.Println("Libs PARAMS: ", "assets/css/"+params["_1"])
    data, err := static.Asset("assets/lib/"+params["_1"])
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(data)
    return string(data)
}

func htmlFiles(params martini.Params) string {
    data, err := static.Asset("assets/views/"+params["_1"])
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(data)
    return string(data)
}
