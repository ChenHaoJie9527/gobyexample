/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2026-03-27 12:13:50
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2026-03-27 12:44:35
 * @FilePath: \gobyexample\examples\constants\constants.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Go 支持字符、字符串、布尔和数值 _常量_ 。
package main

import (
	"fmt"
	"math"
)

// `const` 用于声明一个常量。
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` 语句可以出现在任何 `var` 语句可以出现的地方
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
	fmt.Println(int64(d))

	// 一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。
	// 举个例子，这里的 `math.Sin` 函数需要一个 `float64` 的参数，`n` 会自动确定类型。
	fmt.Println(math.Sin(n))

	const b = 10000000
	fmt.Println(math.Sin(b))

	c := getUser()
	fmt.Println(c)
}

func getUser() string {
	return "John"
}
