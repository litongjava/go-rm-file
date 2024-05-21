package main

import (
  "fmt"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "os"
  "path/filepath"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go-rm-file [path] [name]")
    return
  }
  path := os.Args[1]
  var name = os.Args[2]
  hlog.Info("Path and name:", path, name)

  deleteTaget(path, name)
}

func deleteTaget(path string, name string) {

  err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    stat, _ := os.Stat(path)
    if stat == nil {
      return nil
    }
    if err != nil {
      fmt.Println("error", err)
    }
    if info.IsDir() {
      folderName := info.Name()
      if folderName == name {
        err := os.RemoveAll(path)
        if err != nil {
          hlog.Info("delete faild ", path)
        } else {
          hlog.Info("deleted ", path)
          return nil
        }
      }
    }
    return nil
  })

  if err != nil {
    hlog.Info("Error:", err)
  }

  fmt.Println("Done")
}
