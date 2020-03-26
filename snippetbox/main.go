package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return
    }

    w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Show some Snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create some Snippet"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("starting server on port :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
