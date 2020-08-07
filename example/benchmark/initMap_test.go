package benchmark

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"testing"
)

// GC只回收堆上的内存

func TestTrace(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
}

var lastTotalFreed uint64
var intMap map[int]int
var intMapMap map[int]map[int]int
var cnt = 1024 * 80

// 先delete再nil
func TestMapInit(t *testing.T) {
	printMemStats() //Alloc = 119 TotalAlloc = 119  Just Freed = 0 Sys = 4868 NumGC = 0

	initMap()

	printMemStats() //Alloc = 481 TotalAlloc = 481  Just Freed = 0 Sys = 4868 NumGC = 0
	runtime.GC()
	printMemStats() //Alloc = 423 TotalAlloc = 483  Just Freed = 60 Sys = 4932 NumGC = 1

	log.Println(len(intMap)) //8192
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap)) //0

	printMemStats() //Alloc = 424 TotalAlloc = 485  Just Freed = 0 Sys = 5188 NumGC = 1
	runtime.GC()
	printMemStats() //Alloc = 423 TotalAlloc = 485  Just Freed = 1 Sys = 4996 NumGC = 2

	intMap = nil

	printMemStats() //Alloc = 424 TotalAlloc = 486  Just Freed = 0 Sys = 5252 NumGC = 2
	runtime.GC()
	printMemStats() //Alloc = 111 TotalAlloc = 486  Just Freed = 313 Sys = 4996 NumGC = 3
}

// 先nil再delete
func TestMapInitTwo(t *testing.T) {
	printMemStats() //Alloc = 120 TotalAlloc = 120  Just Freed = 0 Sys = 4868 NumGC = 0

	initMap()

	printMemStats() //Alloc = 483 TotalAlloc = 483  Just Freed = 0 Sys = 6274 NumGC = 0
	runtime.GC()
	printMemStats() //Alloc = 424 TotalAlloc = 485  Just Freed = 60 Sys = 6338 NumGC = 1

	intMap = nil

	printMemStats() //Alloc = 425 TotalAlloc = 486  Just Freed = 0 Sys = 6338 NumGC = 1
	runtime.GC()
	printMemStats() //Alloc = 111 TotalAlloc = 486  Just Freed = 313 Sys = 6402 NumGC = 2

	log.Println(len(intMap)) //0
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap)) //0

	printMemStats() //Alloc = 113 TotalAlloc = 487  Just Freed = 0 Sys = 6402 NumGC = 2
	runtime.GC()
	printMemStats() //Alloc = 111 TotalAlloc = 487  Just Freed = 1 Sys = 6402 NumGC = 3
}

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}

func TestMapInit2(t *testing.T) {
	// 1
	printMemStats()

	// 2
	initMapMap()
	runtime.GC()
	printMemStats()

	// 3
	fillMapMap()
	runtime.GC()
	printMemStats()

	// 4
	log.Println(len(intMapMap))
	for i := 0; i < cnt; i++ {
		delete(intMapMap, i)
	}
	log.Println(len(intMapMap))
	runtime.GC()
	printMemStats()

	// 5
	intMapMap = nil
	runtime.GC()
	printMemStats()
}

func initMapMap() {
	intMapMap = make(map[int]map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMapMap[i] = make(map[int]int, cnt)
	}
}

func fillMapMap() {
	for i := 0; i < cnt; i++ {
		for j := 0; j < cnt; j++ {
			intMapMap[i][j] = j
		}
	}
}
