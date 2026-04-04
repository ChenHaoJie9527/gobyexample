// _Slice_ 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

package main

import "fmt"

func main() {

	// 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 `make`。
	// 这里我们创建了一个长度为 3 的 `string` 类型的 slice（初始值为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 除了基本操作外，slice 支持比数组更丰富的操作。比如 slice 支持内建函数 `append`，
	// 该函数会返回一个包含了一个或者多个新值的 slice。
	// 注意由于 `append` 可能返回一个新的 slice，我们需要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice 还可以 `copy`。这里我们创建一个空的和 `s` 有相同长度的 slice——`c`，
	// 然后将 `s` 复制给 `c`。
	c := make([]string, len(s))
	fmt.Println("创建空切片c: ", c, len(c))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice 支持通过 `slice[low:high]` 语法进行“切片”操作。
	// 例如，右边的操作可以得到一个包含元素 `s[2]`、`s[3]` 和 `s[4]` 的 slice。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 包含从 `s[0]` 到 `s[5]`（不包含 5）的元素。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 包含从 `s[2]`（包含 2）之后的元素。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
	twoD := make([][]int, 3)
	fmt.Println("多维切片", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	test()
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	arr8 := []int{1, 2, 3}
	test7(arr8)
	fmt.Println("arr8:", arr8) // [100, 2, 3]
	test8()
	arr9 := test9([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println("arr9:", arr9)

	//arr10 := []int{1, 2, 3, 4, 5}
	var arr10 []int
	arr11 := cloneInts(arr10)
	fmt.Println("arr10:", arr10)
	fmt.Println("arr11:", arr11)

	arr12 := []int{1, 2, 3}
	arr12 = deleteInts(arr12, 1)
	//arr12 = deleteInts(arr12, 0)
	//arr12 = deleteInts(arr12, 2)
	fmt.Println("arr12:", arr12)

	arr13 := []int{1, 2, 3}
	arr13 = insertAt(arr13, 1, 10)
}

// 基础题
func test() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println("len:", len(arr1)) // 5
	fmt.Println("cap:", cap(arr1)) // 5
}

// 切片题
func test1() {
	arr := []int{10, 20, 30, 40, 50}
	arr1 := arr[1:4] // [20, 30, 40]
	fmt.Println("arr1:", arr1)
}

// append
func test2() {
	var arr []int

	arr = append(arr, 1)

	fmt.Println("arr:", arr)

	arr = append(arr, 2)
	fmt.Println("arr:", arr)

	arr = append(arr, 3)
	fmt.Println("arr:", arr)

	arr = append(arr, 4)
	fmt.Println("arr:", arr)

	arr = append(arr, 5)
	fmt.Println("arr:", arr)
}

// 修改题
func test3() {
	arr := []int{1, 2, 3}
	arr[1] = 99
	fmt.Println("arr:", arr)
}

// 共享底层数组题
func test4() {
	arr := []int{1, 2, 3}
	arr1 := arr[:1]
	fmt.Println("arr1:", arr1)
	arr1[0] = 100
	fmt.Println("arr1:", arr1)
	fmt.Println("arr:", arr)
}

// copy题
func test5() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 100
	fmt.Println("dst:", dst)
	fmt.Println("src:", src)
}

// nil vs empty题
func test6() {
	var a []int
	b := []int{}
	fmt.Println("a:", len(a))
	fmt.Println("b:", len(b))
	fmt.Println("a:", a == nil) // true
	fmt.Println("b:", b == nil) // false
}

// 函数传参题
func test7(arr []int) {
	arr[0] = 100
}

// 预分配题
func test8() {
	arr := make([]string, 0, 10)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
	arr = append(arr, "a", "b", "c")
	fmt.Println("arr:", arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
}

func test9(arr []int) []int {
	dst1 := make([]int, len(arr))
	copy(dst1, arr)

	var arr1 []int

	for _, v := range arr {
		if v%2 == 0 {
			arr1 = append(arr1, v)
		}
	}

	return arr1
}

// 实现 cloneInts 功能
func cloneInts(arr []int) []int {
	if len(arr) == 0 || arr == nil {
		emptyArr := make([]int, 0)
		return emptyArr
	}
	cloneArr := make([]int, len(arr))
	copy(cloneArr, arr)
	cloneArr[1] = 100
	return cloneArr
}

// 实现 deleteInts
func deleteInts(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		dst := make([]int, len(arr))
		return dst
	}

	dst := make([]int, len(arr))
	copy(dst, arr)

	// i=1时 => append(arr[:1], arr[1+1:]...) => append([1], [3]...) => append([1], 3) => [1, 3]
	// i=0时 => append(arr[:0], arr[0+1:]...) => append([], arr[1]...) => append([], 1) => [1]
	// i=2时 => append(arr[:2], arr[2+1:]...) => append([1,2], arr[3]...) => append([1,2], []) => [1,2]
	dst = append(arr[:i], arr[i+1:]...)
	return dst
}

func insertAt(arr []int, i int, val int) []int {
	// TODO: insertAt([1,2,3], 1, 10) => result: [1,10,2,3]
	if i < 0 || i >= len(arr) {
		return arr
	}
	dst := make([]int, 0, len(arr)+1)
	dst = append(dst, arr[:i]...)
	dst = append(dst, val)
	dst = append(dst, arr[i:]...)

	fmt.Println("dst:", dst)
	return dst
}
