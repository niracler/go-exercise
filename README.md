# go-exercise

关于go的练习,以及学习笔记

## 基础

### go get 获取第三方库

方法一：CDN 加速代理

```shell script
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量，以下三选一，首推阿里云

# 1. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

# 2. 官方
go env -w  GOPROXY=https://goproxy.io

# 3. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn
```

方法二：使用自己的代理

```shell script
export http_proxy=socks5://127.0.0.1:1080

go mod tidy
go mod download
go mod vendor
```

### 基本语法

- 变量，go语言中没有这么多花里胡哨的浅拷贝以及深拷贝，这里所有变量都是值传递。
- 选择、循环
- 指针、数组、容器。在函数要修改变量的值的情况下，我们需要将变量的指针传到函数中。但是数组那边要另外讨论

### GOPATH下目录结构

- go build 来编译
- go install 生产 pkg 文件和可执行文件
- go run 直接编译运行

## 內建容器

### 数组

- 数组就是内建容器的一种
- 数组是值类型， [3]int, [5]int, 是两个不一样的数据类型， 函数赋值的时候是值传递
- 调用func f(arr [10]int) 会拷贝数组
- 在go语言中一般不直接使用数组，而是使用切片

### 为什么要用 range

- 意义明确， 美观
- C++：没有类似能力
- Java/Python: 只能 for each value， 不能同时获取 i, v

### 切片(slice)

- 切片在函数赋值中是引用传递， 因为它是对原来数组的视图
- slice的slice，也就是reslice， 可以看后面， 但无法看前面
- slice是基于array上实现出来的
- len 是指切片的长度， 而cap是指内存的分配的长度
- s[i] 不可以超越 len(s), 向后扩展不可以超越底层数组 cap(s)
- 添加元素时如果超过cap， 系统会重新分配更大的底层数组, 两倍两倍地增
- 由于值传递的关系，必须接收append的返回值
- s = append(s, val)
- 拷贝的时候需要先有相应大小的slice才能拷贝

### map 的操作

- 创建:make(map[string]string)
- 获取元素:m[key]
- key不存在时, 获得value类型的初始值
- 用 value, ok := m[key] 来判断是否存在 key
- 用 delete 来删除一个 key
- map 的遍历使用 range 遍历 key , 或者遍历, key, value 对。不保证遍历顺序, 如需顺序, 需要手动对key进行排序
- map 使用哈希表,key必须可以比较相等
- 除了 slice, map, function 的内建类型都可以作为key
- struct 类型不包括上述类型也可以作为 key

## 面向接口

- go语言仅仅支持封装，不支持继承和多态
- go语言没有class，只有struct
- 结构体的创建可以使用自定义工厂函数（要注意返回了局部变量的地址  ）
- duck typing 的概念，当你实现了相应的接口就当你是那个类。
- 组合的思想， go语言的接口是可以组合的， 那是怎么进行组合的呢？
- 可以为结构定义方法
- 可以使用指针作为方法接收者，只有指针才可以改变结构内容
- nil指针也可以调用方法

### 值接收者vs指针接收者

- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者
- 一致性：如有指针接收者，最好都是都是指针接收者
- 值接收者是go语言特有的
- 值/指针接收者均可接收值/指针

### 封装

- 名字一般使用CamelCase
- 首字母大写：public
- 首字母小写：private

### 包

- 每个目录一个包
- main包包含可执行入口
- 为结构定义的方法必须放在同一个包内
- 可以是不同的文件

### 如何扩充系统类型或者别人的类型

- 定义别名
- 使用组合

### duck typing

- 像鸭子走路，像鸭子叫，长得像鸭子，那么就是鸭子
- 描述事物的外部行为而非内部结构
- 严格来说

python 中的 duck typing

- 运行时才知道传入的 retrieve 有没有 get
- 需要注释来说明接口

C++ 中的 duck typing

- 编译时才知道传入的 retriever 有没有 get
- 需要注释来说明接口

java中的类似代码

- 传入的参数必须实现 retrieve 接口
- 不是 duck typing

go 语言的 duck typing

- 同时需要 Readable， Appendable 怎么办？（apache polygene）
- 同时具有 python， C++ 的 duck typing 的灵活性
- 又能具有 java的类型检查

### 函数式编程

- 闭包的概念
- 可以为一个函数实现接口

### 工程化

- 资源管理，错误处理
- 测试和文档，例如这里的表格驱动测试， 然后还能根据注释生成文档。
- 示例代码，测试的时候会运行那个 example 代码
- 性能调优

### 并发编程

- goroutine和channel
- 理解调度器

### 标准库

- fmt, log
- errors
- io, bufio
- time
- net/http, net/rpc
- html/template
- charset, encoding, unicode, utf8
- strings, bytes, strconv
- regexp
- flag
- math
- os
- pprof, 主要用于性能调优
- runtime
- reflect
- testing

### 实战项目

- 从0开始，使用go语言自主搭建简单分布式爬虫
- 单任务版->并发版->分布式版

## leetcode

### array

1. [找字符串中的最大回文段](leetcode/array/longestPalindrome/main.go)
2. [腐烂的橘子](leetcode/array/orangesRotting/main.go)
3. [连续和为x的数](leetcode/array/findContinuousSequence/main.go)
4. [多数元素](leetcode/array/majorityElement/main.go)

### string

1. [最后一个单词的长度](leetcode/string/length_of_last_word/main.go)
2. [字符串的最大公因子](leetcode/string/gcd_of_strings/main.go)

### stack

1. [最小栈](leetcode/stack/min_stack/min_stack.go)

### queue

1. [滑动窗口的最大值](leetcode/queue/sliding_window_maximum/sliding_window_maximum.go)

### hash table

1. [将字符串按照字符出现频率排序](leetcode/hash_table/frequency_sort/frequencySort.go)

### dp

1. [买卖股票的最佳时机](leetcode/dp/best-time-to-buy-and-sell-stock/main.go)

### dfs

1. [找出一棵树的从根节点到叶子的所有路径](leetcode/dfs/btreepaths/main.go)
2. [机器人走路的不同路径](leetcode/dfs/uniquepath1/main.go)
3. [根节点到叶子节点的和为指定数的路径](leetcode/dfs/pathsum/main.go)
4. [判断一棵树是否是平衡树](leetcode/dfs/balancetree/main.go)

### bfs

1. [二叉树的层序遍历](leetcode/bfs/lot/main.go)
1. [判断一棵树是否是对称的](leetcode/bfs/symmetrictree/main.go)

### sort

1. [把数组排成最小的数](leetcode/sort/minNumber/main.go)

## web 框架

> 好像这个就是 go 这边用的比较多的 web 框架，用起来感觉跟 Python 的 Flask 框架差不多。

## 参考文章

- [Google资深工程师深度讲解Go语言](https://coding.imooc.com/class/180.html)
- [go-get命令使用socket代理](http://www.hi-roy.com/2018/10/12/go-get%E5%91%BD%E4%BB%A4%E4%BD%BF%E7%94%A8socket%E4%BB%A3%E7%90%86/)
- [Gin入门实战](https://www.imooc.com/learn/1175)
- [gin文档](https://gin-gonic.com/zh-cn/docs/)
- [Dave Cheney The acme of foolishness](https://dave.cheney.net/resources-for-new-go-programmers)
- [Five suggestions for setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- [KeKe-Li/book](https://github.com/KeKe-Li/book.git)
- [【韩顺平】Go语言零基础教程](https://www.youtube.com/watch?v=ryrlfsFbtVA&list=PLmOn9nNkQxJFWlwItS-iI3C-4jeARUNjq)(重要,我跟我同学立下了flag,一定要看完)
- [Dave Cheney The acme of foolishness(我觉得这个人的博客还是挺棒的)](https://dave.cheney.net/practical-go)
- [Creating Golang WebServer With Echo](https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY)
