package main

import (
	"fmt"
)

type tipo int
var x tipo
type tipotipo tipo
var y tipotipo

func main(){
	fmt.Printf("%T - %v\n",x,x)
	x = 42
	fmt.Println(x)

	y = tipotipo(x)
	fmt.Printf("%T - %v\n",y,y)
}