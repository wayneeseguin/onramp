package main

import (
  "testing"
)

func TestTruth(t *testing.T) {
  if 1 != 1 {
    t.Error("Truth does not exist!")
  }
}

