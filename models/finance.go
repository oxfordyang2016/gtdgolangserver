package models
import(
  "fmt"
  "strings"
  "time"
  //"os/exec"
  //"os"
  //"net/http"
	//"strconv"
"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
"github.com/tidwall/gjson"
)



type(
    Fees  struct{
    gorm.Model
    Name                string
    Money               float64
    Date                string
    Direction           string
    Email               string
    }
  

  )



  



func Getmywealth (c *gin.Context){

  var balance = 1520
  c.JSON(200, gin.H{
						"status": "blog had updated",
						"blance":balance,
                })
}

//record spending   of everyday
func Create_fee(c *gin.Context){

    //---------------get body string-------------
    //https://github.com/gin-gonic/gin/issues/1295
    buf := make([]byte, 1000000)
    num, _ := c.Request.Body.Read(buf)
    reqBody := string(buf[0:num])
//--------------using gjson to parse------------
//https://github.com/tidwall/gjson

emailcookie,err:=c.Request.Cookie("email")
//fmt.Println(emailcookie.Value)
var email string
if err!=nil{
email = c.Request.Header.Get("email")
}else{
fmt.Println(emailcookie.Value)
email =emailcookie.Value
}



inbox := gjson.Get(reqBody, "inbox").String()
fmt.Println(inbox)
money := gjson.Get(reqBody, "money").Float()
direction:=  gjson.Get(reqBody, "direction").String()
date := gjson.Get(reqBody, "date").String()

if strings.Contains(date, "today"){
  // if plantime =="today"{
 loc, _ := time.LoadLocation("Asia/Shanghai")
 //plantimeofanotherforamt :=  time.Now().In(loc)
 //
 date=  time.Now().In(loc).Format("060102")
}
if strings.Contains(date, "tomorrow"){
// if plantime  =="tomorrow"{
 loc, _ := time.LoadLocation("Asia/Shanghai")
//https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
date =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
}

fee := Fees{Name:inbox,Email:email,Direction:direction,Money:money,Date:date}
 //db.Create(&task).Scan(&task)
 db.Create(&fee).Scan(&fee)
    fmt.Println("i am testing the id return")
fmt.Println(fee.Name)
  c.JSON(200, gin.H{
						"status": "blog had updated",
						"blance":12.3,
                })
}

//record the debt   


//record the income 









