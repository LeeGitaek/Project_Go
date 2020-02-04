package main

import "fmt"

//The "fmt" package is imported and it will be used inside the main function to print text to the standard output.
import "math"
import "unsafe"

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
