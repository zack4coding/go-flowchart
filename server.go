package main 

import (
    // "os"
    "fmt"
    "net/http"
    "html/template"
    "go-flowchart/analysis"
    "go-flowchart/analysis/hierarchical"
)

var t = template.Must(template.ParseGlob("assets/*"))
// var a hierarchical.AnalysisData

var analyzeApi = "analyze?path="

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
    var path = r.FormValue("path")
    tree := analysis.ParsePath(path)
    var data hierarchical.LayoutData 
    data = hierarchical.Analyze(tree)
    
    fmt.Println("data:"+data.ToJsonString())
    // jsonFilePath := analysis.WCacheFile(path, data.ToJson())
    w.Header().Set("Content-Type", "text/json")
    w.Write(data.ToJson())
}

func handlePath(w http.ResponseWriter, r *http.Request) {
    
    var path = r.FormValue("path")
    t.ExecuteTemplate(w,"hierarchical.html", analyzeApi+path)
}

func main() {
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/", fs)
    mux.HandleFunc("/analyze", handleAnalyze)
    mux.HandleFunc("/path", handlePath)
    http.ListenAndServe(":8989", mux)

    // analysis
}