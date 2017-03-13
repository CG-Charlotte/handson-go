package main

import (
	"test"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
