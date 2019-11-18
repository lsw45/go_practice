package main

import (
	"fmt"
	"os"
	"testing"
)

func defer_return() int {
	i := 0
	defer func() {
		i++
		fmt.Println(3)
	}()
	defer func() {
		i++
		fmt.Println(2)
	}()

	return i
}

func TestReturn(t *testing.T) {
	fmt.Println(defer_return())
}

func TestPanic(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	defer func() {
		fmt.Println(3)
	}()
	defer func() {
		fmt.Println(2)
	}()

	panic("panic~~~")

	defer func() {
		fmt.Println(1)
	}()
}

func TestExit(t *testing.T) {
	defer func() {
		fmt.Println(3)
	}()
	defer func() {
		fmt.Println(2)
	}()

	fmt.Println("exit 不运行 defer")
	// os.Exit(1)
	os.Environ()

	defer func() {
		fmt.Println(1)
	}()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func TestMain061(t *testing.T) {
	a := 1
	b := 2

	defer calc("1", a, b)

	a = 0
	defer calc("2", a, b)

	b = 1
}

func TestMain062(t *testing.T) {
	a := 1
	b := 2

	defer func() {
		calc("1", a, b)
	}()

	a = 0

	defer func() {
		calc("2", a, b)
	}()

	b = 1
}
