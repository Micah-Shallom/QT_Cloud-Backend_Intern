package main

import "fmt"

func main() {

	int_list := [...]int{1, 2, 3, 4, 5}
	info_list := [...]string{"Shallom Micah", "backend", "Golang"}
	for i, num := range int_list {
		fmt.Printf(" num_array[%d] = ", i)
		fmt.Println(num)
	}
	for i, info := range info_list {
		fmt.Printf("Info_array[%d] = ", i)
		fmt.Println(info)
	}

}

