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
  stmt := "SELECT to_char(CURRENT_TIMESTAMP, 'YYYY-mm-dd HH24:MM:SS')"
  err = conn.QueryRow(ctx.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    os.Exit(1)
  }
  fmt.Fprintf(os.Stdout, "The postgres current timestamp is: %s\n", timestamp)
}


