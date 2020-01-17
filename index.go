package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "encoding/json"
    "os"
)
//Define structure for using in json
type Detail struct {
    Name string
    Size int64
}

func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
    var list []Detail
    
    for _, f := range files {
  
        list=append(list, Detail{f.Name(), f.Size()})
    }


    b, err := json.Marshal(list)

    if err != nil {
        fmt.Println("error:", err)
    }

    os.Stdout.Write(b)
}