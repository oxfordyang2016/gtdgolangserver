package models

import (
	"fmt"

	// "github.com/fatih/color"
	// "encoding/json"
	// "net/http"
	// "github.com/jinzhu/gorm"
	// "strconv"
	"time"
	// // "math"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// //"github.com/gin-contrib/sessions"
	// //_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/lib/pq"
	// "github.com/bradfitz/slice"
	// "math"
	"github.com/tidwall/gjson"
	// "github.com/gomodule/redigo/redis"
)

type Projects struct {
	Name              string
	Alltasksinproject []Tasks
}

type Projectofgoals struct {
	gorm.Model
	Createtime string `json:"createtime" gorm:"default:'unsepcified'"`
	Goal       string `json:"goal"`
	Goalcode   string `json:"goalcode"`
	Project    string `json:"project"`
	Status     string `json:"status" gorm:"default:'unfinished'"`
	Email      string `json:"email"`
	// Sarttime   string `json:"starttime" gorm:"default:'unspecified'"`
	Endtime string `json:"endtime" gorm:"default:'unspecified'"`
	// Details string `json:"details" sql:"type:text;"`
}

func Createprojectofgoal(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	goalcode := gjson.Get(reqBody, "goalcode").String()
	var goalforquery Goalfordbs
	db.Where(&Goalfordbs{Goalcode: goalcode, Email: email}).First(&goalforquery)
	goalname := goalforquery.Name
	projectname := gjson.Get(reqBody, "projectname").String()
	// status := gjson.Get(reqBody, "status").String()
	// projectstatus := "unspecified"
	// marktime  :=  "unspecified"
	// 	if status=="unspecified"||status==""{
	// 		projectstatus = "unspecified"
	// 	}
	// 	if status == "finished"{
	// 		projectstatus = "finished"
	// 		loc, _ := time.LoadLocation("Asia/Shanghai")
	//    marktime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	// 	}
	//这里原来引号出了问题。。。。多出了一个空格
	loc, _ := time.LoadLocation("Asia/Shanghai")
	createtime := time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	var countofproject int
	var project Projectofgoals
	db.Where("Email= ?", email).Where("Project=?", projectname).Where("Goalcode=?", goalcode).Find(&project).Count(&countofproject)
	// db.Where(&Projectofgoals{Goalcode: goalcode, Project: projectname, Email: email}).Count(&countofproject)
	if countofproject > 0 {
		c.JSON(802, gin.H{
			"result": "u have created project",
		})
		return
	}
	projectofgoal := Projectofgoals{Createtime: createtime, Goalcode: goalcode, Goal: goalname, Project: projectname, Email: email, Status: "unspecified"}
	db.Create(&projectofgoal).Scan(&projectofgoal)
	c.JSON(200, gin.H{
		"result": " created project successfully",
	})
}

func Getprojectofgoal(email string) map[string][]string {
	var projects []Projectofgoals
	db.Where("Email= ?", email).Where("Status= ?", "unspecified").Find(&projects)
	var goalmapproject = make(map[string][]string)
	for i := 0; i < len(projects); i++ {
		// allplaces[item.Place] = append(allplaces[item.Place], item)
		goalmapproject[projects[i].Goal] = append(goalmapproject[projects[i].Goal], projects[i].Project)
	}

	color.Yellow("9999999999")
	fmt.Println(len(projects))
	fmt.Println(goalmapproject)
	return goalmapproject
}

func Updateprojectofgoal(c *gin.Context) {

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	// goalname := gjson.Get(reqBody, "goal").String()
	goalcode := gjson.Get(reqBody, "goalcode").String()
	fmt.Println(goalcode)
	projectname := gjson.Get(reqBody, "projectname").String()
	status := gjson.Get(reqBody, "projectstatus").String()
	AllowedStatus := [8]string{"f", "finished", "finish", "g", "giveup", "unfinish", "unfinished", "unspecified"}
	if !itemExists(AllowedStatus, status) {
		c.JSON(905, gin.H{
			"result": "status not allowed",
		})
		return
	}
	color.Yellow("------这是客户端的project status---------")
	fmt.Println(status)
	projectstatus := "unspecified"
	endtime := "unspecified"
	if status == "unspecified" || status == "" {
		projectstatus = "unspecified"
	}
	if status == "finished" {
		projectstatus = "finished"
		loc, _ := time.LoadLocation("Asia/Shanghai")
		endtime = time.Now().In(loc).Format("060102")
	}

	if status == "giveup" {
		projectstatus = "giveup"
		loc, _ := time.LoadLocation("Asia/Shanghai")
		endtime = time.Now().In(loc).Format("060102")
	}

	var projectfromclient Projectofgoals
	db.Where(&Projectofgoals{Goalcode: goalcode, Project: projectname, Email: email}).First(&projectfromclient)
	//   if goalname != "unspecified"{
	// 	projectfromclient.Goal = goalname
	//   }
	if projectname != "unspecified" {
		projectfromclient.Project = projectname
	}

	if projectstatus != "unspecified" {
		projectfromclient.Status = projectstatus
		projectfromclient.Endtime = endtime
	}
	color.Yellow("------我在这里检查传入的projetc信息---------")
	fmt.Println(projectstatus)
	db.Save(&projectfromclient)
	c.JSON(200, gin.H{
		"result": "u have update project",
	})
}

func Getallpojects(c *gin.Context) {
	// buf := make([]byte, 1024)
	// num, _ := c.Request.Body.Read(buf)
	// reqBody := string(buf[0:num])
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	all := Getprojectofgoal(email)
	c.JSON(200, gin.H{
		"result": all,
	})
}

func GetallLinks(c *gin.Context) {
	// buf := make([]byte, 1024)
	// num, _ := c.Request.Body.Read(buf)
	// reqBody := string(buf[0:num])
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	var tasks []Tasks
	db.Where("Email= ?", email).Where("Project=?", "videos").Find(&tasks)
	c.JSON(200, gin.H{
		"task": tasks,
	})
}
