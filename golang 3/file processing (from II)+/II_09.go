// Command-line arguments
// are a common way to parameterize execution of programs.

// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import  (
	"fmt"
	"os"
)	

func main() {

    // `os.Args` provides access to raw command-line
    // arguments. Note that the first value in this slice
    // is the path to the program, and `os.Args[1:]`
    // holds the arguments to the program.
    fmt.Println("command-line arguments with program: \n", os.Args)
    fmt.Println("command-line arguments without program: \n", os.Args[1:])

    // You can get individual args with normal indexing.
    if len(os.Args) > 3  {
		arg := os.Args[3]
		fmt.Println(arg)
	}	

    fmt.Print("command-line contains ", len(os.Args)-1, " argument(s) ")
    fmt.Println("/to say nothing of the path to the program/.")
    for n, arg := range os.Args  {
		fmt.Printf("%3d. %s\n", n, arg)
	}	
}

/*
For such command-line:
  H:\Work. GO\command_line\command_line.exe a1 a2 a3 "a4 a41" a5

the program displays:
command-line arguments with program: 
 [H:\Work. GO\command_line\command_line.exe a1 a2 a3 a4 a41 a5]
command-line arguments without program: 
 [a1 a2 a3 a4 a41 a5]
a3
command-line contains 5 argument(s) /to say nothing of the path to the program/.
  0. H:\Work. GO\command_line\command_line.exe
  1. a1
  2. a2
  3. a3
  4. a4 a41
  5. a5
*/
