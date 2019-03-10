package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
)

/*
	This is a short Go program that demonstrates how a list of similar objects (i.e. different subclass types)
	can be defined in json and then unmarshaled into a Go struct.
	This struct can then be re-marshaled into a json string.

	Look for the related article on medium.com from Alain Drolet.
	"How to Unmarshal an Array of JSON Objects of Different Types into a Go Struct"

	Licensing
	---------
	This code is distributed under a MIT style license.

	Summary:
	--------
	* Keep this copyright notice
	* Do what you want with the code
	* Do not complain if it does not work.

	More specifically
	-----------------
	Copyright (c) February 2019, Alain Drolet

	Permission is hereby granted, free of charge, to any person obtaining a copy of this software
	and associated documentation files (the “Software”), to deal in the Software without restriction,
	including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
	and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
	subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all copies or
	substantial portions of the Software.

	The Software is provided “as is”, without warranty of any kind, express or implied,
	including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement.
	In no event shall the authors or copyright holders be liable for any claim, damages or other liability,
	whether in an action of contract, tort or otherwise, arising from, out of or in connection with the software
	or the use or other dealings in the Software.
*/

// WheelList is the top-level object we want to read from a json string
// It has a name and a list of objects of any wheel subtypes
type WheelList struct {
	Name   string  `json:"name"`
	Wheels []Wheel `json:"wheels"`
}

// Wheel is a second definition for InnerWheel
// It allows Wheel and InnerWheel to have different Un/Marshalers
type Wheel InnerWheel

// InnerWheel is a struct to hold our wheel objects and their sub-types characteristics
// It uses "json" tags to map the json attributes names to the Go struct attribute names
type InnerWheel struct {
	Id     string    `json:"id"`
	Radius int       `json:"radius"`
	Type   WheelType `json:"type"`
	attr   interface{}
}

// SolidAttributes hold attributes specific to the Solid Wheels
type SolidAttributes struct {
	// material used to build the solid wheel
	// We use a string to keep this field open
	// It could be: wood, stone, steel...
	Material string `json:"material"`
}

// SpokedAttributes hold attributes specific to the Spoked Wheels
type SpokedAttributes struct {
	SpokeCount  int `json:"spoke_count"`
	InnerRadius int `json:"inner_radius"`
}

// TireAttributes hold attributes specific to the Tire Wheels
type TireAttributes struct {
	Fill        FillType `json:"fill_type"`
	InnerRadius int      `json:"inner_radius"`
}

// SolidAttr returns a pointer to a SolidAttributes struct.
// The object is read from the private `attr` attribute of the Wheel/InnerWheel.
// It is an error to call this method if Wheel.Type is not equal to WheelType_Solid.
func (w *Wheel) SolidAttr() (*SolidAttributes, error) {
	solidAttr, ok := w.attr.(*SolidAttributes)
	if !ok {
		// This is pretty bad and unexpected,
		// If w is read from json this is almost impossible since the type in attr is created by looking at w.Type
		// If built from code, then that would be a bug to be fixed.
		// Or you just called this method but did not check the type first!
		// The above comments apply to all Wheel.XAttr() methods
		err := fmt.Errorf("Implementation Error: Wheel %s failed assertion of attr field as Solid.", w.Id)
		return nil, err
	}
	return solidAttr, nil
}

// SpokedAttr returns a pointer to a SpokedAttributes struct.
// The object is read from the private `attr` attribute of the Wheel/InnerWheel.
// It is an error to call this method if Wheel.Type is not equal to WheelType_Spoked.
func (w *Wheel) SpokedAttr() (*SpokedAttributes, error) {
	spokedAttr, ok := w.attr.(*SpokedAttributes)
	if !ok {
		err := fmt.Errorf("Implementation Error: Wheel %s failed assertion of attr field as Spoked.", w.Id)
		return nil, err
	}
	return spokedAttr, nil
}

// TireAttr returns a pointer to a TireAttributes struct.
// The object is read from the private `attr` attribute of the Wheel/InnerWheel.
// It is an error to call this method if Wheel.Type is not equal to WheelType_Tire.
func (w *Wheel) TireAttr() (*TireAttributes, error) {
	tireAttr, ok := w.attr.(*TireAttributes)
	if !ok {
		err := fmt.Errorf("Implementation Error: Wheel %s failed assertion of attr field as Tire.", w.Id)
		return nil, err
	}
	return tireAttr, nil
}

//-------------------- Start Wheel Un/Marshalers ----------------------------------

// UnmarshalJSON unmarshals a json string representing a Wheel and its specific subtype attributes.
// It will parse once to get the common attributes: Id, Radius, Type.
// Then based on Type it will parse a second time to get the extra attributes specific to the WheelType in w.
// The unmarshaled attributes are set in the receiver w.
func (w *Wheel) UnmarshalJSON(b []byte) error {
	// populate common attributes using default (InnerWheel) json Unmarshaler
	err := json.Unmarshal(b, (*InnerWheel)(w))
	if err != nil {
		return err
	}

	switch w.Type {
	case WheelType_Solid:
		subTypeAttr := &SolidAttributes{}
		err := json.Unmarshal(b, subTypeAttr)
		if err != nil {
			return err
		}
		w.attr = subTypeAttr
	case WheelType_Spoked:
		subTypeAttr := &SpokedAttributes{}
		err := json.Unmarshal(b, subTypeAttr)
		if err != nil {
			return err
		}
		w.attr = subTypeAttr
	case WheelType_Tire:
		subTypeAttr := &TireAttributes{}
		err := json.Unmarshal(b, subTypeAttr)
		if err != nil {
			return err
		}
		w.attr = subTypeAttr
	default:
		return fmt.Errorf("Wheel.UnmarshalJSON: unexpected type; type = %s (%d)", w.Type.String(), w.Type)
	}

	return nil
}

// MarshalJSON returns the json string representation of the Wheel object in the receiver w.
// To do so it computes the json string from the Wheel common attributes,
// then it computes the json string for the subtype object stored in `attr`
// The returned values is the union of both json strings
func (w *Wheel) MarshalJSON() ([]byte, error) {

	commonStr, comErr := json.Marshal((*InnerWheel)(w))
	if comErr != nil {
		return []byte(""), comErr
	}

	attrStr, attErr := json.Marshal(w.attr)
	if attErr != nil {
		return []byte(""), attErr
	}

	// We expect something like
	// 	common = {"x":"y","p":"q"}
	// 	attr   = {"a":"b","c":"d"}
	// 	the result = {"x":"y","p":"q","a":"b","c":"d"}
	// So we create a string with at each end a curly brace
	// After the opening brace we copy the common string, without its surrounding braces
	// If We have both a common and an attr strings we add a comma
	// Then we add the attr string, again without its surrounding braces
	// Overall very simple, but full support for special cases make it a bit more verbose

	commExists := len(commonStr) > 2
	attrExists := len(attrStr) > 2 && attrStr[0] == '{' // if attr is nil attrStr will be "null"

	buf := make([]byte, 0, (len(commonStr) + len(attrStr) - 1))

	buf = append(buf, byte('{'))

	if commExists {
		buf = append(buf, commonStr[1:len(commonStr)-1]...)
	}

	if commExists && attrExists {
		buf = append(buf, byte(','))
	}

	if attrExists {
		buf = append(buf, attrStr[1:len(attrStr)-1]...)
	}

	buf = append(buf, byte('}'))

	return buf, nil
}

//-------------------- End Wheel Un/Marshalers ------------------------------------

// Go has limited support for enums
// For each enum type in this file we will use:
// * a const block
// * a slice of strings to store names (order matters!)
// * a map to find the int value of an enum from its string
// func and method will be added to make this usable

//-------------------- Start WheelType --------------------------------------------

// WheelType is the list of wheel subtypes we understand
type WheelType int

const (
	WheelType_Solid WheelType = iota
	WheelType_Spoked
	WheelType_Tire
)

var wheelNames = [...]string{
	"solid",
	"spoked",
	"tire",
}

var wheelValues = map[string]WheelType{
	"solid":  WheelType_Solid,
	"spoked": WheelType_Spoked,
	"tire":   WheelType_Tire,
}

// String converts a WheelType value to its string representation
// To convert string to enum value use the function WheelTypeValue
func (t WheelType) String() string {

	// This test will help detect when the wheelNames array is out of sync with the list of WheelType values
	// At the very least you should avoid a panic
	if t < 0 || int(t) >= len(wheelNames) {
		return ""
	}
	return wheelNames[t]
}

// WheelTypeValue take a string as an argument and attempt to return the equivalent WheelType enum value.
// It is the inverse of WheelType.String()
// If it is impossible to convert the string an error is returned, and the value returned should be ignored.
func WheelTypeValue(name string) (WheelType, error) {
	val, found := wheelValues[name]
	if !found {
		// we must return some value even on error, let's pick the default value WheelType_Solid
		return WheelType_Solid, fmt.Errorf("name '%s' not found in enum WheelType", name)
	}

	return val, nil
}

// MarshalJSON converts an enum instance of WheelType to its string representation surrounded by double-quotes.
func (t WheelType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON attempts to unmarshal a quoted json string to its enum value
// An error is returned if the string does not represent a known WheelType value
func (t *WheelType) UnmarshalJSON(b []byte) error {
	// unquote the argument; we could possibly use strconv.Unquote
	var valname string
	err := json.Unmarshal(b, &valname)
	if err != nil {
		return err
	}

	val, err := WheelTypeValue(valname)
	if err != nil {
		return fmt.Errorf("Cannot convert '%s' into a value of the WheelType enum", valname)
	}

	*t = val
	return nil
}

//-------------------- End WheelType ----------------------------------------------

//-------------------- Start FillType ---------------------------------------------
type FillType int

const (
	FillType_Inflated FillType = iota
	FillType_Plain
)

var fillNames = [...]string{
	"inflated",
	"plain",
}

var fillValues = map[string]FillType{
	"inflated": FillType_Inflated,
	"plain":    FillType_Plain,
}

// String converts a FillType value to its string representation
// To convert string to enum value use the function FillTypeValue
func (t FillType) String() string {

	// This test will help detect when the fillNames array is out of sync with the list of FillType values
	// At the very least you should avoid a panic
	if t < 0 || int(t) >= len(fillNames) {
		return ""
	}
	return fillNames[t]
}

// FillTypeValue take a string as an argument and attempt to return the equivalent FillType enum value.
// It is the inverse of FillType.String()
// If it is impossible to convert the string an error is returned, and the value returned should be ignored.
func FillTypeValue(name string) (FillType, error) {
	val, found := fillValues[name]
	if !found {
		// we must return some value even on error, let's pick the default value FillType_Inflated
		return FillType_Inflated, fmt.Errorf("name '%s' not found in enum FillType", name)
	}

	return val, nil
}

// MarshalJSON converts an enum instance of FillType to its string representation surrounded by double-quotes.
func (t FillType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON attempts to unmarshal a quoted json string to its enum value
// An error is returned if the string does not represent a known FillType value
func (t *FillType) UnmarshalJSON(b []byte) error {
	// unquote the argument; we could possibly use strconv.Unquote
	var valname string
	err := json.Unmarshal(b, &valname)
	if err != nil {
		return err
	}

	val, err := FillTypeValue(valname)
	if err != nil {
		return fmt.Errorf("Cannot convert '%s' into a value of the FillType enum", valname)
	}

	*t = val
	return nil
}

//-------------------- End FillType ------------------------------------------------

// myWheelList is an example of a json string that could represent a list of wheel objects
const myWheelListJson = `
{
  "name": "my-wheel-list",
  "wheels": [
    {
      "id": "chariot_wheel_0142",
      "radius": 17,
      "type": "solid",
      "material": "wood"
    },
    {
      "id": "bicycle_wheel_8450",
      "radius": 20,
      "type": "spoked",
      "spoke_count": 30,
      "inner_radius": 19
    },
    {
      "id": "car_wheel_2743",
      "radius": 20,
      "type": "tire",
      "inner_radius": 17,
      "fill_type": "inflated"
    }
  ]
}
`

// Demonstrates how a string can be converted to a Go struct,
// then how the struct can be converted back to a json string.
func main() {

	wheelList := &WheelList{}

	// convert json string in myWheelList to a Go struct
	unMarErr := json.Unmarshal([]byte(myWheelListJson), wheelList)
	if unMarErr != nil {
		fmt.Printf("Unmarshal error: %s\n", unMarErr.Error())
		return
	}

	// If you want to see the internal of the wheelList struct
	// comment-in the next 2 lines and the spew import line at the top
	// make sure you have the github.com/davecgh/go-spew/spew package installed in your Go workspace
	//
	//fmt.Println("\nDump of the Go struct after unmarshaling (output from spew).\n")
	//spew.Dump(wheelList)

	// pretty-print the struct in wheelList to the stdout
	wlJson, marErr := json.MarshalIndent(wheelList, "", "    ")
	if marErr != nil {
		fmt.Printf("Marshal error: %s\n", marErr.Error())
	}

	fmt.Printf("\n-----\nMarshaled List\n%s\n", wlJson)
}
