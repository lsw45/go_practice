package sync_pool

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"testing"
)

func TestNewRoomFrameDataPool(t *testing.T) {
	// 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool := sync.Pool{New: newFunc}
	// New 字段值的作用
	v1 := pool.Get()
	fmt.Printf("v1: %v，count:%v\n", v1, count)
	A1 := pool.Get()
	fmt.Printf("A1: %v，count:%v\n", A1, count)

	// 临时对象池的存取
	pool.Put(newFunc())
	pool.Put(newFunc())
	pool.Put(newFunc()) // 将newFunc发进去，get的时候不执行newFunc了
	pool.Put(8)

	v2 := pool.Get()
	fmt.Printf("v2: %v，count:%v\n", v2, count)

	// 垃圾回收对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()

	v3 := pool.Get()
	fmt.Printf("v3: %v，count:%v\n", v3, count)

	v31 := pool.Get()
	fmt.Printf("v31: %v，count:%v\n", v31, count)

	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("v4: %v，count:%v\n", v4, count)

	v5 := pool.Get()
	fmt.Printf("v5: %v，count:%v\n", v5, count)

	//////////////----------------------
	pool = sync.Pool{New: func() interface{} {
		return "empty string"
	}}
	s := "Hello World"
	pool.Put(s)
	pool.Put("bbbb")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
