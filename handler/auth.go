package handler

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"mysql_exporter/config"
	"net/http"
	"strings"
)

func Auth(handler http.Handler, secrets []config.Auth) http.Handler {
	//return handler

	// 类型转换，将自定义函数转为 http.Handler接口
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// Basic Auth认证
		secret := request.Header.Get("Authorization")
		if !BasicAuth(secret, secrets) {
			response.Header().Set("WWW-Authenticate", `Basic realm=""`)
			response.WriteHeader(401)
			return
		}

		// 调用原始handler方法
		handler.ServeHTTP(response, request)
	})
}

func BasicAuth(secret string, secrets []config.Auth) bool {
	// 无验证
	if secrets == nil || len(secrets) == 0 {
		return true
	}

	// 分割Basic Auth字符串
	nodes := strings.Fields(secret)
	if len(nodes) != 2 {
		return false
	}

	// 解析出明文 username:password
	plaintext, err := base64.StdEncoding.DecodeString(nodes[1])
	if err != nil {
		return false
	}

	// 分割用户名、密码
	nodes = strings.SplitN(string(plaintext), ":", 2)
	if len(nodes) != 2 {
		return false
	}

	// 验证用户名密码
	for _, item := range secrets {
		// 验证用户名
		if item.Username != nodes[0] {
			continue
		}

		// 验证密码
		return bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(nodes[1])) == nil
	}

	return false
}
