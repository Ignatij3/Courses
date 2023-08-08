package arithmetic

import "testing"

func TestCheck(t *testing.T) {
	type test struct {
		seq []int
		ok bool
	}	
	tests := []test {
		// arithmetic progression
		{[]int{2,3}, true},
		{[]int{2,3,4,5,6,7}, true},
		{[]int{2,0,-2,-4,-6}, true},
		{[]int{2,12,22,32}, true},
		// non arithmetic progression
		{[]int{2,3,4,5,7}, false},
		{[]int{2,3,5,7,9}, false},
		{[]int{2}, false},
		{[]int{}, false},
	}		
	for _, ts:= range tests {
		ok, res:= Check(ts.seq...)
		if ok != ts.ok {
			t.Error("Error in the arithmetic progression indication")
		}	 
		if ok { 
			wait := Arithmetic{ts.seq[0], ts.seq[1] - ts.seq[0]}
			if wait != res	{
				t.Errorf("Error in the arithmetic progression construction: want %v, get %v", wait, res)
			}	
		}
	}
}
