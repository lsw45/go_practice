package syncPool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Printf("v：%v\n", v)

	pool.Put(3)
	v1, _ := pool.Get().(int) // 池中有数据，所以没有执行New
	fmt.Printf("v1：%v\n", v1)

	pool.Put(50)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v2, _ := pool.Get().(int)
	fmt.Printf("v2：%v\n", v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
