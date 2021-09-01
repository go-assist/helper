package helper

import "fmt"

func example() {
	sub := TStr.GetBetweenStr(`@abc456%`, "@", "%")
	fmt.Println(sub) // @abc456

	arr := [5]int{1, 2, 3, 4, 5}
	it := 2
	if !TArr.InArray(it, arr) {
		fmt.Printf(" %v not in %v \n", it, arr)
	}
}
