# 视频相关接口

## 在线人数连接

#### 请求URL
- ` ws://域名/api/v1/video/online/ws?vid=视频ID&client_id=随机生成的客户端ID `

#### 返回示例 
```json
{
    "number": 1
}
```

#### 备注
如在线人数变换会向用户发送房间人数


## 通过视频ID获取视频信息

#### 请求URL
- ` http://域名/api/v1/video/get?vid=这里写视频ID `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "video": {
            "vid": 2,
            "title": "测试",
            "cover": "",
            "desc": "什么都没有",
            "created_at": "2022-06-06T08:42:13.525Z",
            "copyright": true,
            "author": {
                "uid": 1,
                "name": "user_1654250698886",
                "sign": "这个人很懒，什么都没有留下1",
                "avatar": "",
                "spacecover": "",
                "gender": 0,
            },
            "resources": [
                {
                    "id": 1,      
                    "title": "",   
                    "url": "",     
                    "duration": 10,
                    "quality": 1080,
                    "status": 0,
                }
            ],
            "clicks": 86
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 

##### 视频信息video
| 参数名     | 类型   | 说明                      |
| :--------- | :----- | ------------------------- |
| vid        | int    | 视频ID                    |
| title      | string | 标题                      |
| cover      | string | 封面URL                   |
| desc       | string | 视频简介                  |
| created_at | time   | 发布时间                  |
| copyright  | bool   | 是否为原创视频            |
| author     | object | 作者信息                  |
| resource   | array  | 视频资源,多个代表多个分集 |
| clicks     | int    | 点击量                    |

##### 作者信息author
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |

##### 视频资源resource
| 参数名   | 类型   | 说明           |
| :------- | :----- | -------------- |
| id       | int    | 分集ID         |
| title    | string | 标题           |
| url      | string | 视频链接       |
| duration | float  | 视频时长       |
| quality  | int    | 资源最大分辨率 |
| status   | int    | 审核状态       |


## 获取视频列表

#### 请求URL
- ` http://域名/api/v1/video/list?page=页码&page_size=内容数量&partition=分区id `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下1",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

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
分区id为0时获取所有分区的视频，分区id为一级分区时获取该分区下所有子分区的视频


## 获取推荐视频

#### 请求URL
- ` http://域名/api/v1/video/recommended?page_size=内容数量 `
  
#### 请求方式
- GET 

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

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


## 搜索视频

#### 请求URL
- `http://域名/api/v1/video/search?keywords=关键词&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下1",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

##### 作者信息author
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |


##### 备注
无


## 通过用户ID获取视频列表

#### 请求URL
- ` http://域名/api/v1/video/user/get?uid=用户id&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "created_at": "2022-06-06T08:42:13.525Z",
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型   | 说明     |
| :----- | :----- | -------- |
| total  | int    | 数量     |
| videos | object | 视频信息 |

##### 视频信息
| 参数名     | 类型   | 说明     |
| :--------- | :----- | -------- |
| vid        | int    | 视频ID   |
| title      | string | 标题     |
| cover      | string | 封面URL  |
| desc       | string | 视频简介 |
| clicks     | int    | 点击量   |
| created_at | time   | 发布时间 |

#### 备注
无


## 上传视频信息

#### 请求URL
- ` http://域名/api/v1/video/info/upload `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名    | 必选 | 类型   | 说明                       |
| :-------- | :--- | :----- | -------------------------- |
| title     | 是   | string | 视频标题                   |
| cover     | 是   | string | 封面图url                  |
| desc      | 否   | string | 视频简介 默认:"什么都没有" |
| copyright | 是   | bool   | 是否为原创视频             |
| partition | 是   | int    | 视频分区                   |

#### 返回示例 

``` json
  {
    "code": 200,
    "data": {
      "vid": 1
    },
    "msg":"ok"
  }
```

#### 返回参数说明 

| 参数名 | 类型 | 说明       |
| :----- | :--- | ---------- |
| vid    | int  | 新视频的ID |


#### 备注
无


## 修改视频信息

#### 请求URL
- ` http://域名/api/v1/video/info/modify `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名    | 必选 | 类型   | 说明           |
| :-------- | :--- | :----- | -------------- |
| vid       | 是   | int    | 视频ID         |
| title     | 是   | string | 视频标题       |
| cover     | 是   | string | 封面图url      |
| desc      | 否   | string | 视频简介       |
| copyright | 是   | bool   | 是否为原创视频 |


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


## 获取视频状态信息

#### 请求URL
- ` http://域名/api/v1/video/status?vid=这里填写视频ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "video": {
            "vid": 1,
            "title": "标题",
            "cover": "封面url",
            "desc": "视频简介",
            "status": 0,
            "partition": 1,
            "copyright": true,
            "resources": [
                {
                    "id": 1,
                    "title": "",
                    "url": "",
                    "duration": 10,
                    "quality": 1080,
                    "status": 0
                }
            ]
        }
    },
    "msg": "ok"
}
```
#### 返回参数说明 

##### 视频信息video
| 参数名    | 类型   | 说明                      |
| :-------- | :----- | ------------------------- |
| vid       | int    | 视频ID                    |
| title     | string | 标题                      |
| cover     | string | 封面URL                   |
| desc      | string | 视频简介                  |
| status    | int    | 视频审核状态              |
| partition | int    | 分区ID                    |
| copyright | bool   | 是否为原创视频            |
| resource  | array  | 视频资源,多个代表多个分集 |

##### 视频资源resource
| 参数名   | 类型   | 说明           |
| :------- | :----- | -------------- |
| id       | int    | 分集ID         |
| title    | string | 标题           |
| url      | string | 视频链接       |
| duration | float  | 视频时长       |
| quality  | int    | 资源最大分辨率 |
| status   | int    | 审核状态       |

#### 备注
无


## 视频提交审核

#### 请求URL
- ` http://域名/api/v1/video/review/submit `
  
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
至少上传一个视频才可以提交


## 删除视频

#### 请求URL
- ` http://域名/api/v1/video/delete `
  
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


## 获取视频列表

#### 请求URL
- ` http://域名/api/v1/video/collect?page=页码&page_size=内容数量&id=收藏夹id `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

####  请求头
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "clicks": 86,
                "created_at": "2022-06-06T08:42:13.525Z",
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名     | 类型   | 说明     |
| :--------- | :----- | -------- |
| vid        | int    | 视频ID   |
| title      | string | 标题     |
| cover      | string | 封面URL  |
| desc       | string | 视频简介 |
| clicks     | int    | 点击量   |
| created_at | time   | 发布时间 |

#### 备注
如果收藏夹是公开的则不需要携带请求头


## 获取上传视频列表

#### 请求URL
- ` http://域名/api/v1/video/upload/get?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token `

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 2000,
    "data": {
        "count": 1,
        "videos": [
            {
                "vid": 2,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "clicks": 10,
				"review": 2000,
                "created_at": "2021-07-16T08:49:54Z",
                "updated_at": "2021-07-16T09:04:47Z"
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型   | 说明     |
| :----- | :----- | -------- |
| total  | int    | 数量     |
| videos | object | 视频信息 |


| 参数名     | 类型   | 说明         |
| :--------- | :----- | ------------ |
| vid        | int    | 视频ID       |
| title      | string | 标题         |
| cover      | string | 封面图url    |
| desc       | string | 简介         |
| clicks     | int    | 视频点击量   |
| status     | int    | 视频审核状态 |
| updated_at | string | 最后更新时间 |

#### 备注
无


## 管理员获取视频列表

#### 请求URL
- ` http://域名/api/v1/video/manage/list?page=页码&page_size=内容数量&partition=分区ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 请求参数
- 见url
- 分区ID默认为所有分区


#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下1",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型  | 说明     |
| :----- | :---- | -------- |
| total  | int   | 数量     |
| videos | array | 视频信息 |


##### 视频信息videos
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

##### 作者信息author
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |

##### 备注
需要审核及以上权限   


## 管理员搜索视频

#### 请求URL
- `http://域名/api/v1/video/manage/search?keywords=关键词&page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下1",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型  | 说明     |
| :----- | :---- | -------- |
| total  | int   | 数量     |
| videos | array | 视频信息 |


##### 视频信息videos
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

##### 作者信息author
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |


##### 备注
需要审核及以上权限


## 管理员删除视频

#### 请求URL
- ` http://域名/api/v1/video/manage/delete `
  
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
需要审核及以上权限


## 获取待审核视频列表

#### 请求URL
- ` http://域名/api/v1/video/manage/review/list?page=页码&page_size=内容数量 `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "total": 1,
        "videos": [
            {
                "vid": 1,
                "title": "测试1",
                "cover": "",
                "desc": "",
                "created_at": "2022-06-06T08:42:13.525Z",
                "copyright": true,
                "author": {
                    "uid": 1,
                    "name": "user_1654250698886",
                    "sign": "这个人很懒，什么都没有留下1",
                    "avatar": "",
                    "spacecover": "",
                    "gender": 0,
                },
                "clicks": 86,
                "partition": 1
            },
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名 | 类型  | 说明     |
| :----- | :---- | -------- |
| total  | int   | 数量     |
| videos | array | 视频信息 |


##### 视频信息videos
| 参数名     | 类型   | 说明           |
| :--------- | :----- | -------------- |
| vid        | int    | 视频ID         |
| title      | string | 标题           |
| cover      | string | 封面URL        |
| desc       | string | 视频简介       |
| created_at | time   | 发布时间       |
| copyright  | bool   | 是否为原创视频 |
| author     | object | 作者信息       |
| partition  | int    | 分区ID         |
| clicks     | int    | 点击量         |

##### 作者信息author
| 参数名     | 类型   | 说明                           |
| :--------- | :----- | ------------------------------ |
| uid        | int    | 用户ID                         |
| name       | string | 用户名                         |
| sign       | string | 个性签名                       |
| avatar     | string | 头像                           |
| spacecover | string | 用户空间封面图                 |
| gender     | int    | 用户性别，0：未知;1：男；2：女 |


##### 备注
需要审核及以上权限


## 获取待审核资源列表

#### 请求URL
- ` http://域名/api/v1/video/manage/review/resource/list?vid=视频ID `
  
#### 请求方式
- GET 

####  请求头
- `Authorization': access_token`

#### 请求参数
- 见url

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "resources": [
            {
                "id": 1,      
                "title": "",   
                "url": "",     
                "duration": 10,
                "quality": 1080,
                "status": 0,
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 
| 参数名   | 类型   | 说明           |
| :------- | :----- | -------------- |
| id       | int    | 分集ID         |
| title    | string | 标题           |
| url      | string | 视频链接       |
| duration | float  | 视频时长       |
| quality  | int    | 资源最大分辨率 |
| status   | int    | 审核状态       |

#### 备注
需要审核及以上权限


## 审核视频 

#### 请求URL
- ` http://域名/api/v1/video/manage/review/video `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明     |
| :----- | :--- | :--- | -------- |
| id     | 是   | int  | 视频ID   |
| status | 是   | int  | 审核状态 |

#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 

需要审核及以上权限


## 审核视频资源

#### 请求URL
- ` http://域名/api/v1/video/manage/review/resource `
  
#### 请求方式
- POST 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 参数

| 参数名 | 必选 | 类型 | 说明     |
| :----- | :--- | :--- | -------- |
| id     | 是   | int  | 视频ID   |
| status | 是   | int  | 审核状态 |

#### 返回示例 

``` json
  {
    "code": 200,
    "data": null,
    "msg":"ok"
  }
```

#### 备注 

需要审核及以上权限
