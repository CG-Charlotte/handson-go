package main

import (
	"testing"

	"cmd"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
