package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// 自定义日志格式
type LogFormatter struct{}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
