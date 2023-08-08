package main

import ( 
    "fmt"
    "os"
)

func main() { 
    f, err := os.Open("nonexistent.file")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

// Output:
// open nonexistent.file: The system cannot find the file specified.
