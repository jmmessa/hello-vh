package main

import (
        "os"
        "net/http"
)

var (
        version = "2.0"
        hostname, err = os.Hostname()
)

func versionHandler(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type","text/plain")
        w.Header().Set("Version",version)
        w.Write([]byte(hostname))
        w.Write([]byte(" is running "))
        w.Write([]byte(version))
        w.Write([]byte("\n"))
}

func handler(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type","text/plain")
        w.Write([]byte("Hello world !\n"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/version", versionHandler)
    http.ListenAndServe(":80", nil)
}

