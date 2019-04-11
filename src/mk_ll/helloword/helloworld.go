package main

import (
	"fmt"
	"strconv"
	"time"
)

//var message = make(chan string)
func sample(ch chan string) {
	for i := 0; i < 9; i++ {
		ch <- "hello imooc!" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}

func sample2(ch chan int) {
	for i := 0; i < 9; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}
func main() {
	//message := make(chan string)
	//go func() {
	//	//fmt.Println("hello imooc!")
	//	message <- "hello imooc!"
	//}()
	//ch := make(chan string)
	//go sample(ch)
	//go sample2(ch)
	//fmt.Println("hello world")
	//for i := 0; i < 18; i++ {
	//	fmt.Println(<-ch)
	//}
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go sample(ch1)
		go sample2(ch2)
	}
	for i:=0;i<1000;i++{
	select {
	case str, ch1Check := <-ch1:
		if !ch1Check {
			fmt.Println("ch1 failed")
		}
		fmt.Println(str)
	case p, ch2Check := <-ch2:
		if !ch2Check {
			fmt.Println("ch2 failed")
		}
		fmt.Println(p)
	}}
	time.Sleep(60 * time.Second)
}
