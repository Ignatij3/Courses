package main

import "fmt"

type (
	Function1D func(float64) float64
	Function2D func(float64, float64) float64
)

func minimum2D(left, right, bottom, top float64, F2 Function2D) (float64, float64) {
	h:= (func (x float64) float64 {
		q:= (func (y float64) float64  {
			return F2(x, y)	
		} )		
		ymin:= minimum1D(bottom, top, q)
		return F2(x, ymin)
	}  )	
	
	xmin:= minimum1D(left, right, h) 
	ymin:= minimum1D(bottom, top, ( func (y float64) float64 {
											return f(xmin, y)  }  )  )					
	return xmin, ymin
}

func minimum1D(start, finish float64, F1 Function1D) float64 {
	// Any procedure of the 1D minimization
	var pmin float64	//minimum point
	// ...
	return pmin  
}

func f(x, y float64) float64 {
	var res float64
	// ...
	return res	
}	  

func main() {
	var left, right, bottom, top float64
	// left, right, bottom, top = ...
	fmt.Println(minimum2D(left, right, bottom, top,  f))
}
