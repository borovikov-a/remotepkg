package remotepkg

import (
	"fmt"
)

func AddAndPrint(a, b int) int {
	res := a + b
	fmt.Printf("This function has same porpose as the operator "+". Result of adding %d to %d is %d.\n", a, b, res)
	return res
}
