# 视频资源接口

## 修改资源标题

#### 请求URL
- ` http://域名/api/v1/resource/title/modify `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型   | 说明     |
| :----- | :--- | :----- | -------- |
| id     | 是   | int    | 资源ID   |
| title  | 是   | string | 资源标题 |

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


## 删除资源

#### 请求URL
- ` http://域名/api/v1/resource/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 否   | int  | 视频id |

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