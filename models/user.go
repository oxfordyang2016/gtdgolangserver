package models

import (
"fmt"
"net/http"
"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"
"github.com/gomodule/redigo/redis"

)
var Yangming int
var db *gorm.DB
var redisDBcon  redis.Conn
type (
	//when username use lowcase,the db will not include the items
	Accounts struct {
		gorm.Model
		Email     string `json:"email"`
		Username     string `json:"username"`
		Password    string    `json:"password"`
	}
)

func init() {
	//open a db connection
	//var a =m add(2,3)
	//fmt.Println(a)
	var err error
	//mysql://dt_admin:dt2016@localhost/dreamteam_db
	db, err = gorm.Open("mysql", "dt_admin:dt2016@/dreamteam_db?charset=utf8&parseTime=True&loc=Local")
        
          //connect database to postgrel

        //postgrel database need to be set install set role and password
        //https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e
        //https://stackoverflow.com/questions/18715345/how-to-create-a-user-for-postgres-from-the-command-line-for-bash-automation
        //https://linuxize.com/post/how-to-install-postgresql-on-centos-7/
      /* 
       db1, err1 := gorm.Open("postgres", "host=127.0.0.1 port=5432  user=yangming  dbname=review password=123456 sslmode=disable")
	if err1 != nil {
	//	panic(err)
           fmt.Println(err1)
	}
        */


         //db, err := gorm.Open("sqlite3", "./yangming.sqlite")
	//defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Accounts{},&TodoModel{},&Tasks{},&Reviewofday{},&Reviewfortimescount{},&Goalfordbs{})
//http://jinzhu.me/gorm/database.html#migration delete database table column
 //db.Model(&Tasks{}).DropColumn("Uer")

   redisDBcon, err = redis.Dial("tcp", "localhost:6379")
   if err != nil {
	 panic(err)
   }


 // パブリッシュ
 

}





func  User(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html",nil)
 }

 func Register(c *gin.Context) {

	 fmt.Println("-----------------")
	 fmt.Println(Yangming)
 	 Email := c.PostForm("email")
 	 Passowrd:= c.DefaultPostForm("password", "anonymous")
 	 Username:= c.PostForm("username")
 	 User1 := Accounts{Email: Email,Username:Username,Password:Passowrd}
 	 fmt.Println(Email,Passowrd,Username)
 	 fmt.Println(User1)
	// db, _ = gorm.Open("mysql", "dt_admin:dt2016@/dreamteam_db?charset=utf8&parseTime=True&loc=Local")
 	 db.Save(&User1)
 	 c.HTML(http.StatusOK, "user.html", nil)
  }


func  Login(c *gin.Context) {
	   //cookie set
	  //store := sessions.NewCookieStore([]byte("secret"))
	  //router.Use(sessions.Sessions("mysession", store))
	    email := c.PostForm("email")
           password:= c.PostForm("password")
            client:=c.PostForm("client")
           
           var userfromdb   Accounts
           db.Where("email = ?", email).First(&userfromdb)
           fmt.Println("================================")
           fmt.Println(userfromdb.Password)
	   if  userfromdb.Password != password{
             c.JSON(http.StatusOK,  gin.H{
                        "status":  "password or email error!",
                })      
                 }            

          

		fmt.Println(client)
			//session := sessions.Default(c)
			//session.Set("count", "yangming")
			//session.Save()
			fmt.Println(email,password,client)
			cookie := &http.Cookie{
							Name:  "username",
							Value: email,
					}
			http.SetCookie(c.Writer, cookie)
			cookie1 := &http.Cookie{
							Name:  "email",
							Value: email,
					}
			http.SetCookie(c.Writer, cookie1)
			cookie2 := &http.Cookie{
							Name:  "logintime",
							Value: "now-nounspecify",
					}
			http.SetCookie(c.Writer, cookie2)
			cookie3 := &http.Cookie{
							Name:  "client",
							Value: client,
					}
			http.SetCookie(c.Writer, cookie3)
					//c.String(http.StatusOK, "0")
			if client == "web"{
				//https://github.com/gin-gonic/gin to redirect
				c.Redirect(http.StatusMovedPermanently, "/v1/inbox")
				//c.Redirect(http.StatusMovedPermanently, "/mainboard")
  		 //c.HTML(http.StatusOK, "user.html", nil)
		 }else{
			 c.JSON(http.StatusOK,  gin.H{
			"status":  "logined",
		})
		 }
  }

func  checkcookie() bool{
	return true
}
