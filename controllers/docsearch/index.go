package docsearch

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/docParser"
    "github.com/gernest/nutz"
    "os/user"
    "encoding/json"
    "strconv"
    //"fmt"
    "github.com/dspencerr/docParser/fileMgr"
)

var bucket = "docsearch"
var dbName string
var db nutz.Storage

func initDB(){
    usr, _ := user.Current()
    dbName = string(usr.HomeDir)+"/go-compliance.db"
    db = nutz.NewStorage(dbName, 0600, nil)
}

type settings struct  {
    Result string
    Sources []string
    Targets []string
}

func GetSettings() string {
    initDB()
    t := db.Get(bucket, "targets")
    s := db.Get(bucket, "sources")

    result := settings{ Result : "success" }

    if t.Error == nil{
        json.Unmarshal(t.Data, &result.Targets)
    }
    if s.Error == nil{
        json.Unmarshal(s.Data, &result.Sources)
    }

    data, _ := json.Marshal(result)
    return string(data)
}

func RunSearch(params martini.Params) string {
    initDB()

    updatePaths("sources", params["source"])
    updatePaths("targets", params["target"])
    parsedDocs, numParsed := docParser.ParseDocs(params["source"], params["target"])

    data, _ := json.Marshal(parsedDocs)


    s := db.Get(bucket, params["source"])
    if s.Error == nil {
        db.Delete(bucket, params["source"])
    }
    db.Create(bucket, params["source"], data)

    return `{"result":"success", "docs":"`+strconv.Itoa(numParsed)+`"}`

    //s := db.Get(bucket, params["source"])
    //return string(s.Data)
}

func GetDataSet(params martini.Params) string {
    initDB()
    s := db.Get(bucket, "sources")
    var sources []string
    json.Unmarshal(s.Data, &sources)
    d := db.Get(bucket, string(sources[0]))

    var data []fileMgr.DocFile
    json.Unmarshal(d.Data, &data)

    start, _ := strconv.Atoi(params["start"])
    length, _ := strconv.Atoi(params["length"])
    end := start + length
    if len(data) <= end {
        end = len(data)
    }
    if start >= end {
        start = end - 1
    }
    result := data[ start:end ]

    res, _ := json.Marshal(result)

    return string(res)

    //return `[{"result":"success"}]`
}

func updatePaths(pathType string, path string){
    r := db.Get(bucket, pathType)
    var paths []string
    if r.Error != nil {
        paths := append(paths, path)
        data, _ := json.Marshal(paths)
        db.Create(bucket, pathType, data)
    } else {
        json.Unmarshal(r.Data, &paths)
        testIn := stringInSlice(path, paths)
        if !testIn {
            paths = append(paths, path)
        }
        data, _ := json.Marshal(paths)
        db.Create(bucket, pathType, data)
    }
}


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


























//func Init(params martini.Params) string {
//
//    bucket = []byte("docsearch")
//
//    source_path := params["source"]
//    target_path := params["target"]
//
//    data := updatePaths(source_path, target_path)
//
//    return data
////    data := docParser.ParseDocs(source_path, target_path)
////
////    jsonData, _ := json.Marshal(data)
////    return string(jsonData);
//
//
//    //return `[{"restul":"finished"}]`
//}
//
//func updatePaths(source string, target string) string{
//
//    //srcKey :=
//
//    //var sources []string
//    //sources = []string{source}
//    //data, _ := json.Marshal(sources)
//
//    boltDb.Update([]byte("tt"), []byte(source), bucket)
//
//    d := boltDb.Find([]byte("tt"), bucket)
//
//    fmt.Println("Length: ", len(d))
//
//    fmt.Println(string(d))
//
////    var n []string
////    fmt.Println("Let's marshal it")
////    json.Unmarshal(d, n)
////    fmt.Println("We unmarshalled it")
////    fmt.Printf("%sn", n)
//
////    if len(val) == 0 {
////    } else {
////    }
//
//    return `[{"restul":"finished"}]`
//}
