package main

// 多參數回傳 -----------------
// func sap(x int, y int) (xpy int, xty int) {
// 	return x + y, x * y
// }

// func main() {
// 	x := 3
// 	y := 4

// 	xpy, xty := sap(x, y)

// 	fmt.Println(xpy, xty)
// }

// 可變參數函數 -----------------
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

// 傳值與傳指標 -----------------
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

// defer -----------------
// 用於確保函式調用在包圍函式結束時執行。這常用於資源清理，例如關閉檔案或釋放鎖。
// with defer
// func ReadWrite() bool {
// 	file.Open("file")
// 	// 做一些工作
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

// 函數作為值、型別 -----------------
