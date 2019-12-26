# go-exercise
关于go的练习

```shell script
export http_proxy=socks5://127.0.0.1:1080
```

```shell script
go mod tidy
go mod download
go mod vendor
```

## leetcode

### dfs

1. [找出一棵树的从根节点到叶子的所有路径](leetcode/dfs/btreepaths/main.go)

2. [机器人走路的不同路径](leetcode/dfs/uniquepath1/main.go)

### bfs

1. [二叉树的层序遍历](leetcode/bfs/lot/main.go)

## web 框架

> 好像这个就是 go 这边用的比较多的 web 框架，用起来感觉跟 Python 的 Flask 框架差不多。

## 参考文章

- [go-get命令使用socket代理](http://www.hi-roy.com/2018/10/12/go-get%E5%91%BD%E4%BB%A4%E4%BD%BF%E7%94%A8socket%E4%BB%A3%E7%90%86/)
- [Gin入门实战](https://www.imooc.com/learn/1175)
- [gin文档](https://gin-gonic.com/zh-cn/docs/)