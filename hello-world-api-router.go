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
    router.HandleFunc("/status", StatusJson)
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
    var l string

    for fscanner.Scan() {
        s := strings.Split(fscanner.Text(), ":")
        if strings.TrimRight(s[0], "\n") == "description" {
            d = s[1]
        } else if strings.TrimRight(s[0], "\n") == "version" {
            v = s[1]
        } else if strings.TrimRight(s[0], "\n") == "lastcommitsha" {
            l = s[1]
        }
    }

    w.Header().Set("Content-Type", "application/json") 
    user := Return {d, v, l} 
    json.NewEncoder(w).Encode(user) 
}