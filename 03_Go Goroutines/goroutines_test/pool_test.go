package goroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)
//berfungsi untuk mengambil data
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func () interface{} {
			return "Goroutines belum mengembalikan data"
		},
	}

	pool.Put("Data 1")
	pool.Put("Data 2")
	pool.Put("Data 3")

	for i := 0; i <10; i++ {
		go func ()  {
			 data := pool.Get()
			 fmt.Println(data)
			 time.Sleep(1*time.Second)
			 pool.Put(data)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}