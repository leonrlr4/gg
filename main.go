package main

import (
	"fmt"
)

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
// }

// struct ----------------------------------------

// type Person struct {
// 	Name string
// 	Age  int
// }

// func olderPerson(p1, p2 Person) (Person, int) {
// 	if p1.Age > p2.Age {
// 		return p1, p1.Age - p2.Age
// 	}
// 	return p2, p2.Age - p1.Age
// }

// func main() {
// 	// 賦值初始化
// 	var tom Person = Person{
// 		Name: "Tom",
// 		Age:  18,
// 	}
// 	// 兩個欄位都寫清楚的初始化
// 	jack := Person{
// 		Name: "Jack",
// 		Age:  20,
// 	}
// 	// 按照 struct 定義順序初始化值
// 	tina := Person{"tina", 21}
// 	// 指標初始化
// 	// var lisa *Person = &Person{
// 	// 	Name: "Lisa",
// 	// 	Age:  19,
// 	// }

// 	tj, difftj := olderPerson(tom, jack)
// 	jt, diffjt := olderPerson(jack, tina)

// 	fmt.Printf("%v is older %v years\n", tj.Name, difftj)
// 	fmt.Printf("%v is older %v years\n", jt.Name, diffjt)
// }

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human      // 匿名欄位，那麼預設 Student 就包含了 Human 的所有欄位
// 	speciality string
// }

// func main() {
// 	// aa := Student{
// 	// 	Human{"Mark", 25, 120},
// 	// 	"Computer Science",
// 	// }

// 	mark := Student{
// 		Human{"mark", 18, 60},
// 		"Computer Science",
// 	}

// 	fmt.Println("His name is ", mark.name)
// 	fmt.Println("His age is ", mark.age)
// 	fmt.Println("His weight is ", mark.weight)
// 	fmt.Println("His speciality is ", mark.speciality)

// 	mark.speciality = "Math"
// 	fmt.Println("His new speciality is ", mark.speciality)
// }

// type Skills []string

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Job struct {
// 	title    string
// 	timeYear int
// }

// type Student struct {
// 	Human // 匿名欄位，struct
// 	Job
// 	Skills     // 匿名欄位，自訂的型別 string slice
// 	int        // 內建型別作為匿名欄位
// 	speciality string
// }

// func main() {
// 	// 初始化學生 Jane
// 	jane := Student{
// 		Human:      Human{"Jane", 35, 100},
// 		Job:        Job{"Cooking", 2},
// 		speciality: "Biology",
// 	}
// 	// 現在我們來存取相應的欄位
// 	fmt.Println("Her name is ", jane.name)
// 	fmt.Println("Her age is ", jane.age)
// 	fmt.Println("Her weight is ", jane.weight)
// 	fmt.Println("Her speciality is ", jane.speciality)
// 	fmt.Println("Her job is ", jane.Job.title)
// 	fmt.Println("She is doing for ", jane.Job.timeYear, " years")

// 	// 我們來修改他的 skill 技能欄位
// 	jane.Skills = []string{"anatomy"}
// 	fmt.Println("Her skills are ", jane.Skills)
// 	fmt.Println("She acquired two new ones ")
// 	jane.Skills = append(jane.Skills, "physics", "golang")
// 	fmt.Println("Her skills now are ", jane.Skills)

// 	// 修改匿名內建型別欄位
// 	jane.int = 3
// 	fmt.Println("Her preferred number is", jane.int)
// }

// type Human struct {
// 	name  string
// 	Age   int
// 	phone string
// }

// type employee struct {
// 	Human
// 	speciality string
// 	phone      string
// }

// func main() {
// 	Bob := employee{
// 		Human{"Bob", 34, "123456"},
// 		"IT",
// 		"654321",
// 	}
// 	fmt.Println("Bob's work phone is:", Bob.phone)
// 	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
// }

// OOP ----------------------------------------
// type rect struct {
// 	l, w float64
// }

// func rtg(r rect) float64 {
// 	return r.l * r.w
// }

// func main() {
// 	r1 := rect{10, 20}
// 	r2 := rect{7, 14}

// 	fmt.Println("r1 Area is ", rtg(r1))
// 	fmt.Println("r2 Area is ", rtg(r2))
// }

// type rectangle struct {
// 	l, w float64
// }

// type circle struct {
// 	r float64
// }

// func (r rectangle) area() float64 {
// 	return r.l * r.w
// }

// func (c circle) area() float64 {
// 	return (c.r * c.r) * math.Pi
// }

// func main() {
// 	r1 := rectangle{10, 20}
// 	r2 := rectangle{7, 14}
// 	c1 := circle{8}
// 	c2 := circle{5}

// 	fmt.Println("Area of r1 is: ", r1.area())
// 	fmt.Println("Area of r2 is: ", r2.area())
// 	fmt.Println("Area of c1 is: ", c1.ar	ea())
// 	fmt.Println("Area of c2 is: ", c2.area())
// }

// method ---------------------------------------------------
// const (
// 	WHITE = iota
// 	BLACK
// 	BLUE
// 	RED
// 	YELLOW
// )

// type Color byte

// type Box struct {
// 	width, height, depth float64
// 	color                Color
// }

// type BoxList []Box //a slice of boxes

// func (b Box) Volume() float64 {
// 	return b.width * b.height * b.depth
// }

// func (b *Box) SetColor(c Color) {
// 	b.color = c
// }

// func (bl BoxList) BiggestColor() Color {
// 	v := 0.00
// 	k := Color(WHITE)
// 	for _, b := range bl {
// 		if bv := b.Volume(); bv > v {
// 			v = bv
// 			k = b.color
// 		}
// 	}
// 	return k
// }

// func (bl BoxList) PaintItBlack() {
// 	for i := range bl {
// 		bl[i].SetColor(BLACK)
// 	}
// }

// func (c Color) String() string {
// 	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
// 	return strings[c]
// }

// func main() {
// 	boxes := BoxList{
// 		Box{4, 4, 4, RED},
// 		Box{10, 10, 1, YELLOW},
// 		Box{1, 1, 20, BLACK},
// 		Box{10, 10, 1, BLUE},
// 		Box{10, 30, 1, WHITE},
// 		Box{20, 20, 20, YELLOW},
// 	}

// 	fmt.Printf("We have %d boxes in our set\n", len(boxes))
// 	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
// 	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
// 	fmt.Println("The biggest one is", boxes.BiggestColor().String())

// 	fmt.Println("Let's paint them all black")
// 	boxes.PaintItBlack()
// 	fmt.Println("The color of the second one is", boxes[1].color.String())

// 	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
// }

// method inheritance -------------------------------------------------------------------
// type Human struct {
// 	name  string
// 	age   int
// 	phone string
// }

// type Student struct {
// 	Human
// 	school string
// 	loan   float32
// }

// type Employee struct {
// 	Human
// 	company string
// 	money   float32
// }

// // Human 物件實現 Sayhi 方法
// func (h *Human) SayHi() {
// 	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
// }

// // Human 物件實現 Sing 方法
// func (h *Human) Sing(lyrics string) {
// 	fmt.Println("lalalalalalala .....", lyrics)
// }

// // Human 物件實現 Guzzle 方法
// // func (h *Human) Guzzle(beerStein string) {
// // 	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
// // }

// // Employee 過載 Human 的 Sayhi 方法
// func (e *Employee) SayHi() {
// 	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n",
// 		e.name,
// 		e.company,
// 		e.phone)
// }

// // Student 實現 BorrowMoney 方法
// // func (s *Student) BorrowMoney(amount float32) {
// // 	s.loan += amount // (again and again and...)
// // }

// // Employee 實現 SpendSalary 方法
// // func (e *Employee) SpendSalary(amount float32) {
// // 	e.money -= amount
// // }

// // 定義 interface
// // type Men interface {
// // 	SayHi()
// // 	Sing(lyrics string)
// // 	Guzzle(beerStein string)
// // }

// // type YoungChap interface {
// // 	SayHi()
// // 	Sing(song string)
// // 	BorrowMoney(amount float32)
// // }

// // type ElderGent interface {
// // 	SayHi()
// // 	Sing(song string)
// // 	SpendSalary(amount float32)
// // }

// // Interface Men 被 Human,Student 和 Employee 實現
// // 因為這三個型別都實現了這兩個方法
// type Men interface {
// 	SayHi()
// 	Sing(lyrics string)
// }

// func main() {
// 	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
// 	paul := Student{Human{"paul", 21, "010120312"}, "HAV", 92222222.11}
// 	sam := Employee{Human{"Sam", 32, "34234234234"}, "google", 12312.213}
// 	// tom := Employee{Human{"Tam", 54, "6536565"}, "amazon", 24351435.565}

// 	//定義 Men 型別的變數 i
// 	// var i Men

// 	//i 能儲存 Student
// 	// i = mike
// 	// fmt.Println("This is Mike, a Student:")
// 	// i.SayHi()
// 	// i.Sing("November rain")

// 	//i 也能儲存 Employee
// 	// i = tom
// 	// fmt.Println("This is tom, an Employee:")
// 	// i.SayHi()
// 	// i.Sing("Born to be wild")

// 	//定義了 slice Men
// 	fmt.Println("Let's use a slice of Men and see what happens")
// 	x := make([]Men, 3)
// 	//這三個都是不同型別的元素，但是他們實現了 interface 同一個介面
// 	x[0], x[1], x[2] = &paul, &sam, &mike

//		for _, value := range x {
//			value.SayHi()
//		}
//	}
//
// ----------------------------------------------
// type Human struct {
// 	name  string
// 	age   int
// 	phone string
// }

// // 透過這個方法 Human 實現了 fmt.Stringer
// func (h Human) String() string {
// 	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
// }

// func main() {
// 	Bob := Human{"Bob", 39, "000-7777-XXX"}
// 	fmt.Println("This Human is : ", Bob)
// }

// interface 變數儲存的型別-------------------------------------------------

// goroutine
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world") //開一個新的 Goroutines 執行
// 	say("hello")    //當前 Goroutines 執行
// }

// func say(s string, wg *sync.WaitGroup) {
// 	defer wg.Done() // 在函數返回時，通知 WaitGroup 這個 goroutine 已經完成
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)            // 增加一個等待的 goroutine
// 	go say("world", &wg) // 開一個新的 Goroutines 執行
// 	say("hello", &wg)    // 當前 Goroutines 執行
// 	wg.Wait()            // 等待所有的 goroutine 完成
// }

// channel----------------------------------------------------------------------
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		// fmt.Println(a)

		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) //後三個元素
	go sum(a[len(a)/2:], c) // =前三個元素
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
