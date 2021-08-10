package main

import (
	"fmt"
	"os"
	"encoding/binary"
	"math"
	"strconv"
)

func main() {
	var b uint64
	data:= []string{"0.0", "1.0", "-1.0", "2.5", "-2.5", "2.8", "-0.3", "0.7", 
					"1.798E+308", "2.23E-308", "2.2E-308", "8E-324", "5.5E-324", "2.5E-324", "2E-324"}
	buff:= make([]byte, 8)
	fout, _ := os.Create("float64.entry")

	for _, s := range data  {
		if x, err:= strconv.ParseFloat(s, 64); err == nil  {
			b = math.Float64bits(x)
			binary.BigEndian.PutUint64(buff, b)
			for _, buffi := range buff  {
				fmt.Fprintf(fout, "%02X ", buffi)
			}
			fmt.Fprintf(fout,"%10s %+28.18E\n", s, x)	
		}	 
	}	
	fout.Close()
	
}
 
