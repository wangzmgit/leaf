# 用户相关接口

## 用户注册

#### 请求URL
- ` http://域名/api/v1/user/register `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`

#### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| email    | 是   | string | 邮箱   |
| password | 是   | string | 密码   |
| code     | 是   | string | 验证码 |

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

## 用户登录（账号）

#### 请求URL
- ` http://域名/api/v1/user/login `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`

#### 参数

| 参数名   | 必选 | 类型   | 说明 |
| :------- | :--- | :----- | ---- |
| email    | 是   | string | 邮箱 |
| password | 是   | string | 密码 |

#### 返回示例 
登录成功返回值
```json
  {
    "code": 200,
    "data": {
      "access_token": "", // v1.1.0之前返回该数据
      "refresh_token": "", // v1.1.0之前返回该数据
      "token": "", // v1.1.0之后返回该数据
    },
    "msg":"ok"
  }
```

连续登录失败三次会返回值
```json
  {
    "code": -1,
    "data": null,
    "msg":"需要人机验证"
  }
```

#### 返回参数说明 

| 参数名        | 类型   | 说明                                              |
| :------------ | :----- | ------------------------------------------------- |
| access_token  | string | 请求使用的token，5分钟内有效 (v1.1.0移除)         |
| refresh_token | string | 获取access_token的token，14天内有效  (v1.1.0移除) |
| token         | string | 请求使用的token，14天内有效 (v1.1.0新增)          |

#### 备注
同一邮箱连续登录失败三次会返回需要人机验证，此时需要等待30分钟或者调用人机验证接口并通过滑块验证才可以继续登录

## 用户登录（邮箱）

#### 请求URL
- ` http://域名/api/v1/user/login/email `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`

#### 参数

| 参数名 | 必选 | 类型   | 说明   |
| :----- | :--- | :----- | ------ |
| email  | 是   | string | 邮箱   |
| code   | 是   | string | 验证码 |

#### 返回示例 
登录成功返回值
```json
  {
    "code": 200,
    "data": {
      "access_token": "", // v1.1.0之前返回该数据
      "refresh_token": "", // v1.1.0之前返回该数据
      "token": "", // v1.1.0之后返回该数据
    },
    "msg":"ok"
  }
```

连续登录失败三次会返回值
```json
  {
    "code": -1,
    "data": null,
    "msg":"需要人机验证"
  }
```

#### 返回参数说明 

| 参数名        | 类型   | 说明                                              |
| :------------ | :----- | ------------------------------------------------- |
| access_token  | string | 请求使用的token，5分钟内有效 (v1.1.0移除)         |
| refresh_token | string | 获取access_token的token，14天内有效  (v1.1.0移除) |
| token         | string | 请求使用的token，14天内有效 (v1.1.0新增)          |


#### 备注
同一邮箱连续登录失败三次会返回需要人机验证，此时需要等待30分钟或者调用人机验证接口并通过滑块验证才可以继续登录


## 发送邮箱验证码

#### 请求URL
- ` http://域名/api/v1/user/code/email `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`

#### 参数

| 参数名 | 必选 | 类型   | 说明     |
| :----- | :--- | :----- | -------- |
| email  | 是   | string | 目标邮箱 |

#### 返回示例 

```json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 
需要先进行人机验证


## 通过用户ID获取用户信息

#### 请求URL
- ` http://localhost:9000/api/v1/user/info/other?uid=用户ID `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "user": {
            "uid": 1,
            "name": "",
            "sign": "",
            "avatar": "",
            "spacecover": "",
            "gender": 1,
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 

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

## 通过用户名获取用户id

#### 请求URL
- ` http://域名/api/v1/user/uid?name=用户名 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
  {
    "code": 200,
    "data": {
      "uid": 1
    },
    "msg":"ok"
}
```

#### 返回参数说明 

| 参数名 | 类型 | 说明   |
| :----- | :--- | ------ |
| uid    | int  | 用户ID |

#### 备注
无

## 修改密码验证

#### 请求URL
- ` http://域名/api/v1/user/resetpwd/check?eamil=用户邮箱 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
}
```

#### 备注
需要先进行人机验证，验证通过后可修改密码


## 修改密码

#### 请求URL
- ` http://域名/api/v1/user/pwd/modify `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`

#### 参数

| 参数名   | 必选 | 类型   | 说明       |
| :------- | :--- | :----- | ---------- |
| email    | 是   | string | 邮箱       |
| password | 是   | string | 新密码     |
| code     | 是   | string | 邮箱验证码 |

#### 返回示例 

```json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 
需要先调用[修改密码验证](#修改密码验证)接口 


## ~~刷新token~~

此接口在v1.1.0中移除移除

#### 请求URL
- ` http://域名/api/v1/user/token/refresh`
  
#### 请求方式
- GET 

#### 请求头
- `Authorization': refresh_token `
- `"content-type": "application/json"`

#### 返回示例 

``` json
  {
    "code": 200,
    "data": {
      "token": "" 
    },
    "msg":"ok"
}
```

#### 返回参数说明 

| 参数名 | 类型   | 说明         |
| :----- | :----- | ------------ |
| token  | string | access_token |

#### 备注
需要refresh_token


## 用户获取个人信息

#### 请求URL
- ` http://域名/api/v1/user/info/get `
  
#### 请求方式
- GET 

#### 请求头
- `Authorization': access_token `
- `"content-type": "application/json"`

#### 返回示例 

``` json
  {
    "code": 200,
    "data": {
      "user_info": {
        "uid": 1,
        "name": "昵称",
        "sign": "个性签名",
        "email": "1****1@qq.com",
        "avatar": "",
        "spacecover": "",
        "gender": 1,
        "role": 0,
        "birthday":"",
        "created_at": ""
      }
    },
    "msg":"ok"
}
```

#### 返回参数说明 

| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| email      | string | 用户邮箱（脱敏后数据）         |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |
| role       | int    | [用户身份](/api/#用户身份)     |
| birthday   | time   | 生日                           |
| created_at | time   | 注册时间                       |

#### 备注
无


## 用户修改个人信息

#### 请求URL
- ` https://域名/api/v1/user/info/modify `
  
#### 请求方式
- POST 

#### 请求头
- `Authorization': access_token `
- `"content-type": "application/json"`

#### 参数

| 参数名   | 必选 | 类型               | 说明                 |
| :------- | :--- | :----------------- | -------------------- |
| avatar   | 是   | string             | 用户头像             |
| name     | 是   | string             | 用户名               |
| birthday | 是   | string（日期格式） | 生日（默认1970-1-1） |
| gender   | 否   | int                | 性别（默认为0未知）  |
| sign     | 否   | string             | 个性签名             |

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


## 用户修改空间封面图

#### 请求URL
- ` https://域名/api/v1/user/cover/modify `
  
#### 请求方式
- POST 

#### 请求头
- `Authorization': access_token `
- `"content-type": "application/json"`

#### 参数

| 参数名     | 必选 | 类型   | 说明      |
| :--------- | :--- | :----- | --------- |
| spacecover | 是   | string | 封面图url |

#### 返回示例 

``` json
{
  "code": 200,
  "data": null,
  "msg": "ok"
}
```

#### 备注
封面图url必须为调用图片上传接口返回的url


## 管理员获取用户列表

#### 请求URL
- ` http://域名/api/v1/user/manage/list?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 请求参数
- 见url
- page默认为1,page_size默认为15，page_size最大值为30

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "users": [
            {
                "uid": 1,
                "name": "昵称",
                "sign": "个性签名",
                "email": "1****1@qq.com",
                "avatar": "",
                "spacecover": "",
                "gender": 1,
                "role": 0,
                "birthday":"",
                "created_at": ""
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| email      | string | 用户邮箱                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |
| role       | int    | [用户身份](/api/#用户身份)     |
| birthday   | time   | 生日                           |
| created_at | time   | 生日                           |

#### 备注
需要管理员权限


## 管理员搜索用户

#### 请求URL
- ` http://域名/api/v1/user/manage/search&keyword=关键词&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 请求参数
- 见url
- page默认为1,page_size默认为15，page_size最大值为30

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "users": [
            {
                "uid": 1,
                "name": "昵称",
                "sign": "个性签名",
                "email": "1****1@qq.com",
                "avatar": "",
                "spacecover": "",
                "gender": 1,
                "role": 0,
                "birthday":"",
                "created_at": ""
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| email      | string | 用户邮箱                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |
| role       | int    | [用户身份](/api/#用户身份)     |
| birthday   | time   | 生日                           |
| created_at | time   | 生日                           |

#### 备注
通过用户名、邮箱、id搜索用户的。需要管理员权限


## 管理员修改用户信息

#### 请求URL
- ` https://域名/api/v1/user/manage/modify `
  
#### 请求方式
- POST 

#### 请求头
- `"content-type": "application/json",`
- - `Authorization': access_token `

#### 参数

| 参数名 | 必选 | 类型   | 说明     |
| :----- | :--- | :----- | -------- |
| id     | 是   | int    | 用户id   |
| name   | 是   | string | 用户名   |
| email  | 是   | string | 邮箱     |
| sign   | 否   | string | 个性签名 |


#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg": "ok"
}
```

##### 备注 
需要管理员权限


## 修改用户权限

#### 请求URL
- ` http://域名/api/v1/user/manage/role/modify `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token `

#### 参数

| 参数名 | 必选 | 类型 | 说明                       |
| :----- | :--- | :--- | -------------------------- |
| id     | 是   | int  | 用户id                     |
| role   | 是   | int  | [用户身份](/api/#用户身份) |
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


## 删除用户

#### 请求URL
- ` http://域名/api/v1/user/manage/delete `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token `

#### 参数

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 用户id |

#### 返回示例 
```json
  {
    "code": 200,
    "data": null,
    "msg": "ok"
  }
```

#### 备注 
需要管理员权限
