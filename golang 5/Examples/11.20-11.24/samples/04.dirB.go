package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
// Выводит список всех файлов и каталогов в каталоге, 
// переданном в командной строке, и во всех его подкаталогах.
// По умолчанию - начиная с текущего каталога.
    var headDir string
    if len(os.Args) == 1  {
        headDir = "."
    }  else  {
        headDir = os.Args[1]
    }
    traceDir(headDir, "")
}

func traceDir (startDir string, prefixStr string)  {
    if files, err := ioutil.ReadDir(startDir); err == nil {
        for _, f := range files {
            if !f.IsDir()  {
                fmt.Println(prefixStr, startDir + "\\" + f.Name())
            }    
        }    
        fmt.Println()
    }
    if files, err := ioutil.ReadDir(startDir); err == nil {
        for _, f := range files {
            if f.IsDir()  {
                traceDir(startDir+"\\"+f.Name(), prefixStr)
            }    
        }    
    }
}    
