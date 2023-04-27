package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  string, textToBeReplaced, replacementText := os.Args[1], os.Args[2], os.Args[3]
  result := strings.ReplaceAll(string, textToBeReplaced, replacementText)
  fmt.Println(result)
}
