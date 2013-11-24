// goTest
package main

// the art of import
// you HAVE TO use, what you import!
import (
	"fmt"
	"time"
)

// globals
var iGlobalMessage string = "i am a global message"

// This is how you declare a global constant
const iGlobalConst string = "i am a global const and CAN NOT be changed"

// the entry point of the program
func main() {
	/*
		MultiLine
		comments

		VERRY goOD samples at: http://blog.golang.org/
	*/

	// print a line to console
	fmt.Println("Hello World!")

	// print formated to console
	fmt.Printf("%s -- %s -- %s<br>\n", time.Now(), iGlobalMessage, iGlobalConst)

	// get user input by line(return)
	fmt.Println("enter new text for iGlobalMessage")
	fmt.Scanln(&iGlobalMessage)
	// string concat with +
	fmt.Println("new global text:" + iGlobalMessage)

	//# playing with vars

	var someString string
	var someStringWithText string = "some text"
	// vars can be autocasted to a type by its assigned value
	someStringToo := "!<>some string<>!"
	// we ALWAYS HAVE to use vars!
	// no var left behind! (the compiler does not like unused vars)
	fmt.Printf("'%s' - '%s' - '%s'\n", someString, someStringWithText, someStringToo)

	var someInt int
	someIntToo := 22
	fmt.Printf("%d -- %d\n", someInt, someIntToo)

	// change values
	someInt = 99
	someIntToo = 33
	fmt.Printf("%d -- %d\n", someInt, someIntToo)

	// increment ints
	someInt++
	someIntToo += 1
	fmt.Printf("%d -- %d\n", someInt, someIntToo)

	var multipleStrings1, multipleStrings2, multipleStrings3 string
	fmt.Printf("%s -- %s -- %s<br>\n", multipleStrings1, multipleStrings2, multipleStrings3)

	multipleStringsToo1, multipleStringsToo2, multipleStringsToo3 := "foo1", "foo2", "foo3"
	fmt.Printf("%s -- %s -- %s<br>\n", multipleStringsToo1, multipleStringsToo2, multipleStringsToo3)

	// you can ONLY instance a var once
	// this will fail:
	//someIntToo := 99

	//# functions

	// function without parameter and return value
	uselessPrint()

	// function with 1 parameter
	uselessPrintParameter(42)

	// function with 2 parameters
	uselessPrintMore("the answer to all:", 23)

	// function with return value
	var uselessReturnTemp string
	uselessReturnTemp = uselessReturn()
	// this would work too:
	//var uselessReturnTemp string = uselessReturn()
	//uselessReturnTempToo := uselessReturn()
	//fmt.Println(uselessReturn())
	fmt.Println(uselessReturnTemp)

	// function with multiple return values
	theAnswerText, theAnserNumber := uselessMultipleReturn()
	// will work too
	/*
		var theAnswerText string
		var theAnserNumber int
		theAnswerText, theAnserNumber = uselessMultipleReturn()
	*/
	fmt.Printf("%s: %s\n", theAnswerText, theAnserNumber)

	//# conditions
	var matchMe int = 23
	if matchMe == 7 {
		fmt.Println("the 7 is all")
	} else if matchMe == 23 {
		fmt.Println("dont trust machines")
	} else {
		fmt.Println("this is not the value i am looking for")
	}

	switch {
	case matchMe == 7:
		fmt.Println("the 7 is all")
	case matchMe == 23:
		fmt.Println("dont trust machines")
	default:
		fmt.Println("this is not the value i am looking for")
	}

	//# array, list of one type with unchangeable length
	var myArray [3]int
	fmt.Println(myArray)
	myArray[1] = 42
	fmt.Println(myArray)

	//# slice, list of one type that can vary in length
	var mySlice []int
	fmt.Println(mySlice)
	mySlice = append(mySlice, 21, 42, 101)
	fmt.Println(mySlice)
	fmt.Println(mySlice[1])

	//# map, key/value pair of two types (dictionary)
	var myMap map[string]int
	// you have to "make" them
	myMap = make(map[string]int)
	// set some values
	myMap["the answer"] = 42
	myMap["half_theTruth"] = 11
	myMap["weTrust"] = 23
	fmt.Println(myMap)
	// access one value
	fmt.Println(myMap["half_theTruth"])
	// change value
	myMap["half_theTruth"] = 21
	fmt.Println(myMap["half_theTruth"])

	// access key->value pair
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// delete key+value
	delete(myMap, "half_theTruth")
	fmt.Println(myMap)
}

// functions (no shit!)
func uselessPrint() {
	fmt.Println("print....")
}

func uselessPrintParameter(answerToAll int) {
	fmt.Println(answerToAll)
}

func uselessPrintMore(someText string, answerToAll int) {
	// we have to cast the int to string before combining them
	fmt.Println(someText + string(answerToAll))
}

func uselessReturn() string {
	var theAnswerToAll string = "the anser to all"
	return theAnswerToAll
}

func uselessMultipleReturn() (string, int) {
	var theAnswerToAll string = "the answer to all is"
	return theAnswerToAll, 42
}
