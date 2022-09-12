package pg

import (
  "context"
  //"strings"
  "fmt"
  "os"
  //"github.com/jackc/pgx/v4"
  //"github.com/jackc/pgx/v4/pgxpool"
  "github.com/jackc/pgx/v5"
  "github.com/jackc/pgx/v5/pgxpool"
  //pgxUUID "github.com/vgarvardt/pgx-google-uuid/v4"
)

var (
  DB *pgxpool.Pool
  config *pgxpool.Config
)

func init() {
  var err error
  config, err = pgxpool.ParseConfig(os.Getenv("PG_URI"))
  if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse PG_URI: '%v'\n", os.Getenv("PG_URI"))
		os.Exit(1)
  }

  config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
    //pgxUUID.Register(conn.TypeMap())
    return nil
  }

  DB, err = pgxpool.NewWithConfig(context.Background(), config)
  if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to pg: %v\n", err)
    os.Exit(1)
  }

  //conn, err := DB.Acquire(context.Background())
  //require.NoError(b, err)
  //defer conn.Release()
}

