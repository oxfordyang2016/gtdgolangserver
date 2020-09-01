package models
import(
  "fmt"
  "reflect"
  //"encoding/json"
  //"log"
  "time"
  "net/http"
	"strconv"
"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
"github.com/bradfitz/slice"
"strings"
//"github.com/go-redis/redis"
//"github.com/garyburd/redigo/redis"

"github.com/tidwall/gjson"

)
type JSON []byte
//var Modeltest int =5
type (
	// TodoModel describes a TodoModel type
	TodoModel struct {
		gorm.Model
    Title2     string `json:"title2"`
		Title1     string `json:"title1"`
		Title     string `json:"title"`
		Completed int    `json:"completed"`

	}

	// transformedTodo represents a formatted Todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

  Tasks struct {
    //gorm.Model this has set id ,cautous!!!http://jinzhu.me/gorm/models.html
		//ID        uint   `json:"id"`
    ID        uint    `gorm:"primary_key"`
    Task      string   `json:"task"`
    User     string `json:"user"`
    Email    string   `json:"email"`
    Place    string   `json:"place"`
    Status  string   `json:"status"`
    Project  string  `json:"project"`
    Plantime  string  `json:"plantime"`
    Finishtime  string `json:"finishtime"`
    Note        string `json:"note"`
    Parentproject  string `json:"parentproject"`
    Ifdissect  string     `json:"ifdissect"`
    AccurateFinishtime  string `json:"AccurateFinishtime"`
    Longitude string `json:"Longitude"`
    Latitude  string `json:"Latitude"`
    Reviewsign string `json:"reviewsign"`
    Score      uint    `json:"score"`
    Reviewdatas string  `json:"reviewdatas"`    


	}

 




     Reviewdata struct{
     "makeuseofthings":","learnewthings":"yes","battlewithlowerbrain":"yes"



     }


     Reviewofday  struct {
     gorm.Model
     Date string
     Scores  uint
     }





    Person struct {
          Name   string
          Emails []string
     }


//{"name":'yangming','children':[]}
     Thinkmapofreview  struct{
       Name                    string
       Children             []Thinkmapofreview
       }


  Projects  struct{
    Name                    string
    Alltasksinproject       []Tasks
    }

    Everyday  struct{
      Name                    string
      Alldays       []Tasks
      }

      Place  struct{
        Name                    string
        Allplaces              []Tasks
        }

 )








var longtitude = "24.24"
var latitude = "47.47"




 // createTodo add a new todo
 func Createtaskfromios(c *gin.Context) {
   emailcookie,_:=c.Request.Cookie("email")
   //fmt.Println(emailcookie.Value)
   email:=emailcookie.Value
   print("22222222222222222222")
   print(email)
   //email := c.PostForm("email")
   inbox := c.PostForm("inbox")
   fmt.Println(inbox)
    project := c.PostForm("project")
   fmt.Println("+++++++++++++++++")
   fmt.Println(project)  
   place := c.PostForm("place")
   fmt.Println("=============")
   fmt.Println(place) 
   plantime := c.PostForm("plantime")
    
        if plantime =="today"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
      //plantimeofanotherforamt :=  time.Now().In(loc)
      //
      plantime =  time.Now().In(loc).Format("060102")
    }
    if plantime  =="tomorrow"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
    }


   if plantime==""{
   plantime ="unspecified" 
} 
   plantime ="unspecified" 
} 
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
    }


   if plantime==""{
   plantime ="unspecified" 
} 
  status := c.PostForm("taskstatus")
   parentproject := c.PostForm("parentproject")
   note := c.PostForm("note")
   ifdissect := c.PostForm("ifdissect")

 if status!="unfinished"{
   clientfinishtime:=  c.PostForm("finishtime")
   fmt.Println("=================")
   fmt.Println(clientfinishtime)
   loc, _ := time.LoadLocation("Asia/Shanghai")
   finishtime :=  time.Now().In(loc)
   if clientfinishtime!="unspecified"{
   task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
   db.Create(&task)
   }else{
   task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
   db.Create(&task)
    }


 }else{
   task := Tasks{Task:inbox,User:email,Finishtime:"unfinished",Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
   db.Create(&task)
 }
 c.JSON(200, gin.H{
     "status":  "posted",
     "message": "u have uploaded info,please come on!",
   })
 	}

















func Googlemapservice(c *gin.Context) {
if latitude == ""{
c.JSON(200, gin.H{
     "lat":  "47.47",
      "long": "23.23",
     "message": "u have uploaded info,please come on!",
   })


} else{


c.JSON(200, gin.H{
     "lat":  latitude,
      "long": longtitude,
     "message": "u have uploaded info,please come on!",
   })


}

}










// createTodo add a new todo
func Createtask(c *gin.Context) {
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  inbox := c.PostForm("inbox")
  fmt.Println(inbox)
  project := c.PostForm("project")
  place := c.PostForm("place")
  plantime := c.PostForm("plantime")
  long :=  c.PostForm("long")
  lat :=  c.PostForm("lat") 
  fmt.Println("+++++++++++++++place info +++++++++++++")   
  longtitude = long
  latitude = lat  
  fmt.Println(long)
  fmt.Println(lat)
  fmt.Println("+++++++++++++++place info +++++++++++++")
  if strings.Contains(plantime, "today"){
       // if plantime =="today"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
      //plantimeofanotherforamt :=  time.Now().In(loc)
      //
      plantime =  time.Now().In(loc).Format("060102")
    }
    if strings.Contains(plantime, "tomorrow"){
     // if plantime  =="tomorrow"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
    }






  status := c.PostForm("taskstatus")
   //status := c.PostForm("taskstatus")
    if status == "f"{status = "finish"}
    if status  == "g"{status = "giveup"}
    if status   == "r"{status =  "replace"}
    if status  == "a"{status = "anotherday"}
  parentproject := c.PostForm("parentproject")
  note := c.PostForm("note")
  ifdissect := c.PostForm("ifdissect")

if status!="unfinished"{
  clientfinishtime:=  c.PostForm("finishtime")
  if clientfinishtime == "y"{
   clientfinishtime = "yesterday"

}

  fmt.Println("=================")
  fmt.Println(clientfinishtime)
  loc, _ := time.LoadLocation("Asia/Shanghai")
  finishtime :=  time.Now().In(loc)
  if clientfinishtime!="unspecified"{
     if strings.Contains(clientfinishtime, "yesterday"){   
     // if clientfinishtime == "yesterday"{
   loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    clientfinishtime =  time.Now().In(loc).AddDate(0, 0,-1).Format("060102")

    }
  task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
  db.Create(&task)
  }else{
  task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
  db.Create(&task)
   }


}else{
  task := Tasks{Task:inbox,User:email,Finishtime:"unfinished",Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
  db.Create(&task)
}
c.JSON(200, gin.H{
    "status":  "posted",
    "message": "u have uploaded info,please come on!",
  })
	}


  // createTodo add a new todo
  func Update(c *gin.Context) {


    //---------------get body string-------------
    //https://github.com/gin-gonic/gin/issues/1295
     buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
   
     
   //--------------using gjson to parse------------
   //https://github.com/tidwall/gjson
   value := gjson.Get(reqBody, "reviewdata")
	fmt.Println(value.String())






    emailcookie,_:=c.Request.Cookie("email")
    fmt.Println(emailcookie.Value)
    email:=emailcookie.Value
    //totalscores:= c.PostForm("totalscoresofsingletask")
    totalscores,_ := strconv.Atoi(c.PostForm("totalscoresofsingletask"))
    fmt.Println("--------------------single task  total scores----------------")
    fmt.Println(totalscores)
    reviewdata := c.PostForm("reviewdata")
    fmt.Println(reviewdata)
    fmt.Println(reflect.TypeOf(reviewdata))
    //---------------------------review algorithms data-------------------------
    inbox := c.PostForm("inbox")
    place := c.PostForm("place")
    fmt.Println(inbox)
    id := c.PostForm("id")
    project := c.PostForm("project")
    finishtime := c.PostForm("finishtime")
    if finishtime == "y"{finishtime = "yesterday"}
    if strings.Contains(finishtime, "yesterday"){
    //if finishtime == "yesterday"{
    
    loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    finishtime =  time.Now().In(loc).AddDate(0, 0,-1).Format("060102")

    }
    
    if finishtime == "today"{

    loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    finishtime =  time.Now().In(loc).Format("060102")

    }





    plantime := c.PostForm("plantime")
    
       if strings.Contains(plantime, "today"){
       // if plantime =="today"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
      //plantimeofanotherforamt :=  time.Now().In(loc)
      //
      plantime =  time.Now().In(loc).Format("060102")
    }
    if strings.Contains(plantime, "tomorrow"){ 
   // if plantime  =="tomorrow"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang 
    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
    }
   






    status := c.PostForm("taskstatus")
    if status == "f"{status = "finish"}
    if status  == "g"{status = "giveup"}
    if status   == "r"{status =  "replace"}
    if status  == "a"{status = "anotherday"}


    parentproject := c.PostForm("parentproject")
    note := c.PostForm("note")
    //status := c.PostForm("taskstatus")
    ifdissect := c.PostForm("ifdissect")
    fmt.Println(status,plantime,project,id,inbox,email)
    var task Tasks
    db.Where("Email= ?", email).First(&task, id)
    fmt.Println(task)
    fmt.Println(task.Email)
    if task.Email!=email{
      c.JSON(200, gin.H{
          "status":  "posted",
          "message": "updated id not exsit",
        })
        //using python design method to return none
      return
    }

     //update a task
    db.Model(&task).Update("reviewdatas", value.String())
    if place!="unspecified"{db.Model(&task).Update("Place", place)}
    if project!="inbox"{db.Model(&task).Update("Project", project)}
    if inbox!="nocontent"{db.Model(&task).Update("Task", inbox)}
    if plantime!="unspecified"{db.Model(&task).Update("Plantime", plantime)}
    if parentproject!="unspecified"{db.Model(&task).Update("Parentproject", parentproject)}
    if ifdissect!="no"{db.Model(&task).Update("Ifdissect", ifdissect)}
    if note!="unspecified"{db.Model(&task).Update("Note", note)}
    if totalscores!=0{
    fmt.Println("--------------------single task  total scores----------------")
    fmt.Println(totalscores)
     db.Model(&task).Update("reviewdatas", value.String())
    db.Model(&task).Update("Score", totalscores)}
    //using it to format time https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
    if status!="unfinished"{
     //locate timezone https://stackoverflow.com/questions/27991671/how-to-get-the-current-timestamp-in-other-timezones-in-golang
      loc, _ := time.LoadLocation("Asia/Shanghai")
      now :=  time.Now().In(loc)

     db.Model(&task).Update("Finishtime",now.Format("060102"))
      //now1 :=  time.Now().In(loc)
      //db.Model(&task).Update("AccurateFinishtime",now1.String()）
      db.Model(&task).Update("Status", status)}
   
    if finishtime!="unspecified"{db.Model(&task).Update("Finishtime", finishtime)}

 c.JSON(200, gin.H{
  			"status":  "posted",
  			"message": "123",
  			"nick": "234",
  		})
  	}




      func  Canvas(c *gin.Context) {
      /*  c.HTML(http.StatusOK, "inbox.html",gin.H{
          "task":"ha",
        })
*/
       fmt.Println("hahhhahhah============")
          		c.HTML(http.StatusOK, "map.html",nil)
       }






    // createTodo add a new todo
    func Test(c *gin.Context) {
      c.JSON(200, gin.H{
    			"status":  "conected........",
    			"message": "welcome to new world",
    			"nick": "234",
    		})
    	}

/*
      func  Mainboard(c *gin.Context) {
          		c.HTML(http.StatusOK, "mainboard.html",nil)
       }

*/


// createTodo add a new todo
 func Updateforios(c *gin.Context) {

   email := c.PostForm("email")
   inbox := c.PostForm("inbox")
   place := c.PostForm("place")
   fmt.Println(inbox)
   id := c.PostForm("id")
   project := c.PostForm("project")
   finishtime := c.PostForm("finishtime")
   plantime := c.PostForm("plantime")

       if plantime =="today"{
     loc, _ := time.LoadLocation("Asia/Shanghai")
     //plantimeofanotherforamt :=  time.Now().In(loc)
     //
     plantime =  time.Now().In(loc).Format("060102")
   }
   if plantime  =="tomorrow"{
     loc, _ := time.LoadLocation("Asia/Shanghai")
   //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
   plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
   }
   status := c.PostForm("taskstatus")
   parentproject := c.PostForm("parentproject")
   note := c.PostForm("note")
   //status := c.PostForm("taskstatus")
   ifdissect := c.PostForm("ifdissect")
   fmt.Println(status,plantime,project,id,inbox,email)
   var task Tasks
   db.Where("Email= ?", email).First(&task, id)
   fmt.Println(task)
   fmt.Println(task.Email)
   if task.Email!=email{
     c.JSON(200, gin.H{
         "status":  "posted",
         "message": "updated id not exsit",
       })
       //using python design method to return none
     return
   }
   if place!="unspecified"{db.Model(&task).Update("Place", place)}
   if project!="inbox"{db.Model(&task).Update("Project", project)}
   if inbox!="nocontent"{db.Model(&task).Update("Task", inbox)}
if plantime!="unspecified"{db.Model(&task).Update("Plantime", plantime)}
if parentproject!="unspecified"{db.Model(&task).Update("Parentproject", parentproject)}
if ifdissect!="no"{db.Model(&task).Update("Ifdissect", ifdissect)}
if note!="unspecified"{db.Model(&task).Update("Note", note)}
//using it to format time https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
if status!="unfinished"{
 //locate timezone https://stackoverflow.com/questions/27991671/how-to-get-the-current-timestamp-in-other-timezones-in-golang
  loc, _ := time.LoadLocation("Asia/Shanghai")
  now :=  time.Now().In(loc)

 db.Model(&task).Update("Finishtime",now.Format("060102"))
  //now1 :=  time.Now().In(loc)
  //db.Model(&task).Update("AccurateFinishtime",now1.String()）
  db.Model(&task).Update("Status", status)}

if finishtime!="unspecified"{db.Model(&task).Update("Finishtime", finishtime)}

c.JSON(200, gin.H{
                    "status":  "posted",
                    "message": "123",
                    "nick": "234",
            })
    }






   












func Inbox(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  
  //in order to write a simple function,i design to add header field in request header
  client:= c.Request.Header.Get("client")
  fmt.Println("+++++++client is++++++++")  
  fmt.Println(client)
  fmt.Println("+++++++client is++++++++")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //use http://doc.gorm.io/crud.html#query
  //the next line is for all tasks 
 //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
 //the next line is for all unfinished task 
 db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)

  //fmt.Println(tasks)
  //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
  //  looptest := "string"
  //fmt.Println(looptest)
  


  //try to STATISTICS for gtd
  //http://doc.gorm.io/crud.html#query
  /*
  db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
//// SELECT count(*) FROM deleted_users;


  */
/*
not reference

db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");



*/



var countofalltasks  int
var countoffinishedtasks  int
var countforunfinishedtasks int
var finishedrate   float64

//reference http://doc.gorm.io/crud.html#query  query with condition
//db.Table("tasks").Where("status = ?", "finish").Count(&countofalltasks)
//db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//i use SELECT COUNT(CustomerID), Country FROM Customers GROUP BY Country; to verify which status  items are there?
db.Table("tasks").Where("Email= ?", email).Where("status = ?","unfinished").Or("status = ?","unfinish").Count(&countforunfinishedtasks)
db.Table("tasks").Where("Email= ?", email).Not("status", []string{"unfinished","unfinish"}).Count(&countoffinishedtasks)//reference not keyword
db.Table("tasks").Where("Email= ?", email).Count(&countofalltasks)

fmt.Println("+++++++++++++")
fmt.Println(countofalltasks)
fmt.Println(countoffinishedtasks)
//https://stackoverflow.com/questions/32815400/how-to-perform-division-in-go
finishedrate = float64(countoffinishedtasks)/float64(countofalltasks)
//strconv.FormatFloat(finishedrate, 'f', -1, 64)
fmt.Println("%.6f",finishedrate)
var finishedratebyend string
//https://gobyexample.com/string-formatting
finishedratebyend = fmt.Sprintf("%.6f", finishedrate)
fmt.Println("+++++++++++++")







c.HTML(http.StatusOK, "inbox.html",gin.H{
   "task":tasks,"finishedrate":finishedratebyend,"countforunfinishedtasks":countforunfinishedtasks,
  })
  }





func Inboxjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
    db.Where("Email= ?", email).Order("id desc").Find(&tasks)

  c.JSON(200, gin.H{
      "task":tasks,
    })

}



func Unfinishedtaskjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
         loc, _ := time.LoadLocation("Asia/Shanghai")
     today :=  time.Now().In(loc).Format("060102")
     tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
   
  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today,tomorrow}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}



func Reviewsjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
  // loc, _ := time.LoadLocation("Asia/Shanghai")
  // today :=  time.Now().In(loc).Format("060102")
  // tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")

  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"review"}).Order("id desc").Find(&tasks)
   if 2<1{
   c.JSON(200, gin.H{
      "task":tasks,
    })

   }
  c.HTML(http.StatusOK, "reviewfromprojectreview.html",gin.H{
   "task":tasks,
  })




}















func Readinglistjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)


  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"readinglist ", "readinglist"}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}






func Habitlistjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)


  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"habit ", "habit","Habit"}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}




func Healthlistjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)


  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"health ", "health","Health"}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}







func Financelistjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)


  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"finance ", "finance","finnace","finnace "}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}















func Todaytaskjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)

     loc, _ := time.LoadLocation("Asia/Shanghai")
  now :=  time.Now().In(loc)
 
 //db.Model(&task).Update("Finishtime",now.Format("060102"))

  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("plantime = ?",now.Format("060102")).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}














    func Review(c *gin.Context) {
      //i use email as identifier
    //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      fmt.Println(email)
      //build search algorithm to get project relationship
      /*
      1.set root project be dm
      2.select datastucture to store
      3.fetch every line to add --------



      */

/*
      var tasks []Tasks
      //email:="yangming1"
      db.Where("Email= ?", email).Find(&tasks)
      alldays:=make(map[string] []Tasks)
      make(map[string]  []string)//{"na
      for _,item :=range tasks{
         alldays[item.Plantime]=append(alldays[item.Plantime],item)
         //alldays[item.Finishtime]=append(alldays[item.Finishtime],item)
      }
      var alleverydays []Everyday
      for k,v := range alldays{
         alleverydays =append(alleverydays,Everyday{k,v})
      }
*/






      c.HTML(http.StatusOK, "review.html", nil)
      	}




func Everydays(c *gin.Context) {
      //i use email as identifier
    //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      client:= c.Request.Header.Get("client")

      //fmt.Println(cookie1.Value)
      var tasks []Tasks
      //email:="yangming1"
       
     if client == "ios"{


      loc, _ := time.LoadLocation("Asia/Shanghai")
     today :=  time.Now().In(loc).Format("060102")
     tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")







      db.Where("Email= ?", email).Where("plantime in (?)", []string{today,tomorrow}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks) 
      }else{
     db.Where("Email= ?", email).Find(&tasks)
      }
       alldays:=make(map[string] []Tasks)
      for _,item :=range tasks{
     
      fmt.Println("++++++++there is a null value+++++++++++++++++++")
      fmt.Println(item.Plantime)
     fmt.Println("++++++++there is a null value+++++++++++++++++++")
    alldays[item.Plantime]=append(alldays[item.Plantime],item)
         //alldays[item.Finishtime]=append(alldays[item.Finishtime],item)
      }
      



var alleverydays []Everyday
      var unspecifiedday  Everyday
      for k,v := range alldays{
if k!="unspecified" {   
alleverydays =append(alleverydays,Everyday{k,v})
      }

      if k=="unspecified"{
       //alleverydays =append(alleverydays,Everyday{k,v})
       unspecifiedday = Everyday{k,v}
    }


}

      slice.Sort(alleverydays, func(i, j int) bool {
return alleverydays[i].Name > alleverydays[j].Name
})


  if unspecifiedday.Name!=""{ 
  alleverydays =append(alleverydays,unspecifiedday)
 }





      fmt.Println("====================")
      k:=alleverydays[0].Alldays
      fmt.Println(k[0].ID)
      fmt.Println(k[0])
      fmt.Println("=====================")

      fmt.Println(alleverydays)
      //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
      //  looptest := "string"
      //fmt.Println(looptest)
     
     if client == "ios"{
     c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "plans": alleverydays})
     }else{
      c.HTML(http.StatusOK, "everyday.html",gin.H{
       "alldays":alleverydays,
      })

 //c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "plans": alleverydays})
}

      	}













        func Finished(c *gin.Context) {
              
            //i use email as identifier
            //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
            
            start := time.Now()
              emailcookie,_:=c.Request.Cookie("email")
              fmt.Println(emailcookie.Value)
              email:=emailcookie.Value
               client:= c.Request.Header.Get("client")
            fmt.Println("+++++++client is++++++++")
           fmt.Println(client)
           fmt.Println("+++++++client is++++++++")




              //fmt.Println(cookie1.Value)
              var tasks []Tasks
              //email:="yangming1"
              db.Where("Email= ?", email).Find(&tasks)
              alldays:=make(map[string] []Tasks)
              
          getdbdatatime := time.Now()
          fmt.Println("from request arrive to finished get data from db  time ")
        //https://stackoverflow.com/questions/45791241/correctly-measure-time-duration-in-go 
         fmt.Println(getdbdatatime.Sub(start))
         fmt.Println("++++++++")



         for _,item :=range tasks{
                // alldays[item.Plantime]=append(alldays[item.Plantime],item)
          if item.Status!="unfinished"{
            if item.Status!="unfinish"{
            alldays[item.Finishtime]=append(alldays[item.Finishtime],item)}}
              }
              var alleverydays []Everyday
              var daybefore180119 Everyday
              var forgotten  Everyday
              for k,v := range alldays{
               if k!="180119before"&&k!="forgotten"{
                 alleverydays =append(alleverydays,Everyday{k,v})
              }
                 if k=="180119before"{
                   daybefore180119 = Everyday{k,v}
                 }
                 if k=="forgotten"{
                 forgotten  = Everyday{k,v}
                 
}



}

//https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field

              slice.Sort(alleverydays, func(i, j int) bool {
       return alleverydays[i].Name > alleverydays[j].Name
   })
             alleverydays =append(alleverydays,daybefore180119)
             alleverydays =append(alleverydays,forgotten)

          operatestructtime := time.Now()
          fmt.Println("from request arrive to finished operate struct time ")
          fmt.Println(operatestructtime.Sub(start))
          fmt.Println("++++++++")
            /*  fmt.Println("====================")
              k:=alleverydays[0].Alldays
              fmt.Println(k[0].ID)
              fmt.Println(k[0])
              fmt.Println("=====================")
             */
              //fmt.Println(alleverydays)
              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
              //  looptest := "string"
              //fmt.Println(looptest)
        var countforfinishedtasks  int
db.Table("tasks").Where("Email= ?", email).Not("status", []string{"unfinished","unfinish"}).Count(&countforfinishedtasks)//reference not keyword

/**
conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer conn.Close()




     
    // add some keys
    if _, err = conn.Do("SET", "k1", "a"); err != nil {
        log.Fatal(err)
    }
    if _, err = conn.Do("SET", "k2", "b"); err != nil {
        log.Fatal(err)
    }
    
    
//https://itnext.io/storing-go-structs-in-redis-using-rejson-dab7f8fc0053

    // for fun, let's leave k3 non-existing

    // get many keys in a single MGET, ask redigo for []string result
    strs, err := redis.Strings(conn.Do("MGET", "k1", "k2", "k3"))
    if err != nil {
        log.Fatal(err)
    }

    // prints [a b ]
    fmt.Println(strs)
   

  b, err := json.Marshal(&alleverydays)
if err != nil {
    return
}


_, err = conn.Do("SET", "testmemories", string(b))
if err != nil {
    return
}


objStr, err := redis.String(conn.Do("GET", "testmemories"))
if err != nil {
    return
}
 databyte:= []byte(objStr)
var alleverydaysfromredis  []Everyday

datafromredis := &alleverydaysfromredis
err = json.Unmarshal(databyte, datafromredis)
if err != nil {
    return
}

fmt.Println(datafromredis)
**/

            if client   == "ios"{


        // c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": datafromredis})
         
//u need to note,the fellowing code ,u can test theperformance of  dealing with json ,when i set different size.
c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays[0:30]})
                returntoclienttime := time.Now()
          fmt.Println("from request arrive to the end time of return json to client ")
        //https://stackoverflow.com/questions/45791241/correctly-measure-time-duration-in-go
         fmt.Println(returntoclienttime.Sub(start))
         fmt.Println("++++++++")   

 }else{

          //   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays})


             c.HTML(http.StatusOK, "pride.html",gin.H{
               "alldays":alleverydays,"countforfinishedtasks":countforfinishedtasks,
              })
returntowebtime := time.Now()
          fmt.Println("from request arrive to the end time of return html to client ")
        //https://stackoverflow.com/questions/45791241/correctly-measure-time-duration-in-go
         fmt.Println(returntowebtime.Sub(start))
         fmt.Println("++++++++")


              	}}






















  func Placebased(c *gin.Context) {
  //i use email as identifier
  //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

 //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  db.Where("Email= ?", email).Find(&tasks)
  allplaces:=make(map[string] []Tasks)
                              for _,item :=range tasks{
                                // alldays[item.Plantime]=append(alldays[item.Plantime],item)
                            //  if item.Place!="unspecified"{
                        //  if item.Status!="unfinish"{
                          //  alldays[item.Finishtime]=append(alldays[item.Finishtime],item)}}
                             allplaces[item.Place]=append(allplaces[item.Place],item)
                             //}
                           }
                             /*var alleverydays []Everyday
                             for k,v := range alldays{
                                alleverydays =append(alleverydays,Everyday{k,v})
                             }
*/
                              var places []Place
                              for k,v := range allplaces{
                                 places =append(places,Place{k,v})
                              }

                              slice.Sort(places, func(i, j int) bool {
                       return places[i].Name < places[j].Name
                   })






                            /*  fmt.Println("====================")
                              k:=alleverydays[0].Alldays
                              fmt.Println(k[0].ID)
                              fmt.Println(k[0])
                              fmt.Println("=====================")
                             */
                             fmt.Println(places)
                              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
                              //  looptest := "string"
                              //fmt.Println(looptest)
                              c.HTML(http.StatusOK, "place.html",gin.H{
                               "places":places,
                              })
                              	}










func Project(c *gin.Context) {

  //the algorithm can be upgrade
              //i use email as identifier
            //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      var tasks []Tasks
      //fmt.Println(cookie1.Value)
              //email:="yangming1"
      db.Where("Email= ?", email).Find(&tasks)
      var projects []string
      for _, item := range tasks {

        projects = append(projects,item.Project)
       }
    //get only project
     var onlyprojects []string
     onlyprojects=append(onlyprojects,projects[0])
     for _,item :=range projects{
         piot:="no"
         for _,item1 :=range onlyprojects{
           if item == item1{piot="yes"}
         }
         if piot=="no"{onlyprojects=append(onlyprojects,item)}
     }
 
    fmt.Println("--------------")
     fmt.Println(onlyprojects)

    //use maps to aviod to design complex algorithm
     allclassproject:=make(map[string] []Tasks)
     for _,item :=range tasks{
        allclassproject[item.Project]=append(allclassproject[item.Project],item)
     }
     var allprojects []Projects
     for k,v := range allclassproject{
        allprojects =append(allprojects,Projects{k,v})

     }

     slice.Sort(allprojects, func(i, j int) bool {
return allprojects[i].Name < allprojects[j].Name
})



     fmt.Println(allclassproject["gtd1"])


              //fmt.Println(tasks)
              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
              //  looptest := "string"
              //fmt.Println(looptest)
      c.HTML(http.StatusOK, "project.html",gin.H{
               "projects":allprojects,
        })
      }







func Projectsjson(c *gin.Context) {

  //the algorithm can be upgrade
              //i use email as identifier
            //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      var tasks []Tasks
      //fmt.Println(cookie1.Value)
              //email:="yangming1"
     // db.Where("Email= ?", email).Find(&tasks)
     loc, _ := time.LoadLocation("Asia/Shanghai")
     today :=  time.Now().In(loc).Format("060102")
     tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")

     


db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime in (?)", []string{today, tomorrow}).Order("id desc").Find(&tasks)
  client:= c.Request.Header.Get("client")
  fmt.Println("+++++++client is++++++++")
  fmt.Println(client)
  fmt.Println("+++++++client is++++++++")


      var projects []string
      for _, item := range tasks {

        projects = append(projects,item.Project)
       }
    //get only project
     var onlyprojects []string
     onlyprojects=append(onlyprojects,projects[0])
     for _,item :=range projects{
         piot:="no"
         for _,item1 :=range onlyprojects{
           if item == item1{piot="yes"}
         }
         if piot=="no"{onlyprojects=append(onlyprojects,item)}
     }

    fmt.Println("--------------")
     fmt.Println(onlyprojects)

    //use maps to aviod to design complex algorithm
     allclassproject:=make(map[string] []Tasks)
     for _,item :=range tasks{
        allclassproject[item.Project]=append(allclassproject[item.Project],item)
     }
     var allprojects []Projects
     for k,v := range allclassproject{
        allprojects =append(allprojects,Projects{k,v})

     }

     slice.Sort(allprojects, func(i, j int) bool {
return allprojects[i].Name < allprojects[j].Name
})



     fmt.Println(allclassproject["gtd1"])
c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "projects": allprojects})

              //fmt.Println(tasks)
              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
              //  looptest := "string"
              //fmt.Println(looptest)
             // c.HTML(http.StatusOK, "project.html",gin.H{
             //  "projects":allprojects,
             //})
      }







// createTodo add a new todo
func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := TodoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []TodoModel
	var _todos []TransformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}










		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
