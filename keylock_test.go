package go_keylock

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewKeyLock(t *testing.T) {
	var keyLock *KeyLock
	keyLock = NewKeyLock()

	keyLock.Lock("shi")
	defer keyLock.Unlock("shi")

	//TODO 需要处理的流程
}

func TestKeyLock(t *testing.T) {
	var wg sync.WaitGroup
	keyLock := NewKeyLock()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			keyLock.Lock(fmt.Sprintf("%v", i))
			fmt.Println("do test", fmt.Sprintf("%v", i))
			time.Sleep(1 * time.Second)
			defer keyLock.Unlock(fmt.Sprintf("%v", i))

		}(1)
		wg.Done()
	}
	wg.Wait()

	fmt.Println(len(keyLock.locks))
	time.Sleep(2 * time.Second)
	fmt.Println(len(keyLock.locks))

	time.Sleep(30 * time.Second)
}

func TestKeyLockClean(t *testing.T) {
	keyLock := NewKeyLock()

	keyLock.Lock("shi")
	keyLock.Lock("jin")
	keyLock.Lock("yu")

	fmt.Println(len(keyLock.locks))

	keyLock.Unlock("shi")
	keyLock.Unlock("jin")
	keyLock.Unlock("yu")
	fmt.Println(len(keyLock.locks))
}

func BenchmarkCombinationParallel(b *testing.B) {
	keyLock := NewKeyLock()

	var keys []string
	for n := 0; n < 100000; n++ {
		keys = append(keys, fmt.Sprintf("key%d", n))
	}
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RunPrintKey(keyLock,keys)
		}
	})

	var i = 0
	for _, v := range keyLock.locks {
		fmt.Println(v)
		i++
	}
	fmt.Println("i=", i)
}

func RunPrintKey(keyLock *KeyLock,keys []string) {
	index := rand.Intn(10)
	key := keys[index]
	keyLock.Lock(key)
	//fmt.Println(key)
	keyLock.Unlock(key)
}

