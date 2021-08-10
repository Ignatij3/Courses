package main

var secret int = 73

func guess(left, right int) {
    // recursive variant
    attempt:= (left + right) / 2
    switch {
	case secret < attempt: 
	    guess(left, attempt-1)   	
    case secret > attempt:
        guess(attempt+1, right)
    case secret == attempt:
        return
    }    
}	

func guess2(left, right int) {
    // cyclic variant
    for {
        attempt:= (left + right) / 2
        switch {
	    case secret < attempt: 
	        right = attempt - 1   	
        case secret > attempt:
            left = attempt + 1
        case secret == attempt:
            return
        }    
    }        
}	

func main() {
    guess(1, 100)
    guess2(1, 100)
}
