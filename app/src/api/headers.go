package api

import (
  "sort"
  "fmt"
  "os"
  "time"
  "net/http"
  "html/template"
  "path/filepath"
)

func apiV1HeadersHandler(w http.ResponseWriter, req *http.Request) {
  timestamp := time.Now().Format(time.RFC3339)

  headerKeys := make([]string, 0, len(req.Header))
  for name := range req.Header {
    headerKeys = append(headerKeys, name)
  }
  sort.Strings(headerKeys)
  headers := make(map[string]string)
  for _, k := range headerKeys {
    for _,v := range req.Header[k] {
      headers[k] = v
    }
  }

  p := filepath.Join(os.Getenv("APP_ROOT"),"public","en","html","v1_headers.html")
  tpl, err := template.ParseFiles(p)
  if err != nil {
    fmt.Fprintf(os.Stderr, "template.ParseFiles err: %v", err)
    return
  }

  tpl.Execute(w, struct { 
    Headers map[string]string
    Timestamp string
  } {
    Headers: headers,
    Timestamp: timestamp,
  })
}

