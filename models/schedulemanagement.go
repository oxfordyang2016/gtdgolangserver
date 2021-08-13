package models

import (

	// "time"
	// "sort"
	//"github.com/bradfitz/slice"

	// "github.com/fatih/color"
	// "github.com/jinzhu/gorm"
	// //"strconv"
	// //"github.com/jinzhu/gorm"

	//"github.com/gin-contrib/sessions"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	// "github.com/tidwall/gjson"
	//"github.com/tidwall/gjson"
)

type Progressbar struct {
	gorm.Model
	Email string
	Name  string
	Url   string
	// 这里是表示进度
	Bar float64
	// 视频还是书籍还是音频
	Typeoftask string
}

// func CreateBar(name string, email string, url string, bar float64, Typeoftask string) {

// 	// var countofproject int
// 	var thisbar Progressbar
// 	// db.Where("Email= ?", email).Where("Project=?", projectname).Where("Goalcode=?", goalcode).Find(&project).Count(&countofproject)
// 	// // db.Where(&Projectofgoals{Goalcode: goalcode, Project: projectname, Email: email}).Count(&countofproject)
// 	// if countofproject > 0 {
// 	// 	c.JSON(802, gin.H{
// 	// 		"result": "u have created project",
// 	// 	})
// 	// 	return
// 	// }
// 	projectofgoal := Progressbar{Createtime: createtime, Goalcode: goalcode, Goal: goalname, Project: projectname, Email: email, Status: "unspecified"}
// 	db.Create(&projectofgoal).Scan(&projectofgoal)
// 	c.JSON(200, gin.H{
// 		"result": " created project successfully",
// 	})

// }
