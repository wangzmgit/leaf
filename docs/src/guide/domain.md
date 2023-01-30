# 域名配置

:::tip
在域名配置之前，需要安装好nginx
:::

## docker部署前端
如果使用docker部署了前端，也就相当于现在有两个项目：
1. 前端项目，运行在9090端口
2. 后端项目，运行在9000端口

配置域名需要将这两个端口转发出去，nginx配置文件如下：
```
server {
    listen       80;
	server_name  localhost; #有域名可以把localhost替换为域名
	client_max_body_size 1024M;

    location / {
        proxy_pass http://127.0.0.1:9090;
		proxy_set_header   Host             $host;
     	proxy_set_header   X-Real-IP        $remote_addr;						
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
    }

	location /api/ {
		proxy_pass http://127.0.0.1:9000;
		proxy_set_header   Host             $host;
     	proxy_set_header   X-Real-IP        $remote_addr;						
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
		proxy_http_version 1.1;
    	proxy_set_header Upgrade $http_upgrade;
    	proxy_set_header Connection "upgrade";
	}
}
```

## 未部署前端
没有使用docker部署前端，需要用nginx部署前端，首先要打包前端
1. 打开 `web\utils\src\global-config.ts`文件，修改`domain`为`""`(空内容)
2. 在构建之前需要先对前端项目进行打包，首先要先安装nodejs和pnpm，然后进入`web`目录，执行`pnpm i`
3. 进入`web\packages\web-client`执行命令`pnpm run build` 将产生的`web`文件夹复制到`/usr/share/nginx/html`
4. 进入`web\packages\manage-client`执行命令`pnpm run build` 将产生的`manage`文件夹复制到`/usr/share/nginx/html`
5. 进入`web\packages\mobile-client`执行命令`pnpm run build` 将产生的`mobile`文件夹复制到`/usr/share/nginx/html`

nginx配置文件如下：
```
server {
    listen       80;
	server_name  localhost; #有域名可以把localhost替换为域名
	client_max_body_size 1024M;

    location / {
        root /usr/share/nginx/html/web;
        index index.html index.htm;
        try_files $uri $uri/ @web;
    }

    # 解决history路由问题
    location @web {
        rewrite ^.*$ /index.html;
    }

    #后台管理
    location /manage/ {
        root /usr/share/nginx/html;
        index index.html index.htm;
        try_files $uri $uri/ @manage;
    }

    # 解决后台管理history路由问题
    location @manage {
        rewrite ^.*$ /manage/index.html;
    }

    #移动端
    location /mobile/ {
        root /usr/share/nginx/html;
        index index.html index.htm;
        try_files $uri $uri/ @mobile;
    }

    # 解决移动端history路由问题
    location @mobile {
        rewrite ^.*$ /mobile/index.html;
    }

    # 转发后端
    location /api/ {
		proxy_pass http://127.0.0.1:9000;
		proxy_set_header   Host             $host;
     	proxy_set_header   X-Real-IP        $remote_addr;						
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
		proxy_http_version 1.1;
    	proxy_set_header Upgrade $http_upgrade;
    	proxy_set_header Connection "upgrade";
	}
}
```