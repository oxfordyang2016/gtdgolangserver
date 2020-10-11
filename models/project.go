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
"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
// //"github.com/gin-contrib/sessions"
// //_ "github.com/jinzhu/gorm/dialects/mysql"
// _ "github.com/jinzhu/gorm/dialects/postgres"
// _ "github.com/lib/pq"
// "github.com/bradfitz/slice"
// "math"
"github.com/tidwall/gjson"
// "github.com/gomodule/redigo/redis"
)




type Projectofgoals  struct {
	gorm.Model
	Createtime string   `json:"createtime"`
	Goal    string   `json:"goal"`     
	Project    string   `json:"project"`
	Status    string   `json:"status"`
	Email    string   `json:"email"`
	Marktime    string   `json:"marktime"`
	// Details string `json:"details" sql:"type:text;"`
 }





 func Createprojectofgoal(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie,_:=c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email:=emailcookie.Value
	goalname := gjson.Get(reqBody, "goal").String()
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
	createtime :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	var countofproject int
	db.Where(&Projectofgoals{Goal:goalname, Project: projectname,Email:email}).Count(&countofproject)
	if countofproject >0{
		c.JSON(200, gin.H{
			"result":"u have created project",
		  })
		  return
	}
	projectofgoal := Projectofgoals{Createtime:createtime,Goal:goalname,Project:projectname,Email:email,Status:"unspecified",Marktime:"unspecified"}
	db.Create(&projectofgoal).Scan(&projectofgoal)
	c.JSON(200, gin.H{
		"result":"u have created project",
	  })
}




func Updateprojectofgoal(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie,_:=c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email:=emailcookie.Value
	goalname := gjson.Get(reqBody, "goal").String()
	projectname := gjson.Get(reqBody, "projectname").String()
	status := gjson.Get(reqBody, "status").String()
	projectstatus := "unspecified"
	marktime  :=  "unspecified"
	if status=="unspecified"||status==""{
		projectstatus = "unspecified"
	}
	if status == "finished"{
		projectstatus = "finished"
		loc, _ := time.LoadLocation("Asia/Shanghai") 
   marktime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	}

	if status == "giveup"{
		projectstatus = "giveup"
		loc, _ := time.LoadLocation("Asia/Shanghai") 
   marktime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	}

	  var  projectfromclient Projectofgoals
	  db.Where(&Projectofgoals{Goal:goalname, Project: projectname,Email:email}).First(&projectfromclient)
	//   if goalname != "unspecified"{
	// 	projectfromclient.Goal = goalname
	//   } 
	  if projectname != "unspecified"{
        projectfromclient.Project = projectname
	  }  

	  
	  if projectstatus != "unspecified"{
		projectfromclient.Status  =projectstatus
		projectfromclient.Marktime = marktime
	  }
	  
	  db.Save(&projectfromclient)
	c.JSON(200, gin.H{
		"result":"u have created project",
	   })
}




