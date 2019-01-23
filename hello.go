// define the main package (potentially starting point of go program)
package main

// Import modules from the std lib
import "fmt"
import "strconv"

var yeet = "yeet"

var c, python, java string = "c", "python", "java"

// Create a function that returns the sum of two numbers
func addition() int {
    num1 := 10
    num2 := 30
    return num1 + num2
}

// Add two numbers passed as arguments
// if two args are of the same type, the type can just be placed at the end of
// the func declaration addNums(x, y int)
func addNums(x int, y int) int {
    return x + y
}

// Swap values to demonstrate that multiple values can be returned from a go function
func swap(x, y string) (string, string) {
    return y, x
}

// Demonstrate the use of naked returns, where the returns are inferred where you define the
// the types of the values being returned (good for short functions, get confusing with longer ones)
func split(sum int)(x, y int) {
    x = sum / 2
    y = sum / 2
    return 
}

// This is the main function
func main() {
    // Simple variable declaration and output to std out
    // type is inferred using the := operator
    msg := "hello there go!\n"
    fmt.Printf(msg)

    // Use strconv module to convert integers to strings (and probably other types)
    fmt.Printf(strconv.Itoa(addition()) + "\n")
    fmt.Printf(strconv.Itoa(addNums(10, 10)) + "\n")

    // Pass two values into swap, get two back
    bye, hi := swap("hi", "bye")
    fmt.Printf(bye + " " + hi + "\n") 

    // Some more messing around
    x, y := split(100)
    fmt.Printf(strconv.Itoa(x) + strconv.Itoa(y))
}
