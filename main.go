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
