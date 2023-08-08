package main

import ( 
    "fmt"
    "os"
)

func main() { 
    f, err := os.Open("./nonexistent/file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

// Output:
// open ./nonexistent/file.txt: The system cannot find the path specified.
