package main

import (  
    "fmt"
    "time"
)

func repchar(symbol rune, amount int) {  
    for i := 0; i < amount; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Printf("%c ", symbol)
    }
}

func main() {  
    go repchar('A', 10)
    go repchar('B', 7)
    time.Sleep(1500 * time.Millisecond)
    fmt.Println(" - main terminated")
}

/* результаты четырёх запусков :
B A B A B A A B A B B A B A A A A  - main terminated
A B B A A B A B A B B A B A A A A  - main terminated
B A A B B A A B B A B A A B A A A  - main terminated
A B B A B A B A B A A B A B A A A  - main terminated
*/


