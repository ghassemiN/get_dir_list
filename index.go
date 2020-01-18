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
)


//Define structure for using in json
type Detail struct {
    Name string
    Size int64
    Hash string
}


//Generate Hash of files
func hash_file_md5(filePath string) (string, error) {
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
    var list []Detail

    //read a directory
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    //get detail from each elemnt in the dir
    for _, f := range files {
        
        hash, err := hash_file_md5(f.Name())
        if err == nil {
            fmt.Println(hash)
        }
        list=append(list, Detail{f.Name(), f.Size(), hash})
  
    }

    //convert to json
    b, err := json.Marshal(list)

    if err != nil {
        fmt.Println("error:", err)
    }
    os.Stdout.Write(b)

}