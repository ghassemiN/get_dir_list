package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "encoding/json"
    "os"
    "crypto/md5"
    "encoding/hex"
    "io"
    "net/http"
)


//Define structure for using in json
type Detail struct {
    Name string
    Size int64
    Hash string
}


//Generate Hash of files
func hashFileMd5(filePath string) (string, error) {
    var returnMD5String string
    file, err := os.Open(filePath)
    if err != nil {
        return returnMD5String, err
    }
    defer file.Close()
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return returnMD5String, err
    }
    hashInBytes := hash.Sum(nil)[:16]
    returnMD5String = hex.EncodeToString(hashInBytes)
    return returnMD5String, nil

}


func main() {
    http.HandleFunc("/", servePage)
    http.ListenAndServe(":8080", nil)
}


func servePage(writer http.ResponseWriter, reqest *http.Request) {

    var list []Detail

    //read a directory
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    //get detail from each elemnt in the dir
    for _, f := range files {
        
        hash, err := hashFileMd5(f.Name())
        if err == nil {
            fmt.Println(hash)
        }
        list=append(list, Detail{f.Name(), f.Size(), hash})
  
    }  
    b, err := json.Marshal(list)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }

  writer.Header().Set("Content-Type", "application/json")
  writer.Write(b)
}