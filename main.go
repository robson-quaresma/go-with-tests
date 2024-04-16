package main

import (
	"os"

	"github.com/robson-quaresma/go-with-tests/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}
