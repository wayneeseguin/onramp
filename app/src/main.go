package main

import (
  "fmt"
  "os"

  api "onramp/api"
)

func main() {
  err := api.Api()
  if err != nil {
    fmt.Fprintf(os.Stderr, "api.Api() err: %v\n", err)
    os.Exit(1)
  }
}

