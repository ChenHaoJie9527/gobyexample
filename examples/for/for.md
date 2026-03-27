# Go `for` 循环教程

`for` 是 Go 语言里唯一的循环结构。它既可以扮演 `while` 的角色，也可以写成经典的三段式循环，还可以直接写成无限循环。

这篇文档会从常见写法、应用场景、注意事项和主流用法几个角度，帮你把 `for` 用熟。

## 1. `for` 的基本形式

Go 里的 `for` 主要有三种写法：

### 1.1 条件循环

这种写法最像其他语言里的 `while`。

```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}
```

适合：

- 不确定循环次数，但知道结束条件
- 轮询状态
- 读取数据直到满足某个退出条件

### 1.2 三段式循环

这是最经典的计数循环写法。

```go
package main

import "fmt"

func main() {
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
}
```

适合：

- 固定次数重复执行
- 遍历一段数字区间
- 批量重试

### 1.3 无限循环

如果省略所有条件，`for` 就变成无限循环。

```go
package main

import "fmt"

func main() {
	for {
		fmt.Println("loop")
		break
	}
}
```

适合：

- 常驻任务
- 事件循环
- 配合 `select` 处理并发消息

## 2. 常见应用场景

### 2.1 遍历固定范围

当你只想做 N 次事情时，三段式 `for` 最直接。

```go
for i := 0; i < 3; i++ {
	fmt.Println("retry", i)
}
```

典型场景：

- 重试请求
- 生成测试数据
- 运行固定轮次的任务

### 2.2 持续等待条件满足

条件循环很适合“等到某个状态变化再退出”。

```go
for !ready {
	time.Sleep(100 * time.Millisecond)
}
```

典型场景：

- 等待服务启动
- 轮询外部状态
- 读取直到 EOF

### 2.3 过滤数据

`continue` 能让你跳过不需要的元素。

```go
for n := 0; n <= 5; n++ {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
```

典型场景：

- 跳过无效数据
- 跳过空行
- 跳过已处理项

### 2.4 配合并发模型

在 Go 里，`for` 常常和 `goroutine`、`channel`、`select` 一起使用。

```go
for {
	select {
	case msg := <-ch:
		fmt.Println("got:", msg)
	case <-done:
		return
	}
}
```

典型场景：

- worker 常驻消费
- 监听消息队列
- 服务内部事件分发

## 3. 需要注意的坑

### 3.1 忘记退出条件

无限循环如果没有 `break`、`return` 或阻塞点，很容易卡死。

```go
for {
	fmt.Println("never ends")
}
```

注意：

- 这是最常见的死循环来源
- 写之前先想好退出条件

### 3.2 `continue` 让后续逻辑跳过

如果更新状态的代码写在 `continue` 后面，它就永远不会执行。

```go
for n := 0; n < 5; n++ {
	if n%2 == 0 {
		continue
	}
	// 这行之后的逻辑只会处理奇数
	fmt.Println(n)
}
```

注意：

- 必须执行的清理、计数、记录日志，尽量放在 `continue` 之前

### 3.3 循环变量被闭包捕获

这是 Go 初学者很容易踩的坑，尤其在 `go func()` 里。

```go
for i := 0; i < 3; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

问题：

- goroutine 里拿到的可能不是你以为的那个 `i`

更安全的写法：

```go
for i := 0; i < 3; i++ {
	i := i
	go func(v int) {
		fmt.Println(v)
	}(i)
}
```

### 3.4 边界条件写错

`<` 和 `<=` 的差别很小，但后果很大。

```go
for i := 0; i < 3; i++ {
	fmt.Println(i)
}
```

注意：

- `i < 3` 会输出 `0, 1, 2`
- `i <= 3` 会输出 `0, 1, 2, 3`

### 3.5 忽略 `for range`

遍历数组、切片、map、string、channel 时，Go 更推荐 `for range`。

```go
nums := []int{1, 2, 3}
for idx, n := range nums {
	fmt.Println(idx, n)
}
```

原因：

- 更简洁
- 更符合 Go 的惯用写法
- 少写索引逻辑，少出边界错误

## 4. 主流用法总结

在实际 Go 项目里，`for` 最常见的几种用法是：

1. 用 `for range` 遍历切片、数组、字符串、map、channel。
2. 用三段式 `for` 做计数循环。
3. 用 `for {}` 加 `select` 写事件循环或 worker。
4. 用 `continue` 跳过无效数据，用 `break` 提前结束循环。

## 5. 一个完整例子

下面这个例子把常见写法放在一起，适合你边看边跑。

```go
package main

import "fmt"

func main() {
	fmt.Println("条件循环")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("三段式循环")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("无限循环")
	for {
		fmt.Println("loop once")
		break
	}

	fmt.Println("continue 示例")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

## 6. 记忆口诀

- 不知道次数，用条件循环
- 知道次数，用三段式循环
- 要一直跑，用无限循环
- 遍历数据，优先 `for range`
- 跳过当前项，用 `continue`
- 提前结束，用 `break`

如果你愿意，下一步我也可以继续帮你把这份 `for.md` 扩展成“`for` 和 `range` 的对比教程”，这样会更适合你现在跟着 `PROGRESS.md` 往下学。
