# 开始

## ~~双token登录方案~~
在v1.1.0及之后的版本中，双token方案已经被移除，文档中携带的access_token全部改为token
![双token登录方案](/token.png)

## 稿件状态码
| 状态                | 状态码 | 含义             |
| ------------------- | ------ | ---------------- |
| AUDIT_APPROVED      | 0      | 审核通过         |
| CREATED_VIDEO       | 100    | 成功创建视频     |
| VIDEO_PROCESSING    | 200    | 视频转码中       |
| SUBMIT_REVIEW       | 300    | 用户提交审核     |
| WAITING_REVIEW      | 500    | 等待审核         |
| WRONG_VIDEO_INFO    | 2100   | 视频信息存在问题 |
| WRONG_VIDEO_CONTENT | 2200   | 视频内容存在问题 |
| PROCESSING_FAIL     | 2300   | 视频处理失败     |


## 用户身份
| 角色    | 角色编号 | 含义     |
| ------- | -------- | -------- |
| root    | 3        | 超级管理 |
| admin   | 2        | 管理     |
| auditor | 1        | 审核     |
| user    | 0        | 普通用户 |

角色权限向下兼容


