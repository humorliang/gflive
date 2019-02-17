# gflive
## API设计
### API设计：用户
* 创建(注册)用户： URL:/user Method:POST SC:201,400,500
* 用户登录： URL:/user/:username Method:POST SC:200,400,500
* 获取用户基本信息： URL:/user/：username Method:GET SC:200,400,401(验证不通过),403(验证通过没权限),500
* 用户注销： URL:/user Method:DELETE SC:204(no content),400,401,403,500
### API设计：用户资源
* List all videos: URL:/user/:username/videos Method:GET, SC:200,400,500
* Get one video: URL:/user/:username/videos/:vid-id Method:GET, SC:200,400,500
* Delete one video: URL:/user/:username/videos/vid-id Method:DELETE, SC:204,400,401,403,500
### API设计：评论
* Show comments: URL:/videos/:vid-id/comments Method:GET, SC:200,400,500
* Post a comments: URL:/videos/:vid-id/comments Method:POST, SC:201,400,500
* Delete a comments: URL:/videos/:vid-id/comment/:comment-id Method:DELETE, SC:204,400,401,403,500


## 数据库设计

用户表users
```sql
id UNSIGNED INT,PRIMARY KEY,AUTO_INCREMENT,
login_name VARCHAR(64),UNIQUE KEY,
pwd TEXT
```
视频资源video_info
```sql
id VARCHAR(64),PRIMARY KEY,NOT NULL,
author_id UNSIGNED INT,
name TEXT,
dispaly_ctime TEXT,
create_time DATETIME
```
评论表comments
```sql
id VARCHAR(64),PRIMARY KEY,NOT NULL,
video_id VARCHAR(64),
author_id UNSIGNED INT,
content TEXT,
time DATETIME
```
会话表sessions
```sql
session_id TINYTEXT,PRIMARY KEY,NOT NULL,
TTL TINYTEXT,#过期时间
login_name VARCHAR(64)
```