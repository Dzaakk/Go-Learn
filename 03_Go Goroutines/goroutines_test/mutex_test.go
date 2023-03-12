package goroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//Mutual Exclusion
/*
Berfungsi untuk mengatasi race condition
melakukan locking dan unlocking
sehingga hanya 1 goroutine yang diperbolehkan pada saat lock
setelah unlock baru goroutine selanjutnya diperbolehkan melakukan unlock
- Proses lebih lambat
*/
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

//RWMutex
type BankAccount struct {
	RWMUTEX sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMUTEX.Lock()
	account.Balance = account.Balance + amount
	account.RWMUTEX.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMUTEX.RLock()
	balance := account.Balance
	account.RWMUTEX.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i <100; i++ {
		go func ()  {
			for j := 0; j <100; j++{
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 *time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}