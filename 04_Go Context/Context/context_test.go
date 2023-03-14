package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

//context with value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}
func TestContextWithCancelLeak(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())

	destination := CreateCounterLeak()
	fmt.Println("Total Goroutine Ketika Counter Berjalan : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Total Goroutine Setelah Proses Selesai : ", runtime.NumGoroutine())
}

//Not leak
func CreateCounterNotLeak(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancelNotLeak(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterNotLeak(ctx)
	fmt.Println("Total Goroutine Ketika Counter Berjalan : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	//mengirim sinyal cancel ke context

	time.Sleep(1 * time.Second)
	fmt.Println("Total Goroutine Setelah Proses Selesai : ", runtime.NumGoroutine())
}

//timeout
func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterNotLeak(ctx)
	fmt.Println("Total Goroutine Ketika Counter Berjalan : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine Setelah Proses Selesai : ", runtime.NumGoroutine())
}

//deadline (waktunya fix)
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()
	// dijalankan dari waktu saat ini sampai dengan 5 detik kedepan

	destination := CreateCounterNotLeak(ctx)
	fmt.Println("Total Goroutine Ketika Counter Berjalan : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine Setelah Proses Selesai : ", runtime.NumGoroutine())
}
