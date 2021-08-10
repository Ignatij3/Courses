/*package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed( time.Now().UnixNano())
    var bytes int

    for i:= 0 ; i < 10 ; i++{ 
        bytes = rand.Intn(100)
        fmt.Println(bytes)
        }
}*/


package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Create("h.txt") // creating...
    if err != nil {
        fmt.Printf("error creating file: %v", err)
        return
    }
    defer f.Close()
    for i := 0; i < 10; i++ { // Generating...
        _, err = f.WriteString(fmt.Sprintf("%d\n", i)) // writing...
        if err != nil {
            fmt.Printf("error writing string: %v", err)
        }
    }
}


/*package main

import (
    crypto_rand "crypto/rand"
    "encoding/binary"
    math_rand "math/rand"
)

func init() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}*/
