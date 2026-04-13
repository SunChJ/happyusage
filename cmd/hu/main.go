package main

import (
	"os"

	happyusage "github.com/SunChJ/happyusage"
)

func main() {
	os.Exit(happyusage.Main("hu", os.Args[1:]))
}
