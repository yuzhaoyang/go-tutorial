package main

import (
	"fmt"
	"sync"
	"time"
)

func worker0(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	/* wg跟踪10个goroutine */
	size := 10
	wg.Add(size)
	/* 开启10个goroutine并发执行
	   这里10个goroutine在闭包里都用了变量i，等闭包真正执行的时候，i的值
	   已经改变，所以这种实现方式不安全
	*/
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			worker0(i)
		}()
	}
	/* Wait会一直阻塞，直到wg的计数器为0*/
	wg.Wait()
	fmt.Println("end")
}
