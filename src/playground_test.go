package main

import (
	"testing"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
