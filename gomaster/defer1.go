package main

import "fmt"

func main() {
	deferExample()
	fmt.Println("Returned from deferExample.")
	deferPanicExample()
	fmt.Println("Returned from deferPanicExample.")
	deferRecurseExample()
	fmt.Println("Returned from deferRecurseExample.")
}

func deferPanicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic", r)
		}
	}()
	fmt.Println("Call recurse.")
	recursePanic(0)
	fmt.Println("Return normally from recurse.")
}

func recursePanic(count int) {
	if count > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", count))
	}
	defer fmt.Println("Defer log in recurse", count)
	fmt.Println("Print line in recurse", count)
	recursePanic(count + 1)
}

func deferExample() {
	defer fmt.Println("Deferred log.")
	fmt.Println("Print line.")
}

func deferRecurseExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic", r)
		}
	}()
	fmt.Println("Call recurse.")
	recursePanic(0)
	fmt.Println("Return normally from recurse.")
}

func recurse(count int) {
	if count > 3 {
		fmt.Println("Stopping recursion!")
		return
	}
	defer fmt.Println("Defer log in recurse", count)
	fmt.Println("Print line in recurse", count)
	recurse(count + 1)
}
