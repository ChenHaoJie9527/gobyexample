# Go `switch` 分支结构教程

这个示例展示了 Go 里非常常用的一种多分支控制结构：`switch`。

和 `if/else if/else` 相比，`switch` 更适合处理“一个值对应多个分支”或者“多个条件按顺序匹配”的场景。Go 还在这个语法上提供了无表达式 `switch` 和类型开关，所以它在实际项目里很常见。

## 1. 这个文件里用到了什么技术

`switch.go` 里主要用了下面几种技术：

### 1.1 基本值匹配

```go
i := 2
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}
```

这里展示了最基础的 `switch` 用法：

- `switch` 后面跟一个表达式
- `case` 按值匹配
- 命中某个分支后执行对应代码

这类写法通常用来替代一长串 `if/else if`。

### 1.2 一个 `case` 匹配多个值

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}
```

这里展示了两个重点：

- 一个 `case` 可以用逗号匹配多个值
- `default` 用来处理兜底逻辑

这种写法在状态分类、枚举处理、权限角色判断里很常见。

### 1.3 无表达式 `switch`

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}
```

这种写法等价于一组 `if/else if/else`，只是形式更整齐。

特点是：

- `switch` 后面不写表达式
- 每个 `case` 本身就是一个布尔条件
- 从上到下匹配，第一个满足条件的分支会被执行

### 1.4 类型开关

```go
whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}
```

这是 Go 里比较有代表性的高级用法：`type switch`。

它的作用是：

- 判断接口值的动态类型
- 在不同类型下走不同逻辑
- 在各个 `case` 中拿到对应类型的变量

在处理中间件、通用工具、序列化、错误分类等场景时会见到这种写法。

## 2. 主流开发里是怎么使用 `switch` 的

在真实项目里，`switch` 往往不是拿来打印数字，而是用来组织复杂但结构清晰的分支逻辑。

### 2.1 处理状态机或枚举值

```go
switch order.Status {
case "pending":
    handlePending(order)
case "paid":
    handlePaid(order)
case "cancelled":
    handleCancelled(order)
default:
    return fmt.Errorf("unknown status: %s", order.Status)
}
```

常见场景：

- 订单状态流转
- 任务状态处理
- 消息类型分发
- 角色或权限判断

### 2.2 根据区间条件分类

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 60:
    fmt.Println("C")
default:
    fmt.Println("D")
}
```

这种无表达式 `switch` 很适合：

- 分数评级
- 风险等级分类
- 请求限流等级判断
- 配置命中策略

### 2.3 类型分发

```go
switch v := data.(type) {
case string:
    fmt.Println("string:", v)
case []byte:
    fmt.Println("bytes:", len(v))
default:
    fmt.Printf("unknown: %T\n", v)
}
```

常见场景：

- 通用接口入参处理
- 自定义日志或调试输出
- JSON 或消息体的动态解析
- 错误类型细分处理

### 2.4 替代过长的 `if/else if`

当多个条件本质上是在做“多选一”时，主流 Go 代码通常更倾向用 `switch`，因为结构更紧凑，也更容易维护。

## 3. 这个例子对应的 Go 惯用法

### 3.1 `switch` 默认只执行命中的一个分支

Go 的 `switch` 和很多语言不同，默认不会自动往下穿透。

```go
switch n {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

命中 `case 1` 后，只会执行这一段，不会继续执行 `case 2`。

这让 Go 的 `switch` 更安全，也更不容易产生遗漏 `break` 的 bug。

### 3.2 分支多时优先考虑 `switch`

如果代码像这样：

```go
if status == "pending" {
    // ...
} else if status == "paid" {
    // ...
} else if status == "cancelled" {
    // ...
} else {
    // ...
}
```

通常改成 `switch` 会更清晰。

### 3.3 `default` 不是必须，但很多时候建议写

如果你能确定所有分支都覆盖到了，可以不写 `default`。但在业务开发里，很多时候保留兜底逻辑更安全，尤其是：

- 外部输入不可控
- 枚举值未来可能扩展
- 状态来自数据库或第三方系统

### 3.4 类型开关适合接口值，不适合滥用

`type switch` 很有用，但如果一个系统里到处都在判断具体类型，往往说明抽象边界可能还不够好。主流 Go 代码仍然更提倡通过接口能力组织行为，而不是到处做类型分派。

## 4. 需要注意的坑

### 4.1 Go 的 `switch` 默认不会自动贯穿

有些开发者从 C、Java、JavaScript 转过来，容易误以为 `case` 会自动继续执行后续分支。

在 Go 中不会这样，除非你显式写 `fallthrough`。

### 4.2 `fallthrough` 要谨慎使用

Go 支持 `fallthrough`，但它会无条件进入下一个分支，而不是重新判断下一个 `case` 条件。

这意味着它很容易让逻辑变绕，除非你特别确定需要这么做，否则大多数业务代码不建议使用。

### 4.3 `case` 的顺序很重要

在无表达式 `switch` 里，条件是从上到下依次匹配的，第一个命中就结束。

错误示例思路：

```go
switch {
case score >= 60:
    fmt.Println("pass")
case score >= 90:
    fmt.Println("excellent")
}
```

如果 `score` 是 `95`，它会先命中 `score >= 60`，后面的更精确条件就永远到不了。

### 4.4 `default` 缺失可能让异常状态被静默忽略

对于核心业务状态，如果没有 `default`，一旦输入超出预期，程序可能什么都不做，问题会更难定位。

### 4.5 类型开关只能用于接口值

像下面这种写法的前提是变量是接口类型：

```go
switch v := data.(type) {
case string:
    fmt.Println(v)
}
```

如果不是接口值，就不能做类型开关。

### 4.6 不要把 `switch` 写成超长业务脚本

当一个 `switch` 的每个分支里都塞进大量逻辑时，可读性会迅速下降。更常见的工程做法是：

- `switch` 只做分发
- 具体逻辑拆到独立函数
- 让每个分支保持短小清晰

## 5. 一个更贴近业务的例子

```go
package main

import "fmt"

func main() {
    level := "warn"

    switch level {
    case "info":
        fmt.Println("record normal log")
    case "warn":
        fmt.Println("send warning notification")
    case "error":
        fmt.Println("create incident")
    default:
        fmt.Println("unknown level")
    }
}
```

这类写法常见于：

- 日志级别分发
- 事件类型处理
- 命令路由
- 消息消费者中的 handler 选择

## 6. 主流用法总结

在 Go 的日常开发中，`switch` 最常见的使用方式可以总结为：

1. 用值 `switch` 处理状态、枚举和固定分类。
2. 用多值 `case` 合并相同行为的分支。
3. 用无表达式 `switch` 替代较长的 `if/else if` 条件链。
4. 用类型开关处理接口值的不同具体类型。
5. 用 `default` 做兜底，避免异常输入被悄悄忽略。

## 7. 记忆要点

- `switch` 适合多分支的“多选一”逻辑
- Go 默认不会自动 `fallthrough`
- 无表达式 `switch` 很适合条件分类
- 类型开关用于接口值的动态类型判断
- 分支太多时，让 `switch` 负责分发，把具体逻辑拆出去

如果你愿意，下一步我也可以继续把 `arrays` 或 `range` 这类后续示例整理成同样风格的文档，方便你连续学习。
