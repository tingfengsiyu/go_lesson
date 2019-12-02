package main

//map类型 k-v
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 1
	a[1] = 11
	a[2] = 13
	a[3] = 14
	a[4] = 15
	a[5] = 16
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()
		}(a)

	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

func testRWlock() {
	var a map[int]int
	var count int32
	a = make(map[int]int, 5)
	a[0] = 1
	a[1] = 11
	a[2] = 13
	a[3] = 14
	a[4] = 15
	a[5] = 16
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[1] = rand.Intn(100)
			rwlock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			rwlock.RLock()
			fmt.Println(a)
			rwlock.RUnlock()
			atomic.AddInt32(&count, 1)

		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) //原子操作
}
func main() {
	//testMap()
	testRWlock()
}
