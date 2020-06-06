package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	name := os.Getenv("USER")
	course := "Yeet"
	fmt.Println(name)
	fmt.Println("Hello from", runtime.GOOS)
	changeCourse(&course)
	fmt.Println(course)

}

func changeCourse(course *string) {
	*course = "Yote"
	fmt.Println("Changing course to", *course)
}
