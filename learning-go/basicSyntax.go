package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
)

func main() {
	structs()
}

func structs() {
	// woah oo in go or yes, no, maybe
	type person struct {
		firstName   string
		surnameName string
		fullName    string
		age         int
	}

	//var Rossy person
	//Rossy := new(person)

	var name = "Louis"
	var surname = "Ross"

	Rossy := person{
		firstName:   name,
		surnameName: surname,
		fullName:    (name + " " + surname),
		age:         22,
	}

	fmt.Println("Name is:", Rossy.fullName)
	Rossy.fullName = "Rossy167"
	fmt.Println("Name is:", Rossy.fullName)

	// references as always, fast etc

}

func maps() {
	// maps are cheap, they are references and therefore bad for concurrency (which btw, hype for learning i've never done that before)
	// map[keyType]valueType
	stuff := make(map[string]int)
	stuff["Rossy"] = 167
	stuff["Louis"] = 22

	bigStuff := map[string]int{
		"Example":      12,
		"otherExample": 10,
		"yh":           7,
		"np":           5,
		"aaa":          109,
		"yeuayd":       289,
	}

	fmt.Println(stuff["Louis"])
	fmt.Println(bigStuff)

	// map is unordered list, yo hey, basically think dicts in python
	for key, value := range bigStuff {
		fmt.Println(key, "is", value)
	}

	// to delete
	delete(bigStuff, "aaa")
}

func misc() {
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

}

func slicesAndArrays() {
	// slices are built on top of arrays
	// they are references to an array
	// which means they are natively passed by reference
	// defining a slice, overload to not define capacity ez
	stuff := make([]string, 5, 10)
	fmt.Printf("len is %d and capacity is %d \n", len(stuff), cap(stuff))
	otherStuff := []string{"thing", "yeah", "blegh"}

	fmt.Println(otherStuff)
	otherStuff = append(otherStuff, "yeet")
	fmt.Println(otherStuff)
	fmt.Println("added", otherStuff[3])

	// ^ will add the capacity if it is already full, does it 4, 8, 16, 32 etc
}

func conditionals() {
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
}

func errors() {
	// errors, gonna return "no such file or directory" if can't find
	_, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println("Error was", err)
	} else {
		fmt.Println("let's fucking gooooo")
	}
}

func forLoops() {
	// for loops <expression>
	for i := 10; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	//for loop range
	stuff := []string{"yes", "no", "maybe"}
	otherStuff := []string{"aaa", "yes", "lemons"}
	for _, i := range stuff {
		for _, x := range otherStuff {
			if i == x {
				fmt.Println(x, "is", i)
			}
		}
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
