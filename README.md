> 启发于[这篇博客](https://shusunny.github.io/sunnyblog/blockchain/simple-blockchain.html)

# 介绍

这个项目是一个图书管理系统，使用（简易） 区块链作为存储形式存储所有书籍数据，以保证数据安全性。

# 使用方法

现在代码到本地，使用 `go get` 安装依赖后，运行 `main.go`，程序成功运行在 3000 端口后，使用 `curl` 命令模拟出 http 的 POST 请求。  

具体使用方法：

```bash
curl http://localhost:3000/newBook \
  -H "Content-Type: application/json" \
  -d '{"title": "All about Lily Chou-chou", "author":"Shunji Iwai", 
"isbn":"9787544244398","publish_date":"2009-07-01"}'
```

这将会返回一个书目 ID，有了书目ID 之后，再发送借阅记录

```bash
curl http://localhost:3000/borrowBook \
  -H "Content-Type: application/json" \
  -d '{"book_id": "BOOK_ID", "user": "USER", 
"checkout_date":"2024-08-29"}'
```

就可以在本地浏览器 3000 端口看见借阅记录已经被录入区块当中了。

# TaskList

1. [x] 后端实现
2. [ ] 前端实现
3. [ ] Mysql 数据库实现
4. [ ] 测试
    1. [ ] 数据合法性检测
    2. [ ] 书库查重
5. [ ] 部署上线
    1. [ ] Docker
