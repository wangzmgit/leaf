# 文件上传相关接口

## 视频封面上传接口

#### 请求URL
- ` http://域名/api/v1/upload/image `
  
#### 请求方式
- POST 

#### 请求头
- `Authorization': access_token `

#### 参数
| 参数名 | 必选 | 类型 | 说明     |
| :----- | :--- | :--- | -------- |
| image  | 否   | file | 图片文件 |

- 在postman中使用form-data类型进行测试

#### 返回示例 

``` json
  {
    "code": 200,
    "data": {
	    "url":"",
    },
    "msg":"ok"
  }
```

#### 返回参数说明 

| 参数名 | 类型   | 说明    |
| :----- | :----- | ------- |
| url    | string | 图片url |

#### 备注
无

    
## 视频上传接口

#### 请求URL
- ` http://域名/api/v1/upload/video/:vid `
  
#### 请求方式
- POST 

#### 请求头
- `Authorization': access_token `

#### 参数

| 参数名 | 必选 | 类型   | 说明     |
| :----- | :--- | :----- | -------- |
| vid    | 否   | string | 视频id（在URL中）   |
| video  | 否   | file   | 视频文件 |

- 在postman中使用form-data类型进行测试

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