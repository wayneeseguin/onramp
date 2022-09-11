package main

import (
  "context"
  //"strings"
  "fmt"
  "os"
  //"github.com/jackc/pgx/v4"
  "net/http"

  "github.com/jackc/pgx/v4/pgxpool"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "<html><body>%s</body>!!!</html>", "Hello, Business World!!!")
  fmt.Fprintf(os.Stdout, "Hello, Business World!!!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, h := range headers {
      fmt.Fprintf(w, "%v: %v", name, h)
    }
  }
}

func main() {
  pool, err := pgxpool.Connect(context.Background(), os.Getenv("PG_URI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to pg: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

  var timestamp string
  stmt := "SELECT to_char(CURRENT_TIMESTAMP, 'YYYY-mm-dd HH24:MM:SS')"
  err = pool.QueryRow(context.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    os.Exit(1)
  }
  fmt.Fprintf(os.Stdout, "The postgres current timestamp is: %s\n", timestamp)

  http.HandleFunc("/", hello)
  http.HandleFunc("/headers", headers)

  port := os.Getenv("PORT")
  err = http.ListenAndServe(fmt.Sprintf(":%s",port), nil)
  if err != nil {
    fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
    os.Exit(1)
  }
}
