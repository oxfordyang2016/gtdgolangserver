package models
import(
  "fmt"
  "reflect"
  "github.com/fatih/color"
"regexp"
 "net/url"
 "io/ioutil"
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
    Userid    int    `json:"userid"`
    Task      string   `json:"task"`
    User     string `json:"user"`
    Email    string   `json:"email"`
    Place    string   `json:"place"`
    Status  string   `json:"status"`
    Goal  string   `json:"goal"`
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
    Devotedtime int    `json:"devotedtime"`
    Priority    int    `json:"priority"`
    Reviewdatas string  `json:"reviewdatas" sql:"type:text;"`    
    Tasktags string `json:"tasktags" sql:"type:text;"`
    Goalcoefficient float64   `json:"goalcoefficient"`
    Mark_finished_time_switch int  `json:"mark_finished_time_switch"` 
    First_finish_timestamp string   `json:"First_finish_timestamp"` 
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




//  // createTodo add a new todo
//  func Createtaskfromios(c *gin.Context) {
//    emailcookie,_:=c.Request.Cookie("email")
//    //fmt.Println(emailcookie.Value)
//    email:=emailcookie.Value
//    print("22222222222222222222")
//    print(email)
//    //email := c.PostForm("email")
//    inbox := c.PostForm("inbox")
//    fmt.Println(inbox)
//     project := c.PostForm("project")
//    fmt.Println("+++++++++++++++++")
//    fmt.Println(project)  
//    place := c.PostForm("place")
//    fmt.Println("=============")
//    fmt.Println(place) 
//    plantime := c.PostForm("plantime")
    
//         if plantime =="today"{
//       loc, _ := time.LoadLocation("Asia/Shanghai")
//       //plantimeofanotherforamt :=  time.Now().In(loc)
//       //
//       plantime =  time.Now().In(loc).Format("060102")
//     }
//     if plantime  =="tomorrow"{
//       loc, _ := time.LoadLocation("Asia/Shanghai")
//     //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
//     plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
//     }


//    if plantime==""{
//    plantime ="unspecified" 
// }

 
//  if   plantime !="unspecified" {
//     //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
//      loc, _ := time.LoadLocation("Asia/Shanghai") 
//    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
//     }


//    if plantime==""{
//    plantime ="unspecified" 
// } 

//   status := c.PostForm("taskstatus")
//    parentproject := c.PostForm("parentproject")
//    note := c.PostForm("note")
//    ifdissect := c.PostForm("ifdissect")

//  if status!="unfinished"{
  

//     clientfinishtime:=  c.PostForm("finishtime")
   
//   only_finishtime:= clientfinishtime

//    fmt.Println("=================")
//    fmt.Println(clientfinishtime)
//    loc, _ := time.LoadLocation("Asia/Shanghai")
//    finishtime :=  time.Now().In(loc)
//    if clientfinishtime!="unspecified"{
//    task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
//    db.Create(&task)
//    }else{
//    only_finishtime = finishtime.Format("060102")
//    task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
//    db.Create(&task)
//     }

    
// /*
// compute single day review data

// */
//  Check_reviewdaylog(only_finishtime,email)
//  computeinfo := Compute_singleday(only_finishtime,email)
//  fmt.Println(computeinfo)

//  }else{
//    task := Tasks{Task:inbox,User:email,Finishtime:"unfinished",Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
//    db.Create(&task)
//  }
//  c.JSON(200, gin.H{
//      "status":  "posted",
//      "message": "u have uploaded info,please come on!",
//    })
//  	}











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
// func Createtask(c *gin.Context) {
 
//    fmt.Println("+++++++++++++++++++ i am invoked in create task++++++++++++++++++++++")



//    emailcookie,_:=c.Request.Cookie("email")
//   fmt.Println(emailcookie.Value)
//   email:=emailcookie.Value
//   inbox := c.PostForm("inbox")
//   tasktags := c.PostForm("tasktags")
//   fmt.Println("=====task  tags=========")
//   fmt.Println(tasktags) 
//   fmt.Println("======*************========")
//   fmt.Println(inbox)
  
//   project := c.PostForm("project")
//   place := c.PostForm("place")
//   plantime := c.PostForm("plantime")
//   long :=  c.PostForm("long")
//   lat :=  c.PostForm("lat") 
//   goal :=  c.PostForm("goal") 
//   devotedtime,_:= strconv.Atoi(c.PostForm("timedevotedto_a_task"))
//   fmt.Println("+++++++++++++++devoted time info +++++++++++++") 
//   fmt.Println(devotedtime)
//   fmt.Println("+++++++++++++++place info +++++++++++++")   
//   longtitude = long
//   latitude = lat  
//   fmt.Println(long)
//   fmt.Println(lat)
  
//   fmt.Println("+++++++++++++++place info +++++++++++++")
//   if strings.Contains(plantime, "today"){
//        // if plantime =="today"{
//       loc, _ := time.LoadLocation("Asia/Shanghai")
//       //plantimeofanotherforamt :=  time.Now().In(loc)
//       //
//       plantime =  time.Now().In(loc).Format("060102")
//     }
//     if strings.Contains(plantime, "tomorrow"){
//      // if plantime  =="tomorrow"{
//       loc, _ := time.LoadLocation("Asia/Shanghai")
//     //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
//     plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
//     }






//   status := c.PostForm("taskstatus")
//    //status := c.PostForm("taskstatus")
//     if status == "f"{status = "finish"}
//     if status  == "g"{status = "giveup"}
//     if status   == "r"{status =  "replace"}
//     if status  == "a"{status = "anotherday"}
//   parentproject := c.PostForm("parentproject")
//   note := c.PostForm("note")
//   ifdissect := c.PostForm("ifdissect")







// if status!="unfinished"{
//   clientfinishtime:=  c.PostForm("finishtime")
//   if clientfinishtime == "y"{
//    clientfinishtime = "yesterday"

// }

//   fmt.Println("=================")
//   fmt.Println(clientfinishtime)
//   loc, _ := time.LoadLocation("Asia/Shanghai")
//   finishtime :=  time.Now().In(loc)


//   if clientfinishtime!="unspecified"{
//      if strings.Contains(clientfinishtime, "yesterday"){   
//      // if clientfinishtime == "yesterday"{
//    loc, _ := time.LoadLocation("Asia/Shanghai")
//     //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
//     clientfinishtime =  time.Now().In(loc).AddDate(0, 0,-1).Format("060102")

//     }
//   task := Tasks{Note:note,Ifdissect:ifdissect,Devotedtime:devotedtime,Goal:goal,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
//   db.Create(&task)



// //for adding review data for db when create a task
// Check_reviewdaylog(clientfinishtime,email)
// return_info:= Compute_singleday(clientfinishtime,email)
// fmt.Println(return_info)
//   }else{

//   task := Tasks{Note:note,Ifdissect:ifdissect,Devotedtime:devotedtime,Goal:goal,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
//   db.Create(&task)
  
// Check_reviewdaylog(finishtime.Format("060102"),email)
// return_info:= Compute_singleday(finishtime.Format("060102"),email)
// fmt.Println(return_info)


//  }


// }else{
//   task := Tasks{Task:inbox,User:email,Devotedtime:devotedtime,Finishtime:"unfinished",Goal:goal,Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
//   db.Create(&task)
// }






// c.JSON(200, gin.H{
//     "status":  "posted",
//     "message": "u have uploaded info,please come on!",
//   })
// 	}


func getindex(s string) int{
  alphabet :=[26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
 index:=28
  for k:=0;k<26;k++ {
   if s == alphabet[k]{
    index =  k
   }
  
 }
 return index
}




// createTodo add a new todo
func Createtaskfromsiri(c *gin.Context) {
 fmt.Println("+++++++++++++++++++ i am invoked in create task++++++++++++++++++++++")

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
  tasktags := gjson.Get(reqBody, "tasktags").String()
clientfinishtime:=  gjson.Get(reqBody, "finishtime").String()
   plantime := gjson.Get(reqBody, "plantime").String()

  client_status := gjson.Get(reqBody, "taskstatus").String()



   fmt.Println("------------------")
   fmt.Println(client_status)



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




task := Tasks{Task:inbox,Status:client_status,Project:"inbox",Goal:"Fight against fate",Finishtime:clientfinishtime,Email:email,User:email,Plantime:plantime,Tasktags:tasktags}
 //db.Create(&task).Scan(&task)
 db.Create(&task).Scan(&task)
    fmt.Println("i am testing the id return")
fmt.Println(task.ID)

if len(plantime)!=6{
         // if plantime =="today"{
          loc, _ := time.LoadLocation("Asia/Shanghai")
          //plantimeofanotherforamt :=  time.Now().In(loc)
          //
          plantime =  time.Now().In(loc).Format("060102")
}
// 先将查数据库中是否有评价数据的空，如果没有先创建，没有这一行会引起大bug
Check_reviewdaylog(plantime,email)
var score =   Compute_singleday(plantime,email)

fmt.Println("真成绩是")

fmt.Println(score)

c.JSON(200, gin.H{
    "status":  "posted",
    "score":score, 
   "message": "u have uploaded info,please come on!",
  })
        }




// createTodo add a new todo
func CreatetaskbyJSON(c *gin.Context) {
 fmt.Println("+++++++++++++++++++ i am invoked in create task++++++++++++++++++++++")

    //---------------get body string-------------
    //https://github.com/gin-gonic/gin/issues/1295
     buf := make([]byte, 1000000)
        num, _ := c.Request.Body.Read(buf)
        reqBody := string(buf[0:num])
   //--------------using gjson to parse------------
   //https://github.com/tidwall/gjson
  value := gjson.Get(reqBody, "reviewdata")
fmt.Println(value.String())

  emailcookie,err:=c.Request.Cookie("email")
  //fmt.Println(emailcookie.Value)
  var email string
  if err!=nil{
    email = c.Request.Header.Get("email")
  }else{
    fmt.Println(emailcookie.Value)
    email =emailcookie.Value
  }

  //检测客户端是从哪里过来的这里主要是为了解决从Siri过来的东西
  // client := c.Request.Header.Get("client")  



  admin := c.Query("admin")
  if admin =="xxx"{
    email ="yang756260386@gmail.com"
  }  
  inbox := gjson.Get(reqBody, "inbox").String()
  inboxlist := gjson.Get(reqBody, "inboxlist")
  devotedtime:= gjson.Get(reqBody, "timedevotedto_a_task").Int()
   fmt.Println(devotedtime)
   
  //这里如果是从Siri上传过来，我先获取tag，然后匹配可能存在的时间


   tasktagtextfromsiri := gjson.Get(reqBody, "tasktagtextfromsiri").String()
   re := regexp.MustCompile(`投入[\p{Han}]*[0-9]+分钟`)
  if  re.MatchString(tasktagtextfromsiri) == true{
  d :=   fmt.Sprintf("%q\n",re.Find([]byte(tasktagtextfromsiri)))
  re1 := regexp.MustCompile(`[0-9]+`)
  d1 := fmt.Sprintf("%s",re1.Find([]byte(d)))
  fmt.Println(d1)
   i, err := strconv.Atoi(d1)
     if err!= nil{
fmt.Println(err)
}else{
devotedtime = int64(i)
}

}
  
   




 // tasktags1 := gjson.Get(reqBody, "tasktags")
  // fmt.Println(tasktags1)
  tasktags := gjson.Get(reqBody, "tasktags").String()

  fmt.Println("------------------yangming /is here----------")
  fmt.Println(tasktags)




  goalcode_fromgtdcli := gjson.Get(reqBody, "goalcode").String()

  reviewalgodata := gjson.Get(reqBody, "reviewalgo").String()
  fmt.Println("=====task  tags=========")
  fmt.Println(tasktags)
  fmt.Println(reviewalgodata)
  fmt.Println("======*************========")
  fmt.Println(inbox)

  project := gjson.Get(reqBody, "project").String()
  

  

  //create a goal and save it to table
   alphabet :=[26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  
   //numbertable := [26]int{0,1, 2, 3, 4, 5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}
  var goalcount = 0
  var goalcountforsamegoal = 0
  var goalsforemail []Goalfordbs 
  if project == "goal"{
    db.Where("Email= ?", email).Where("Name=?",inbox).Find(&goalsforemail).Count(&goalcountforsamegoal)
    if goalcountforsamegoal >0{
      c.JSON(200, gin.H{
        "status":  "posted",
        "message": "the goal had ben created,not repeated youself",
      })
      return
    }else{

   
   db.Where("Email= ?", email).Find(&goalsforemail).Count(&goalcount)
   if goalcount ==0{
    goalfromclient := Goalfordbs{Name:inbox,Email:email,Goalcode:"aaa"}
    db.Create(&goalfromclient).Scan(&goalfromclient)
   }else{
    goalcode := goalsforemail[len(goalsforemail)-1].Goalcode
    a,b,c := string(goalcode[0]),string(goalcode[1]),string(goalcode[2])
    if c!= "z"{
      c = alphabet[getindex(c)+1]
    }else{
      if b!="z"{
        c = "a"
        b = alphabet[getindex(b)+1]
      }else{
        c= "a"
        b= "a"
        a = alphabet[getindex(a)+1]
      }
    }
    goalfromclient := Goalfordbs{Name:inbox,Email:email,Goalcode:fmt.Sprintf("%s%s%s ",a,b,c)}
    db.Create(&goalfromclient).Scan(&goalfromclient)
   } 

  }
  
  }
  //taglight := gjson.Get(reqBody, "taglight").String()
  place := gjson.Get(reqBody, "place").String()
  plantime := gjson.Get(reqBody, "plantime").String()
  long :=  gjson.Get(reqBody, "long").String()
  lat :=  gjson.Get(reqBody, "lat").String()
  goal :=  gjson.Get(reqBody, "goal").String()
 fmt.Println("------------测试目标是否运行正常------------")
  fmt.Println(goal)
    
  task_priority := 0
//检查客户端是哪里注意Siri  
client :=  gjson.Get(reqBody, "client").String()
  
  //deal with gtd commandline client info uploaded!
 if client == "gtdcli" {
     //gtd cli to pass goal code
  if goalcode_fromgtdcli !="xxx"{
    var goalsforemail []Goalfordbs 
    db.Where("Email= ?", email).Where("Goalcode=?",goalcode_fromgtdcli).Find(&goalsforemail).Count(&goalcountforsamegoal)
    goal = goalsforemail[0].Name
    task_priority =  goalsforemail[0].Priority
  }else{
    goal = "no goal"
  } 
 }

  if goal !="no goal"{
    var goalsforemail []Goalfordbs 
    db.Where("Email= ?", email).Where("Name=?",goal).Find(&goalsforemail)
    task_priority = goalsforemail[0].Priority
  }
  goalcoefficient :=  Get_goal_coffient(goal,email)


 // if goal !=="no goal"{}


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




  status := gjson.Get(reqBody, "taskstatus").String()
   //status := c.PostForm("taskstatus")
    if status == "f"{status = "finish"}
    if status  == "g"{status = "giveup"}
    if status   == "r"{status =  "replace"}
    if status  == "a"{status = "anotherday"}
  parentproject := gjson.Get(reqBody, "parentproject").String()
  note := gjson.Get(reqBody, "note").String()
  ifdissect := gjson.Get(reqBody, "ifdissect").String()





  var taskid uint;

  mark_finished_time_firstly_switch := 0
  var mark_finished_time_firstly_timestamp = ""
// give timestamp for marking time firstly the task finished 
if status!="unfinished"{
  // fmt.Println(clientfinishtime)
  loc, _ := time.LoadLocation("Asia/Shanghai")
  mark_finished_time_firstly_switch = 1
  mark_finished_time_firstly_timestamp =  time.Now().In(loc).Format("2006-01-02T15:04:05")
}



if status!="unfinished"{
  clientfinishtime:=  gjson.Get(reqBody, "finishtime").String()
  if clientfinishtime == "y"{
   clientfinishtime = "yesterday"

}


  if clientfinishtime == "today"{
   clientfinishtime = "today"

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


          if strings.Contains(clientfinishtime, "today"){
     // if clientfinishtime == "yesterday"{
   loc, _ := time.LoadLocation("Asia/Shanghai")
    //https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
    clientfinishtime =  time.Now().In(loc).Format("060102")
    
    }





    fmt.Println("-----------i am try to print inbox list-------------")
    fmt.Println(len(inboxlist.Array()))
   if len(inboxlist.Array())>1{
     for _,inbox := range inboxlist.Array(){
      // Mark_finished_time_switch int  `json:"mark_finished_time_switch"` 
     
      task := Tasks{Note:note,Ifdissect:ifdissect,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Priority:task_priority,Devotedtime:int(devotedtime),Goalcoefficient:goalcoefficient,Goal:goal,Parentproject:parentproject,Task:inbox.String(),User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
      db.Create(&task).Scan(&task)
      fmt.Println("i am testing the id return")
      fmt.Println(task.ID)
      taskid = task.ID
     }   

   }else{
    task := Tasks{Note:note,Ifdissect:ifdissect,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Priority:task_priority,Devotedtime:int(devotedtime),Goalcoefficient:goalcoefficient,Goal:goal,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
    db.Create(&task).Scan(&task)
    fmt.Println("i am testing the id return")
fmt.Println(task.ID)
taskid = task.ID
   }
 

   //for adding review data for db when create a task
   Check_reviewdaylog(clientfinishtime,email)
   return_info:= Compute_singleday(clientfinishtime,email)
   fmt.Println(return_info)
  }else{

    if len(inboxlist.Array())>1{
      for _,inbox := range inboxlist.Array(){
       //task := Tasks{Note:note,Ifdissect:ifdissect,Goal:goal,Parentproject:parentproject,Task:inbox.String(),User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
       task := Tasks{Note:note,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Ifdissect:ifdissect,Priority:task_priority,Goalcoefficient:goalcoefficient,Goal:goal,Devotedtime:int(devotedtime),Parentproject:parentproject,Task:inbox.String(),User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
       db.Create(&task).Scan(&task)
       fmt.Println("i am testing the id return")
       fmt.Println(task.ID)
       taskid = task.ID
      }   
 
    }else{
      task := Tasks{Note:note,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Ifdissect:ifdissect,Priority:task_priority,Goalcoefficient:goalcoefficient,Devotedtime:int(devotedtime),Goal:goal,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
     //task := Tasks{Note:note,Ifdissect:ifdissect,Goal:goal,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
     db.Create(&task).Scan(&task)
     fmt.Println("i am testing the id return")
    fmt.Println(task.ID)
   taskid = task.ID
    }

  Check_reviewdaylog(finishtime.Format("060102"),email)
  return_info:= Compute_singleday(finishtime.Format("060102"),email)
  fmt.Println(return_info)


  }


  }else{

    if len(inboxlist.Array())>1{
      for _,inbox := range inboxlist.Array(){
       //task := Tasks{Note:note,Ifdissect:ifdissect,Goal:goal,Parentproject:parentproject,Task:inbox.String(),User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
       //task := Tasks{Note:note,Ifdissect:ifdissect,Goal:goal,Parentproject:parentproject,Task:inbox.String(),User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
       //db.Create(&task).Scan(&task)
       task := Tasks{Task:inbox.String(),User:email,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Priority:task_priority,Finishtime:"unfinished",Goalcoefficient:goalcoefficient,Devotedtime:int(devotedtime),Goal:goal,Status:status,Email:email,Place:place,Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
      db.Create(&task).Scan(&task)
       
       
       fmt.Println("i am testing the id return")
       fmt.Println(task.ID)
       taskid = task.ID
      }   
 
    }else{
  task := Tasks{Task:inbox,User:email,Mark_finished_time_switch:mark_finished_time_firstly_switch,First_finish_timestamp:mark_finished_time_firstly_timestamp,Priority:task_priority,Finishtime:"unfinished",Goal:goal,Goalcoefficient:goalcoefficient,Devotedtime:int(devotedtime),Status:status,Email:email,Place:place,Project:project, Plantime:plantime,Tasktags:tasktags,Reviewdatas:reviewalgodata}
  db.Create(&task).Scan(&task)
     fmt.Println("i am testing the id return")
    fmt.Println(task.ID)
   taskid = task.ID
    }
  }


  if len(plantime)!=6{
    // if plantime =="today"{
     loc, _ := time.LoadLocation("Asia/Shanghai")
     //plantimeofanotherforamt :=  time.Now().In(loc)
     //
     plantime =  time.Now().In(loc).Format("060102")
}
// 先将查数据库中是否有评价数据的空，如果没有先创建，没有这一行会引起大bug
Check_reviewdaylog(plantime,email)
var score =   Compute_singleday(plantime,email)
fmt.Println("真成绩是")

fmt.Println(score)
//Print the HTTP response status.




 s := fmt.Sprintf("%f", score)



     ttsclienttext := "AI女娲在陆家嘴为你播报，评价算法的分数为"+ s 
  // ttsclient(ttsclienttext)
 //这里分别向前段推送语音合成数据
 ttsclient(ttsclienttext)
 //这里是向前段推送图像数据
 visiualdata2websocket(ttsclienttext)



  c.JSON(200, gin.H{
    "taskid": taskid,
    "score":score,
    "status":  "posted",
    "message": "u have uploaded info,please come on!",
  })
        }



//向腾讯语音合成推流

 func ttsclient(text string){
   //访问flask服务器，去合成音频流
  ttsurl := "http://localhost:5050/pcm?%s"
  var rq = url.Values{}
  rq.Add("text",text)
  resp, err := http.Get(fmt.Sprintf(ttsurl, rq.Encode()))
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
  defer resp.Body.Close()
if err != nil {
   panic(err)
}
defer resp.Body.Close()
//Print the HTTP response status.

fmt.Println("Response status:", resp.Status)
 }


//向echart推送图像服务

func visiualdata2websocket(text string){
  echarturl := "http://127.0.0.1:3030/pushtreedatetoweb"
  // ttsurl := "http://localhost:5050/pcm?%s"
  // var rq = url.Values{}
  // rq.Add("text",text)
  resp, err := http.Get(fmt.Sprintf(echarturl))
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
  defer resp.Body.Close()
  if err != nil {
    panic(err)
   }
defer resp.Body.Close()
//Print the HTTP response status.
color.Red("-------测试服务器是否工作正常--------")
fmt.Println("Response status:", resp.Status)
 }












  // Update a task
  func Update(c *gin.Context) {


    //---------------get body string-------------
    //https://github.com/gin-gonic/gin/issues/1295
     buf := make([]byte, 1024000)
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
    isreview := gjson.Get(reqBody, "isreview").String() 
    goal := gjson.Get(reqBody, "goal").String() 
    reviewdata := gjson.Get(reqBody, "reviewdata").String()
    fmt.Println(reviewdata)
    taglight := gjson.Get(reqBody, "taglight").String()
    tasktags := gjson.Get(reqBody, "tasktags").String()
    //gtd-cli to update the review algo data
    reviewalgolight := gjson.Get(reqBody, "reviewalgolight").String()
    reviewalgodata := gjson.Get(reqBody, "reviewalgo").String()
    goalcode_fromgtdcli := gjson.Get(reqBody, "goalcode").String()
    devotedtime:= gjson.Get(reqBody, "timedevotedto_a_task").Int()
    fmt.Println("---------------------tasktags info -------------------")
    fmt.Println(taglight)
    fmt.Println(tasktags)
    fmt.Println(reviewalgolight)
    fmt.Println(reviewalgodata)
    
    fmt.Println("-------------------------------------------------------")


     fmt.Println(reflect.TypeOf(reviewdata))
    //---------------------------review algorithms data-------------------------
    inbox := gjson.Get(reqBody, "inbox").String()
    place := gjson.Get(reqBody, "place").String()
    fmt.Println(inbox)
    
   // id := c.PostForm("id")
    id := gjson.Get(reqBody, "id").String()
   //检查更新id是否存在
   



    fmt.Println("-------------yangming------------------")
    fmt.Println(reflect.TypeOf(id))
    project := gjson.Get(reqBody, "project").String()
    finishtime := gjson.Get(reqBody, "finishtime").String()
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





    plantime := gjson.Get(reqBody, "plantime").String()
    
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
   






    status := gjson.Get(reqBody, "taskstatus").String()
    if status == "f"{status = "finish"}
    if status  == "g"{status = "giveup"}
    if status   == "r"{status =  "replace"}
    if status  == "a"{status = "anotherday"}


    parentproject := gjson.Get(reqBody, "parentproject").String()
   // note := c.PostForm("note")
    
    note := gjson.Get(reqBody, "note").String()
    fmt.Println(note)



    //status := c.PostForm("taskstatus")
    ifdissect := gjson.Get(reqBody, "ifdissect").String()
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


    //deal with goal client update
    if goalcode_fromgtdcli!=""&& goalcode_fromgtdcli=="xxx"{
      goal="unspecified"

    }


    var task_priority = 0
    client :=  gjson.Get(reqBody, "client").String()
    
    //deal with gtd commandline client info uploaded!
   if client == "gtdcli" {
     fmt.Println("---------------------+++++----------------------------")
       //gtd cli to pass goal code
       fmt.Println(goalcode_fromgtdcli)
       fmt.Println("---------------------+++++----------------------------")
    if goalcode_fromgtdcli !="xxx"{
      var goalsforemail []Goalfordbs 
      db.Where("Email= ?", email).Where("Goalcode=?",goalcode_fromgtdcli).Find(&goalsforemail)
      goal = goalsforemail[0].Name
      task_priority = goalsforemail[0].Priority
    }else{
      goal = "unspecified"
    } 
   }
    
   goalcoefficient :=  Get_goal_coffient(goal,email)

   if goal!="unspecified"{
fmt.Println("目标的权限等级错误")  
 fmt.Println(goal) 
  if goal!= "no goal"&&goal!=""{
    var goalsforemail []Goalfordbs 
    db.Where("Email= ?", email).Where("Name=?",goal).Find(&goalsforemail)
    task_priority = goalsforemail[0].Priority
}else{
   fmt.Println("目标的权限等级错误")
}    
db.Model(&task).Update("Priority", task_priority) 
    db.Model(&task).Update("Goal", goal)}

    if place!="unspecified"{
     
      db.Model(&task).Update("Place", place)}
    if project!="inbox"{db.Model(&task).Update("Project", project)}
    if inbox!="nocontent"{db.Model(&task).Update("Task", inbox)}
    if plantime!="unspecified"{db.Model(&task).Update("Plantime", plantime)}
    if parentproject!="unspecified"{db.Model(&task).Update("Parentproject", parentproject)}
    if ifdissect!="no"{db.Model(&task).Update("Ifdissect", ifdissect)}
    if note!="unspecified"{db.Model(&task).Update("Note", note)}
    if taglight == "yes"{db.Model(&task).Update("Tasktags", tasktags)}
    if reviewalgolight == "yes" {db.Model(&task).Update("Reviewdatas", reviewalgodata)}
    if devotedtime != 0  {db.Model(&task).Update("Devotedtime", int(devotedtime))}
    //EVERYTIME TO COMPUTE
    db.Model(&task).Update("Goalcoefficient", goalcoefficient)
    if isreview=="yes"{db.Model(&task).Update("reviewdatas", reviewdata)}
    
    //always to update tasktags,i think ther is a room to optimatical
    //db.Model(&task).Update("tasktags", tasktags)
    if totalscores!=0{
       fmt.Println("--------------------single task  total scores----------------")
       fmt.Println(totalscores)
       db.Model(&task).Update("reviewdatas", value.String())
       db.Model(&task).Update("Score", totalscores)}
       //using it to format time https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
   


fmt.Println("-----------i am printing -----status-------------")
fmt.Println(status)
fmt.Println("-----------i am printing -----status-------------")
  


//if the task fisrtly finished,update the mark switch and mark  the time stamp
if status!="unfinished"{
  if task.Mark_finished_time_switch == 0{
    // Mark_finished_time_switch int  `json:"mark_finished_time_switch"` 
    // First_finish_timestamp float64   `json:"First_finish_timestamp"`
    loc, _ := time.LoadLocation("Asia/Shanghai")
    // mark_finished_time_firstly_switch = 1
    mark_finished_time_firstly_timestamp :=  time.Now().In(loc).Format("2006-01-02T15:04:05")
    db.Model(&task).Update("Mark_finished_time_switch", 1)
    db.Model(&task).Update("First_finish_timestamp", mark_finished_time_firstly_timestamp)
  }
}


 if status!="unfinished"{
   db.Model(&task).Update("Finishtime", finishtime)
   //locate timezone https://stackoverflow.com/questions/27991671/how-to-get-the-current-timestamp-in-other-timezones-in-golang
      loc, _ := time.LoadLocation("Asia/Shanghai")
      now :=  time.Now().In(loc)
     db.Model(&task).Update("Finishtime",now.Format("060102"))
      //now1 :=  time.Now().In(loc)

//db.Model(&task).Update("Finishtime", finishtime)
      //db.Model(&task).Update("AccurateFinishtime",now1.String()


     db.Model(&task).Update("Status", status)

Check_reviewdaylog(now.Format("060102"),email)
return_info:= Compute_singleday(now.Format("060102"),email)
fmt.Println(return_info)



  }else{
    loc, _ := time.LoadLocation("Asia/Shanghai")
    now :=  time.Now().In(loc)
   db.Model(&task).Update("Finishtime",now.Format("060102"))
    //now1 :=  time.Now().In(loc)

//db.Model(&task).Update("Finishtime", finishtime)
    //db.Model(&task).Update("AccurateFinishtime",now1.String()

    fmt.Println("---------------i am confed in unfeid  even if user only update the one part the algo will update the score------------------------------")
    db.Model(&task).Update("Status", status)

   Check_reviewdaylog(now.Format("060102"),email)
   return_info:= Compute_singleday(now.Format("060102"),email)
fmt.Println(return_info)

  } 
  if finishtime!="unspecified"{db.Model(&task).Update("Finishtime", finishtime)


Check_reviewdaylog(finishtime,email)
return_info:= Compute_singleday(finishtime,email)
fmt.Println(return_info)
}




//Check_reviewdaylog(finishtime,email)
//return_info:= Compute_singleday(finishtime,email)
//fmt.Println(return_info)
loc, _ := time.LoadLocation("Asia/Shanghai")
today :=  time.Now().In(loc).Format("060102")
// 先将查数据库中是否有评价数据的空，如果没有先创建，没有这一行会引起大bug
Check_reviewdaylog(today,email)
var scorefotoday =   Compute_singleday(plantime,email)

//推送语音到客户端
s := fmt.Sprintf("%f", scorefotoday)
ttsclienttext := "AI女娲在陆家嘴为你播报，评价算法的分数为"+ s 
//ttsclient(ttsclienttext)

//这里分别向前段推送语音合成数据
  ttsclient(ttsclienttext)
  //这里是向前段推送图像数据
  visiualdata2websocket(ttsclienttext)
 c.JSON(200, gin.H{
        "status":  "posted",
        "score":scorefotoday,
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
//Find(&users)
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
  //db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Where("goal = ?", nil).Not("plantime", []string{today,tomorrow}).Order("id desc").Find(&tasks)
  
  var sql = fmt.Sprintf(`select * from tasks where status in("unfinish", "unfinished") and project not in ("goal") and (goal in ("no goal","") or goal is null) and plantime not in (%s,%s) and email ='%s' order by id desc`,today,tomorrow,email)
  db.Raw(sql).Scan(&tasks)
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
   client:= c.Request.Header.Get("client")
 fmt.Println(client)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
  // loc, _ := time.LoadLocation("Asia/Shanghai")
  // today :=  time.Now().In(loc).Format("060102")
  // tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")

  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("project in (?)", []string{"review"}).Order("id desc").Find(&tasks)
   if (client == "commandline"||client =="ios"){
   c.JSON(200, gin.H{
      "task":tasks,
    })

   }else{
  c.HTML(http.StatusOK, "reviewfromprojectreview.html",gin.H{
   "task":tasks,
  })

}


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




func Clock(c *gin.Context) {
c.HTML(http.StatusOK, "time.html", nil)
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
  //  })
   // })
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






func Tomorrowtaskjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  loc, _ := time.LoadLocation("Asia/Shanghai")
  tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
 

  //today :=  time.Now().In(loc).Format("060102")
 //tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")


 //db.Model(&task).Update("Finishtime",now.Format("060102"))

  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("plantime = ?",tomorrow).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}









func Yesterdaytaskjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  loc, _ := time.LoadLocation("Asia/Shanghai")
  yesterday :=  time.Now().In(loc).AddDate(0, 0, -1).Format("060102")
  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)
 

  //today :=  time.Now().In(loc).Format("060102")
 //tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")


 //db.Model(&task).Update("Finishtime",now.Format("060102"))

  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("plantime = ?",yesterday).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
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
      
      reviewtype := c.Query("reviewtype")
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




      fmt.Println(reviewtype)
     if reviewtype == "statistics"{
   fmt.Println("------------------------------------------------------")

    c.HTML(http.StatusOK, "reviewalgo.html", nil)
     }else if reviewtype == "goals"{

      c.HTML(http.StatusOK, "goals.html", nil)

     }else{

      c.HTML(http.StatusOK, "review.html", nil)
}      
	}




func Everydays(c *gin.Context) {
      //i use email as identifier
    //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,err:=c.Request.Cookie("email")
      var email string
      if err!=nil{
        email = c.Request.Header.Get("email")
      }else{
        fmt.Println(emailcookie.Value)
        email =emailcookie.Value
      }
      

      client:= c.Request.Header.Get("client")

      //fmt.Println(cookie1.Value)
      var tasks []Tasks
      //email:="yangming1"
       
     if client == "ios"{


      loc, _ := time.LoadLocation("Asia/Shanghai")
     today :=  time.Now().In(loc).Format("060102")
     tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
     yesterday :=  time.Now().In(loc).AddDate(0, 0, -1).Format("060102")






      db.Where("Email= ?", email).Where("plantime in (?)", []string{yesterday,today,tomorrow}).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks) 
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
              db.Where("Email= ?", email).Not("status", []string{"unfinished","unfinish","giveup","g"}).Find(&tasks)
              //alldays:=make(map[string] []Tasks)
              
              getdbdatatime := time.Now()
              fmt.Println("from request arrive to finished get data from db  time ")
             //https://stackoverflow.com/questions/45791241/correctly-measure-time-duration-in-go 
             fmt.Println(getdbdatatime.Sub(start))
             fmt.Println("++++++++")



              var alleverydays = Sort_tasksbyday(tasks)
            

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
db.Table("tasks").Where("Email= ?", email).Not("status", []string{"unfinished","unfinish","giveup","g"}).Count(&countforfinishedtasks)//reference not keyword

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

if(client   == "ios" || client == "clientforjson"||client == "weekpridejson"||client == "todaypridejson"||client =="yesterdaypridejson"){


        // c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": datafromredis})
         
//u need to note,the fellowing code ,u can test theperformance of  dealing with json ,when i set different size.
if (client == "clientforjson"){
c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays})
}else if(client == "weekpridejson"){
  fmt.Println("=========i am week pride=================")
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays[0:6]})
}else if(client == "todaypridejson"){
  fmt.Println("=========i am day pride=================")
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays[0:1]})
}else if(client == "yesterdaypridejson"){
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays[1:2]})
}else{
  //if (len(alldays)<30){
    if (len(alleverydays)<30){
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays})
  }else{
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays[0:30]})
  }

}



    returntoclienttime := time.Now()
    fmt.Println("from request arrive to the end time of return json to client ")
        //https://stackoverflow.com/questions/45791241/correctly-measure-time-duration-in-go
         fmt.Println(returntoclienttime.Sub(start))
         fmt.Println("++++++++")   

 }else{

          //   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "memories": alleverydays})


             c.HTML(http.StatusOK, "pridetimeline.html",gin.H{
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
