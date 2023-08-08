package square
 
import (
	"testing"
	"log"
	"os"
)
	
func TestPackageSquare(t *testing.T) {
	f, _ := os.Create("log.txt")
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.Lmsgprefix)
	defer f.Close()
	
	log.SetPrefix("TestPackageSquare: ")
    log.Println("Start.")

    result := Square(32)
    if result != 1024 {
        t.Error("result should be 1024, got", result)
    }
	
	log.SetPrefix("PerfectSquare1: ")
    ok, result := PerfectSquare1(1025)	// !!!
    if !ok || result != 32 {
        log.Printf("result should be <true 32>, got <%t %d>", ok, result)
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
    if ok, _ = PerfectSquare1(1000); ok {
        t.Error("result should be false, got true")
    }

    ok, result = PerfectSquare2(1024)
    if !ok || result != 32 {
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
    if ok, _ = PerfectSquare2(1000); ok {
        t.Error("result should be false, got true")
    }

    ok, result = PerfectSquare3(1024)
    if !ok || result != 32 {
        t.Error("result should be <true 32>, got <", ok, result, ">")
    }
    if ok, _ = PerfectSquare3(1000); ok {
        t.Error("result should be false, got true")
    }
	
	log.SetPrefix("TestPackageSquare: ")
    log.Println("Finish.")
}
