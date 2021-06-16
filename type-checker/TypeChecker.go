package main

import (
	"errors"
	"fmt"
)

func typeCheck(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return "string", nil
	case int, int32, int64:
		return "int", nil
	case float32, float64:
		return "float", nil
	case bool:
		return "bool", nil
	default:
		return "{}", errors.New("unsupported type passed")
	}
}

func main() {
	result, _ := typeCheck(-5)
	fmt.Println("-5 :", result)
	result, _ = typeCheck(5)
	fmt.Println("5 :", result)
	result, _ = typeCheck("yum")
	fmt.Println("yum :", result)
	result, _ = typeCheck(true)
	fmt.Println("true :", result)
	result, _ = typeCheck(float32(3.14))
	fmt.Println("3.14 :", result)
	result, _ = typeCheck(struct{}{})
	fmt.Println("struct:", result)
}
