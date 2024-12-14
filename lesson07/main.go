package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func LotterySystem(db *gorm.DB) {
	//1.生成随机数，选择奖品
	num := rand.Intn(2)
	var result string
	switch num {
	case 0:
		fmt.Println("很遗憾，您没有中奖")
		result = "NO"
	case 1:
		fmt.Println("恭喜，您中奖了！")
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

	//2.创建数据库，结构体
	//*********修改这下面一行以连接自己的数据库**********
	dsn := "root:G794q028@tcp(127.0.0.1:3306)/game?charset=utf8&mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Lottery01{})
	fmt.Println("******************抽奖系统******************")

	for {
		fmt.Println("请输入您想进行的功能（1.抽奖）（2.查询）（3.删除）（4.退出）（5.其他操作）：")
		var input int
		fmt.Scan(&input)
		//输入1进行抽奖
		if input == 1 {

			LotterySystem(db)

		} else if input == 2 {
			//4.实现通过id查询中奖

			fmt.Println("请输入你要查询第几次的抽奖结果：")
			var n int
			fmt.Scan(&n)
			var result Lottery01
			//var result []Lottery01
			db.First(&result, n)
			if result.ID == 0 || db.Error != nil {
				fmt.Println("****没有找到记录****")
			} else {
				fmt.Println(result.ID, ":", result.Results) // 打印找到的记录
			}
		} else if input == 3 {
			//5.实现通过id删除抽奖结果
			fmt.Println("请输入你要删除第几次的抽奖结果：")
			var n int
			fmt.Scan(&n)
			var result Lottery01
			db.First(&result, n)

			if result.ID == 0 || db.Error != nil {
				fmt.Println("****没有找到记录****")
			} else {
				db.Delete(&result)
				fmt.Printf("删除成功，您删除了第%d次抽奖的结果\n", n)
			}

		} else if input == 4 {
			break
		} else if input == 5 {
			fmt.Println("（1.删除所有抽奖数据并退出）2.（查看所有抽奖记录）")
			var input02 int
			fmt.Scan(&input02)
			if input02 == 1 {
				// 删除整个表（包括表结构）
				db.Exec("DROP TABLE IF EXISTS `Lottery01`")
				fmt.Println("所有数据已经删除...")
				time.Sleep(3 * time.Second)
				break

			} else if input02 == 2 {
				var results []Lottery01     // 假设 Lottery01 是你的结构体类型
				result := db.Find(&results) // 查询所有记录并映射到 results 切片
				if result.Error != nil {
					fmt.Println("查询出错:", result.Error)
				} else {
					for _, r := range results {
						fmt.Printf("ID: %d, Results: %s\n", r.ID, r.Results)
					}
				}
			}

		}

	}

	//5.通过powershell进行选择新建抽奖，查询抽奖，删除抽奖
}
