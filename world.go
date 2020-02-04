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
