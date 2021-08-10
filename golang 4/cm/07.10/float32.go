package main

import (
	"fmt"
	"os"
	"encoding/binary"
	"math"
	"strconv"
)

func main() {
	var b uint32
	data:= []string{"0.0", "1.0", "-1.0", "2.5", "-2.5", "2.8", "-0.3", "0.7", 
					"3E-39", "2.5E-42", "5E-45", "1.5E-45", "1.2E-45", "1E-45"}
	buff:= make([]byte, 4)
	fout, _ := os.Create("float32.entry")

	for _, s := range data  {
		if x, err:= strconv.ParseFloat(s, 32); err == nil  {
			b = math.Float32bits(float32(x))
			binary.BigEndian.PutUint32(buff, b)
			for _, buffi := range buff  {
				fmt.Fprintf(fout, "%02X ", buffi)
			}
			fmt.Fprintf(fout,"%9s %+24.14E\n", s, float32(x))	
		}	 
	}	
	fout.Close()
	
}
 
