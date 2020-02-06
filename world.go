package main

import "fmt"

//The "fmt" package is imported and it will be used inside the main function to print text to the standard output.
import "math"
import "unsafe"
import "time"
import "sync"

/*
What kind of variable type?

bool
Numeric Types
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64
complex64, complex128
byte
rune
string

*/

/*
---------------- structure ---------------
What is a structure?
A structure is a user defined type which represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than maintaining each of them as separate types.

For instance a employee has a firstName, lastName and age. It makes sense to group these three properties into a single structure employee.
*/
type image struct {
	data map[int]int
}

type Address struct {
	city, state string
}

type Employee struct {
	FirstName   string
	LastName    string
	age, salary int
	address     Address
}
type Empl struct {
	name     string
	salary   int
	currency string
}
type myInt int

/*

What is an interface?
The general definition of an interface in the Object oriented world is "interface defines the behaviour of an object".
메서드들의 집합체라고 할 수 있음.
interface는 타입이 구현해야 되는 메서드 원형들을 정의한다.

*/
type VowelsFinder interface {
	FindVowels() []rune
}

func (a myInt) add(b myInt) myInt {
	return a + b
}

func (e Empl) displaySalary() {
	fmt.Println("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	//The main is a special function. The program execution starts from the main function.
	var (
		name string = "richard"
		year int    = 2020
		tall int    = 0
	)

	var age int
	fmt.Println("hello world my age is ", age)
	age = 23
	fmt.Println("my new age is ", age)
	var width, height int = 100, 50
	// if int would be dropped , It's possible to run your code.
	fmt.Println("width = ", width, "height =", height)
	fmt.Println("name = ", name, ", year =", year, ", tall = ", tall)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is", c)

	d := true
	e := false
	f := d && e
	fmt.Println("f true && false :", f)
	g := d || e
	fmt.Println("g true || false : ", g)

	var h int = 85
	var i int = 10

	fmt.Println("value of h is", h, "and i is", i)
	fmt.Println("type of h is %T, size of h is %d", h, unsafe.Sizeof(h))   //type and size of h
	fmt.Println("\ntype of i is %T, size of i is %d", i, unsafe.Sizeof(i)) //type and size of i

	//int + float64 not allowed
	//sum := i + int(j) //j is converted to int

	var j float64 = float64(i)
	fmt.Println("j", j)

	const con = 48 // reassignment not allowed . math function not allowed.
	var names = "Sam"
	fmt.Println("type %T value %v", names, names)
	// %T is printed as a type. %v is printed as a value.
	type myString string
	var customName myString = "Sam" //allowed // main.myString type , value is sam.
	fmt.Println("type %T value %v", customName, customName)
	//customName = defaultName //not allowed

	const n = 5
	var intVar int = n
	var int32Var int32 = n
	var float64Var float64 = n
	var complex64Var complex64 = n
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var z int = 10
	var x int = 20
	v := sum(z, x)
	fmt.Println("total sum = ", v)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Println("Area %f Perimeter %f", area, perimeter)
	/*
		func functionname(parametername type) returntype {
			//function body
		}
	*/

	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	/*
		if statement; condition {
		}

		for loop
		for initialisation; condition; post {
		}


	*/
	for i := 1; i <= 10; i++ {

		if i > 5 {
			break
			// or continue
		}
		fmt.Println(" %d", i)
	}

	s := 5
	for i := 0; i < s; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*
		infinite loop
		The syntax for creating an infinite loop is,

		for {
		}
	*/

	yum := 4
	//switch num := number(); { //num is not a constant
	switch yum { // it's possible to put multiple expressions in case
	case 1:
		fmt.Println("it's 1")
	case 2:
		fmt.Println("it's 2")
	case 3:
		fmt.Println("it's 3")
	case 4:
		fmt.Println("correct")
	default:
		fmt.Println("none")
	}

	var arr [3]int // array length 3
	arr[0] = 2
	arr[1] = 3
	arr[2] = 4
	fmt.Println(arr)
	fmt.Println("arr len = ", len(arr))
	//a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	ar := [3]int{12}
	fmt.Println(ar)

	y := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(y)
	var k [3][2]string
	k[0][0] = "apple"
	k[0][1] = "samsung"
	k[1][0] = "microsoft"
	k[1][1] = "google"
	k[2][0] = "AT&T"
	k[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(k)
	/*
		slice using golang

		  a := [5]int{76, 77, 78, 79, 80}
		    var b []int = a[1:4] //creates a slice from a[1] to a[3]
		    fmt.Println(b)

	*/
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
		// index + 1 from 2 to 4
	}
	fmt.Println("array after", darr)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	//How to create a map?

	//personSalary := make(map[string]int)
	personSalary := make(map[string]int)
	if personSalary == nil {
		fmt.Println("map is nil . going to make one.")
		personSalary = make(map[string]int)
	}
	personSalary = make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["mike"] = 15000
	personSalary["richard"] = 90000
	fmt.Println("salary contents of map", personSalary)

	/*
		it's also allowed the way.

		personSalary := map[string]int {
		        "steve": 12000,
		        "jamie": 15000,
		    }
		    personSalary["mike"] = 9000
		    fmt.Println("personSalary map contents:", personSalary)


		how to delete element of map >
			delete(personSalary, "steve")

		length of the map >
			len(personSalary)

	*/
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	if map1["one"] == map2["one"] && map1["two"] == map2["two"] {
		fmt.Println("map1 is equal to map2")
	} else {
		fmt.Println("map1 is not equal to map2")
	}

	/*
		map1 := map[string]int{
		        "one": 1,
		        "two": 2,
		    }

		    map2 := map1

		    if map1 == map2 {
		    }

		map comparison code is not working its like this.

	*/

	/* String */

	str := "Hello World"
	fmt.Println(str)

	printBytes(str)
	printCharacters(str)

	printChars(str)
	printCharsAndBytes(str)

	gu := "hello"
	fmt.Println(mutate([]rune(gu)))

	/* [Go] to Pointer */
	/* A pointer is a variable which stores the memory address of another variable. */
	number_b := 255
	var number_a *int = &number_b

	fmt.Printf("Type of a is %T\n", number_a) // Type of a is *int
	fmt.Println("address of b is", number_a)  // 0xc0001341c8

	ni := 25
	var nb *int
	if nb == nil {
		// The zero value of a pointer is nil.
		fmt.Println("b is ", nb) // b is nil
		nb = &ni
		fmt.Println("b after initialization is ", nb) // b after initialization is 0xc000016208

	}
	hum := hello()
	fmt.Println("returning pointer from a function : result : = ", hum)
	fmt.Println("returning pointer from a function : result : = ", *hum)
	*hum++
	fmt.Println("result+1 : = ", *hum)

	test_arr := [3]int{90, 86, 94}
	modify(&test_arr)
	fmt.Println(test_arr) // [90 86 94] -> [90 86 10]

	/* [Go] to Structure */
	// structure code is at top of code. /|\
	//									  |
	/*

		type Address struct {
		    city, state string
		}
		type Employee struct{
			FirstName 	  string
			LastName  	  string
			age,salary 	  int
			address 	  Address
		}
	*/

	emp1 := Employee{
		FirstName: "Richard",
		LastName:  "Lee",
		age:       24,
		salary:    500,
	}
	fmt.Println("emp1 : ", emp1) // emp1 :  {Richard Lee 24 500}
	////zero valued structure is allowed.
	//Employee 4 {  0 0}  <--- example

	fmt.Println("emp1 firstname ->", emp1.FirstName)
	fmt.Println("emp1 lastname ->", emp1.LastName)
	fmt.Println("emp1 age ->", emp1.age)
	fmt.Println("emp1 salary ->", emp1.salary)
	//Accessing individual fields of a struct

	emp8 := &Employee{"Sam", "Anderson", 55, 6000, Address{city: "LA", state: "CA"}}
	fmt.Println("First Name:", (*emp8).FirstName)
	fmt.Println("Age:", (*emp8).age)
	/* Pointers to a struct
	It is also possible to create pointers to a struct. */

	var emp3 Employee
	emp3.FirstName = "Justin"
	emp3.LastName = "Bieber"
	emp3.age = 26
	emp3.salary = 500000000
	emp3.address = Address{
		city:  "Los Angeles",
		state: "California",
	}
	fmt.Println("emp3 firstname ->", emp3.FirstName)
	fmt.Println("emp3 lastname ->", emp3.LastName)
	fmt.Println("emp3 age ->", emp3.age)
	fmt.Println("emp3 salary ->", emp3.salary)
	fmt.Println("emp3 address ->", emp3.address.city)
	fmt.Println("emp3 address state ->", emp3.address.state)
	/*
		image1 := image{data: map[int]int{
			0: 155,
		}}
		image2 := image{data: map[int]int{
			0: 155,
		}}

		hence image1 and image2 cannot be compared.
		If you run this program, compilation will fail with error
	*/

	/* Methods in Go  */

	/*
		func(t Type) function name(parameter list){

		}
	*/

	//emp1.displaySalary()
	emp5 := Empl{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp5.displaySalary()

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)

	/* Interface code is at the top of code. */

	/* Concurrency in Go */
	/*
		Concurrency means multiple computations are happening at the same time.
		Concurrency is an inherent part of the Go programming language.
		Concurrency is the capability to deal with lots of things at once. It's best explained with an example.
	*/

	/*
		What are Goroutines?

		Goroutines are functions or methods that run concurrently with other functions or methods.
		Goroutines can be thought of as light weight threads.
		The cost of creating a Goroutine is tiny when compared to a thread.
		Hence its common for Go applications to have thousands of Goroutines running concurrently.

		This program only outputs the text main function.
		What happened to the Goroutine we started?
		We need to understand two main properties of go routines to understand why this happens.

		When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
		The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
	*/
	done := make(chan bool)

	fmt.Println("Main going to call hello go goroutine")
	go hi(done)
	<-done
	fmt.Println("Main received data")
	//time.Sleep(1 * time.Second)
	fmt.Println("main function")

	var wait sync.WaitGroup
	wait.Add(2)
	// create wait group , waiting for 2 (go routine)

	go func() {
		defer wait.Done() // call Done() after the end.
		fmt.Println("Hello World man yo !")
	}()

	go func(msg string) {
		defer wait.Done() // call Done() after the end.
		fmt.Println(msg)
	}("Hi man")

	wait.Wait() // wait until all routine be finished.

	go numbers()
	go alphabet()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	//1 a 2 3 b 4 c 5 d e

	/*

		What are channels

		Channels can be thought as pipes using which Goroutines communicate.
		Similar to how water flows from one end to another in a pipe,
		data can be sent from one end and received from the another end using channels.
	*/

	var lee chan int
	if lee == nil {
		fmt.Println("channel lee is nil , going to define it.")
		lee = make(chan int)
		fmt.Printf("Type of lee is  %T \n", lee)
	}

	/*

		Sending and receiving from channel

		The syntax to send and receive data from a channel are given below,

		data := <- a // read from channel a
		a <- data // write to channel a

		for more information , you can see on the issue page.
	*/
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech // read from channel sqrch, cubech
	fmt.Println("Final output", squares+cubes)

}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
	//data write to channel squareop
}
func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func hi(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true // write to channel done
}

func modify(arr *[3]int) {
	arr[2] = 10
}

//Returning pointer from a function
func hello() *int {
	i := 5
	return &i
}

func mutate(g []rune) string {
	runes := []rune(g)
	runes[0] = 'a' //any valid unicode character within single quote is a rune
	return string(runes)
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		//It's gonna be printed Unicode UT8-encoded values of "Hello World"
		fmt.Printf("%x ", s[i])
	}
	fmt.Println(" ")
}

func printCharacters(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println(" ")
}

func printChars(s string) {
	runes := []rune(s)
	/*
		A rune is a builtin type in Go and it's the alias of int32.
		rune represents a Unicode code point in Go. It does not matter how many bytes the code point occupies,
		it can be represented by a rune. Lets modify the above program to print characters using a rune.
	*/
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	//H e l l o   W o r l d
}

func printCharsAndBytes(s string) {
	fmt.Printf("\n")
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
	/*
		output

		H starts at byte 0
		e starts at byte 1
		l starts at byte 2
		l starts at byte 3
		o starts at byte 4
		  starts at byte 5
		W starts at byte 6
		o starts at byte 7
		r starts at byte 8
		l starts at byte 9
		d starts at byte 10
	*/
}

/* sum function , */
func sum(price int, no int) int {
	var total_price = price * no
	return total_price
}

/* It has multiple return type  */
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*
area and perimeter are the named return values in the below function.

func rectProps(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

init function , Every package can contain a init function.
func init(){

}
*/
func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
