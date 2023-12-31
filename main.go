package main

// 多參數回傳 -------------------------------------------------------------------
// func sap(x int, y int) (xpy int, xty int) {
// 	return x + y, x * y
// }

// func main() {
// 	x := 3
// 	y := 4

// 	xpy, xty := sap(x, y)

// 	fmt.Println(xpy, xty)
// }

// 可變參數函數 -------------------------------------------------------------
// func sum(arg ...int) int {
// 	total := 0
// 	for _, num := range arg {
// 		total += num
// 	}
// 	return total
// }

// func main() {
// 	fmt.Println(sum(1, 2, 3))
// 	fmt.Println(sum(1, 2))
// 	fmt.Println(sum(1))
// 	fmt.Println(sum([]int{1, 2, 3, 4}...))
// }

// 傳值與傳指標 ---------------------------------------------------------------------
// -- 傳值
// func addOne(a int) int {
// 	a += 1
// 	return a
// }

// func main() {
// 	x := 3
// 	fmt.Println("inital x:", x)

// 	x1 := addOne(x)
// 	fmt.Println("add one x1:", x1)

// 	fmt.Println("x after add one x:", x)
// }

// -- 傳指標 傳遞記憶體位址 &v, 接受型別改為指標 *int, 針對指標操作 *a += 1, 回傳新值 return *a
// --- https://willh.gitbook.io/build-web-application-with-golang-zhtw/02.0/02.3#chuan-zhi-yu-chuan-zhi-biao
// func addOne(a *int) int {
// 	*a += 1
// 	return *a
// }

// func main() {
// 	x := 3
// 	fmt.Println("inital x, x:", x)

// 	x1 := addOne(&x)
// 	fmt.Println("add one x, x1:", x1)

// 	fmt.Println("x after add one, x:", x)
// }

// defer ------------------------------------
// 用於確保函式調用在包圍函式結束時執行。這常用於資源清理，例如關閉檔案或釋放鎖。
// with defer
// func ReadWrite() bool {
// 	file.Open("file")
// 	// 做一些事
// 	if failureX {
// 		file.Close()
// 		return false
// 	}

// 	if failureY {
// 		file.Close()
// 		return false
// 	}

// 	file.Close()
// 	return true
// }
// without defer
// func ReadWrite() bool {
// 	file.Open("file")
// 	defer file.Close()
// 	if failureX {
// 		return false
// 	}
// 	if failureY {
// 		return false
// 	}
// 	return true
// }

// 函數作為值、型別 ----------------------------------------
// type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
// type testInt func(int) bool

// func isOdd(number int) bool {
// 	return number%2 != 0
// }

// func isEven(number int) bool {
// 	return number%2 == 0
// }

// func filter(arr []int, f testInt) []int {
// 	result := []int{}
// 	for _, v := range arr {
// 		if f(v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 7}
// 	fmt.Println("slice = ", slice)
// 	odd := filter(slice, isOdd) // 韓式當做值來傳遞了
// 	fmt.Println("Odd elements of slice are: ", odd)
// 	even := filter(slice, isEven) // 函式當做值來傳遞了
// 	fmt.Println("Even elements of slice are: ", even)
// // }

// // 指针类型----------------------------------------------------
// func main() {

// 	// 基本数据类型的内存布局
// 	var i int = 3
// 	// 指向 i 的地址是 &i
// 	fmt.Println("i 的地址=", &i)

// 	// 下面的 var ptr *int = &i
// 	// 1. ptr 是一个指针变量
// 	// 2. ptr 的类型 *int
// 	// 3. ptr 本身的值是 &i
// 	var ptr *int = &i
// 	fmt.Printf("ptr = %v\n", ptr)
// 	// ptr 的地址
// 	fmt.Printf("ptr 的地址=%v\n", &ptr)
// 	fmt.Printf("ptr 指向的值=%v\n", *ptr)
// 	// 输出结果
// 	// i 的地址= 0xc00001a0f8
// 	// ptr = 0xc00001a0f8
// 	// ptr 的地址=0xc00000e030
// 	// ptr 指向的值=3
// }

// func main() {
// 	var a, b int = 1, 2
// 	var e, f = 123, "hell0"
// 	fmt.Printf("%T,%T,%T,%T\n", a, b, e, f)
// }
// // enum in go
// func main() {
// 	const (
// 		a = iota
// 		b
// 		c
// 	)

// 	fmt.Println(a, b, c)
// }
