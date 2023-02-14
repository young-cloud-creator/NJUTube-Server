# Goto2023

字节跳动第五届青训营大项目

## 注意事项

- 项目依赖于OpenCV 4.7.0，请自行安装，安装方法参考https://gocv.io/getting-started/
- `service/publish.go`文件定义了服务器地址，运行时请自行修改`serverAddr`变量为服务器地址
  - 该变量将影响返回给客户端的视频和封面地址
- `repository/db_setup.go`文件的InitDB函数中的`dsn`定义了MySQL数据库地址，运行时请自行修改
  - 数据库中需要存在名为`douyin`的database，其中的表结构详见`数据表结构`部分
- `service/publish.go`文件定义了视频和封面文件存放的子目录，如有需要，可自行修改`PublicDir`、`VideoDir`和`CoverDir`变量
- 数据库表结构已导出到项目目录下的`douyin.sql`文件，如有需要请自行导入

## 项目选题

极简抖音服务端，初步目标是实现互动方向（即视频Feed流、视频投稿、用户登陆注册、用户信息查询、点赞和评论功能）的API

## 代码结构介绍

代码的组织结构如下，基本按照[github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)中描述的架构进行设计

![image](https://user-images.githubusercontent.com/84324349/217787742-6d8274a4-a8af-44d9-aad8-8ab968323247.png)

```
Goto2023
│
├── structs        // 项目常用的结构体 ✅
│   └── common.go     //  一些基本的结构体 如User、Video、Comment等 ✅
│
├── controller     // 直接与API使用者交互 负责组织和发送Response、验证token等
│   ├── comment.go    // 评论相关API
│   ├── favorite.go   // 点赞相关API
│   ├── feed.go       // Feed流相关API ✅
│   ├── publish.go    // 发布相关API 包括发布视频、已发布视频列表等API ✅
│   └── user.go       // 用户相关API 如注册、登陆、用户信息等 ✅
│
├── service        // 被controller使用 负责具体业务逻辑实现
│   ├── comment.go
│   ├── favorite.go
│   ├── feed.go       // feed API相关的具体业务逻辑 包括时间戳转换、视频列表数据处理等 ✅
│   ├── publish.go    // publish API相关的具体业务逻辑 包括投稿视频封面截取、投稿视频列表数据处理等 ✅
│   └── user.go       // user API相关的具体业务逻辑 包括密码加密、用户名密码验证等 ✅
│
├── repository     // 被service使用 负责底层数据库操作
│   ├── common.go     // gorm model 用于表示数据库中的各张表 ✅
│   ├── db_setup.go   // 负责连接到数据库 ✅
│   ├── video.go      // 视频信息存储与查询 ✅
│   └── user.go       // user表相关数据库操作 ✅
│
├── security       // 安全相关的函数 例如token生成和验证 ✅
│   └── jwt.go        // JSON Web Token（JWT）生成和验证 ✅
│
├── main.go        // 入口 负责调用相关方法初始化数据库、初始化路由和运行服务器 ✅
├── router.go      // 路由初始化 负责设置静态文件系统和API路由 ✅
│
├── go.mod
└── go.sum

```

## 数据表结构介绍

数据库的配置位于repository/db_setup.go，运行项目后，将会连接到root:12345678@127.0.0.1:3306的名为douyin的MySQL数据库。数据库有user、video、favorite和comment四张表，分别存储用户信息、视频信息、点赞信息和评论信息，下面是这几张表的结构。

```
user

+--------+--------------+------+-----+---------+----------------+
| Field  | Type         | Null | Key | Default | Extra          |
+--------+--------------+------+-----+---------+----------------+
| id     | bigint       | NO   | PRI | NULL    | auto_increment |
| name   | varchar(128) | YES  |     | NULL    |                |
| passwd | varchar(512) | YES  |     | NULL    |                |
+--------+--------------+------+-----+---------+----------------+

注：为了数据安全考虑，此处的passwd字段存储的是经过加密的密码；id为主键，设置了自增属性，无需额外考虑ID的唯一性
```

```
video

+-------------+---------------+------+-----+---------+----------------+
| Field       | Type          | Null | Key | Default | Extra          |
+-------------+---------------+------+-----+---------+----------------+
| id          | bigint        | NO   | PRI | NULL    | auto_increment |
| author      | bigint        | YES  |     | NULL    |                |
| title       | varchar(1024) | YES  |     | NULL    |                |
| play_url    | varchar(1024) | YES  |     | NULL    |                |
| cover_url   | varchar(1024) | YES  |     | NULL    |                |
| create_time | datetime      | YES  |     | NULL    |                |
+-------------+---------------+------+-----+---------+----------------+

注：id为主键，设置了自增属性，无需额外考虑ID的唯一性
```

```
favorite

+----------+--------+------+-----+---------+-------+
| Field    | Type   | Null | Key | Default | Extra |
+----------+--------+------+-----+---------+-------+
| user_id  | bigint | YES  |     | NULL    |       |
| video_id | bigint | YES  |     | NULL    |       |
+----------+--------+------+-----+---------+-------+
```

```
comment

+-------------+---------------+------+-----+---------+----------------+
| Field       | Type          | Null | Key | Default | Extra          |
+-------------+---------------+------+-----+---------+----------------+
| id          | bigint        | NO   | PRI | NULL    | auto_increment |
| user_id     | bigint        | YES  |     | NULL    |                |
| video_id    | bigint        | YES  |     | NULL    |                |
| content     | varchar(2048) | YES  |     | NULL    |                |
| create_time | datetime      | YES  |     | NULL    |                |
+-------------+---------------+------+-----+---------+----------------+

注：id为主键，设置了自增属性，无需额外考虑ID的唯一性
```
