# 分区相关接口

## 获取分区列表

#### 请求URL
- ` http://域名/api/v1/partition/get `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
  "code": 200,
  "data": {
    "partitions": [
      {
        "id": 2,
        "content": "动画",
        "parent_id": 0
      },
      {
        "id": 9,
        "content": "MAD·AMV",
        "parent_id": 0
      }
    ]
  },
  "msg": "ok"
}
```

## 新增分区

#### 请求URL
- ` http://域名/api/v1/partition/add `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名   | 必选 | 类型   | 说明       |
| :------- | :--- | :----- | ---------- |
| content  | 是   | string | 分区名     |
| parentId | 否   | int    | 所属分区id |

#### 返回示例 

```json
  {
    "code": 200,
    "data": null,
    "msg": "ok"
  }
```

#### 备注 
需要超级管理员权限


## 删除分区

#### 请求URL
- ` http://域名/api/v1/partition/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 分区id |

#### 返回示例 

```json
  {
    "code": 200,
    "data": null,
    "msg": "ok"
  }
```

#### 备注 
需要超级管理员权限
