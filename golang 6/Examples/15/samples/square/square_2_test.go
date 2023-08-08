package square
 
import "testing"
	
func TestSquare(t *testing.T) {
    result := Square(32)
    if result != 1024 {
        t.Error("result should be 1024, got", result)
    }
}

func TestPerfectSquare1(t *testing.T) {
    ok, result := PerfectSquare1(1024)
    if !ok || result != 32 {
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
}

func TestPerfectSquare1_no(t *testing.T) {
    if ok, _ := PerfectSquare1(1000); ok {
        t.Error("result should be false, got true")
    }
}

func TestPerfectSquare2(t *testing.T) {
    ok, result := PerfectSquare2(1024)
    if !ok || result != 32 {
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
}

func TestPerfectSquare2_no(t *testing.T) {
    if ok, _ := PerfectSquare2(1000); ok {
        t.Error("result should be false, got true")
    }
}

func TestPerfectSquare3(t *testing.T) {
    ok, result := PerfectSquare3(1024)
    if !ok || result != 32 {
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
}

func TestPerfectSquare3_no(t *testing.T) {
    if ok, _ := PerfectSquare3(1000); ok {
        t.Error("result should be false, got true")
    }
}

