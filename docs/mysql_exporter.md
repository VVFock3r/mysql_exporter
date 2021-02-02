##### MySQL 配置             
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  host  | string |  false  |        | mysql地址             |
|  port  | int |  true  |   3306  | mysql端口             |
|  username  | string |  false  |     | mysql用户名             |
|  password  | string |  false  |     | mysql密码             |
|  db  | string |  true  |  information_schema   | 数据库名             |
|  charset  | string |  true  |  utf8mb4   | 客户端字符集             |
|  timezone  | string |  true  |  Local   | 时区             |
|  conn_timeout  | string |  true  |  3s   | 连接超时时间             |
|  read_timeout  | string |  true  |  3s   | 读超时时间             |




#####  Web端 配置

###### Addr 监听地址
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  host  | string |  true  |  0.0.0.0      | 监听地址             |
|  port  | int |  true  |  9100      | 监听端口             |
|  path  | string |  true  |  /metrics/      | metrics路径             |
###### Auth 基础认证
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  username  | string |  true  |        | Basic Auth用户名             |
|  password  | string |  true  |        | Basic Auth密码             |
> Tips:  
>&nbsp;&nbsp;&nbsp;(1) 密码使用bcrypt加密，可以使用`htpasswd -B -n  <用户名>` 生成密文  
>&nbsp;&nbsp;&nbsp;(2) 支持填写多组用户名密码
###### SSL
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  cert_file  | string |  true  |        | SSL公钥文件 |
|  key_file  | string |  true  |        | SSL私钥文件 |
> 开启SSL后会默认会使用HTTP/2，这是由Go语言的http包决定的  
> 在响应头中可以看到HTTP版本信息，如果Chrome查看不方便可以使用Firefox




##### Log 配置

###### global 全局配置  
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  level  | string |  true  |  warning      | 日志等级，可选值：debug/error/warn/warning/info |
###### file 输出到文件
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  enabled  | bool |  true  |  false      | 文件日志开关 |
|  filename  | string |  true  |        | 文件名|
|  max_age  | int |  true  |        | 文件最多保存多少天，过期自动删除|
|  max_size  | int |  true  |        | 每个日志文件保存的最大尺寸(M) |
|  max_backups  | int |  true  |        | 最多保存多少个备份|
|  compress  | bool |  true  |  | 是否压缩|
###### stdout 输出到控制台标准输出
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  enabled  | bool |  true  |  false      | 控制台标准输出日志开关 |
###### stderr 输出到控制台标准错误输出
|  字段  | 类型   | 可以为空| 默认值 | 描述                  |
| :----: | :----: | :----:  | :----: | :----:                |
|  enabled  | bool |  true  |  false      | 控制台标准错误输出日志开关 |






