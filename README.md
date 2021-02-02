## My Prometheus MySQL Exporter
基于Go语言实现的Prometheus MySQL Exporter，因为只是学习缘故，所以只实现了几个最简单的指标监控，比如慢查询等  
除了指标监控外，还有一些其他功能：
* 支持从文件和环境变量读取配置
* 支持日志记录，可选择输出到文件/stdout/stderr，可选择日志保留时长等常用配置
* Basic Auth认证，支持配置多个用户名密码
* 支持HTTPS配置
* 有一个简单的落地页(首页

### 下载
`git clone https://github.com/VVFock3r/mysql_exporter.git`

### 配置数据库
`/etc/conf/mysql_exporter`  
```
mysql:  
   host: 192.168.73.129  
   port: 3306
   username: root
   password: 123456
```
### 运行服
`go run main.go`

### 访问测试
`http://ip:9100/`

### 屏幕截图

![screen_capture](https://raw.github.com/VVFock3r/mysql_exporte/docs/screen_capture.png)

