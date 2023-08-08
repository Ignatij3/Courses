package main

import "fmt"

func ExampleShowMap() {
    x := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range x {
        fmt.Printf("key=%s value=%d\n", key, value)
    }

    // Unordered output:
    // key=a value=1
    // key=b value=2
    // key=c value=3
}
