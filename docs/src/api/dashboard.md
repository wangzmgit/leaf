# 后台仪表盘相关接口

:::tip 
该部分接口为v1.2版本新增接口
:::

## 获取仪表盘卡片数据

#### 请求URL
- ` http://域名/api/v1/dashboard/card/data `
  
#### 请求方式
- GET 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "data": {
            "user_count": 95,
            "video_count": 42,
            "review_video_count": 0,
            "new_user_count": 0,
            "new_video_count": 0
        }
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名             | 类型   | 说明         |
| :----------------- | :----- | ------------ |
| user_count         | number | 用户数       |
| video_count        | number | 视频数       |
| review_video_count | number | 待审核视频数 |
| new_user_count     | number | 新增用户数   |
| new_video_count    | number | 新增视频数   |

#### 备注 

需要审核及以上权限

## 获取新增用户和视频的趋势数据

#### 请求URL
- ` http://域名/api/v1/dashboard/trend `
  
#### 请求方式
- GET 

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "data": [
            {
                "user": 0,
                "video": 0,
                "date": "2023/10/13"
            },
            {
                "user": 0,
                "video": 0,
                "date": "2023/10/12"
            },
            {
                "user": 0,
                "video": 0,
                "date": "2023/10/11"
            },
            {
                "user": 0,
                "video": 0,
                "date": "2023/10/10"
            },
            {
                "user": 0,
                "video": 0,
                "date": "2023/10/09"
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名 | 类型   | 说明       |
| :----- | :----- | ---------- |
| user   | number | 新增用户数 |
| video  | number | 新增视频数 |
| date   | string | 日期       |

#### 备注 

需要审核及以上权限


## 获取分区视频数量数据

#### 请求URL
- ` http://域名/api/v1/dashboard/partition?partition_id=分区ID `
  
#### 请求方式
- GET 

#### 请求参数
- 见url

####  请求头
- `"content-type": "application/json",`
- `Authorization': access_token`

#### 返回示例 

``` json
{
    "code": 200,
    "data": {
        "data": [
            {
                "content": "MAD·AMV",
                "video_count": 5
            },
            {
                "content": "MMD·3D",
                "video_count": 1
            },
            {
                "content": "番剧二创",
                "video_count": 1
            },
            {
                "content": "综合",
                "video_count": 2
            },
            {
                "content": "特摄",
                "video_count": 0
            }
        ]
    },
    "msg": "ok"
}
```

#### 返回参数说明 

| 参数名      | 类型   | 说明         |
| :---------- | :----- | ------------ |
| content     | string | 分区名       |
| video_count | number | 分区视频数量 |

#### 备注 

需要审核及以上权限
