# 轮播图相关接口

## 获取轮播图

#### 请求URL
- ` http://域名/api/v1/carousel/get `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "carousels": [
            {
                "id": 1,
                "img": "123",
                "url": "",
                "title": "",
                "color": "",
                "created_at": "2021-08-01T10:39:25Z"
            },
        ]
    },
    "msg": "ok"
}
```
#### 返回参数说明
| 参数名     | 类型   | 说明               |
| :--------- | :----- | ------------------ |
| id         | int    | 轮播图id           |
| img        | string | 图片地址           |
| url        | string | 单击图片前往的地址 |
| title      | string | 标题               |
| color      | string | 主题色             |
| created_at | time   | 发布时间           |

#### 备注 
无



## 上传轮播图信息

#### 请求URL
- ` http://域名/api/v1/carousel/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型   | 说明                |
| :----- | :--- | :----- | ------------------- |
| img    | 是   | string | 轮播图url           |
| url    | 否   | string | 单击轮播图打开的url |
| title  | 否   | string | 标题                |
| color  | 否   | string | 主题色              |

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


    
## 删除轮播图

#### 请求URL
- ` http://域名/api/v1/carousel/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明     |
| :----- | :--- | :--- | -------- |
| id     | 是   | int  | 轮播图id |

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