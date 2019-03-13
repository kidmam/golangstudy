package main

import (
	"fmt"
	"github.com/openacid/slim/index"
	"strings"
)

type Data string

func (d Data) Read(offset int64, key string) (string, bool) {
	kv := strings.Split(string(d)[offset:], ",")[0:2]
	if kv[0] == key {
		return kv[1], true
	}
	return "", false
}

func main() {
	// `data` is a sample external data.
	// In order to let SlimIndex be able to read data, `data` should have
	// a `Read` method:
	//     Read(offset int64, key string) (string, bool)
	data := Data("Aaron,1,Agatha,1,Al,2,Albert,3,Alexander,5,Alison,8")

	// keyOffsets is a prebuilt index that stores key and its offset in data accordingly.
	keyOffsets := []index.OffsetIndexItem{
		{Key: "Aaron", Offset: 0},
		{Key: "Agatha", Offset: 8},
		{Key: "Al", Offset: 17},
		{Key: "Albert", Offset: 22},
		{Key: "Alexander", Offset: 31},
		{Key: "Alison", Offset: 43},
	}

	// Create an index
	st, err := index.NewSlimIndex(keyOffsets, data)
	if err != nil {
		fmt.Println(err)
	}

	// Lookup by SlimIndex
	v, found := st.Get2("Alison")
	fmt.Printf("key: %q\n  found: %t\n  value: %q\n", "Alison", found, v)

	v, found = st.Get2("foo")
	fmt.Printf("key: %q\n  found: %t\n  value: %q\n", "foo", found, v)

	// Output:
	// key: "Alison"
	//   found: true
	//   value: "8"
	// key: "foo"
	//   found: false
	//   value: ""
}
