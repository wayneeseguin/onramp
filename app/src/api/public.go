package api

import (
  "fmt"
  "os"
  "net/http"
  "path/filepath"
)

func cssHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(os.Stdout, "cssHandler() req.URL.Path: %v\n", req.URL.Path)
  cssPath := filepath.Join(os.Getenv("APP_ROOT"),"public", req.URL.Path)
  //http.Handle("/css", http.FileServer(http.Dir(css)))
  http.ServeFile(w, req, cssPath)
}

func jsHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(os.Stdout, "jsHandler() req.URL.Path: %v\n", req.URL.Path)
  jsPath := filepath.Join(os.Getenv("APP_ROOT"),"public", req.URL.Path)
  http.ServeFile(w, req, jsPath)
}

