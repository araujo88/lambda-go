package main

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/araujo88/lambda-go/pkg/core"
	"github.com/araujo88/lambda-go/pkg/predicate"
	"github.com/araujo88/lambda-go/pkg/tuple"
	"github.com/araujo88/lambda-go/pkg/utils"
)

func main() {
	// Example: Doubling the values in a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}
	doubled := core.Map(intSlice, func(x int) int { return x * 2 })
	stringNumbers := core.Map(intSlice, strconv.Itoa)
	sum := core.Foldr(func(x int, acc int) int { return x + acc }, 0, intSlice)
	greaterThan2 := predicate.Filter(intSlice, func(x int) bool {
		if x > 2 {
			return true
		} else {
			return false
		}
	})

	fmt.Println("Head:", utils.Head(intSlice))
	fmt.Println("Tail:", utils.Tail(intSlice))
	fmt.Println("Doubled Integers:", doubled)
	fmt.Println(stringNumbers)
	fmt.Println("Sum of integers:", sum)
	fmt.Println("Filtering elements greater than 2:", greaterThan2)

	// Example: Converting strings to uppercase in a slice of strings
	stringSlice := []string{"hello", "world"}
	uppercased := core.Map(stringSlice, strings.ToUpper)
	concat := core.Foldr(func(x string, acc string) string { return x + acc }, "", stringSlice)
	fmt.Println("Head:", utils.Head(stringSlice))
	fmt.Println("Tail:", utils.Tail(stringSlice))
	fmt.Println("Uppercased Strings:", uppercased)
	fmt.Println("Concatenated string:", concat)

	intSlice2 := []int{1, 2, 3}
	charSlice := []string{"a", "b", "c"}
	newSlice := tuple.Zip(intSlice2, charSlice)
	fmt.Println("Zipped slice:", newSlice)
}
