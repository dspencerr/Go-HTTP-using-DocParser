package batch


type DocArray struct  {
    Key string
    Data []DocFile
}


type DocFile struct {
    Path string
    Revision string
}

type CsvExport struct {
    Source string
    Target string
    Name string
}


type FileOpen struct {
    File string
}
