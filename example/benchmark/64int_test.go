package benchmark

import (
	"sync/atomic"
	"testing"
)

//go test -run=alloc_test.go -bench="." -benchtime="3s" -cpuprofile profile_cpu.out
//go tool pprof profile_cpu.out

//go test -bench=BenchmarkPad2 -trace=reference.out
//go tool trace reference.out

//go test 64int_test.go -test.bench=".*"

// pad 结构的 x y z 会被并发的执行原子操作
type pad struct {
	x uint64 // 8byte
	y uint64 // 8byte
	z uint64 // 8byte
}

func (s *pad) increase() {
	atomic.AddUint64(&s.x, 1)
	atomic.AddUint64(&s.y, 1)
	atomic.AddUint64(&s.z, 1)
}

func (s *pad2) increase() {
	atomic.AddUint64(&s.x, 1)
	atomic.AddUint64(&s.y, 1)
	atomic.AddUint64(&s.z, 1)
}

// 一级缓存又分为数据缓存和指令缓存，他们都由高速缓存行组成，对于X86架构的CPU来说，高速缓存行一般是32/64个字节，
// 现在的CPU一般都是32K的一级缓存。
// 当CPU需要读取一个变量时，该变量所在的以32/64字节分组的内存数据将被一同读入高速缓存行，一次性将访问频繁的32/64字节数据对齐后读入高速缓存中，
// 减少CPU高级缓存与低级缓存、内存的数据交换。
// cpu.CacheLinePadSize = 64，现代处理器每个 L1 缓存一般拥有 32 * 1024 / 64 = 512 条缓存行
// 通过 CPU 缓存行大小，计算出需要填充的实际字节数。如下，uint64是8字节，需要56字节填充，就是一个完整的缓存行。

// pad 结构的 x y z 会被并发的执行原子操作
type pad2 struct {
	x uint64 // 8byte
	_ [56]byte
	y uint64 // 8byte
	_ [56]byte
	z uint64 // 8byte
	_ [56]byte
}

func BenchmarkPad(b *testing.B) {
	b.ReportAllocs()
	s := pad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.increase()
		}
	})
}

func BenchmarkPad2(b *testing.B) {
	b.ReportAllocs()
	s := pad2{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.increase()
		}
	})
}

/*
go test -run=64int_test.go -bench="." -count=3
goos: windows
goarch: amd64
BenchmarkPad-4          30000000                52.3 ns/op             0 B/op          0 allocs/op
BenchmarkPad-4          30000000                56.4 ns/op             0 B/op          0 allocs/op
BenchmarkPad-4          30000000                55.7 ns/op             0 B/op          0 allocs/op
BenchmarkPad2-4         30000000                33.6 ns/op             0 B/op          0 allocs/op
BenchmarkPad2-4         50000000                33.3 ns/op             0 B/op          0 allocs/op
BenchmarkPad2-4         50000000                33.2 ns/op             0 B/op          0 allocs/op
PASS
ok      _/D_/workspace/go_workspace/go_practise/example/benchmark       12.672s
*/
