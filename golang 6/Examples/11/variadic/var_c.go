package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

func main() { 
	nn:= []int{89, 90, 95} 
    find(89, nn...)                 // type of nums is []int    
                                    // 89 found at index 0 in [89 90 95] 
                                    //
	nn = []int{56, 67, 45, 90, 109} 
    find(45, nn...)                 // type of nums is []int
                                    // 45 found at index 2 in [56 67 45 90 109]
                                    //
    find(78, []int{38, 56, 98}...)  // type of nums is []int
                                    //  78 not found in  [38 56 98]                       
                                    //
}
