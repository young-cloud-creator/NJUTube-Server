# Goto2023

字节跳动第五届青训营大项目

## 代码结构介绍

代码的组织结构如下，基本按照[github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)中描述的架构进行设计

![image](https://user-images.githubusercontent.com/84324349/217787742-6d8274a4-a8af-44d9-aad8-8ab968323247.png)

```
Goto2023
│
├── controller     // 直接与API使用者交互 负责组织和发送Response、验证token等
│   ├── comment.go    // 评论相关API
│   ├── common.go     // 一些基本的数据结构 如User、Video、Comment等
│   ├── favorite.go   // 点赞相关API
│   ├── feed.go       // Feed流相关API
│   ├── publish.go    // 发布相关API 包括发布视频、已发布视频列表等API
│   └── user.go       // 用户相关API 如注册、登陆、用户信息等
│
├── service        // 被controller使用 负责具体业务逻辑实现
│   ├── comment.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   └── user.go
│
├── repository     // 被service使用 负责底层数据库操作
│   └── db_setup.go   // 负责连接到数据库
│
├── main.go        // 入口 负责调用相关方法初始化数据库、初始化路由和运行服务器
├── router.go      // 路由初始化 负责设置静态文件系统和API路由
│
├── go.mod
└── go.sum
```
