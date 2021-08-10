package main

import  ( 
    "fmt"
    "unsafe"
)
    
type (
	Slice struct {
        start   unsafe.Pointer
        len     int
        cap     int
	}
)	

func sliceInfo(s []int) Slice  {
    return *(*Slice)(unsafe.Pointer(&s))
}

func main() {
	b:= make([]int, 0, 1)
	fmt.Println(b)          // []  
	fmt.Println(sliceInfo(b))           // {0xc000048058 0 1}
	b = append(b, 7)
	fmt.Println(b)          // [7]  
	fmt.Println(sliceInfo(b))           // {0xc000048058 1 1}
	b = append(b, 7)
	fmt.Println(b)          // [7 7]  
	fmt.Println(sliceInfo(b))           // {0xc0000480b0 2 2}
	b = append(b, 7)
	fmt.Println(b)          // [7 7 7]  
	fmt.Println(sliceInfo(b))           // {0xc0000460e0 3 4}
	fmt.Println()  
	c:= cap(b)
	for i:= 0; i<1000000; i++  {
		b = append(b, i)                // {0xc000078040 5 8}
		if cap(b) != c  {               // {0xc00007a080 9 16}
			fmt.Println(sliceInfo(b))   // {0xc000034100 17 32}
			c = cap(b)                  // {0xc000086000 33 64}
		}                               // {0xc000088000 65 128}
	}			                        // {0xc00008a000 129 256}
                                        // {0xc00008c000 257 512}
                                        // {0xc00008e000 513 1024}
                                        // {0xc000090000 1025 1280}
                                        // {0xc00009a000 1281 1696}
                                        // {0xc0000a4000 1697 2304}
                                        // {0xc0000b6000 2305 3072}
                                        // {0xc0000bc000 3073 4096}
                                        // {0xc0000c4000 4097 5120}
                                        // {0xc0000ce000 5121 7168}
                                        // {0xc0000dc000 7169 9216}
                                        // {0xc0000ee000 9217 12288}
                                        // {0xc000106000 12289 15360}
                                        // {0xc000124000 15361 19456}
                                        // {0xc00014a000 19457 24576}
                                        // {0xc00017a000 24577 30720}
                                        // {0xc0001b6000 30721 38912}
                                        // {0xc000202000 38913 49152}
                                        // {0xc000262000 49153 61440}
                                        // {0xc0002da000 61441 76800}
                                        // {0xc000370000 76801 96256}
                                        // {0xc000082000 96257 120832}
                                        // {0xc00016e000 120833 151552}
                                        // {0xc00044e000 151553 189440}
                                        // {0xc0005c0000 189441 237568}
                                        // {0xc000082000 237569 296960}
                                        // {0xc00044e000 296961 371712}
                                        // {0xc000082000 371713 464896}
                                        // {0xc00044e000 464897 581632}
                                        // {0xc0008be000 581633 727040}
                                        // {0xc000e4a000 727041 909312}
                                        // {0xc00044e000 909313 1136640}
}
