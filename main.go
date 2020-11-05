package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Printf("%v: %s\n",time.Now(),"compiling....")
	panic("This is a testing stage failed message.")
}
