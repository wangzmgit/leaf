# 手动部署

:::tip
运行项目需要预先安装`golang(>1.19)`,`mysql`,`redis`,`mongodb`,`nginx`,`nodejs`，
手动部署仅部署后端，前端部分参考 [域名配置](/guide/domain#未部署前端)
:::

### 配置后端文件
将`data/config/application.yml`文件复制到`server/config/`中，打开 `application.yml`文件，根据其中的注释修改配置文件并保存

### 启动后端
设置go代理
```sh
go env -w GOPROXY=https://goproxy.cn,direct
```
编译
```sh
 go build -o 给程序取个名字 main.go
```
后台运行
```sh
nohup ./给程序取的名字 >logs/nohup.log 2>&1
```

### 开启端口
如果使用云服务并且使用IP访问的话，需要在防火墙配置打开配置文件`port`所设置的端口
