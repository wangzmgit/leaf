# 弹幕接口

    
## 发送弹幕

#### 请求URL
- ` http://域名/api/v1/danmaku/send `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token `


#### 参数

| 参数名 | 必选 | 类型   | 说明                   |
| :----- | :--- | :----- | ---------------------- |
| vid    | 是   | int    | 视频id                 |
| part   | 否   | int    | 分P (默认1)            |
| time   | 是   | int    | 时间                   |
| type   | 是   | int    | 类型 0滚动 1顶部 2底部 |
| color  | 是   | string | 颜色                   |
| text   | 是   | string | 内容                   |

#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 
无

    
## 获取弹幕

#### 请求URL
- ` http://域名/api/v1/danmaku/list?vid=视频ID?part=分P `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "danmaku": [
            {
                "time": 10,
                "type": 0,
                "color": "#fff",
                "text": "私信测试"
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名 | 类型   | 说明     |
| :----- | :----- | -------- |
| time   | int    | 进度     |
| type   | int    | 弹幕类型 |
| color  | string | 弹幕颜色 |
| text   | string | 弹幕内容 |

#### 备注 
无
