package main

import (
	"fmt"
)

func main() {

	/*var iMap map[string]int
	    iMap["k1"] = 12
		iMap["k2"] = 13*/ //assignment to entry in nil map

	//map literal
	/*iMap := map[string]int {
		"k1": 12,
		"k2": 13,
	}*/

	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13

	fmt.Println("iMap", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)
	delete(anotherMap, "k1")
	//delete(anotherMap, "k1")
	//delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	_, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does Not Exist")
	}

	for key, value := range iMap {
		fmt.Println(key, value)
	}

	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap:", aMap)
	//aMap["test"] = 1 //assignment to entry in nil map
}
