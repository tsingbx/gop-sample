package main

import (
	"fmt"
	"gop-sample/cpkag"
	ab "gop-sample/cpkag/b"
	"gop-sample/cpkag/b/ss"
	"gop-sample/cpkag/k"
	"gop-sample/gpkag"
)
//line a.gop:3:1
func a() {
//line a.gop:4:1
	fmt.Println("a")
}
//line a.gop:7:1
func Ab() {
//line a.gop:8:1
	fmt.Println("ab")
}
//line b.gop:12
func main() {
//line b.gop:12:1
	bb := 3
//line b.gop:13:1
	Ab()
//line b.gop:14:1
	ab.Ab()
//line b.gop:15:1
	ab.Ac()
//line b.gop:16:1
	ss.Ab()
//line b.gop:17:1
	fmt.Println(bb)
//line b.gop:18:1
	k.K()
//line b.gop:19:1
	k.Kk()
//line b.gop:20:1
	k.Ab()
//line b.gop:22:1
	a()
//line b.gop:23:1
	cpkag.P()
//line b.gop:24:1
	cpkag.Gg()
//line b.gop:25:1
	gpkag.G()
//line b.gop:26:1
	fmt.Println("b")
//line b.gop:27:1
	fmt.Println("b")
//line b.gop:28:1
	fmt.Println("b")
//line b.gop:29:1
	fmt.Println("b")
//line b.gop:30:1
	fmt.Println("b")
//line b.gop:31:1
	fmt.Println("b")
//line b.gop:32:1
	fmt.Println("b")
}
