package main

import (
	"fmt"
	"gop-sample/cpkag"
	ab "gop-sample/cpkag/b"
	"gop-sample/cpkag/b/ss"
	"gop-sample/cpkag/k"
	"gop-sample/cpkag/k/utils"
	utils1 "gop-sample/cpkag/utils"
	utils2 "gop-sample/cpkag/utils/utils"
	"gop-sample/gpkag"
	utils3 "gop-sample/utils"
	utils4 "github.com/liuscraft/testgomod/utils"
)

const _ = true
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
//line b.gop:18
func main() {
//line b.gop:18:1
	bb := 3
//line b.gop:19:1
	Ab()
//line b.gop:20:1
	ab.Ab()
//line b.gop:21:1
	ab.Ac()
//line b.gop:22:1
	ss.Ab()
//line b.gop:23:1
	ss.Bs()
//line b.gop:24:1
	fmt.Println(bb)
//line b.gop:25:1
	k.K()
//line b.gop:26:1
	k.Kk()
//line b.gop:27:1
	k.Ab()
//line b.gop:28:1
	utils4.TestCsgo()
//line b.gop:29:1
	utils1.TestCsgo()
//line b.gop:30:1
	utils2.TestCsgo()
//line b.gop:31:1
	utils3.TestCsgo()
//line b.gop:32:1
	utils.TestCsgo()
//line b.gop:33:1
	utils.Kutils()
//line b.gop:35:1
	a()
//line b.gop:36:1
	cpkag.P()
//line b.gop:37:1
	cpkag.Gg()
//line b.gop:38:1
	gpkag.G()
//line b.gop:39:1
	fmt.Println("b")
//line b.gop:40:1
	fmt.Println("b")
//line b.gop:41:1
	fmt.Println("b")
//line b.gop:42:1
	fmt.Println("b")
//line b.gop:43:1
	fmt.Println("b")
//line b.gop:44:1
	fmt.Println("b")
//line b.gop:45:1
	fmt.Println("b")
}
