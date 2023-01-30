FROM golang:1.19.3-alpine
WORKDIR /server
COPY . .

# #构建后端和安装环境
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o app main.go 

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add ffmpeg  

EXPOSE 9000

CMD ./app
