package main

import (
	"fmt"
	"os"
	color "otaviocosta2110/goEx/src/middleware"
)


func main() {
  dir := "."

  d, err := os.Open(dir)

  if err != nil{
    fmt.Println("Error: ", err)
    return
  }

  defer d.Close()

  files, err := d.Readdir(-1)

  for _, file := range files {
    if file.IsDir() {
      fmt.Println(color.Blue, file.Name(), color.Reset)
    } else {
      fmt.Println(file.Name(), color.Reset)
    }
  }
}
