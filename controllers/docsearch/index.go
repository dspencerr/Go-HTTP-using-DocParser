package docsearch

import (
    "github.com/codegangsta/martini"
    "github.com/dspencerr/docParser"
    "github.com/gernest/nutz"
    "os/user"
    "encoding/json"
    "strconv"
    "github.com/dspencerr/docParser/fileMgr"
    "os"
    "github.com/dspencerr/ComplianceApp/batch"
    //"github.com/dspencerr/docParser/fileWriter"
    "strings"
    "io/ioutil"
    "github.com/skratchdot/open-golang/open"
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

    if !updatePaths("sources", params["source"]) {
        return `{"result":"fail", "error":"Source directory not valid directory", "path":"source"}`
    }
    if !updatePaths("targets", params["target"]) {
        return `{"result":"fail", "error":"Target directory not valid directory", "path":"target"}`
    }
    parsedDocs, numParsed := docParser.ParseDocs(params["source"], params["target"])

    data, _ := json.Marshal(parsedDocs)

    s := db.Get(bucket, params["source"])
    if s.Error == nil {
        db.Delete(bucket, params["source"])
    }
    db.Create(bucket, params["source"], data)

    return `{"result":"success", "docs":"`+strconv.Itoa(numParsed)+`"}`
}

func GetDataSet(params martini.Params) string {
    initDB()

    if info, err := os.Stat(params["path"]); err != nil || !info.IsDir() {
        return `{"result":"fail", "error":"Source directory not valid directory", "path":"source"}`
    }

    d := db.Get(bucket, string(params["path"]))

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
    if len(data) == 0 {
        return `{"result":"empty"}`
    }

    result := data[ start:end ]
    res, _ := json.Marshal(result)

    return `{"data":`+string(res)+`, "total":"`+strconv.Itoa(len(data))+`"}`
    //return `[{"result":"success"}]`
}

func updatePaths(pathType string, path string) bool {
    if info, err := os.Stat(path); err != nil || !info.IsDir() {
        return false
    }
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
    return true
}


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func UpdateRevision(doc fileMgr.DocFile) string {
    initDB()

    d := db.Get(bucket, string(doc.Key))


    var docsFromBucket []fileMgr.DocFile
    json.Unmarshal(d.Data, &docsFromBucket)
    for indx, _ := range docsFromBucket {
        if docsFromBucket[indx].Path == doc.Path {
            //fmt.Println("WE IN: ", doc.Path)
            docsFromBucket[indx].Revision = doc.Revision
        }
    }

    data, _ := json.Marshal(docsFromBucket)
    db.Delete(bucket, string(doc.Key))
    db.Create(bucket, string(doc.Key), data)

    return `{"status":"success"}`
}

func UpdateRevisionBatch(obj batch.DocArray) string {
    initDB()

    dbBlog := db.Get(bucket, string(obj.Key))
    var docsFromBucket []fileMgr.DocFile
    json.Unmarshal(dbBlog.Data, &docsFromBucket)
    for dbIndx, _ := range docsFromBucket {
        for prIndx, _ := range obj.Data{

            if docsFromBucket[dbIndx].Path == obj.Data[prIndx].Path {
                //fmt.Println("WE IN: ", doc.Path)
                docsFromBucket[dbIndx].Revision = obj.Data[prIndx].Revision
            }

        }
    }

    data, _ := json.Marshal(docsFromBucket)
    db.Delete(bucket, string(obj.Key))
    db.Create(bucket, string(obj.Key), data)

    return `{"status":"success"}`
}


func ExportToCSV(obj batch.CsvExport) string {

    initDB()

    dbBlog := db.Get(bucket, string(obj.Source))
    var docsFromBucket []fileMgr.DocFile
    json.Unmarshal(dbBlog.Data, &docsFromBucket)

    writeDataToCSV(docsFromBucket, obj)

    return `{"status":"success"}`
}


func writeDataToCSV(docs []fileMgr.DocFile, obj batch.CsvExport){
    var dataToWrite string
    for x, doc := range docs{
        if x == 0{
            dataToWrite += "Name    Type    Revision"+"\n"
        }
        name := strings.TrimSuffix(doc.Name, ".zip")
        revision := strings.Replace(doc.Revision, "\n", " ", -1)
        dataToWrite += name+"\t"+doc.Type+"\t"+revision+"\n"
    }

    byteData := []byte(dataToWrite)
    //fmt.Println(byteData)

    file := obj.Target + "/" + obj.Name + ".txt"
    //fmt.Println(file)

    ioutil.WriteFile(file, byteData, 0777)
}


func OpenFileInApp(obj batch.FileOpen){
    open.Run(obj.File)
}
