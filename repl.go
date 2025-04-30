package main

import (
  "bufio"
  "fmt"
  "os"
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
