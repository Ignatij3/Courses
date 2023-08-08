package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	actual := helloWorld()
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestHello(t *testing.T) {
	expected := "Hello, Go!"
	actual := hello("Go")
	if actual != expected {
		t.Error("Test failed")
	}
}
