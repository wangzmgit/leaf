import{_ as s,c as n,o as a,a as l}from"./app.6fababfc.js";const D=JSON.parse('{"title":"域名配置","description":"","frontmatter":{},"headers":[{"level":2,"title":"未使用docker部署前端","slug":"未使用docker部署前端","link":"#未使用docker部署前端","children":[]},{"level":2,"title":"未使用docker部署前端(v1.2.0之前)","slug":"未使用docker部署前端-v1-2-0之前","link":"#未使用docker部署前端-v1-2-0之前","children":[]},{"level":2,"title":"docker部署的前端","slug":"docker部署的前端","link":"#docker部署的前端","children":[]}],"relativePath":"guide/domain.md"}'),e={name:"guide/domain.md"},p=l(`<h1 id="域名配置" tabindex="-1">域名配置 <a class="header-anchor" href="#域名配置" aria-hidden="true">#</a></h1><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>在域名配置之前，需要安装好nginx</p></div><h2 id="未使用docker部署前端" tabindex="-1">未使用docker部署前端 <a class="header-anchor" href="#未使用docker部署前端" aria-hidden="true">#</a></h2><ol><li>从github中下载最新发行版中的<code>web.zip</code>文件</li><li>将文件解压到<code>/usr/share/nginx/html</code> 目录内，解压完成后确认<code>/usr/share/nginx/html</code> 目录内有<code>web</code>、<code>manage</code>、<code>mobile</code>三个文件夹</li><li>修改<code>data/upload/config.web.json</code>中的配置，可以修改网站名，上传文件大小等信息</li><li>按下方提供的nginx配置文件配置nginx</li><li>重启nginx</li></ol><p>nginx配置文件如下：</p><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-palenight"><code><span class="line"><span style="color:#A6ACCD;">server {</span></span>
<span class="line"><span style="color:#A6ACCD;">    listen       80;</span></span>
<span class="line"><span style="color:#A6ACCD;">	server_name  localhost; #有域名可以把localhost替换为域名</span></span>
<span class="line"><span style="color:#A6ACCD;">	client_max_body_size 1024M;</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    location / {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html/web;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @web;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @web {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    #后台管理</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /manage/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @manage;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决后台管理history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @manage {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /manage/index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    #移动端</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /mobile/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @mobile;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决移动端history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @mobile {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /mobile/index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 转发后端</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /api/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_pass http://127.0.0.1:9000;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_set_header   Host             $host;</span></span>
<span class="line"><span style="color:#A6ACCD;">     	proxy_set_header   X-Real-IP        $remote_addr;						</span></span>
<span class="line"><span style="color:#A6ACCD;">        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_http_version 1.1;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Upgrade $http_upgrade;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Connection &quot;upgrade&quot;;</span></span>
<span class="line"><span style="color:#A6ACCD;">	}</span></span>
<span class="line"><span style="color:#A6ACCD;">}</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span></code></pre></div><h2 id="未使用docker部署前端-v1-2-0之前" tabindex="-1">未使用docker部署前端(v1.2.0之前) <a class="header-anchor" href="#未使用docker部署前端-v1-2-0之前" aria-hidden="true">#</a></h2><p>没有使用docker部署前端，需要用nginx部署前端，首先要打包前端</p><ol><li>打开 <code>web\\utils\\src\\global-config.ts</code>文件，修改<code>domain</code>为<code>&quot;&quot;</code>(空内容)</li><li>在构建之前需要先对前端项目进行打包，首先要先安装nodejs和pnpm，然后进入<code>web</code>目录，执行<code>pnpm i</code></li><li>进入<code>web\\packages\\web-client</code>执行命令<code>pnpm run build</code> 将产生的<code>web</code>文件夹复制到<code>/usr/share/nginx/html</code></li><li>进入<code>web\\packages\\manage-client</code>执行命令<code>pnpm run build</code> 将产生的<code>manage</code>文件夹复制到<code>/usr/share/nginx/html</code></li><li>进入<code>web\\packages\\mobile-client</code>执行命令<code>pnpm run build</code> 将产生的<code>mobile</code>文件夹复制到<code>/usr/share/nginx/html</code></li><li>按下方提供的nginx配置文件配置nginx</li><li>重启nginx</li></ol><p>nginx配置文件如下：</p><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-palenight"><code><span class="line"><span style="color:#A6ACCD;">server {</span></span>
<span class="line"><span style="color:#A6ACCD;">    listen       80;</span></span>
<span class="line"><span style="color:#A6ACCD;">	server_name  localhost; #有域名可以把localhost替换为域名</span></span>
<span class="line"><span style="color:#A6ACCD;">	client_max_body_size 1024M;</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    location / {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html/web;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @web;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @web {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    #后台管理</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /manage/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @manage;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决后台管理history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @manage {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /manage/index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    #移动端</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /mobile/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">        root /usr/share/nginx/html;</span></span>
<span class="line"><span style="color:#A6ACCD;">        index index.html index.htm;</span></span>
<span class="line"><span style="color:#A6ACCD;">        try_files $uri $uri/ @mobile;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 解决移动端history路由问题</span></span>
<span class="line"><span style="color:#A6ACCD;">    location @mobile {</span></span>
<span class="line"><span style="color:#A6ACCD;">        rewrite ^.*$ /mobile/index.html;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    # 转发后端</span></span>
<span class="line"><span style="color:#A6ACCD;">    location /api/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_pass http://127.0.0.1:9000;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_set_header   Host             $host;</span></span>
<span class="line"><span style="color:#A6ACCD;">     	proxy_set_header   X-Real-IP        $remote_addr;						</span></span>
<span class="line"><span style="color:#A6ACCD;">        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_http_version 1.1;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Upgrade $http_upgrade;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Connection &quot;upgrade&quot;;</span></span>
<span class="line"><span style="color:#A6ACCD;">	}</span></span>
<span class="line"><span style="color:#A6ACCD;">}</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span></code></pre></div><h2 id="docker部署的前端" tabindex="-1">docker部署的前端 <a class="header-anchor" href="#docker部署的前端" aria-hidden="true">#</a></h2><p>如果使用docker部署了前端，也就相当于现在有两个项目：</p><ol><li>前端项目，运行在9090端口</li><li>后端项目，运行在9000端口</li></ol><p>配置域名需要将这两个端口转发出去，nginx配置文件如下：</p><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-palenight"><code><span class="line"><span style="color:#A6ACCD;">server {</span></span>
<span class="line"><span style="color:#A6ACCD;">    listen       80;</span></span>
<span class="line"><span style="color:#A6ACCD;">	server_name  localhost; #有域名可以把localhost替换为域名</span></span>
<span class="line"><span style="color:#A6ACCD;">	client_max_body_size 1024M;</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">    location / {</span></span>
<span class="line"><span style="color:#A6ACCD;">        proxy_pass http://127.0.0.1:9090;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_set_header   Host             $host;</span></span>
<span class="line"><span style="color:#A6ACCD;">     	proxy_set_header   X-Real-IP        $remote_addr;						</span></span>
<span class="line"><span style="color:#A6ACCD;">        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;</span></span>
<span class="line"><span style="color:#A6ACCD;">    }</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">	location /api/ {</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_pass http://127.0.0.1:9000;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_set_header   Host             $host;</span></span>
<span class="line"><span style="color:#A6ACCD;">     	proxy_set_header   X-Real-IP        $remote_addr;						</span></span>
<span class="line"><span style="color:#A6ACCD;">        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;</span></span>
<span class="line"><span style="color:#A6ACCD;">		proxy_http_version 1.1;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Upgrade $http_upgrade;</span></span>
<span class="line"><span style="color:#A6ACCD;">    	proxy_set_header Connection &quot;upgrade&quot;;</span></span>
<span class="line"><span style="color:#A6ACCD;">	}</span></span>
<span class="line"><span style="color:#A6ACCD;">}</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span></code></pre></div>`,16),o=[p];function c(t,r,i,C,A,d){return a(),n("div",null,o)}const h=s(e,[["render",c]]);export{D as __pageData,h as default};
