package main
 
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "bufio"
    "os"
    "strings"
 
    "github.com/gorilla/mux"
)
 
type Return struct { 
    Description    string     `json:"description"` 
    Version  string  `json:"version"` 
    Lastcommitsha  string  `json:"lastcommitsha"` 
}

func main() {
 
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/status", Status)
    router.HandleFunc("/statusjson", StatusJson)
    log.Fatal(http.ListenAndServe(":8080", router))

}
 
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World")
}

func StatusJson(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("./metadata")
    fscanner := bufio.NewScanner(file)
    var d string
    var v string

    for fscanner.Scan() {
        s := strings.Split(fscanner.Text(), ":")
        if strings.TrimRight(s[0], "\n") == "description" {
            d = s[1]
        } else if strings.TrimRight(s[0], "\n") == "version" {
            v = s[1]
        }
    }
    fmt.Println("StatusJson")
    fmt.Println(d)
    fmt.Println(v) 

    w.Header().Set("Content-Type", "application/json") 
    user := Return {d, v, "7h7h7h"} 
    json.NewEncoder(w).Encode(user) 
}

func Status(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("./metadata")
    fscanner := bufio.NewScanner(file)
    var status string
    var desc string
    var version string

    for fscanner.Scan() {
        status += fscanner.Text() + "\n"
        s := strings.Split(fscanner.Text(), ":")
        if strings.TrimRight(s[0], "\n") == "description" {
            desc = s[1]
        } else if strings.TrimRight(s[0], "\n") == "version" {
            version = s[1]
        }
    }
    fmt.Println(desc)
    fmt.Println(version) 

    fmt.Fprintf(w, "%s", status)
}