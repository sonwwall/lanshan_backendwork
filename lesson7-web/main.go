package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
)

func newlottery(db *gorm.DB, c *gin.Context) {
	//1.生成随机数，选择奖品
	num := rand.Intn(2)
	var result string
	switch num {
	case 0:
		c.JSON(200, gin.H{"msg": "很遗憾，您没有中奖"})
		result = "NO"
	case 1:
		c.JSON(200, gin.H{"msg": "恭喜，您中奖了"})
		result = "YES"
	}
	//3.存入数据库
	db.Create(&Lottery01{Results: result})

}

type Lottery01 struct {
	gorm.Model
	Results string
}

func main() {

	router := gin.Default()

	// 开启 CORS 支持

	// 配置 CORS，允许来自 http://localhost:63342 的请求
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342", "http://localhost:8080"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},         // 允许的 HTTP 方法，包括 OPTIONS
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "num"},  // 允许的请求头
		AllowCredentials: true,                                                        // 允许携带凭证（如 Cookies）
	}))

	//************修改这下面一行以连接自己的数据库**********//
	dsn := "root:G794q028@tcp(127.0.0.1:3306)/game?charset=utf8&mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Lottery01{})

	router.POST("/home",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "欢迎来到抽奖系统"})
		},
		func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "请选择你想进行的功能"})
		},
	)
	router.POST("/lottery", func(c *gin.Context) {
		method, _ := c.GetQuery("method")

		if method == "newlottery" {
			newlottery(db, c)
		} else if method == "queryByid" {
			// 第一步：先提示用户输入查询次数
			c.JSON(200, gin.H{"msg": "请输入你想查询的次数"})
		} else if method == "delete" {
			c.JSON(200, gin.H{"msg": "请输入你删除第几次的结果	"})
		} else if method == "delete_all" {
			db.Exec("TRUNCATE TABLE `lottery01`")
			c.JSON(200, gin.H{"msg": "所有数据已经被删除"})
			return
		} else if method == "query_all" {
			var results []Lottery01
			// 查询所有记录
			if err := db.Find(&results).Error; err != nil {
				c.JSON(500, gin.H{"msg": "查询出错"})
				return
			}

			// 格式化返回数据
			var formattedResults []map[string]interface{}
			for _, result := range results {
				formattedResults = append(formattedResults, map[string]interface{}{
					"id":      result.ID,
					"results": result.Results,
					"created": result.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间
				})
			}

			// 返回 JSON 响应
			c.JSON(200, gin.H{
				"msg":    "查询成功",
				"result": formattedResults,
			})
			return
		}
	})
	router.POST("/lottery/query", func(c *gin.Context) {

		num := c.GetHeader("num")
		if num == "" {
			c.JSON(400, gin.H{"msg": "请求头缺少 num 参数"})
			return
		}

		var result Lottery01
		// 使用 err 来捕获查询错误
		err := db.First(&result, num).Error

		if err != nil {
			// 如果查询失败，返回 404 和错误信息
			if err == gorm.ErrRecordNotFound {
				// 数据未找到
				c.JSON(404, gin.H{"msg": "没有找到相关记录"})
			} else {
				// 其他数据库错误
				c.JSON(500, gin.H{"msg": "查询失败，请稍后再试！"})
				fmt.Println("数据库查询错误:", err.Error())
			}
			return
		}

		// 查询成功，返回查询内容
		content := []interface{}{
			result.ID,
			":",
			result.Results,
		}
		c.JSON(200, gin.H{"msg": content})

	})

	router.POST("/lottery/delete", func(c *gin.Context) {
		num := c.GetHeader("num")
		if num == "" {
			c.JSON(400, gin.H{"msg": "请求头缺少 num 参数"})
			return
		}
		var result Lottery01
		// 使用 err 来捕获查询错误
		err := db.First(&result, num).Error

		if err != nil {
			// 如果查询失败，返回 404 和错误信息
			if err == gorm.ErrRecordNotFound {
				// 数据未找到
				c.JSON(404, gin.H{"msg": "没有找到相关记录"})
			} else {
				// 其他数据库错误
				c.JSON(500, gin.H{"msg": "查询失败，请稍后再试！"})
				fmt.Println("数据库查询错误:", err.Error())
			}
			return
		}
		db.Delete(&result)
		c.JSON(200, gin.H{"msg": "删除成功"})

	})

	router.Run(":8080")
}
