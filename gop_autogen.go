package main

import (
	"fmt"
	"gop-sample/cpkag"
	ab "gop-sample/cpkag/b"
	"gop-sample/cpkag/b/ss"
	"gop-sample/cpkag/k"
	"gop-sample/cpkag/utils"
	"gop-sample/gpkag"
	utils1 "gop-sample/utils"
	utils2 "github.com/liuscraft/testgomod/utils"
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
//line b.gop:16
func main() {
//line b.gop:16:1
	bb := 3
//line b.gop:17:1
	Ab()
//line b.gop:18:1
	ab.Ab()
//line b.gop:19:1
	ab.Ac()
//line b.gop:20:1
	ss.Ab()
//line b.gop:21:1
	ss.Bs()
//line b.gop:22:1
	fmt.Println(bb)
//line b.gop:23:1
	k.K()
//line b.gop:24:1
	k.Kk()
//line b.gop:25:1
	k.Ab()
//line b.gop:26:1
	utils2.TestCsgo()
//line b.gop:27:1
	utils.TestCsgo()
//line b.gop:28:1
	utils1.TestCsgo()
//line b.gop:30:1
	a()
//line b.gop:31:1
	cpkag.P()
//line b.gop:32:1
	cpkag.Gg()
//line b.gop:33:1
	gpkag.G()
//line b.gop:34:1
	fmt.Println("b")
//line b.gop:35:1
	fmt.Println("b")
//line b.gop:36:1
	fmt.Println("b")
//line b.gop:37:1
	fmt.Println("b")
//line b.gop:38:1
	fmt.Println("b")
//line b.gop:39:1
	fmt.Println("b")
//line b.gop:40:1
	fmt.Println("b")
}
