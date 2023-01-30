# 消息相关接口

## 获取公告

#### 请求URL
- ` http://域名/api/v1/message/announce/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "announces": [
            {
                "id": 1,
                "title": "测试",
                "content": "123",
                "url": "123",
                "created_at": "2021-07-29T13:46:21Z",
                "important": false
        
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名    | 类型  | 说明     |
| :-------- | :---- | -------- |
| total     | int   | 总数量   |
| announces | array | 公告列表 |


##### 公告内容announces
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| id         | int    | 公告ID         |
| title      | string | 标题           |
| content    | string | 内容           |
| url        | string | 指向的链接     |
| created_at | time   | 发布时间       |
| important  | bool   | 是否为重要公告 |


#### 备注
无

## 添加公告

#### 请求URL
- ` http://域名/api/v1/message/announce/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token `

#### 参数

| 参数名    | 必选 | 类型   | 说明           |
| :-------- | :--- | :----- | -------------- |
| title     | 是   | string | 标题           |
| content   | 是   | string | 内容           |
| url       | 否   | string | 指向的地址     |
| important | 是   | bool   | 是否为重要公告 |

#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 
需要管理员及以上权限


## 管理员删除公告

#### 请求URL
- ` http://域名/api/v1/message/announce/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 否   | int  | 公告id |

#### 返回示例 
``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

##### 备注 
需要管理员及以上权限


## 获取点赞消息

#### 请求URL
- ` http://域名/api/v1/message/like/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "messages": [
            {
                "created_at": "2021-07-29T13:46:21Z",
                "user": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0
                },
                "video": {
                    "vid": 2,
                    "title": "测试",
                    "cover": "",
                    "desc": "什么都没有",
                    "clicks": 86,
                    "created_at": "2021-07-29T13:46:21Z"
                }
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| created_at | time   | 点赞时间     |
| video      | int    | 视频信息     |
| user       | string | 点赞用户信息 |


##### 视频信息video
| 参数名     | 类型   | 说明                                 |
| :--------- | :----- | ------------------------------------ |
| vid        | int    | 视频ID                               |
| title      | string | 标题                                 |
| cover      | string | 封面URL                              |
| desc       | string | 视频简介                             |
| created_at | time   | 发布时间                             |
| clicks     | int    | 点击量(点击量不是实时点击量，不准确) |


##### 点赞用户信息user
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |

#### 备注
无


## 获取@消息

#### 请求URL
- ` http://域名/api/v1/message/at/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "messages": [
            {
                "created_at": "2021-07-29T13:46:21Z",
                "user": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0
                },
                "video": {
                    "vid": 2,
                    "title": "测试",
                    "cover": "",
                    "desc": "什么都没有",
                    "clicks": 86,
                    "created_at": "2021-07-29T13:46:21Z"
                }
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| created_at | time   | 点赞时间     |
| video      | int    | 视频信息     |
| user       | string | 点赞用户信息 |


##### 视频信息video
| 参数名     | 类型   | 说明                                 |
| :--------- | :----- | ------------------------------------ |
| vid        | int    | 视频ID                               |
| title      | string | 标题                                 |
| cover      | string | 封面URL                              |
| desc       | string | 视频简介                             |
| created_at | time   | 发布时间                             |
| clicks     | int    | 点击量(点击量不是实时点击量，不准确) |


##### 点赞用户信息user
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |

#### 备注
无


## 获取回复消息

#### 请求URL
- ` http://域名/api/v1/message/at/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "messages": [
            {
                "created_at": "2021-07-29T13:46:21Z",
                "content": "",
                "target_reply_content": "",
                "root_content": "",
                "comment_id": "",
                "user": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0
                },
                "video": {
                    "vid": 2,
                    "title": "测试",
                    "cover": "",
                    "desc": "什么都没有",
                    "clicks": 86,
                    "created_at": "2021-07-29T13:46:21Z"
                }
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名               | 类型   | 说明                       |
| :------------------- | :----- | -------------------------- |
| created_at           | time   | 点赞时间                   |
| content              | string | 回复内容                   |
| target_reply_content | string | 被回复的那一条回复的内容   |
| root_content         | string | 回复所在的那一条评论的内容 |
| comment_id           | string | 评论ID                     |
| video                | int    | 视频信息                   |
| user                 | string | 点赞用户信息               |


##### 视频信息video
| 参数名     | 类型   | 说明                                 |
| :--------- | :----- | ------------------------------------ |
| vid        | int    | 视频ID                               |
| title      | string | 标题                                 |
| cover      | string | 封面URL                              |
| desc       | string | 视频简介                             |
| created_at | time   | 发布时间                             |
| clicks     | int    | 点击量(点击量不是实时点击量，不准确) |


##### 点赞用户信息user
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |

#### 备注
无


## 私信websocket

#### 请求URL
- ` ws://域名/api/v1/message/whisper/ws?token=这里写access_token `

#### 返回示例 
```json
{
    "fid": 1,//来自用户ID
    "content": "",
}
```

#### 备注
返回消息使用了base64编码


## 获取私信列表

#### 请求URL
- ` http://域名/api/v1/message/whisper/list `
  
#### 请求方式
- POST 

#### 请求头
- `Authorization': access_token `

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "messages": [
            {
                "created_at": "2021-07-29T13:46:21Z",
                "user": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0
                },
                "status": true,
            },
        ]
    },
    "msg": "ok"
}
```


#### 返回参数说明 
| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| created_at | time   | 点赞时间     |
| user       | string | 点赞用户信息 |
| status     | bool   | 是否已读     |

##### 用户信息user
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |


#### 备注
无


## 获取私信详细信息

#### 请求URL
- ` http://域名/api/v1/message/whisper/details?fid=目标用户ID&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

####  请求头
- `Authorization: access_token`

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "messages": [
            {
                "fid": 2,
                "from_id": 1,
                "content": "123",
                "created_at": "2021-07-29T10:54:42Z"
            },
            {
                "fid": 2,
                "from_id": 1,
                "content": "456",
                "created_at": "2021-07-29T10:55:31Z"
            }
        ],
    },
    "msg": "ok"
}

```

##### 返回参数说明 
| 参数名     | 类型   | 说明       |
| :--------- | :----- | ---------- |
| fid        | int    | 关联用户ID |
| from_id    | int    | 发送者id   |
| content    | string | 消息内容   |
| created_at | time   | 发送时间   |

#### 备注
无


## 发送私信

#### 请求URL
- ` http://域名/api/v1/message/whisper/send `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名  | 必选 | 类型   | 说明       |
| :------ | :--- | :----- | ---------- |
| fid     | 是   | int    | 目标用户id |
| content | 是   | string | 内容       |

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


## 私信已读

#### 请求URL
- ` http://域名/api/v1/message/whisper/read `
  
#### 请求方式
- POST 

####  请求头
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明       |
| :----- | :--- | :--- | ---------- |
| id     | 是   | int  | 目标用户id |

#### 返回示例 

``` json
{
    "code": 200,
    "data": null,
    "msg": "ok"
}
```

#### 备注
无