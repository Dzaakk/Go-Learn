package goroutines_test

import (
	"fmt"
	"testing"
	"time"
)

func Hello() {
	fmt.Println("Hello")
}

func TestCreateGoroutine(t *testing.T) {
	go Hello()
	fmt.Println("Ups")
	
	time.Sleep(1* time.Second)
}
