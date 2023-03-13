package goroutines_test

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}
//Berfungsi untuk melakukan sebuah function hanya satu kali
func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	group.Add(1)
	for i := 0; i < 100; i++ {
		go func() {

			once.Do(OnlyOnce)

		}()
	}
	group.Done()
	group.Wait()
	fmt.Println(counter)
}
