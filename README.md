## My Prometheus MySQL Exporter
基于Go语言实现的Prometheus MySQL Exporter，因为只是学习缘故，所以只实现了几个最简单的指标监控，比如慢查询等  
除了指标监控外，还有一些其他功能：
* 支持从文件和环境变量读取配置
* 支持日志记录，可选择输出到文件/stdout/stderr，可选择日志保留时长等常用配置
* Basic Auth认证，支持配置多个用户名密码
* 支持HTTPS配置，项目带了自签证书用于测试，使用方法
  * 服务端打开SSL配置，使用项目默认带的证书，域名为：example.com
  * 客户端导入client.crt证书
  * 域名解析为服务器IP
  * 访问测试https://example.com:port
* 有一个简单的落地页(首页)

### 下载
`git clone https://github.com/VVFock3r/mysql_exporter.git`

### 配置数据库
`/etc/conf/mysql_exporter.yml`  

```
mysql:  
   host: 192.168.73.129  
   port: 3306
   username: root
   password: 123456
```
### 运行服
方式一：编译并运行：`go run main.go`

方式二：编译好的二进制文件：`./main`



### 访问测试
`http://ip:9100/`

### 屏幕截图

![image](https://raw.github.com/VVFock3r/mysql_exporter/main/docs/screen_capture.png)

