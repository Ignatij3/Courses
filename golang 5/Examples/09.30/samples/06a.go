package main

import "fmt"

func process(arr [3]int) {  
    fmt.Println(arr)     // [2020 2021 2022]
    for i, _ := range arr {
		arr[i] *= 2
	}
    fmt.Println(arr)     // [4040 4042 4044]
}

func main() {  
    a := [3]int{2020, 2021, 2022}
    fmt.Println(a)       // [2020 2021 2022]
    process(a)
    fmt.Println(a)       // [2020 2021 2022]
}
