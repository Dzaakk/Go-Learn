package goroutines_test

import (
	"fmt"
	"testing"
	"time"
)
//gunakan ticker untuk aksi yang berulang
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()
	//ambil data dari channel
	for time := range ticker.C {
		fmt.Println(time)
	} 
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	
	for time := range channel{
		fmt.Println(time)
	}
}


