package main

import (
	"fmt"
	"time"
)

var x string = "hy saya fadhil"

const y string = "ada apa"

func main() {

	fmt.Printf("hello, world\n")
	fmt.Println(x)
	x = "saya bukan fadhil"
	fmt.Println(x)

	fmt.Println(y)
	//y = "kenapa"    =====>  tidak bisa
	//fmt.Fprintln(y)

	fmt.Println("what time is it? It's", time.Now())
}
