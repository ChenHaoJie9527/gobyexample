# Go `if/else` 条件分支教程

这个示例展示了 Go 里最基础的条件分支写法：`if`、`if/else`、`else if`，以及 `if` 前置短变量声明。

如果你刚开始学 Go，这个例子非常重要，因为几乎所有业务代码都会频繁使用条件判断来决定程序接下来该做什么。

## 1. 这个文件里用到了什么技术

`if-else.go` 里主要用了下面几种技术：

### 1.1 基本条件判断

```go
if 7%2 == 0 {
    fmt.Println("7 is even")
} else {
    fmt.Println("7 is odd")
}
```

这里包含了两个点：

- 用 `%` 取模判断奇偶
- 用 `if/else` 根据布尔表达式选择分支

### 1.2 只有 `if` 没有 `else`

```go
if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
}
```

这说明在 Go 里，`else` 不是必须的。只在“条件成立时才做事”的场景下，单独使用 `if` 更直接。

### 1.3 多分支判断

```go
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
```

这里展示了：

- `else if` 用于连续分支判断
- `if` 前可以先写一个短变量声明
- 这个变量只在当前 `if / else if / else` 语句块中有效

### 1.4 布尔表达式驱动流程

Go 的条件判断必须是明确的布尔值，比如：

- `7%2 == 0`
- `num < 0`
- `num < 10`

Go 不支持把整数、字符串直接当条件使用，这和一些动态语言不同。

## 2. 主流开发里是怎么使用 `if/else` 的

在真实项目里，`if/else` 最常见的用途不是判断奇偶，而是控制业务流程。

### 2.1 参数校验

```go
if userID == "" {
    return errors.New("userID is required")
}
```

常见场景：

- API 入参校验
- 配置检查
- 请求字段合法性判断

### 2.2 错误处理

Go 项目里最常见的 `if`，通常是这个：

```go
if err != nil {
    return err
}
```

这也是 Go 最主流的风格之一：先处理异常路径，让正常流程自然往下走。

### 2.3 按业务状态分支

```go
if order.IsPaid {
    ship(order)
} else {
    waitForPayment(order)
}
```

常见场景：

- 用户是否登录
- 订单是否支付
- 资源是否存在
- 功能开关是否开启

### 2.4 配合短变量声明缩小作用域

```go
if v, ok := data[key]; ok {
    fmt.Println(v)
}
```

这是 Go 很常见的惯用法，适合：

- 读取 map 时判断 key 是否存在
- 类型断言后判断是否成功
- 调用返回值后只在局部范围内使用临时变量

## 3. 这个例子对应的 Go 惯用法

从主流 Go 风格来看，这个示例背后有几个值得记住的习惯：

### 3.1 条件不用圆括号

Go 里这样写：

```go
if x > 0 {
    fmt.Println(x)
}
```

而不是这样：

```go
if (x > 0) {
    fmt.Println(x)
}
```

圆括号不是必须的，但花括号必须写。

### 3.2 优先让逻辑保持简单

Go 社区很强调“清晰胜过巧妙”。条件判断能写简单就不要写复杂，能提前返回就不要层层嵌套。

例如：

```go
if err != nil {
    return err
}

process()
```

通常比下面这种更符合 Go 风格：

```go
if err == nil {
    process()
} else {
    return err
}
```

### 3.3 分支太多时考虑 `switch`

如果判断条件越来越多，`else if` 链很长，通常可以考虑改成 `switch`，代码会更清楚。

## 4. 需要注意的坑

### 4.1 条件必须是布尔值

错误写法：

```go
if 1 {
    fmt.Println("invalid")
}
```

在 Go 中这会直接编译失败。条件表达式必须是 `bool`。

### 4.2 花括号不能省略

错误写法：

```go
if x > 0
    fmt.Println(x)
```

Go 强制要求使用花括号，即使只有一行代码也不能省略。

### 4.3 `if` 里的短变量作用域有限

像示例中的：

```go
if num := 9; num < 0 {
    fmt.Println(num)
}
```

这里的 `num` 只能在当前条件分支链里使用，离开这段 `if` 后就访问不到了。

### 4.4 不要把分支写得太深

例如多层 `if` 套多层 `if`，会让代码越来越难读。Go 更推荐：

- 先处理异常分支
- 尽早 `return`
- 保持 happy path 清晰

### 4.5 不要滥用 `else`

如果 `if` 分支里已经 `return`、`break` 或 `continue` 了，后面通常不需要再套一个 `else`。

例如更推荐：

```go
if err != nil {
    return err
}

return nil
```

而不是：

```go
if err != nil {
    return err
} else {
    return nil
}
```

### 4.6 变量遮蔽要小心

短变量声明很方便，但如果在复杂代码里重复使用 `:=`，可能会不小心创建一个新的同名变量，导致你以为修改了外层变量，实际没有。

## 5. 一个更贴近业务的例子

下面这个例子更像真实开发中的用法：

```go
package main

import "fmt"

func main() {
    age := 20

    if age < 0 {
        fmt.Println("invalid age")
    } else if age < 18 {
        fmt.Println("minor")
    } else {
        fmt.Println("adult")
    }
}
```

这类写法经常出现在：

- 用户等级判断
- 权限控制
- 状态分类
- 风险分级

## 6. 主流用法总结

在 Go 的日常开发中，`if/else` 最常见的使用方式可以总结为：

1. 用 `if` 做参数校验和错误处理。
2. 用 `if/else` 根据业务状态走不同流程。
3. 用 `else if` 处理少量连续分支。
4. 用 `if` 前置短变量声明缩小变量作用域。
5. 当分支很多时，优先考虑 `switch`，避免长链式判断。

## 7. 记忆要点

- 条件必须是 `bool`
- 圆括号可省略，花括号不能省略
- `if` 前可以先声明变量
- 短变量只在当前分支链中有效
- Go 更推荐提前返回，少写没必要的 `else`
