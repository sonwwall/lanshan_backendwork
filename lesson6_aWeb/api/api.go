package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lesson6_aWeb/dao"
	"net/http"
	"strings"
	"time"
)

// JWTClaims 定义JWT的声明
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT
func GenerateToken(username string) (string, error) {
	key := []byte("my_secret_key") // 你应该使用一个更安全的密钥

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWTClaims, error) {
	key := []byte("my_secret_key")

	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func register(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	// 重复则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}

func login(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	// 正确则登录成功，生成JWT
	token, err := GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Failed to generate token",
		})
		return
	}
	// 返回结果和JWT
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
		"token":   token,
	})
}

func changePassword(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "No token provided",
		})
		return
	}

	// 验证token格式（应该以"Bearer "开头）
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Invalid token format",
		})
		return
	}

	// 解析token
	claims, err := ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Invalid token",
		})
		return
	}

	// 获取请求中的用户名和新密码
	username := c.PostForm("username")
	newPassword := c.PostForm("new_password")

	// 检查用户名是否匹配
	if claims.Username != username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Username in token does not match the one provided",
		})
		return
	}

	// 验证旧密码
	if dao.SelectPasswordFromUsername(username) != c.PostForm("password") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}

	// 更新密码
	dao.AddUser(username, newPassword)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Password changed successfully!",
	})
}
