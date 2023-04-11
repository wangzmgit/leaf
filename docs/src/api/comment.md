# 评论回复相关接口

## 获取评论回复

#### 请求URL
- ` http://域名/api/v1/comment/get?vid=视频ID&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "comments": [
            {
                "id": 1,
                "content": "测试",
                "author": {
                    "uid": 1,
                    "name": "",
                    "sign": "",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 1
                },
                "reply": [
                    {
                        "id": 3,
                        "content": "1",
                        "author": {
                            "uid": 1,
                            "name": "",
                            "sign": "",
                            "avatar": "",
                            "spacecover": "",
                            "gender": 1
                        },
                        "at": [],
                        "created_at": "2022-06-22T21:08:25Z"
                    }
                ],
                "at": [],
                "created_at": "2022-06-20T13:42:40.625Z"
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型       | 说明         |
| :--------- | :--------- | ------------ |
| id         | int        | 用户ID       |
| content    | string     | 评论内容     |
| author     | object     | 评论作者信息 |
| created_at | time       | 发布时间     |
| at         | array[int] | @用户的ID    |


##### 回复reply
| 参数名     | 类型       | 说明         |
| :--------- | :--------- | ------------ |
| id         | int        | 用户ID       |
| content    | string     | 评论内容     |
| author     | object     | 评论作者信息 |
| reply      | array      | 回复(0-2条)  |
| created_at | time       | 发布时间     |
| at         | array[int] | @用户的ID    |


##### 作者信息author
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


## 获取回复详情

##### 请求URL
- ` http://域名/api/v1/comment/reply/get?cid=评论ID&page=页码&page_size=内容数量 `
  
##### 请求方式
- GET 

##### 请求参数
- 见url

##### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "replies": [
            {
                "id": 4,
                "content": "1",
                "author": {
                    "uid": 1,
                    "name": "",
                    "sign": "",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 1
                },
                "at": [],
                "created_at": "2022-06-22T21:08:25Z"
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型       | 说明         |
| :--------- | :--------- | ------------ |
| id         | int        | 用户ID       |
| content    | string     | 评论内容     |
| author     | object     | 评论作者信息 |
| reply      | array      | 回复(0-2条)  |
| created_at | time       | 发布时间     |
| at         | array[int] | @用户的ID    |


##### 作者信息author
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

## 发布评论

#### 请求URL
- ` http://域名/api/v1/comment/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名  | 必选 | 类型          | 说明          |
| :------ | :--- | :------------ | ------------- |
| vid     | 是   | int           | 视频ID        |
| content | 是   | string        | 评论内容      |
| at      | 是   | array[string] | @的用户名数组 |


#### 返回示例 
``` json
  {
    "code": 200,
    "data": {
        "id": "63a11c56e48a9fe01ff95937",
    },
    "msg":"ok"
  }
```

#### 返回参数说明 
| 参数名 | 类型   | 说明   |
| :----- | :----- | ------ |
| id     | string | 评论ID |


#### 备注
无


## 发布回复

#### 请求URL
- ` http://域名/api/v1/comment/reply/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名       | 必选 | 类型          | 说明                           |
| :----------- | :--- | :------------ | ------------------------------ |
| vid          | 是   | int           | 视频ID                         |
| content      | 是   | string        | 评论内容                       |
| parentID     | 是   | string        | 评论ID                         |
| at           | 是   | array[string] | @的用户名数组                  |
| replyUserID  | 是   | int           | 回复用户的ID                   |
| replyContent | 是   | string        | 回复内容(不包含 ‘回复 @xxx: ’) |


#### 返回示例 
``` json
  {
    "code": 200,
    "data": {
        "id": "63a11c56e48a9fe01ff95937",
    },
    "msg":"ok"
  }
```

#### 返回参数说明 
| 参数名 | 类型   | 说明   |
| :----- | :----- | ------ |
| id     | string | 回复ID |


#### 备注
无


## 删除评论

#### 请求URL
- ` http://域名/api/v1/comment/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 评论ID |

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


## 删除回复

#### 请求URL
- ` http://域名/api/v1/comment/reply/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名    | 必选 | 类型 | 说明   |
| :-------- | :--- | :--- | ------ |
| commentID | 是   | int  | 评论ID |
| replyID   | 是   | int  | 回复ID |

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