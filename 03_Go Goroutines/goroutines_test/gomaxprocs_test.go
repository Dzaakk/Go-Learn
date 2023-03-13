package goroutines_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU = ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread = ", totalThread)
	//untuk mengubah jumlah thread
	//runtime.GOMAXPROCS(20)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine = ", totalGoroutine)
}
