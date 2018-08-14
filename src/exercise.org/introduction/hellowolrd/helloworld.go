// package  declaration is  the first statement in a Go source file.
package main // executable commands must always use package main.

import (
	"fmt"

	"exercise.org/introduction/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
