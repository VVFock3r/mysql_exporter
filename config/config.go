package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Options struct {
	MySQL  MySQL  `mapstructure:"mysql"`
	Web    Web    `mapstructure:"web"`
	Logger Logger `mapstructure:"logger"`
}

// mysql
type MySQL struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Db          string `mapstructure:"db"`
	CharSet     string `mapstructure:"charset"`
	TimeZone    string `mapstructure:"timezone"`
	ConnTimeout string `mapstructure:"conn_timeout"`
	ReadTimeout string `mapstructure:"read_timeout"`
}

// web
type Web struct {
	Addr Addr   `mapstructure:"addr"`
	Auth []Auth `mapstructure:"auth"`
	SSL  SSL    `mapstructure:"ssl"`
}

type Addr struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Path string `mapstructure:"path"`
}

type Auth struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type SSL struct {
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

// log
type Logger struct {
	Global Global       `mapstructure:"global"`
	File   FileLogger   `mapstructure:"file"`
	Stdout StdoutLogger `mapstructure:"stdout"`
	Stderr StderrLogger `mapstructure:"stderr"`
}

type Global struct {
	Level string
}

type FileLogger struct {
	Enabled    bool   `mapstructure:"enabled"`
	FileName   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type StdoutLogger struct {
	Enabled bool `mapstructure:"enabled"`
}
type StderrLogger struct {
	Enabled bool `mapstructure:"enabled"`
}

func notEmptyCheck(items map[string]string) error {
	for key, value := range items {
		if value == "" {
			return fmt.Errorf(fmt.Sprintf("%s %s", key, "cannot be empty"))
		}
	}
	return nil
}

func ParseConfig(path string) (*Options, error) {
	conf := viper.New()

	// 设置默认参数
	conf.SetDefault("mysql.port", 3306)
	conf.SetDefault("mysql.db", "information_schema")
	conf.SetDefault("mysql.charset", "utf8mb4")
	conf.SetDefault("mysql.timezone", "Local")
	conf.SetDefault("mysql.conn_timeout", "3s")
	conf.SetDefault("mysql.read_timeout", "3s")

	conf.SetDefault("web.addr.host", "0.0.0.0")
	conf.SetDefault("web.addr.port", 9100)
	conf.SetDefault("web.addr.path", "/metrics/")

	conf.SetDefault("logger.global.level", "warning")

	// 支持从环境变量解析参数
	conf.AutomaticEnv()
	conf.SetEnvPrefix("MySQL_EXPORTER")

	// 读取配置
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	// 将配置映射到结构体
	options := &Options{}
	if err := conf.Unmarshal(options); err != nil {
		return nil, err
	}

	// 检查必须参数是否有值
	if err := notEmptyCheck(map[string]string{
		"mysql.host":     options.MySQL.Host,
		"mysql.username": options.MySQL.Username,
		"mysql.password": options.MySQL.Password,
	}); err != nil {
		return nil, err
	}

	// 数据类型转换

	return options, nil
}
