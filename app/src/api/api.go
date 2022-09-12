package api

import (
  "sort"
  "fmt"
  "os"
  "time"
  "net/http"
  "html/template"
  "path/filepath"

  "context"
  "onramp/pg"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
  var (
    err error
    timestamp string
  )
  stmt := "SELECT CURRENT_TIMESTAMP::varchar"
  err = pg.DB.QueryRow(context.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    return
  }

  p := filepath.Join(os.Getenv("APP_ROOT"),"public","html","onramp.html")
  t, err := template.ParseFiles(p)
  if err != nil {
    fmt.Fprintf(os.Stderr, "template.ParseFiles err: %v", err)
    return
  }

  t.Execute(w, struct { 
    Message string
    Timestamp string
    PGTimestamp string 
  } {
    Message: "Hello, Business World!!!",
    Timestamp: time.Now().Format(time.RFC3339),
    PGTimestamp: timestamp,
  })
}

func apiV1HeadersHandler(w http.ResponseWriter, req *http.Request) {
  timestamp := time.Now().Format(time.RFC3339)

  fmt.Fprintf(w, "<html><body><h1>Request Headers</h1><hr/><table><tr><th>Header</th><th>Value</th></tr>")

  headerKeys := make([]string, 0, len(req.Header))
  for name := range req.Header {
    headerKeys = append(headerKeys, name)
  }
  sort.Strings(headerKeys)
  for _, k := range headerKeys {
    for _,h := range req.Header[k] {
      fmt.Fprintf(w, "<tr><td>%v</td><td>%v</td></tr>", k, h) //req.Header[k])
    }
  }
  fmt.Fprintf(w, "</table></body><hr/><footer><div>processed_at: %s</div></footer></html>", timestamp)
}

func Api() (err error) {
  http.HandleFunc("/", rootHandler)
  http.HandleFunc("/api/v1/headers", apiV1HeadersHandler)

  err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
  if err != nil {
    fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
    os.Exit(1)
  }

  return
}

