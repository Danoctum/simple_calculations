// Package simple_calculations provides some simple calculations
package simple_calculcations

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds the two input integers and returns the result.
// For more information, see [mathisfun].
//
// [mathisfun]: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
