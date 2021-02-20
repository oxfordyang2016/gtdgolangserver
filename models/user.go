package models

import (
	"fmt"

	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//"github.com/gin-contrib/sessions"

	"strings"

	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/tidwall/gjson"
)

var Yangming int
var db *gorm.DB
var redisDBcon redis.Conn
var testlogger *log.Logger
var crontab *cron.Cron

type (
	//when username use lowcase,the db will not include the items
	Accounts struct {
		gorm.Model
		Email                               string `json:"email"`
		Username                            string `json:"username"`
		Password                            string `json:"password"`
		DreamforfutureDeviceTokenforiphone  string `json:"DreamforfuturedeviceTokenforiphone"`
		DreamforfutureDeviceTokenforipad    string `json:"DreamforfuturedeviceTokenforipad"`
		ThinkingforlifeDeviceTokenforipad   string `json:"thinkingforlifedeviceTokenforipad "`
		ThinkingforlifeDeviceTokenforiphone string `json:"thinkingforlifedeviceTokenforiphone"`
	}
)

func init() {
	//open a db connection
	//var a =m add(2,3)
	//fmt.Println(a)
	// 首先创建一个日志文件
	file, fileerr := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileerr != nil {
		log.Fatal(fileerr)
	}
	testlogger = log.New(file, "测试日志: ", log.LstdFlags|log.Lshortfile)

	var err error
	crontab = cron.New()
	//mysql://dt_admin:dt2016@localhost/dreamteam_db
	db, err = gorm.Open("mysql", "root:123456@/dreamteam_db?charset=utf8&parseTime=True&loc=Local")
	db.BlockGlobalUpdate(true)
	// db.LogMode(true)
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
	//try
	//db, err := gorm.Open("sqlite3", "./yangming.sqlite")
	//defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Migrate the schema
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Principledetails{}, &Tasktagfortimescount{}, &Projectofgoals{}, &Taskexecutelog{}, &Principlecodewithtasktag{}, &Principlefordbs{}, &Accounts{}, &Tasks{}, &Reviewofday{}, &Reviewfortimescount{}, &Goalfordbs{}, &Fees{})
	//http://jinzhu.me/gorm/database.html#migration delete database table column
	//db.Model(&Tasks{}).DropColumn("Uer")

	redisDBcon, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}

	// パブリッシュ

}

var Code = make(map[string]int)

type EmailVerify struct {
	email string
	code  int
}

func SendEmail(email string) (int, error) {
	//smtp.PlainAuth()
	// 参数1：Usually identity should be the empty string, to act as username
	// 参数2：username
	//参数3：password
	//参数4：host
	auth := smtp.PlainAuth("", "thinking_for_life@163.com", "ELACZHTZLMLVJYDF", "smtp.163.com")
	to := []string{email}
	//发送随机数为验证码
	// Seed uses the provided seed value to initialize the default Source to a
	// deterministic state. If Seed is not called, the generator behaves as
	// if seeded by Seed(1). Seed values that have the same remainder when
	// divided by 2^31-1 generate the same pseudo-random sequence.
	// Seed, unlike the Rand.Seed method, is safe for concurrent use.
	rand.Seed(time.Now().Unix())
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	num := rand.Intn(10000)
	//发送内容使用base64 编码，单行不超过80字节，需要插入\r\n进行换行
	//The msg headers should usually include
	// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
	// messages is accomplished by including an email address in the to
	// parameter but not including it in the msg headers.
	str := fmt.Sprintf("From:thinking_for_life@163.com\r\nTo:%s\r\nSubject:注册验证码verifycode\r\n\r\nThinkingforlife\r\n验证码是%d \r\n感恩信任\r\n祝您身体健康，平安喜乐，成就伟大事业！", email, num) //邮件格式
	msg := []byte(str)
	err := smtp.SendMail("smtp.163.com:25", auth, "thinking_for_life@163.com", to, msg)

	if err != nil {
		log.Println("发送邮件出错")
		log.Println(err)
	}
	// 使用字符串进行存储

	var veryfyInfo = fmt.Sprintf("%s/%s/%s", email, strconv.Itoa(num), "time")
	set_status, err1 := redis.String(redisDBcon.Do("SET", email, veryfyInfo))
	log.Println(set_status)
	if err1 != nil {
		log.Println(err1)
		log.Println("设置信息出错")
		return -1, err1
	}

	return num, nil
}

func EmailGenerateCode(c *gin.Context) {
	buf := make([]byte, 1000000)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	//--------------using gjson to parse------------
	//https://github.com/tidwall/gjson
	Email := gjson.Get(reqBody, "email").String()
	// 这里需要检查验证码
	generatedNum, err := SendEmail(Email)
	if err != nil {

		log.Println("发送邮件出错了哟")
		c.JSON(http.StatusOK, gin.H{
			"info": " email error",
		})
	}
	log.Println(generatedNum)
	c.JSON(http.StatusOK, gin.H{
		"info": "code had been sent",
	})
}

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", nil)
}

func Register(c *gin.Context) {

	// fmt.Println("-----------------")
	// fmt.Println(Yangming)
	// Email := c.PostForm("email")
	// Passowrd := c.DefaultPostForm("password", "anonymous")
	// Username := c.PostForm("username")

	buf := make([]byte, 1000000)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	//--------------using gjson to parse------------
	//https://github.com/tidwall/gjson
	Email := gjson.Get(reqBody, "email").String()
	// 在这里检测email是否存在
	var numofaccounts = 0
	db.Table("accounts").Where("Email= ?", Email).Count(&numofaccounts)
	if numofaccounts > 0 {
		c.JSON(409, gin.H{
			"info": "该邮件已经被注册!",
		})
		return
	}

	Password := gjson.Get(reqBody, "password").String()
	// Username := gjson.Get(reqBody, "username").String()
	verifycode := gjson.Get(reqBody, "verifycode").String()
	client := gjson.Get(reqBody, "client").String()
	savedveryinfo, readerr := redis.String(redisDBcon.Do("GET", Email))
	if readerr != nil {
		log.Println(readerr)
		log.Println("redis 读取数据发生错误")
		// return -1, readerr
	}
	var info = strings.Split(savedveryinfo, "/")
	if verifycode != info[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"info": "验证码不正确!",
		})
		return
	}
	log.Println("读取reids的邮件验证保存信息", info)

	//发送email到用户，要求用户进行验证然后

	User1 := Accounts{Email: Email, Username: Email, Password: Password}
	fmt.Println(Email, Password, Email)
	fmt.Println(User1)
	// db, _ = gorm.Open("mysql", "dt_admin:dt2016@/dreamteam_db?charset=utf8&parseTime=True&loc=Local")
	db.Save(&User1)
	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "register ok!",
	// })
	// return
	token, err := GenarateJwt(Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"info": "服务器内部错误",
		})
		return
	}
	if client == "clientv2" {
		cookie := &http.Cookie{
			Name:  "token",
			Value: token,
		}
		http.SetCookie(c.Writer, cookie)
		//最好在这里设置token信息
		c.JSON(http.StatusOK, gin.H{
			"info": "register ok!",
		})
		return
	}
	c.HTML(http.StatusOK, "user.html", nil)
}

func Login1(c *gin.Context) {
	//cookie set
	//store := sessions.NewCookieStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))
	buf := make([]byte, 1000000)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	email := c.PostForm("email")
	fmt.Println(email)
	password := c.PostForm("password")
	client := c.PostForm("client")
	devicetoken := c.PostForm("devicetoken")
	devicetype := c.PostForm("devicetype")

	datatype := c.Request.Header.Get("Content-Type")
	if strings.Contains(datatype, "application/json") {
		email = gjson.Get(reqBody, "email").String()
		fmt.Println(email)
		password = gjson.Get(reqBody, "password").String()
		client = gjson.Get(reqBody, "client").String()
	}

	//    如果传上来的数据是application/json

	//首先查询然后，更新
	//   这里仅考虑单个设备的情况
	var userfromdb Accounts
	db.Where("email = ?", email).First(&userfromdb)

	fmt.Println(userfromdb)
	if client == "dreamforlifeforiphone" {
	}
	if client == "dreamforlifeforiphone" {
	}
	if client == "thinkingforlife" {
		if devicetype == "ipad" {
			if userfromdb.ThinkingforlifeDeviceTokenforipad != devicetoken {
				fmt.Println("-----thinkingforlife  with ipad login--------")
				db.Model(&userfromdb).Update("ThinkingforlifeDeviceTokenforipad", devicetoken)
			}
		}
	}
	if client == "thinkingforlifeforiphone" {
	}
	if client == "mapforlifeformac" {
	}
	if client == "mapforlifeforwindows" {
	}

	fmt.Println("================================")
	fmt.Println(userfromdb.Password)

	if userfromdb.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{
			"info": "password or email error!",
		})
		return
	}

	fmt.Println(client)
	//session := sessions.Default(c)
	//session.Set("count", "yangming")
	//session.Save()
	fmt.Println(email, password, client)
	cookie := &http.Cookie{
		Name:  "username",
		Value: email,
	}
	http.SetCookie(c.Writer, cookie)
	cookie1 := &http.Cookie{
		Name:   "email",
		Value:  email,
		Domain: "www.blackboxo.top",
		Path:   "/",
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
	if client == "web" {
		//https://github.com/gin-gonic/gin to redirect
		c.Redirect(http.StatusMovedPermanently, "/v1/inbox")
		//c.Redirect(http.StatusMovedPermanently, "/mainboard")
		//c.HTML(http.StatusOK, "user.html", nil)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"info": "logined",
		})
	}
}

func checkcookie() bool {
	return true
}

// Logout doc
// @Summary 登出用户系统
// @Tags 用户系统
// @Description 登出
// @Produce json
// @Success 200 {string} string ""
// @Router /logout  [get]
func Logout(c *gin.Context) {
	cookie1 := &http.Cookie{
		Name:  "token",
		Value: "",
	}
	http.SetCookie(c.Writer, cookie1)

	cookie2 := &http.Cookie{
		Name:  "email",
		Value: "",
	}
	http.SetCookie(c.Writer, cookie2)
	cookie3 := &http.Cookie{
		Name:  "username",
		Value: "",
	}
	http.SetCookie(c.Writer, cookie3)
	c.JSON(http.StatusOK, gin.H{
		"info": "logouted",
	})
}

// Login doc
// @Summary 登陆
// @Tags 用户系统
// @Description 登入
// @Accept json
// @Produce json
// @Param param body main.JSONParams true {"email":"yang75@qq.com","password":"1992"}
// @Success 200 {string} string "{"msg":"login success"}"
// @Failure 400 {string} string "{"msg": "user or password error"}"
// @Router /login [post]
func Login(c *gin.Context) {
	//cookie set
	//store := sessions.NewCookieStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	//https://github.com/gin-gonic/gin/issues/1295
	buf := make([]byte, 1000000)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	//--------------using gjson to parse------------
	//https://github.com/tidwall/gjson
	email := gjson.Get(reqBody, "email").String()
	password := gjson.Get(reqBody, "password").String()
	client := gjson.Get(reqBody, "client").String()

	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// client := c.PostForm("client")
	// devicetoken := c.PostForm("devicetoken")
	//首先查询然后，更新
	//   这里仅考虑单个设备的情况
	var userfromdb Accounts
	db.Where("email = ?", email).First(&userfromdb)
	// if (userfromdb.DeviceToken != devicetoken){
	// db.Model(&userfromdb).Update("Devicetoken", devicetoken)
	// }

	fmt.Println("================================")
	fmt.Println(userfromdb.Password)

	if userfromdb.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{
			"info": "password or email error!",
		})
		return
	}

	fmt.Println(client)
	//session := sessions.Default(c)
	//session.Set("count", "yangming")
	//session.Save()
	fmt.Println(email, password, client)
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

	token, err := GenarateJwt(email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"info": "服务器内部错误",
		})
		return
	}

	cookie4 := &http.Cookie{
		Name:  "token",
		Value: token,
	}
	http.SetCookie(c.Writer, cookie4)

	//c.String(http.StatusOK, "0")
	if client == "web" {
		//https://github.com/gin-gonic/gin to redirect
		// c.Redirect(http.StatusMovedPermanently, "/v1/inbox")
		// c.Redirect(http.StatusMovedPermanently, "http://baidu.com")
		c.JSON(http.StatusOK, gin.H{
			"info": "logined",
		})
		// c.Abort()
		//c.Redirect(http.StatusMovedPermanently, "/mainboard")
		//c.HTML(http.StatusOK, "user.html", nil)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"info": "logined",
		})
	}
}
