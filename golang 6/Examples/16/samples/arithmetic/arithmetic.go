// Around arithmetic progression
package arithmetic

type Arithmetic struct {
	start int
	diff int
}	

// Check reports whether a series seq forms arithmetic progression.
// If yes, then builds an arithmetic progression, the initial 
// segment of which is this series
func Check(seq ...int) (ok bool, sequence Arithmetic) {
	if len(seq) <= 1 { 
		ok = false
		return 
	}
	d:= seq[1] - seq[0]
	for i:= 1; i < len(seq) -1; i++ {
		if seq[i+1] - seq[i] != d {
			ok = false
			return 
		}	
	}
	return true, Arithmetic{seq[0], d}
}

