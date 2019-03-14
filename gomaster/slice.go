package main

import "fmt"

func main() {

	distros := make([]string, 1, 5)
	fmt.Println("length = ", len(distros))
	fmt.Println("capacity = ", cap(distros))

	books := []struct {
		id        int
		name      string
		available bool
	}{
		{1, "Linux-Book", true},
		{3, "Windows-Book", false},
		{5, "Mac-Book", true},
	}
	fmt.Println("Number of books = ", len(books))

	var mySlice []string
	if mySlice == nil {
		fmt.Println("mySlice is empty!")
	}

	yourSlice := make([]string, 0)
	if yourSlice != nil {
		fmt.Println("yourSlice is not empty!", yourSlice)
	}

	names := []string{}
	myNames := append(names, "Nick")
	myNames = append(myNames, "Rita", "Pam")
	fmt.Println(myNames)

	duplicateNames := make([]string, len(myNames))
	copy(duplicateNames, myNames)
	fmt.Println(duplicateNames)

	count := 1
	var multi = make([][]int, 4)
	for i := 0; i < 4; i++ {
		multi[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			multi[i][j] = count
			count++
		}
	}
	fmt.Println(multi)

	for index, value := range myNames {
		fmt.Println(index, " = ", value)
	}
}
