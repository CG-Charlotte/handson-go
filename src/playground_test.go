package main

import (
	"testing"
	"main"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
