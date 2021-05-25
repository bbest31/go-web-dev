package main

import "fmt"

type person struct {
	fname string
	lname string
	Age   uint
}

type programmer struct {
	person
	Language string
	Backend  bool
	Frontend bool
}

// any type that implements the speak method implicitly implments the human interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// func for person struct
func (p person) speak() {
	fmt.Println(p.fname, `says, "Hello there."\n`)
}

func (pr programmer) speak() {
	fmt.Printf(`My name is %v and I know how to code in %v\n`, pr.fname, pr.Language)
}

func main() {
	xi := []int{1, 2, 3}
	fmt.Println(xi)

	m := map[string]int{
		"Brandon": 27,
		"Emma":    23,
	}

	fmt.Println(m)

	p1 := person{"Brandon", "Best", 27}

	fmt.Println(p1)
	// p1.speak()

	pr1 := programmer{
		p1,
		"Golang",
		true,
		true,
	}

	// pr1.speak()

	saySomething(pr1)
	saySomething(p1)
}
