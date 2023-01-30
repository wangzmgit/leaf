# 收藏夹相关接口

## 获取收藏夹列表

#### 请求URL
- ` http://域名/api/v1/collection/list `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
      "collections": [
        {
          "id": 2,
          "name": "测试1",
          "cover": "",
          "desc": "",
          "open": true,
          "created_at": "2021-07-16T08:49:54Z",
        }
      ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| id         | int    | 收藏夹id     |
| name       | string | 收藏夹名称   |
| cover      | string | 收藏夹封面图 |
| desc       | string | 简介         |
| open       | bool   | 是否公开     |
| created_at | string | 创建时间     |

#### 备注
无


## 获取收藏夹信息

#### 请求URL
- ` http://域名/api/v1/collection/info?id=收藏夹id `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "collection": {
            "id": 2,
            "name": "测试1",
            "cover": "",
            "desc": "",
            "open": true,
            "created_at": "2021-07-16T08:49:54Z",
        },
        "author": {
            "uid": 1,
            "name": "user_1654250698886",
            "sign": "这个人很懒，什么都没有留下",
            "avatar": "",
            "spacecover": "",
            "gender": 0,
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| id         | int    | 收藏夹id     |
| name       | string | 收藏夹名称   |
| cover      | string | 收藏夹封面图 |
| desc       | string | 简介         |
| open       | bool   | 是否公开     |
| created_at | string | 创建时间     |


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
对于公开收藏夹不需要登录携带请求头，对于非公开收藏夹需要登录，且仅限创建者可以访问。


## 新建收藏夹

#### 请求URL
- ` http://域名/api/v1/collection/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型   | 说明       |
| :----- | :--- | :----- | ---------- |
| name   | 是   | string | 收藏夹名称 |

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


## 修改收藏夹信息

#### 请求URL
- ` http://域名/api/v1/collection/modify `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型   | 说明                |
| :----- | :--- | :----- | ------------------- |
| id     | 否   | int    | 收藏夹id            |
| name   | 是   | string | 收藏夹名称          |
| cover  | 否   | string | 收藏夹封面图        |
| desc   | 否   | string | 简介                |
| open   | 否   | bool   | 是否公开(默认false) |

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


## 删除收藏夹

#### 请求URL
- ` http://域名/api/v1/collection/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明     |
| :----- | :--- | :--- | -------- |
| id     | 是   | int  | 收藏夹id |

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

