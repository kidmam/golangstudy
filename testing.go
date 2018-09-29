package main

import (
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	myint, ok := quick.Value(reflect.TypeOf(123), rand.New(rand.NewSource(time.Now().Unix())))
	mystring, ok := quick.Value(reflect.TypeOf("abc"), rand.New(rand.NewSource(time.Now().Unix())))
	mybool, ok := quick.Value(reflect.TypeOf(true), rand.New(rand.NewSource(time.Now().Unix())))
}
