package main 
  
import "fmt"
  
func main() { 
	slc0:= make([]int,5)
	fmt.Println(slc0)

    // Creating slices 
    slc1:= []string{"one", "two", "three", "four","five", "six", "seven", "eight"} 
    var slc2 []string
    slc3:= make([]string, 5) 
    slc4:= []string{"eleven", "twelve", "thirteen", "fourteen"} 
  
    // Before copying 
    fmt.Println("Slice_1:", slc1) 
    fmt.Println("Slice_2:", slc2) 
    fmt.Println("Slice_3:", slc3) 
    fmt.Println("Slice_4:", slc4) 
  
    // Copying the slices 
    copy_1 := copy(slc2, slc1) 
    fmt.Println("\nSlice:", slc2) 
    fmt.Println("Total number of elements copied:", copy_1) 
  
    copy_2 := copy(slc3, slc1) 
    fmt.Println("\nSlice:", slc3) 
    fmt.Println("Total number of elements copied:", copy_2) 
  
    copy_3 := copy(slc3, slc4) 
    fmt.Println("\nSlice:", slc3) 
    fmt.Println("Total number of elements copied:", copy_3) 
  
    copy_4:= copy(slc1, slc4) 
    fmt.Println("\nSlice:", slc1) 
    fmt.Println("Total number of elements copied:", copy_4) 
      
    copy_5:= copy(slc1, slc1[3:]) 
    fmt.Println("\nSlice:", slc1) 
    fmt.Println("Total number of elements copied:", copy_5) 
      
// [0 0 0 0 0]
// Slice_1: [one two three four five six seven eight]
// Slice_2: []
// Slice_3: [    ]
// Slice_4: [eleven twelve thirteen fourteen]
//
// Slice: []
// Total number of elements copied: 0
//
// Slice: [one two three four five]
// Total number of elements copied: 5
//
// Slice: [eleven twelve thirteen fourteen five]
// Total number of elements copied: 4
//
// Slice: [eleven twelve thirteen fourteen five six seven eight]
// Total number of elements copied: 4
//
// Slice: [fourteen five six seven eight six seven eight]
// Total number of elements copied: 5

} 
