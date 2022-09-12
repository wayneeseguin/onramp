package api

import (
  "sort"
  "fmt"
  "os"
  "net/http"
  "time"

  "context"
  "onramp/pg"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
  var timestamp string
  stmt := "SELECT CURRENT_TIMESTAMP::varchar"
  err := pg.DB.QueryRow(context.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    return
  }
  fmt.Fprintf(os.Stdout, "Request at postgres timestamp: %s\n", timestamp)

  fmt.Fprintf(w, "<html><body><h1>%s</h1></body><footer>%s</footer></html>", "Hello, Business World!!!", timestamp)
}

func apiHeadersHandler(w http.ResponseWriter, req *http.Request) {
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
  http.HandleFunc("/api/v1/headers", apiHeadersHandler)

  err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
  if err != nil {
    fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
    os.Exit(1)
  }

  return
}

