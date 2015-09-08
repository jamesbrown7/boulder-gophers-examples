package main

import (
	"fmt"
	"reflect"
)

func add(o1, o2 interface{}) (sum interface{}, err error) {
	if reflect.TypeOf(o1).Name() != reflect.TypeOf(o2).Name() {
		return nil, fmt.Errorf("types are not the same! type of o1='%s', o2='%s'",
			reflect.TypeOf(o1).Name(), reflect.TypeOf(o2).Name())
	}

	switch reflect.TypeOf(o1).Name() {
	case "string":
		sum = o1.(string) + o2.(string)
	case "int":
		sum = o1.(int) + o2.(int)
	}

	return
}

func main() {
	sVal, err := add("test", "hub")
	if err != nil {
		fmt.Printf("Could not add two strings: err='%s'", err.Error())
	}

	fmt.Printf("sum of \"test\" and \"hub\" is: '%s'\n", sVal)

	iVal, err := add(2, 5)
	if err != nil {
		fmt.Printf("Could not add two ints: err='%s'", err.Error())
	}

	fmt.Printf("sum of 2 and 5 is: %d\n", iVal)
}
