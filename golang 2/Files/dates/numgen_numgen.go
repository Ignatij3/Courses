package main

import  (
	"fmt"
	"encoding/binary"
	"math/rand"
	"os"
)
	
func main()  {
	if len(os.Args) != 3  {
		incorrectCommandLine()
		return
	}	
	f, err := os.Create(os.Args[2])
	if err != nil  { 
		fmt.Println(err)
		return
	}
	defer f.Close()	
	
	rand.Seed(20040225)

	var n int32
	buff := make([]byte, 4)
	switch os.Args[1]  {
	case "-b":
		for  {
			n = rand.Int31n(100000000)
			if n < 100  {
				break
			}
			binary.LittleEndian.PutUint32(buff, uint32(n))
			_, err = f.Write(buff)
			if err != nil  {
				fmt.Println(err)
				return	
			}
		}		
	case "-t":
		for  {
			n = rand.Int31n(100000000)
			if n < 100  {
				break
			}
			fmt.Fprintln(f, n)	
		}		
	default:
		incorrectCommandLine()
		return
	}	
	
}

func incorrectCommandLine()  {
	fmt.Println("Usage: numgen.exe output_type output_file")
	fmt.Println(" output_type:")  		
	fmt.Println("    -b  binary output")
	fmt.Println("    -t  text output; one number per line")
	return
}	
