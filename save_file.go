package main

import (
    "fmt"
    "log"
    "os"
    "io/ioutil"
    "net/http"
)

func main() {

  response, err := http.Get("http://frameworks-staging.act.org/ims/case/v1p0/CFPackages/5156aa6c-3bbb-11e9-abf4-1b53fdf9e1d4.json")

  if err != nil {
      fmt.Print(err.Error())
      os.Exit(1)
  }

  responseData, err := ioutil.ReadAll(response.Body)
  if err != nil {
      log.Fatal(err)
  }

    file, err := os.Create("file.json")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()


    fmt.Fprintf(file,string(responseData))
  }
