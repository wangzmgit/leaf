# 点赞收藏相关接口

## 获取点赞收藏数据

#### 请求URL
- ` http://域名/api/v1/archive/stat?vid=视频ID `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "stat": {
            "collect": 1,
            "like": 1
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名  | 类型 | 说明     |
| :------ | :--- | -------- |
| collect | int  | 收藏数量 |
| like    | int  | 点赞数量 |

#### 备注
无


## 点赞

#### 请求URL
- ` http://域名/api/v1/archive/like `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 视频ID |

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
    

## 取消点赞

#### 请求URL
- ` http://域名/api/v1/archive/cancel/like `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token `

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 视频ID |

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


## 获取点赞状态

#### 请求URL
- ` http://域名/api/v1/archive/has/like?vid=视频ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
      "like": true
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型 | 说明     |
| :----- | :--- | -------- |
| like   | bool | 是否点赞 |

#### 备注
无


## 添加收藏

#### 请求URL
- ` http://域名/api/v1/archive/collect `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名     | 必选 | 类型       | 说明             |
| :--------- | :--- | :--------- | ---------------- |
| vid        | 是   | int        | 视频ID           |
| addList    | 否   | array[int] | 添加的收藏夹数组 |
| cancelList | 否   | array[int] | 移除的收藏夹数组 |

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


## 获取已收藏的文件夹

#### 请求URL
- ` http://域名/api/v1/archive/collect/collected?vid=视频ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
      "collectionIds": [1, 2]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名        | 类型  | 说明               |
| :------------ | :---- | ------------------ |
| collectionIds | array | 视频所在的收藏夹id |

#### 备注
无

## 获取收藏状态

#### 请求URL
- ` http://域名/api/v1/archive/has/collect?vid=视频ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
      "collect": true
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名  | 类型 | 说明     |
| :------ | :--- | -------- |
| collect | bool | 是否收藏 |

#### 备注
无
