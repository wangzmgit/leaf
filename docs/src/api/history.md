# 历史记录相关接口

## 记录历史记录

#### 请求URL
- ` http://域名/api/v1/history/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`


#### 参数

| 参数名 | 必选 | 类型  | 说明        |
| :----- | :--- | :---- | ----------- |
| vid    | 是   | int   | 视频id      |
| part   | 否   | int   | 分P (默认1) |
| time   | 是   | float | 时间        |

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


## 获取历史记录

#### 请求URL
- ` http://域名/api/v1/history/video/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "history": [
            {
                "time": 10.6,
                "part": 1,
                "video": {
                    "vid": 2,
                    "title": "测试",
                    "cover": "",
                    "desc": "什么都没有",
                    "clicks": 86,
                    "created_at": "2021-07-29T13:46:21Z"
                },
                "created_at": "2022-06-06T08:42:13.525Z"
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型  | 说明        |
| :--------- | :---- | ----------- |
| int        | int   | 历史记录id  |
| time       | float | 进度        |
| part       | int   | 分P (默认1) |
| created_at | time  | 创建时间    |
| video      | int   | 视频信息    |


##### 视频信息video
| 参数名     | 类型   | 说明                                 |
| :--------- | :----- | ------------------------------------ |
| vid        | int    | 视频ID                               |
| title      | string | 标题                                 |
| cover      | string | 封面URL                              |
| desc       | string | 视频简介                             |
| created_at | time   | 发布时间                             |
| clicks     | int    | 点击量(点击量不是实时点击量，不准确) |

#### 备注 
无


## 获取播放进度

#### 请求URL
- ` http://域名/api/v1/history/progress/get?vid=视频ID `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "progress": {
            "time": 10.6,
            "part": 1
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型  | 说明        |
| :--------- | :---- | ----------- |
| time       | float | 进度        |
| part       | int   | 分P (默认1) |

#### 备注 
无
