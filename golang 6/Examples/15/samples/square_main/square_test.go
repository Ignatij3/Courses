package main
 
import "testing"
 
func TestSquare(t *testing.T) {
    result := Square(32)
    if result != 1024 {
        t.Error("result should be 1024, got", result)
    }
}
