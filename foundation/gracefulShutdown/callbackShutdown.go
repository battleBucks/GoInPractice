package main

import (
        "fmt"
        "net/http"
        "os"
)

func main() {
        http.HandleFunc("/", homepage)
        http.HandleFunc("/shutdown", shutdown)
        http.ListenAndServe(":8080", nil)
}
func homepage(res http.ResponseWriter, req *http.Request) {
        if req.URL.Path != "/" {
                http.NotFound(res, req)
                return
        }
        fmt.Fprint(res, "The homepage")
}
func shutdown(res http.ResponseWriter, req *http.Request) {
        os.Exit(0)
}
