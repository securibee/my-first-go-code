package main

import (
  "testing"
  "os"
)

func TestHello(t *testing.T) {
  t.Run("say 'Hello World' if no argument provided", func(t *testing.T) {
    got := Hello("")
    want := "Hello World"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })
  t.Run("saying hello to provided argument", func(t *testing.T) {
    got := Hello("Bob")
    want := "Hello Bob"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  t.Run("saying hello to host", func(t *testing.T) {
    host, _ := os.Hostname()
    got := Hello("host")
    want := "Hello " + host

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

}
