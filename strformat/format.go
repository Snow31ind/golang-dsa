package strformat

import "fmt"

func PrintStructWithName(v any) {
	fmt.Printf("%+v\n", v)
}

func PrintType(v any) {
	fmt.Printf("%T\n", v)
}

func PrintBoolean(v bool) {
	fmt.Printf("%t\n", v)
}
