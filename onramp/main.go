package main

import (
  ctx "context"
  "fmt"
  "os"
  "github.com/jackc/pgx/v4"

  //"github.com/jackc/pgx/v4/pgxpool"
)


func main() {
  println("Hello, Business World!!!")

	conn, err := pgx.Connect(ctx.Background(), os.Getenv("PG_URI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to pg: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx.Background())

  var timestamp string
  stmt := "SELECT CURRENT_TIMESTAMP::string"
  err = conn.QueryRow(ctx.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    os.Exit(1)
  }
  fmt.Fprintf(os.Stdout, "The current timestamp from pg is: %s", timestamp)
}


