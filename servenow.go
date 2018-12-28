package servenow

import (
  "log"
  "net/http"
)


func ServeNow() {
  log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
