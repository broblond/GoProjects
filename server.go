package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe("127.0.0.1:3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Server", "A Go Web Server")
  w.WriteHeader(200)
}
