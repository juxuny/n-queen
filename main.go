package main

import (
	"fmt"
	"flag"
)

var (
	N = 10
	num = 1
	countChan = make(chan bool, 100)
	count int64
	ii int64 = 0
	queue = make(chan int64, 100)
	max int64 = 0
)

func init() {
	flag.IntVar(&N, "n", 4, "the number of queen")
	flag.IntVar(&num, "g", 5, "the number of goruntines")
	flag.Parse()
}

func counter() {
	for {
		b := <- countChan
		if b {
			count ++
		}
		ii ++
		if ii >= max {
			break
		}
	}
}

func check(k int64) bool {
	a := make([]bool, N)
	b := make([]int64, N)
	for i := 0; i < N; i++ {
		m := k % int64(N)
		if a[m] {
			return false
		}
		a[m] = true
		for j := 0; j < i; j++ {
			if int64(i-j) == m-b[j] || int64(-i+j) == m-b[j] {
				return false
			}
		}
		b[i] = m
		k = k / int64(N)
	}
	return true
}

func main() {
	max = 1
	for i := 0; i < N; i++ {
		max *= int64(N)
	}
	go func () {
		var i int64 = 0
		for i < max {
			queue <- i
			i++
		}
	}()
	for k := 0; k < num; k++ {
		go func () {
			for {
				countChan <- check(<- queue)
			}
		}()
	}
	counter()
	fmt.Println(count)
}