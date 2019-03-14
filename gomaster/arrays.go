package main

import "fmt"

func main() {
	var distros [5]string
	var ids [5]int

	distros[0] = "Ubuntu"
	distros[1] = "CentOS"
	distros[2] = "RedHat"
	distros[3] = "Debian"
	distros[4] = "OpenBSD"

	ids[0] = 1
	ids[1] = 2
	ids[2] = 3
	ids[3] = 4
	ids[4] = 5

	_ = distros[1]
	_ = ids[3]

	fmt.Println(distros)
	fmt.Println(ids)

	fmt.Println("distros[2] = ", distros[2])
	fmt.Println("ids[2] = ", ids[2])

}
