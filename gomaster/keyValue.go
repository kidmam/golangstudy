package main

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k+ == nil) {
		DATA[k] = n
	}
	return false
}

func main() {

}
