
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head><title>Greetings to the world</title></head><body><h1>new version of  the Jenkins X golang http example</h1></body></html>\n")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
