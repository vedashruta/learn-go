package arrays

import "fmt"

func Arrays() {
	var intArray1 []int = []int{1, 2, 3, 4, 5, 6, 7}
	intArray2 := [3]int{1, 2, 3}
	intArray3 := [...]int{1, 2, 3, 4, 5}
	intArray4 := [...]int{4: -1}
	fmt.Println("array 1", intArray1)
	fmt.Println("array 2", intArray2)
	fmt.Println("array 3", intArray3)
	fmt.Println("array 4", intArray4)

	//built-in append
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	z := append(x, y)
	fmt.Println("array x ", x)
	fmt.Println("array y", y)
	fmt.Println("array z", z)
	fmt.Println("length of z array", len(z))
	fmt.Println("capacity of z array", cap(z))
	zz := z[3:4]
	fmt.Println("array zz", zz)
	fmt.Println("length of zz array", len(zz))
	fmt.Println("capacity of zz array", cap(zz))
}

func append(x, y []int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow
		z = z[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
