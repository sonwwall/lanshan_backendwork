package dao

// 假数据库，用 map 实现
var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func AddUser(username, password string) {
	database[username] = password
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}
