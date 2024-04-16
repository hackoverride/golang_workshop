package examples

import (
	"errors"
	"fmt"
	"time"

	"gogame/workshop/tools"
)

type Employee struct {
	Name string
	Age  int
}

func (e Employee) ToString() string {
	return fmt.Sprintf("%s is %d years old.", e.Name, e.Age)
}

func (e Employee) Birthday() {

}

// BasicDataTypes demonstrates the declaration and use of basic data types in Go.
func BasicDataTypes() {
	fmt.Println("Basic Data Types:")

	var myInt int = 42
	myInt2 := 42
	fmt.Printf("Integers: %d, %d\n", myInt, myInt2)

	var myFloat float64 = 3.14159265359
	myString := "Hello, Cap!"
	fmt.Printf("Float and String: %f, %s\n", myFloat, myString)
	tools.Pause()
}

// ArraysAndSlices demonstrates how to work with arrays and slices in Go.
func ArraysAndSlices() {
	fmt.Println("Arrays and Slices:")

	var myArray [3]int = [3]int{1, 2, 3}
	stringArray := [3]string{"Hello", "Go", "Workshop"}
	anotherArray := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Arrays:", myArray, stringArray, anotherArray)
	fmt.Println("Array Length:", len(anotherArray))

	var mySlice []int
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("Slice:", mySlice)
	tools.Pause()
}

// MapsAndStructs demonstrates how to use maps and structs in Go.
func MapsAndStructs() {
	fmt.Println("Maps and Structs:")

	myMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map:", myMap)

	allEmployees := []Employee{
		{Name: "Jørgen", Age: 29},
		{Name: "Henning", Age: 39},
		{Name: "Thor", Age: 42},
	}
	fmt.Println("Employees:", allEmployees)
	tools.Pause()
}

// ComplexStructs demonstrates the use of nested structs, maps, and slices.
func ComplexStructs() {
	fmt.Println("Complex Structs:")

	type Company struct {
		Name       string
		Employees  []Employee
		Department map[string]int
		Meta       map[string]string
		Created    time.Time
	}

	myCompany := Company{
		Name: "Go Workshop",
		Employees: []Employee{
			{Name: "Jørgen", Age: 29},
			{Name: "Henning", Age: 39},
			{Name: "Thor", Age: 42},
		},
		Department: map[string]int{"HR": 3, "IT": 5},
		Meta:       map[string]string{"location": "Copenhagen", "country": "Denmark"},
		Created:    time.Now(),
	}

	fmt.Println("Company:", myCompany)
	tools.Pause()
}

// Pointers demonstrates how to use pointers in Go.
func Pointers() {
	fmt.Println("Pointers:")

	var myPointer *int = new(int)
	*myPointer = 42

	fmt.Println("Pointer Value:", *myPointer)
	fmt.Println("Pointer Address:", myPointer)

	myInt := 42
	myPointer = &myInt

	fmt.Println("Pointer to myInt:", *myPointer)
	fmt.Println("Pointer Address:", myPointer)
	tools.Pause()
}

// BasicFunction demonstrates a simple function with parameters and a return value.
func BasicFunction(name string, age int) string {
	return fmt.Sprintf("%s is %d years old.", name, age)
}

// HighOrderFunction demonstrates a function that takes another function as a parameter.
func HighOrderFunction(f func(int) int, val int) int {
	return f(val)
}

// Square is a simple function to square a number.
func Square(x int) int {
	return x * x
}

// Closure demonstrates the use of closures in Go.
func Closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ExampleFunctions calls and demonstrates all the functions defined above.
func ExampleFunctions() {
	// Using BasicFunction
	result := BasicFunction("Alice", 30)
	fmt.Println(result)
	tools.Pause()

	// Using HighOrderFunction with Square as the argument
	squared := HighOrderFunction(Square, 4)
	fmt.Println("Square of 4 is:", squared)
	tools.Pause()

	// Demonstrating a closure
	counter := Closure()
	fmt.Println("First call to counter:", counter())
	fmt.Println("Second call to counter:", counter())
	fmt.Println("Third call to counter:", counter())
	tools.Pause()
}

// Example function calls all demonstration functions.
func Example() {
	BasicDataTypes()
	ArraysAndSlices()
	MapsAndStructs()
	ComplexStructs()
	Pointers()
	ExampleFunctions()
}

func ExampleForTest(a, b int) (int, error) {
	result := a + b
	if result == 0 {
		return 0, errors.New("result is zero, which is not allowed")
	}
	return result, nil
}
