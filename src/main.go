package main

import (
	"fmt"
	"os"
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

  for _, files := range files {
    fmt.Println(files.Name())
  }
  // println(files)
}
