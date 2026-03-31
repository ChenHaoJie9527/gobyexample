// _Slice_ 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 3)
	fmt.Println("arr:", arr)

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
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	part1 := []int{1, 2, 3, 4, 5}
	part2 := part1[1:4]
	fmt.Println("part1:", part1) // [1, 2, 3, 4, 5]
	fmt.Println("part2:", part2) // [2, 3, 4]

	// part1和part2共享底层数组，所以修改part2会影响part1
	part2[0] = 100
	fmt.Println("part1:", part1) // 为什么是[1, 100, 3, 4, 5]，而不是[100, 2, 3, 4, 5]？
	fmt.Println("part2:", part2) // 下标为0的元素被改成100，[100, 3, 4]

	// ------------------------------------------------------------
	test()
}

func test() {
	// 练习1：创建并打印：创建一个 []int 类型的 slice，长度为 5，然后打印它。
	nums := make([]int, 5)
	// var nums = [5]int{}
	fmt.Println("nums:", nums)

	// 练习 2：赋值和取值: 创建一个长度为 3 的 []string，依次赋值为 "Go"、"is"、"fun"，然后打印第二个元素
	str := make([]string, 3)
	str[0] = "Go"
	str[1] = "is"
	str[2] = "fun"
	fmt.Println("str[1]:", str[1])

	// 练习 3：用 append 追加元素: 创建一个空的 []int，依次追加 10、20、30，最后打印整个 slice
	arr := make([]int, 0)
	arr = append(arr, 10)
	arr = append(arr, 20)
	arr = append(arr, 30)
	fmt.Println("arr:", arr)

	// 练习4 切出一段子 slice
	list := []int{1, 2, 3, 4, 5, 6}
	part1 := list[1:4] // [2, 3, 4]
	fmt.Println("part1:", part1)
	part2 := list[:3] // [1, 2, 3]
	fmt.Println("part2:", part2)
	part3 := list[3:] // [4, 5, 6]
	fmt.Println("part3:", part3)

	// 练习5：复制一份独立数据：创建一个新的 dst，把 src 的内容复制过去，然后修改 dst[0]，要求 src 不受影响
	src := []int{1, 2, 3}
	src1 := make([]int, len(src))
	copy(src1, src) // copy: 复制一份数据给src1，src1和src是两个不同的slice，所以修改src1不会影响src
	src1[0] = 100
	fmt.Println("src:", src)
	fmt.Println("src1:", src1)

	// 练习6：遍历求和：创建一个 []int{1, 2, 3, 4, 5}，使用 range 遍历并求和
	list1 := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range list1 {
		sum += value
	}
	fmt.Println("sum:", sum)

	// 练习7：筛选偶数：创建一个 []int{1, 2, 3, 4, 5, 6}，使用 range 遍历并筛选出所有偶数
	list2 := []int{1, 2, 3, 4, 5, 6}
	even := []int{}

	for _, value := range list2 {
		if value%2 == 0 {
			even = append(even, value)
		}
	}
	fmt.Println("even:", even)

	// 练习8：理解共享底层数据：创建一个 []int{10, 20, 30, 40, 50}，然后修改 arr3[0]，观察 arr2 是否变化
	arr2 := []int{10, 20, 30, 40, 50}
	arr3 := arr2[1:4] // [20, 30, 40]
	arr3[0] = 100
	fmt.Println("arr2:", arr2) // [10, 100, 30, 40, 50]
	fmt.Println("arr3:", arr3) // [100, 30, 40]

	// 练习 10：自己实现一个简单删除元素逻辑: 创建一个 []int{10, 20, 30, 40, 50}，然后删除下标为 2 的元素，得到 [10, 20, 40, 50]
	arr4 := []int{10, 20, 30, 40, 50}
	// arr5 := append([]int{}, arr4...)
	// arr6 := append(arr5[:2], arr5[3:]...)
	// fmt.Println("arr6:", arr6)
	arr5 := make([]int, len(arr4))
	copy(arr5, arr4)
	arr5 = append(arr5[:2], arr5[3:]...)
	fmt.Println("arr5:", arr5)
}
