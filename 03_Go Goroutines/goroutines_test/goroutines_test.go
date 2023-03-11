package goroutines_test

import (
	"fmt"
	"testing"
)

func Hello() {
	fmt.Println("Hello")
}

func TestCreateGoroutine(t *testing.T) {
	Hello()
	fmt.Println("Ups")
	
}
