package main

import "fmt"
//
//func main() {
//	numbers := []int{1, 2, 3, 4, 5}
//	print("all     ", numbers)
//	fromTwo := numbers[2:]
//	print("from  2 ", fromTwo)
//	print("until 2 ", numbers[:2])
//	print("middle  ", numbers[2:3])
//	numbers = append(numbers, 6)
//	print("all    ", numbers)
//	print("from 2 ", fromTwo)
//}
//func print(name string, nrs []int) {
//	fmt.Print(name, "len(", len(nrs), ") cap(", cap(nrs), "):")
//	for _, nr := range nrs {
//		fmt.Print(" ", nr)
//	}
//	fmt.Println()
//}

func main() {
	numbers := [5]int{1,2,3,4,5}

	dump("numbers  ", numbers[:])
	dump("from  2  ", numbers[2:])
	dump("until 2  ", numbers[:2])
	middle := numbers[2:3]
	dump("middle   ", middle)
	fmt.Println("Appending 6")
	newNumbers := append(middle, 6)
	dump("numbers  ",  numbers[:])
	dump("new nrs  ",  newNumbers)
}

func dump(name string, nrs []int) {

	fmt.Print(name, "len(", len(nrs), ") cap(", cap(nrs),"):")
	for _, nr := range nrs {
		fmt.Print(" ", nr)
	}
	fmt.Println()
}
