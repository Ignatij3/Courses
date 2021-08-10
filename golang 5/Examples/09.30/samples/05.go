package main

import  (
	"fmt"
	"unsafe"
)	

type tag struct {
	order 	byte
	article	*string
	size 	int
	active 	bool
}	 

func main() {  
	s:= "sample struct"
    v := tag{11, &s, 123456, true}
    fmt.Println("v:", v)
    fmt.Printf("v. address: %p  size: %d   alignment: %d\n", &v, unsafe.Sizeof(v), unsafe.Alignof(v))
    fmt.Println("  order: size =", unsafe.Sizeof(v.order), "offset= ", unsafe.Offsetof(v.order))
    fmt.Println("article: size =", unsafe.Sizeof(v.article), "offset= ", unsafe.Offsetof(v.article))
    fmt.Println("   size: size =", unsafe.Sizeof(v.size), "offset= ", unsafe.Offsetof(v.size))
    fmt.Println(" active: size =", unsafe.Sizeof(v.active), "offset= ", unsafe.Offsetof(v.active))
	w:= v
	w.order++
	fmt.Println("v:", v, "address =", unsafe.Pointer(&v))
	fmt.Println("w:", w, "address =", unsafe.Pointer(&w))

/*								
v: {11 0xc00002c1f0 123456 true}
v. address: 0xc0000044c0  size: 32  unsafe.Alignof(v)
  order: size = 1 offset=  0
article: size = 8 offset=  8
   size: size = 8 offset=  16
 active: size = 1 offset=  24
v: {11 0xc00002c1f0 123456 true} address = 0xc0000044c0
w: {12 0xc00002c1f0 123456 true} address = 0xc000004520
*/
}
