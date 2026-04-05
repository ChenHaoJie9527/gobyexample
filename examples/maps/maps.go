// _map_ 是 Go 内建的[关联数据类型](http://zh.wikipedia.org/wiki/关联数组)
// （在一些其他的语言中也被称为 _哈希(hash)_ 或者 _字典(dict)_ ）。

package main

import (
	"fmt"
)

func main() {

	// 要创建一个空 map，需要使用内建函数 `make`：`make(map[key-type]val-type)`。
	m := make(map[string]int)

	// 使用典型的 `name[key] = val` 语法来设置键值对。
	m["k1"] = 7
	m["k2"] = 13

	// 打印 map。例如，使用 `fmt.Println` 打印一个 map，会输出它所有的键值对。
	fmt.Println("map:", m)

	// 使用 `name[key]` 来获取一个键的值。
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 内建函数 `len` 可以返回一个 map 的键值对数量。
	fmt.Println("len:", len(m))

	// 内建函数 `delete` 可以从一个 map 中移除键值对。
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。
	// 这可以用来消除 `键不存在` 和 `键的值为零值` 产生的歧义，
	// 例如 `0` 和 `""`。这里我们不需要值，所以用 _空白标识符(blank identifier)_ _ 将其忽略。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 你也可以通过右边的语法在一行代码中声明并初始化一个新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	mapsFn()

	isSet()
	isGroup()
	isCountBy()
	hasKey()
}

func mapsFn() {
	words := []string{"foo", "bar", "baz", "baz"}
	counts := getCounts(words)
	fmt.Println("counts:", counts) // [bar:1 baz:2 foo:1]
	fmt.Println("bar出现的次数是：", counts["bar"])
	fmt.Println("foo出现的次数是：", counts["foo"])
	fmt.Println("baz出现的次数是：", counts["baz"])
}

// 计数器 - 统计出现的次数
func getCounts(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++ // count[w]++ 能成立，是因为不存在时默认拿到零值 0。
	}
	return counts
}

// map 集合
func isSet() {
	set := make(map[string]struct{})
	set["go"] = struct{}{}
	set["rust"] = struct{}{}
	fmt.Println("set: ", set)

	if _, ok := set["go"]; ok {
		fmt.Println("go 集合")
	}

}

// group 分组
func isGroup() {
	type User struct {
		Name string
		City string
	}

	users := []User{
		{Name: "A", City: "Beijing"},
		{Name: "B", City: "Shanghai"},
		{Name: "C", City: "Beijing"},
	}

	group := make(map[string][]User)
	for _, u := range users {
		group[u.City] = append(group[u.City], u)
	}

	fmt.Println("group:", group)
}

// 练习：统计单词频次
func isCountBy() {
	words := []string{"go", "java", "go", "rust", "go", "java"}
	counts := map[string]int{}

	for _, word := range words {
		counts[word]++
	}
	fmt.Println("counts:", counts)
}

// 练习：判断 key 是否存在
func hasKey() {
	scores := map[string]int{"foo": 1, "bar": 2}

	if _, ok := scores["foo"]; ok {
		fmt.Println("foo", scores["foo"])
	}
	if _, ok := scores["bar"]; ok {
		fmt.Println("bar", scores["bar"])
	}
}

type StringSet map[string]struct{}

// Add 练习：添加到字符串集合
func (s StringSet) Add(value string) {
	s[value] = struct{}{}
}

// Remove 练习：删除指定字符串
func (s StringSet) Remove(value string) {
	delete(s, value)
}

// Has 练习：检查指定字符串
func (s StringSet) Has(value string) bool {
	_, ok := s[value]
	return ok
}

// Len 练习：查看当前集合长度
func (s StringSet) Len() int {
	return len(s)
}
