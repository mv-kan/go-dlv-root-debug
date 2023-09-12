package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NeedRootAccess() {
	file := "/you_can_remove_it"
	fmt.Printf("Trying to write to file \"%s\"\n", file)
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(file, d1, 0777)
	check(err)

	fmt.Printf("\"%s\": write success\n", file)
	os.Remove(file)
	fmt.Printf("\"%s\": deletion success\n", file)
}
