# Homework of service computing -- simple cloudgo 


## 参考资料
潘老师博客：https://blog.csdn.net/pmlpml/article/details/78539261

## 运行环境以及命令

go version go1.11 windows/amd64

进入 main.go 文件所在的目录
```shell
go run main.go
```
默认会在 8080 端口搭建简单的服务器，监听请求, 具体的测试见下文。

## 目录结构

│  main.go
│  readme.md
│  testweb.exe
│
├─assets
│  ├─css
│  │      main.css
│  │
│  ├─images
│  └─js
│          hello.js
│
├─service
│      apitest.go
│      home.go
│      login.go
│      phone_sale.go
│      server.go
│
└─templates
        index.html
        moblie_sales.html

## 静态文件服务

利用 Http包提供的 Hanlder 简单的静态文件服务。指定静态文件存放的目录之后就可以使用，本次实验指定的目录时`assets`.


## 提供简单的api

返回简单的json，方便用户使用 `js` 脚本和服务器交互。对应的路径是 `/api/test`

对应代码：[apitest.go](./servie/apitest.go)

## 使用模板 

使用模板来生成对应的 html 文件,其中模板文件存放于 `./tempaltes`。本次实验在两个部分使用了模板，第一个部分是 `index.html`，另一个部分是在练习返回手机销售信息的部分。模板引擎使用的封装了`golang/text/template`的`render`包。

对应代码：[phone_sale.go](./service/phone_sale.go)

## 处理表单数据

### 测试方法： 

在 `localhost:8080/index` 中输入相关的用户名和密码，把在后端重新生成页面传回。把登陆用户的信息显示在页面上, 初始化时 Null, 没有密码检验。（出于测试简单，实际应用中不会出现这种用法。）

对应代码：[login.go](./service/login.go)

## optinal assignment: custom middleware

negroni 实际上就是对 HTTP/Net 的一个封装，简化利用 http包搭建服务器的流程，提供一些更易于使用的接口。同时让中间件的构造和应用更加简单。

>Negroni is an idiomatic approach to web middleware in Go. It is tiny, non-intrusive, and encourages use of net/http Handlers.

我们可以通过下面的方式来构造。


## optional assignment: mux 类解析 // todo

这个类是用来分析方便路由请求的。

1. 关于路由中的路径匹配规则

    - 优先级受什么影响
        - 是否和声明的顺序相关
        
        - 共同前缀的路由如何处理（比如 /api, /api/abc，应该响应哪一个的handler）。

### Router 


### route
