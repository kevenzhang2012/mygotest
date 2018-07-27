package main

import "fmt"
//channel 是 Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或多个goroutine之间传递消息。
//channel 是进程内的通信方式，因此 通过 channel 传递对象  的过程和   调用函数时的   参数  传递行为比较一致，
//比如也可以传递指针等。 如果需要夸进程通信，我们建议用分布式系统的方法来解决，
//比如使用Socket 或者 Http等通信协议。  GO语言 对于网络方面 也有非常完善的支持。
//channel 是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。如果对Unix管道有所了解的话，
//就不难理解channel，可以将其认为是一种类型安全的管道。

// make(chan type)
//ch ,


func main() {
	message := make(chan string)

	go func() {message <- "ping"}()

	msg := <- message
	fmt.Println(msg)
}

