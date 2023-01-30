# Docker部署


:::tip 
在使用docker安装之前，需要安装好docker,docker-compose
:::

## 后端部署

### 1.配置后端文件
打开 `data/config/application.yml`文件，根据其中的注释修改配置文件并保存

### 2.构建
执行以下命令
```
docker-compose build
```

### 3.启动后端
执行以下命令
```
docker-compose up -d
```

如需停止项目执行
```
docker-compose stop 
```

## 前端部署

前端提供了docker部署，但不是必要的，如果需要配置域名建议直接跳转到[域名配置](/guide/domain#未部署前端)

### 1.配置后端文件
打开 `web\utils\src\global-config.ts`文件，修改`domain`为后端域名

**注意：仅域名不包含http、斜线、反斜线等内容，没有域名可以使用可以为ip:端口**

### 2.构建

docker构建提供了两种构建方式，docker打包+部署内包含了前端项目的构建，但是对服务器性能要求较高，
docker仅部署对服务器要求相对低，但需要手动打包前端

#### 2.1 使用docker打包+部署

进入web目录执行以下命令
```
docker build -f Dockerfile-build -t "leaf-web" .
```

#### 2.2 使用docker仅部署
1. 在构建之前需要先对前端项目进行打包，首先要先安装nodejs和pnpm，然后进入`web`目录，执行
```
pnpm i
```

2. 在`web`目录中创建`dist`文件夹
3. 进入`web\packages\web-client`执行命令`pnpm run build` 将产生的`web`文件夹复制到`dist`
4. 进入`web\packages\manage-client`执行命令`pnpm run build` 将产生的`manage`文件夹复制到`dist`
5. 进入`web\packages\mobile-client`执行命令`pnpm run build` 将产生的`mobile`文件夹复制到`dist`
6. 在`web`目录下执行以下命令
```
docker build -t "leaf-web" .
```

### 3.启动前端
```
docker run -itd --name leafWeb -p 9090:9090 "leaf-web"
```

如需停止项目执行
```
docker stop leafWeb
```


