package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go-rm-file [path] [suffix]")
    return
  }
  path := os.Args[1]
  suffix:=os.Args[2]
  fmt.Println("Path,suffix", path,suffix)

  err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }

    filename := info.Name()
    if !info.IsDir() && strings.HasSuffix(filename, "."+suffix) {
      err = os.Remove(filePath)
      if err != nil {
        fmt.Println("delete fail:{}", filename)
        return err
      } else {
        fmt.Println("deleted", filename)
      }
    }

    return nil
  })

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Done")
}
