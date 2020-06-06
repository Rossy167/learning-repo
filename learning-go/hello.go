package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

func main() {

	// using an env from os
	name := os.Getenv("USER")
	fmt.Println(name)

	// pointers etc
	course := "Yeet"
	changeCourse(&course)
	fmt.Println(course)

	// getting type of a var
	fmt.Println(reflect.TypeOf(name))

	// woah calling some runtime info
	fmt.Println("Hello from", runtime.GOOS)

	//calling variadic func
	no := vary(1, 2, 4, 5, 6, 7, 10)
	fmt.Println(no)

	// i mean yh, condtions are very normal, tho i don't like how else formats
	if true {
		fmt.Println("true is true")
	} else {
		println("this code will never be ran")
	}

	// i doubt i'll be using else ifs but yh
	if 1 > 2 {
		fmt.Println("this code shouldn't be ran")
	} else if 2 > 1 {
		fmt.Println("1 is smaller than 2 scrub")
	} else {
		fmt.Println("wtf")
	}

	// switches!!1 <3 rip python
	x := 1
	switch x {
	case 2:
		fmt.Println("nah")
	case 3:
		fmt.Println("nah")
	case 4:
		fmt.Println("nah")
	case 1, 6:
		fmt.Println("yh")
		fallthrough
	default:
		//fallthrough will run the one directly below... weird command ^
		fmt.Println("wot")
	}

	//errors, gonna return "no such file or directory" if can't find
	_, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println("Error was", err)
	} else {
		fmt.Println("let's fucking gooooo")
	}

}

// using pointers
func changeCourse(course *string) {
	*course = "Yote"
	fmt.Println("Changing course to", *course)
}

// variadic function
func vary(no ...int) int {

	high := no[0]
	for _, i := range no {
		if i > high {
			high = i
		}
	}
	return high
}
