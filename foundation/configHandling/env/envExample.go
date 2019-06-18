package main

import (
        "fmt"
        "net/http"
        "os"
)

func main() {
        fmt.Println(os.Getenv("Port"))
        http.HandleFunc("/", homePage)
        http.ListenAndServe(":"+os.Getenv("Port"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
        if req.URL.Path != "/" {
                http.NotFound(res, req)
                return
        }
        fmt.Fprint(res, "The homepage.")
}
