package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart()  {
   
for {
   
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print(">")

  scanner.Scan()
  text := scanner.Text()

  fmt.Println("echoing:", text)

  }

}

func cleanInput(str string) []string {
  lowered := strings.ToLower(str)
  words := strings.Fields(lowered)
  return words
  
}
