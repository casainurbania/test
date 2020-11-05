package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Printf("%v: %s\n",time.Now(),"compiling....")
	if time.Now().Unix() %2==0{
		panic("fault test")
	}
}
