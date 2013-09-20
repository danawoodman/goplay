package main

import(
  "net/http"
  "io"
  "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "hello world\n")
}

func main() {
  http.HandleFunc("/hello", HelloServer)
  err := http.ListenAndServe(":1234", nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
