package main

import  (
	"fmt"
	"errors"
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

	var (
		year uint16
		month byte
		day byte
	)	
	buff := make([]byte, 4)


	for  {
		year = uint16(rand.Int31n(420) + 1600)
		month = byte(rand.Int31n(12) + 1)
		day, err = randDay(year, month)
		if err != nil  {
		fmt.Println(err)
			return	
		}

		switch os.Args[1]  {
		case "-b":
			buff[0] = byte(year >> 8)
			buff[1] = byte(year & 0xff)
			buff[2] = month
			buff[3] = day
			_, err = f.Write(buff)
			if err != nil  {
				fmt.Println(err)
				return	
			}
		case "-t":
			_, err = fmt.Fprintf(f, "%4d %2d %2d\n", year, month, day)	
			if err != nil  {
				fmt.Println(err)
				return	
			}		
		default:
			incorrectCommandLine()
			return
		}
		if rand.Int31n(50000) == 0  { break }		
	}
}

func randDay(year uint16, month byte) (day byte, err error)  {
	switch 	month  {
	case 1, 3, 5, 7, 8, 10, 12:
		day = byte(rand.Int31n(31)) + 1
	case 4, 6, 9, 11:
		day = byte(rand.Int31n(30)) + 1
	case 2:
		if leapYear(year)  {
			day = byte(rand.Int31n(29)) + 1
		}  else  {
			day = byte(rand.Int31n(28)) + 1
		}
	default:
		err = errors.New("Invalid month")
	}	
	return
}
	
func leapYear(year uint16) bool  {
	if year % 4 != 0  {
		return false	
	}  else  
	if year % 100 != 0  {
		return true
	}  else  {
	// year % 100 == 0
		century:= year / 100
		return century % 4  == 0
	}			
}		
	
func incorrectCommandLine()  {
	fmt.Println("Usage: dategen.exe output_type output_file")
	fmt.Println(" output_type:")  		
	fmt.Println("    -b  binary output")
	fmt.Println("    -t  text output; one number per line")
	return
}	
