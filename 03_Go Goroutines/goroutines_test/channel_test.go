package goroutines_test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Data dimasukan ke channel"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

//Channel as Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ini Data"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

//Channel hanya mengirim/memasukan data
func OnlyIn(channel chan<- string){
	time.Sleep(2* time.Second)
	channel <- "data dimasukan ke channel"
}
//Channel hanya menerima data
func OnlyOut(channel <-chan string){
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
 channel := make(chan string)

 go OnlyIn(channel)
 go OnlyOut(channel)

 time.Sleep(3 * time.Second)
 close(channel)
}