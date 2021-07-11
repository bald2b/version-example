package main

import "fmt"

var version = "unknown"
var gitBranch = "unknown"
var buildDate = "unknown"

func main() {
	fmt.Printf("Version: %s, branch: %s, date: %s\n", version, gitBranch, buildDate)
}
