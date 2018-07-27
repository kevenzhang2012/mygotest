package main
import (
"fmt"
"reflect"
)

func main(){
    type BitFlag int
    const  Active BitFlag = 1 << iota
    const Send
    const Receive
flag := Active |send
fmt.Println(flag)

}
