# Leaf弹幕视频网站

## 项目介绍
Leaf弹幕视频网站是一个基于Go+Vue的前后端分离项目。Web 端使用 Vue + NaiveUI , 后端使用 Golang + Gin + Gorm进行开发。

### 特点
- 实现了文件存储在服务器本地或阿里云、腾讯云、七牛云对象存储功能
- 实现了视频自动处理分辨率并将视频切片(MPEG-DASH)功能


## 项目文档

`开发交流群: 909847398`

`文档：` http://leaf.kuukaa.fun/

## 相关项目

- Vue-WPlayer弹幕播放器(https://github.com/wangzmgit/vue-wplayer)
- GO滑块拼图生成(https://github.com/wangzmgit/jigsaw)
- GO对象存储(https://github.com/wangzmgit/unioss)

## 项目结构
```
├── data docker挂载文件
├── docker-compose.yml
├── server 后端
│   ├── api
│   ├── cache 缓存相关
│   ├── common  一些常量
│   ├── config 后端配置文件
│   │   └── application.yml
│   ├── db 数据库相关
│   ├── Dockerfile
│   ├── domain
│   │   ├── dto
│   │   ├── model
│   │   ├── resp 返回内容
│   │   ├── valid 数据校验
│   │   └── vo
│   ├── go.mod
│   ├── go.sum
│   ├── initialize 初始化
│   ├── logger
│   │   └── logger.go
│   ├── logs 存储日志文件
│   ├── main.go
│   ├── middleware 中间件
│   │   ├── auth.go 权限中间件
│   │   └── cors.go 跨域中间件
│   ├── routes 路由
│   ├── service
│   ├── static 静态文件
│   │   ├── casbin 权限相关
│   │   └── jigsaw 滑块拼图相关
│   ├── test 测试文件
│   ├── upload 用户上传文件
│   │   ├── image
│   │   └── video
│   ├── util
│   └── ws websocket相关
└── web 前端
    ├── apis api相关内容
    ├── components 公共组件
    ├── icons 图标
    ├── package.json
    ├── packages
    │   ├── manage-client 后台管理端项目
    │   ├── mobile-client 移动端项目
    │   └── web-client web端项目
    ├── pnpm-lock.yaml
    ├── pnpm-workspace.yaml
    └── utils 

```

## 截图

|                              Web端                               |                                                             |
| :--------------------------------------------------------------: | :---------------------------------------------------------: |
|       ![Web端登录](https://leaf.kuukaa.fun/web_login.png)        |     ![Web端首页](https://leaf.kuukaa.fun/web_home.png)      |
|       ![Web端上传](https://leaf.kuukaa.fun/web_upload.png)       |     ![Web端视频](https://leaf.kuukaa.fun/web_video.png)     |
|     ![Web端个人中心](https://leaf.kuukaa.fun/web_space.png)      |    ![Web端消息](https://leaf.kuukaa.fun/web_message.png)    |
|                            后台管理端                            |                                                             |
|     ![管理端登录](https://leaf.kuukaa.fun/manage_login.png)      | ![管理端视频管理](https://leaf.kuukaa.fun/manage_video.png) |
| ![管理端轮播图管理](https://leaf.kuukaa.fun/manage_carousel.png) |                                                             |
|                              移动端                              |                                                             |
|     ![移动端登录](https://leaf.kuukaa.fun/mobile_login.jpg)      |   ![移动端视频](https://leaf.kuukaa.fun/mobile_video.jpg)   |
|    ![移动端评论](https://leaf.kuukaa.fun/mobile_comment.jpg)     |                                                             |




