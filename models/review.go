
package models

import (
"fmt"
"github.com/fatih/color"
"encoding/json"
"net/http"
"github.com/jinzhu/gorm"
"strconv"
"time"
// "math"
//"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"
"github.com/bradfitz/slice"
"math"
"github.com/tidwall/gjson"
"github.com/gomodule/redigo/redis"
)


//from now on, i only to modify this file to add gtd review standard,that includes, detail struct and if control and totalscore and review struct instance data

//json is that it will be changed to this string json in db
//https://stackoverflow.com/questions/26327391/json-marshalstruct-returns
type Reviewdatadetail struct{
  Totalscore    float64 `json:"totalscore"`
  Averagescoreofhistory   float64 `json:"averagescoreofhistory"`
  Patience      float64  `json:"patience"`
  Attackactively      float64  `json:"attackactively"`
  Usebrain      float64   `json:"usebrain"`
  Useprinciple      float64   `json:"useprinciple"`
  Battlewithlowerbrain float64   `json:"battlewithlowerbrain"`
  Learnnewthings float64     `json:"learnnewthings"`
  Makeuseofthingsuhavelearned float64    `json:"makeuseofthingsuhavelearned"`
  Difficultthings float64  `json:"difficultthings"`
  Challengethings float64  `json:"challengethings"`
  Threeminutes    float64   `json:"threeminutes"`
  Getlesson       float64    `json:"getlesson"`
  Learntechuse    float64    `json:"learntechuse"` 
  Thenumberoftasks_score  float64    `json:"thenumberoftasks_score"`
  Self_discipline_score   float64    `json:"self_discipline_score"`
  Serviceforgoal_score  float64    `json:"serviceforgoal_score"`
  Onlystartatask float64       `json:"onlystartatask_score" sql:"size:999999"`
  Atomadifficulttask  float64    `json:"atomadifficulttask"`
  Alwaysprofit       float64     `json:"alwaysprofit"` 
  Markataskimmediately float64   `json:"markataskimmediately"`
  Doanimportantthingearly float64  `json:"doanimportantthingearly"`
  Buildframeandprinciple    float64 `json:"buildframeandprinciple"`
  Acceptfactandseektruth    float64 `json:"acceptfactandseektruth"`
  Acceptpain                float64 `json:"acceptpain"`
  Solveakeyproblem                float64 `json:"solveakeyproblem"`
  Depthfirstsearch                float64 `json:"depthfirstsearch"`
  Noflinch                float64 `json:"noflinch"`
  Setarecord             float64 `json:"setarecord"`
  Conquerthefear             float64 `json:"conquerthefear"`
  Executeability_score             float64 `json:"executeability_score"`
}




type Reviewofday  struct {
  gorm.Model
  Date string   `json:"date"`
  Email    string   `json:"email"`     
  Details string `json:"details" sql:"type:text;"`
   }


 //add coeffient   get table
  //  type Reviewofday  struct {
  //   gorm.Model
  //   Date string   `json:"date"`
  //   Email    string   `json:"email"`     
  //   Details string `json:"details" sql:"type:text;"`
  //    }









type Reviewfortimescount struct {
    gorm.Model
     Patience int   `json:"patiencenumber"`
     Email    string   `json:"email"`     
     //Details string `json:"details" sql:"type:text;"`
     Date string   `json:"date"`
     //Patience      int  `json:"patience"`
     Usebrain      int   `json:"usebrain"`
     Useprinciple      int   `json:"useprinciple"`
     Battlewithlowerbrain int   `json:"battlewithlowerbrain"`
    Learnnewthings int     `json:"learnnewthings"`
    Makeuseofthingsuhavelearned int    `json:"makeuseofthingsuhavelearned"`
    Difficultthings int  `json:"difficultthings"`
    Challengethings int  `json:"challengethings"`
   Threeminutes    int   `json:"threeminutes"`
   Getlesson       int    `json:"getlesson"`
Learntechuse    int    `json:"learntechuse"` 
Thenumberoftasks_score  int    `json:"thenumberoftasks_score"`
Serviceforgoal_score  int    `json:"serviceforgoal_score"`
Onlystartatask int       `json:"onlystartatask_score" sql:"size:999999"`
Atomadifficulttask  int    `json:"atomadifficulttask"`
Attackactively      int  `json:"attackactively"`
Alwaysprofit       int     `json:"alwaysprofit"` 
Markataskimmediately int   `json:"markataskimmediately"`
Doanimportantthingearly int  `json:"doanimportantthingearly"`  
Buildframeandprinciple    int `json:"buildframeandprinciple"`
Acceptfactandseektruth    int `json:"acceptfactandseektruth"`
Acceptpain                int `json:"acceptpain"`  
Solveakeyproblem                int `json:"solveakeyproblem"` 
Depthfirstsearch                int `json:"depthfirstsearch"`
Noflinch                   int  `json:"noflinch"` 
Setarecord             int `json:"setarecord"`
Conquerthefear          int `json:"conquerthefear"` 
Self_discipline_number   int    `json:"self_discipline_bumber"`
}







func Reviewalgorithmjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
var email string
emailcookie,err:=c.Request.Cookie("email")
if err!=nil{
  email = c.Request.Header.Get("email")
}else{
  fmt.Println(emailcookie.Value)
  email =emailcookie.Value
}

// 获取url中的参数
   daycount := c.Query("days")
   counts, _:= strconv.Atoi(daycount)


   var reviewdays []Reviewofday
   db.Where("email =  ?", email).Order("date").Find(&reviewdays)
  if counts >1{
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays[len(reviewdays)-counts:],
    })
  }



  //if u set the len,u will get the size of slice

  if len(reviewdays)<63{
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }else{

  //这里设置算反馈的日期
    reviewdays = reviewdays[len(reviewdays)-61:]
  
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }
 

}




func Reviewalgorithmjsonforyangming(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
var email string
// emailcookie,err:=c.Request.Cookie("email")
// if err!=nil{
//   email = c.Request.Header.Get("email")
// }else{
//   fmt.Println(emailcookie.Value)
//   email =emailcookie.Value
// }
email = "yang756260386@gmail.com"
// 获取url中的参数
   daycount := c.Query("days")
   counts, _:= strconv.Atoi(daycount)


   var reviewdays []Reviewofday
   db.Where("email =  ?", email).Order("date").Find(&reviewdays)
  if counts >1{
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays[len(reviewdays)-counts:],
    })
  }



  //if u set the len,u will get the size of slice

  if len(reviewdays)<63{
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }else{

  //这里设置算反馈的日期
    reviewdays = reviewdays[len(reviewdays)-61:]
  
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }
 

}




func Errorlog(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  //emailcookie,_:=c.Request.Cookie("email")
  emailcookie,err:=c.Request.Cookie("email")
  //fmt.Println(emailcookie.Value)
  var email string
   if err!=nil{
     email = c.Request.Header.Get("email")
   }else{
     fmt.Println(emailcookie.Value)
     email =emailcookie.Value
   }  





  var errors []Tasks
  //.Not("status", []string{"finished","f","finish","giveup","g"})
  db.Where("email =  ?", email).Where("project =  ?", "error").Order("id").Find(&errors)
  
  c.JSON(200, gin.H{
      "errorlog":errors,
    })

}




func Search(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)
  var keywords = c.Query("keywords")
  var search []Tasks
  db.Where("email =  ?", email).Where("task LIKE ?", "%"+keywords+"%").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&search)
  
  c.JSON(200, gin.H{
      "search":search,
    })

}






func goalcompare(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)
  var keywords = c.Query("keywords")
  var search []Tasks
  db.Where("email =  ?", email).Where("task LIKE ?", "%"+keywords+"%").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&search)
  
  c.JSON(200, gin.H{
      "search":search,
    })

}








func Searchwithtags(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,err:=c.Request.Cookie("email")
  //fmt.Println(emailcookie.Value)
  var email string
   if err!=nil{
     email = c.Request.Header.Get("email")
   }else{
     fmt.Println(emailcookie.Value)
     email =emailcookie.Value
   }
  //fmt.Println(cookie1.Value)
  var keywords = c.Query("keywords")
  
  var status = c.Query("status")
  
  fmt.Println(status)
  fmt.Println(keywords)
  var search []Tasks
  //var s string = "12312sf"
  querystring := "select * from tasks where status not in ('finished','finish','giveup','g') and  email =" +`"`+ email +`" `+ " and tasktags REGEXP "+"'"+`"`+keywords+`"`+":[ ]{0,1}"+`"yes"`+"'"
  querystring2 := "select * from tasks where status  in ('finished','finish') and  email =" +`"`+ email +`" `+ " and tasktags REGEXP "+"'"+`"`+keywords+`"`+":[ ]{0,1}"+`"yes"`+"'"
  //qurystring = fmt.Sprintf("select * from tasks where tasktags REGEXP '%s %s %s",s,"123123")
 // select * from tasks where tasktags REGEXP  '"hardtag":"yes"'\G;
  //db.Where("email =  ?", email).Where("task LIKE ?", "%"+keywords+"%").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&search)
  fmt.Println(querystring)
  if status == "unfinished"{db.Raw(querystring).Scan(&search)}
  if status == "f"{
    
    db.Raw(querystring2).Scan(&search)}
  c.JSON(200, gin.H{
      "search":search,
    })

}







func Problemssystem(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  var problems []Tasks
  db.Where("email =  ?", email).Where("project =  ?", "problem").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&problems)
  
  c.JSON(200, gin.H{
      "problems":problems,
    })

}



func Deadlinesystem(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  // emailcookie,_:=c.Request.Cookie("email")
  // fmt.Println(emailcookie.Value)
  // email:=emailcookie.Value
  email:="yang756260386@gmail.com"
  //fmt.Println(cookie1.Value)
  var daysfordeadline = Getmonthallday()
  var tasksofdeadline []Tasks
  db.Where("Email= ?",email).Not("status", []string{"finished","f","finish","giveup","g"}).Where("deadline IN (?)",daysfordeadline).Order("id").Find(&tasksofdeadline)
 fmt.Println(tasksofdeadline)
  c.JSON(200, gin.H{
      "deadlinefortasks":tasksofdeadline,
    })

}










func Questionssystem(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  var questions []Tasks
  db.Where("email =  ?", email).Where("project =  ?", "question").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&questions)
  
  c.JSON(200, gin.H{
      "questions":questions,
    })

}











func Reviewalgorithmjsonforios(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  email:= "yang756260386@gmail.com"
  //fmt.Println(cookie1.Value)

  var reviewdays []Reviewofday
  db.Where("email =  ?", email).Order("date").Find(&reviewdays)

  c.JSON(200, gin.H{
      "reviewdata":reviewdays,
    })

}


//图像部分

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




func Reviewfromyangming(c *gin.Context) {
//i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
// emailcookie,_:=c.Request.Cookie("email")
// fmt.Println(emailcookie.Value)
email:="yang756260386@gmail.com"
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

c.HTML(http.StatusOK, "reviewalgoforyangming.html", nil)
}else if reviewtype == "goals"{

c.HTML(http.StatusOK, "goals.html", nil)

}else{

c.HTML(http.StatusOK, "review.html", nil)
}      
}








func Reviewforios(c *gin.Context) {

  c.HTML(http.StatusOK, "reviewalgoforios.html",nil)

}


//this api was used to prepare the data of review
func Reviewforstastics(c *gin.Context){
  //get 7 days review datas 
  emailcookie,err:=c.Request.Cookie("email")
  var email string
   if err!=nil{
     email = c.Request.Header.Get("email")
   }else{
     fmt.Println(emailcookie.Value)
     email =emailcookie.Value
   }
 
  //fmt.Println(cookie1.Value)
  count_need_bystastics_from_client := c.Query("days")
  counts, _:= strconv.Atoi(count_need_bystastics_from_client)





type Result struct {
    Name string
}

loc, _ := time.LoadLocation("Asia/Shanghai")
var resultofgoal_tofinishintoday  []Result
today :=  time.Now().In(loc).Format("060102")
   db.Raw(`SELECT name  FROM goalfordbs  WHERE email ="`+email+`"`+" and goalstatus not in ("+`"giveup","g","finished","finish"`+`) and `+ ` name   NOT IN (SELECT goal  FROM tasks  WHERE finishtime=` +`"`+today+`"`+` and email =`+`"`+email+`"`+`);`).Scan(&resultofgoal_tofinishintoday)
















 // db.Where("email =  ?", email).Order("date").Find(&reviewsfortimescount)

//how many things do u had finished in theses days?
var countfortasks = 0
db.Table("tasks").Where("Email= ?", email).Where("status =?","giveup").Count(&countfortasks)
//how many things was correlated to goals?
//how many the times u had devoted to goals?
//which goals do u care specially?
//how urgent your goal is?
//

var tasks []Tasks
//email:="yangming1"
db.Where("Email= ?", email).Not("status", []string{"unfinished","unfinish","giveup","g"}).Find(&tasks)
var alleverydays = Sort_tasksbyday(tasks)
var tasksbydays []Everyday

// -1 表示昨天 1表示今天
fmt.Println("// -1 表示昨天 1表示今天")
fmt.Println(counts)

if counts == -1{
  tasksbydays = alleverydays[1:2]
}else{
  tasksbydays = alleverydays[0:counts]
}
//fmt.Println(tasksbydays)
var plannedtask_today_count = 0
var plannedtask_yesterday_count = 0
var plannedtask_same_with_finished_today_count = 0
var plannedtask_same_with_finished_yesterday_count = 0
//https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
yesterdaytime :=  time.Now().In(loc).AddDate(0, 0,-1).Format("060102")
todaytime :=  time.Now().In(loc).AddDate(0,0,0).Format("060102")
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",todaytime).Count(&plannedtask_today_count)
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",yesterdaytime).Count(&plannedtask_yesterday_count)
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",todaytime).Where("finishtime =?",todaytime).Count(&plannedtask_same_with_finished_today_count)
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",yesterdaytime).Where("finishtime =?",yesterdaytime).Count(&plannedtask_same_with_finished_yesterday_count)


var goal_devotedtime = make(map[string]int)
var alltasks_count = 0
var all_time_u_had_devoted_inthe_time_range = 0
var  alltime_goal_oriented = 0 
//这段时间范围内，每一个工作日你投入的时间是由接下来这个数组决定的 

var  devotedtime_for_goal_in_everyday []int
for  _,item :=range tasksbydays{
  alltasks_count =  alltasks_count+len(item.Alldays)
  //接下来在循环每一天的任务
  var day_devotedtime = 0//每天投入的时间
  for _,item1 := range item.Alldays{
         day_devotedtime = day_devotedtime +item1.Devotedtime
    all_time_u_had_devoted_inthe_time_range = all_time_u_had_devoted_inthe_time_range + item1.Devotedtime
    if item1.Goal!="no goal"{
     
      if val, ok := goal_devotedtime[item1.Goal]; ok {
        goal_devotedtime[item1.Goal] =val +item1.Devotedtime
      }else{
        goal_devotedtime[item1.Goal] = item1.Devotedtime
      }
      alltime_goal_oriented = alltime_goal_oriented + item1.Devotedtime
    } 
  } 
  //把每天投入的时间加入到合理的区间里面
  devotedtime_for_goal_in_everyday = append(devotedtime_for_goal_in_everyday,day_devotedtime)
}
// fmt.Printf("the task length is %d",len(tasksbydays))
// fmt.Printf("theses task counts is %d",alltasks_count)
// fmt.Printf("u had devoted %d  minutes in the time range",all_time_u_had_devoted_inthe_time_range)
// fmt.Printf("u had devoted %d  minutes in the time range for goal",alltime_goal_oriented)
  var reviewsfortimescount []Reviewfortimescount


tomorrowtime :=  time.Now().In(loc).AddDate(0, 0,1).Format("060102")

  
//这里其实会出现功能性的bug，当用户不小心更新了以后的日子
db.Where("email =  ?", email).Where("date < ?", tomorrowtime).Order("date").Find(&reviewsfortimescount)


fmt.Println("--------")
fmt.Println(reviewsfortimescount)

  if (len(reviewsfortimescount)-counts < 0){
    c.JSON(200, gin.H{
      "errorcode":1101,
      "msg":"u need more days to use the functions",
    })
   return
  }
  lengthofreviewsfortimescount := len(reviewsfortimescount)
  var reviewdata []Reviewfortimescount
 
  //for yesterday
  if counts == -1{
    reviewdata = reviewsfortimescount[lengthofreviewsfortimescount-2:lengthofreviewsfortimescount-1]
  }else{
    weekstart := lengthofreviewsfortimescount - counts
    reviewdata = reviewsfortimescount[weekstart:]
  }
 
  

  fmt.Println(reviewdata)
  
fmt.Println("---------------")

//   for _,item :=range reviewdata{
  //   var detailofday = item.Details 
  //   challengetag := gjson.Get(detailofday, "challengetag").String()
  //   fmt.Println(challengetag)  
  // }
  c.JSON(200, gin.H{
      "errorcode":1102,
      "dayscount":len(tasksbydays),
      "alltasks_count":alltasks_count,
      "devotedtime":all_time_u_had_devoted_inthe_time_range,
      "devotedtime_oriented":alltime_goal_oriented,
      "yesterday_planed_task_count":plannedtask_yesterday_count,
      "today_planed_task_count":plannedtask_today_count,
      "plannedtask_same_with_finished_today_count":plannedtask_same_with_finished_today_count,   
      "plannedtask_same_with_finished_yesterday_count":plannedtask_same_with_finished_yesterday_count, 
      "goaltime":goal_devotedtime,
      "reviewdata":reviewdata,
      "resultofgoal_tofinishintoday":resultofgoal_tofinishintoday,
      "devotedtime_for_goal_in_everyday":devotedtime_for_goal_in_everyday,
    })
}



func pow(x, n int) int {
  ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
  for n != 0 {
      if n%2 != 0 {
          ret = ret * x
      }
      n /= 2
      x = x * x
  }
  return ret
}
// --------------------- 
// 作者：陈鹏万里 
// 来源：CSDN 
// 原文：https://blog.csdn.net/qq245671051/article/details/70342047 
// 版权声明：本文为博主原创文章，转载请附上博文链接！




func Task_execute_priority_table_review(date string,email string) float64{
  //get the finished task in the day by the accurate finished time
  var tasks_finished_pure []Tasks
  db.Where("Email= ?", email).Where("plantime =  ?", date).Where("finishtime =  ?", date).Order("first_finish_timestamp").Find(&tasks_finished_pure)
  
  var year = "20"+date[0:2]
  var month = date[2:4]
  var day = date[4:6]
  
  pivot := year+"-"+month+"-"+day+"T10:00:01"
  fmt.Printf("-------lay out is %s------\n",pivot)
  t, err := time.Parse("2006-01-02T15:04:05", pivot)
  if err != nil {
      fmt.Println(err)
  }
  var  time_pivot =  t.Unix()
  
  fmt.Println(time_pivot)
   
  //I am about to reorder the tasks completed before ten o'clock.
  var tasks_finished []Tasks
  var task_finished_before_10am []Tasks
  var task_finished_after_10am  []Tasks
   // time.Now().UnixNano()
   //https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
  for _,task  :=range tasks_finished_pure{
    t, _ := time.Parse("2006-01-02T15:04:05", task.First_finish_timestamp)
    fmt.Println("-------******---------")
    fmt.Println(task.First_finish_timestamp)
    fmt.Println(time_pivot)
    fmt.Println(t.Unix())
    if t.Unix() < time_pivot{
      task_finished_before_10am =  append(task_finished_before_10am,task)
    }else{
      task_finished_after_10am =  append(task_finished_after_10am,task)
    }
  }
  //slice sort
  slice.Sort(task_finished_before_10am[:], func(i, j int) bool {
    return task_finished_before_10am[i].Priority > task_finished_before_10am[j].Priority
}) 



//https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go/29688973
   //tasks_finished = task_finished_before_10am + task_finished_after_10am 
  tasks_finished =  append(tasks_finished, append(task_finished_before_10am, task_finished_after_10am...)...)

  
  //get the optimal excute order
  var tasks_planed []Tasks
  db.Where("Email= ?", email).Where("plantime =  ?", date).Order("priority").Find(&tasks_planed)
  fmt.Printf("-----------------the planed tasks-----%d---------\n",len(tasks_planed))
  var optimal_order_by_paln = 0.0
  var optiaml_order = ""
  for i,task :=range tasks_planed{
     optiaml_order =  optiaml_order + strconv.Itoa(task.Priority)
     optimal_order_by_paln = optimal_order_by_paln + float64(task.Priority)*math.Pow10(i)
  }
  //get the real execute order
  var real_order_by_execute = 0.0
  var real_order = ""
  for i,task :=range tasks_finished{
    real_order =  real_order + strconv.Itoa(task.Priority)
    real_order_by_execute = real_order_by_execute + float64(task.Priority)*math.Pow10(i)
 }

 var gap_between_reality_and_ideal = len(tasks_planed) - len(tasks_finished)

 real_order_by_execute = real_order_by_execute*math.Pow10(gap_between_reality_and_ideal)  
fmt.Printf("------------------------------------------------\n")
fmt.Println(gap_between_reality_and_ideal)
fmt.Printf("real order include justify is %s \n",real_order)
fmt.Printf("optiaml order is %s \n",optiaml_order)
fmt.Println(time_pivot)
fmt.Println(tasks_finished_pure)
fmt.Println(task_finished_before_10am)
fmt.Println(task_finished_after_10am)
fmt.Println(tasks_finished)
fmt.Printf("optiaml order score is %f",optimal_order_by_paln)
fmt.Printf("real order is score is %f",real_order_by_execute)
fmt.Printf("------------------------------------------------\n")

 
var  task_execute_priority_table_review float64 = 0.1
task_execute_priority_table_review = float64(real_order_by_execute)/float64(optimal_order_by_paln)
return task_execute_priority_table_review
}






// createTodo add a new todo
func Reviewscore_today(c *gin.Context) {
  fmt.Println("+++++++++++++++++++ i am invoked in create task++++++++++++++++++++++")
 
     //---------------get body string-------------
     //https://github.com/gin-gonic/gin/issues/1295
      buf := make([]byte, 1000000)
         num, _ := c.Request.Body.Read(buf)
         reqBody := string(buf[0:num])
    //--------------using gjson to parse------------
    //https://github.com/tidwall/gjson
   
   //emailcookie,_:=c.Request.Cookie("email")
   //fmt.Println(emailcookie.Value)
   emailcookie,err:=c.Request.Cookie("email")
   var email string
   if err!=nil{
     email = c.Request.Header.Get("email")
   }else{
     fmt.Println(emailcookie.Value)
     email =emailcookie.Value
   }


   inbox := gjson.Get(reqBody, "inbox").String()
  fmt.Println(inbox)
   loc, _ := time.LoadLocation("Asia/Shanghai")
    plantime :=  time.Now().In(loc).Format("060102")
 // 先将查数据库中是否有评价数据的空，如果没有先创建，没有这一行会引起大bug
 Check_reviewdaylog(plantime,email)
 var score =   Compute_singleday(plantime,email)
 
 c.JSON(200, gin.H{
     "status":  "posted",
     "score":score, 
    "message": "u have uploaded info,please come on!",
   })
         }











func   Check_reviewdaylog(date string,email string){
/*
check if date row was created in reviewday table,if it is no,the function will create it

*/

var reviewday Reviewofday 
var reviewfortimecount  Reviewfortimescount
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewday)
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewfortimecount)
if reviewday.Date!=date{
db.Create(&Reviewofday{Date: date,Email:email,Details:"no"})
}else{
fmt.Println("===========the record has been created in the past==========")
}

if reviewfortimecount.Date!=date{
  db.Create(&Reviewfortimescount{Date: date,Email:email})
  }else{
  fmt.Println("===========the record has been created in the past==========")
  }
}









//执行能力评价部分
func  Check_execute(date string,email string) float64{
  /*
  check if date row was created in reviewday table,if it is no,the function will create it
  
  */
  
   //先获取当天的任务
    var tasksoftoday []Tasks
    var countoftasksoftoday int
    // var tasksoffinished []Tasks
    db.Where("plantime = ?", date).Where("email =  ?", email).Find(&tasksoftoday).Count(&countoftasksoftoday)
    //今天计划了几件事，完成了几件事情轻的比例
   var finishedtaskcount = 0
    for _,item :=range tasksoftoday{
      if (item.Status == "finish"){
      finishedtaskcount+=1
      }
    }
    color.Blue(string("--------检查是否执行函数颜色----------"))
    fmt.Println(finishedtaskcount)
    fmt.Println(countoftasksoftoday)
    var execute1 float64 = float64(finishedtaskcount)/float64(countoftasksoftoday)
    var toalscoreofexecute = execute1
    //判断几乎是开始时间和实际执行时间是否一样
    return  toalscoreofexecute 
  }














//compute total_scores of someday

func Compute_singleday(date string,email string) float64{
//获取执行能力细节
executeability_score  := Check_execute(date,email) 



  //https://tour.golang.org/basics/10
fmt.Println("------------ i am here to compute the single day---------------------------")
var tasks []Tasks
//email := "yang756260386@gmail.com"
var brainuse_score,makeuseofthethingsuhavelearned_score,difficultthings_score,threeminutes_score,getlesson_score,learntechuse_score,battlewithlowerbrain_score,patience_score,learnnewthings_score float64 = 0,0,0,0,0,0,0,0,0 
var serviceforgoal_score,onlystartatask_score float64 = 0,0
var atomadifficulttask_score,alwaysprofit_score float64 = 0,0
var doanimportantthingearly_score,markataskimmediately_score float64 = 0,0
var challengetag_score float64= 0
var challengetag_number = 0
var atomtag_score float64= 0
db.Where("Email= ?", email).Where("finishtime =  ?", date).Not("status", []string{"unfinished","unfinish","giveup"}).Order("id desc").Find(&tasks)

var taskcount_score float64

var countoffinishedtasks int

var countofgivenuptasks int
var count_makeplanfortomorrow  = 0 

//for times stastics
var patiencenumber = 0 
var battlewithlowerbrainnumber = 0
var usebrainnumber = 0
var buildframeandprinciple_score float64 =0 
var buildframeandprinciplenumber =0 
var markataskimmediately_number = 0
var alwaysprofit_number = 0
var threeminutes_number = 0
var difficultthings_number = 0
var learnnewthings_number = 0
var atomadifficulttask_number = 0 
var makeuseofthethingsuhavelearned_number = 0
var doanimportantthingearly_number =0
var serviceforgoal_number = 0
var acceptfactandseektruth_score  float64= 0
var acceptfactandseektruth_number = 0
var acceptpain_score float64 = 0
var acceptpain_number = 0
var solveakeyproblem_score float64= 0
var solveakeyproblem_number = 0
var attackactively_number = 0
var attackactively_score float64= 0
var useprinciple_number = 0
var useprinciple_score float64= 0
var dfs_number = 0
var dfs_score float64= 0
var noflinch_number = 0
var noflinch_score float64= 0
var conquerthefear_score float64= 0
var conquerthefear_number  = 0
var setarecord_score float64= 0
var setarecord_number  = 0
var self_discipline_score float64= 0
var self_discipline_number  = 0
// var self_discipline_coffient = 0.5

db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Count(&countoffinishedtasks)
db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Where("status =?","giveup").Count(&countofgivenuptasks)
db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Where("status =?","giveup").Count(&countofgivenuptasks)
// db.Where(&Tasks{Finishtime : date, Task:"make plan for tomorrow on "+date,Status:"finish"}).Count(&count_makeplanfortomorrow)
//count_makeplanfortomorrow
fmt.Println("-----=======-------++++++++++++++=-----======----------------")
fmt.Println(countofgivenuptasks)
//task basic count
taskcount_score =  float64(1* (countoffinishedtasks - countofgivenuptasks) +countofgivenuptasks*0)
 var pain_coeffient float64 = 0.5
 var difficult_coeffient float64 = 0.5
 var atom_coffient = 0.5 
 

 for _,item :=range tasks{
fmt.Println("------------i had been into loop----------------")

/*
给执行时间 参数
1.当执行时间比计划时间在5分钟以内就加上
2.不确定为什么在结束时间

*/

//转换获取计划时间
var starttime_plan = item.Starttime 
var endtime_plan = item.Endtime 
var plantime = item.Plantime
var starttime_exe = item.Starttime_exe
var endtime_exe = item.Endtime_exe

fmt.Println(endtime_plan,endtime_exe)
if starttime_exe !="unspecified"&&starttime_plan!="unspecified"{
  //为分钟级别1131的精度进行设计
  if len(starttime_plan) ==4{
  year, _ := strconv.Atoi("20"+plantime[0:2])
  month,_ := strconv.Atoi(plantime[2:4])
  day,_ := strconv.Atoi(plantime[4:6])
  hour,_ := strconv.Atoi(starttime_plan[0:2])
  minute,_ := strconv.Atoi(starttime_plan[2:4])
  date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
  //获取的是typeint   64
  starttime_planunix := date.Unix()*1000
  starttime_exeunix,_  := strconv.Atoi(starttime_exe)
  
  gap := int64(starttime_exeunix) - starttime_planunix
  
  if gap < 5000{
   self_discipline_number = self_discipline_number + 1
   self_discipline_score = float64((self_discipline_score + 0.25))
  }
   
  fmt.Printf("date is :%s \n", date)
  date = time.Date(2018, 01, 12, 22, 51, 48, 324359102, time.UTC)
  fmt.Printf("date is :%s", date)
}
 }


if endtime_exe !="unspecified"&&endtime_plan!="unspecified"{
  //为分钟级别1131的精度进行设计
  if len(endtime_plan)==4{
  year, _ := strconv.Atoi("20"+plantime[0:2])
  month,_ := strconv.Atoi(plantime[2:4])
  day,_ := strconv.Atoi(plantime[4:6])
  hour,_ := strconv.Atoi(endtime_plan[0:2])
  minute,_ := strconv.Atoi(endtime_plan[2:4])
  date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
  //获取的是typeint   64
  endtime_planunix := date.Unix()*1000
  endtime_exeunix,_  := strconv.Atoi(starttime_exe)
  
  gap := endtime_planunix-int64(endtime_exeunix) 
  
  if gap > 3000{
   self_discipline_number = self_discipline_number + 1
   self_discipline_score = float64((self_discipline_score + 0.25))
  }
   
  fmt.Printf("date is :%s \n", date)
  date = time.Date(2018, 01, 12, 22, 51, 48, 324359102, time.UTC)
  fmt.Printf("date is :%s", date)
}
}






var jsonoftasktags = item.Tasktags
if  challengetag := gjson.Get(jsonoftasktags, "challengetag").String();challengetag=="yes"{
challengetag_score  = float64((challengetag_score + 5))*item.Goalcoefficient
challengetag_number  = challengetag_number + 1

fmt.Println("-----=======-------++++++++++杨明在这里++++=-----======----------------")
fmt.Println(item)
}



//this  tom tag work for  tasks tagged with atom and finished!
if  atomtag := gjson.Get(jsonoftasktags, "atomtag").String();atomtag=="yes"{
  atomtag_score  = float64((atomtag_score + 5))*item.Goalcoefficient
  if atom_coffient < 1{
    atom_coffient =  atom_coffient + 0.5
  }else{
    atom_coffient =  atom_coffient + 0.2
  }
  
  }


var json = item.Reviewdatas
//var goalcoffient = 1item.goalcoffient
fmt.Println(json)
fmt.Println("------------i had been into loop----------------")
if  brainuse := gjson.Get(json, "brainuse").String();brainuse=="yes"{
fmt.Println(brainuse)
brainuse_score = float64(brainuse_score +5)*(item.Goalcoefficient)
usebrainnumber = usebrainnumber +1
 } 

if  buildframeandprinciple_from_client := gjson.Get(json, "buildframeandprinciple").String();buildframeandprinciple_from_client=="yes"{
  //fmt.Println(brainuse)
  buildframeandprinciple_score = float64(buildframeandprinciple_score +5)*item.Goalcoefficient
  buildframeandprinciplenumber = buildframeandprinciplenumber +1
   } 

   if  useprinciple_from_client := gjson.Get(json, "useprinciple").String();useprinciple_from_client=="yes"{
    //fmt.Println(brainuse)
    useprinciple_score = float64(useprinciple_score +10)*(item.Goalcoefficient)
        useprinciple_number = useprinciple_number +1
     } 

   
     if  noflinch := gjson.Get(json, "noflinch").String();noflinch=="yes"{
      //fmt.Println(brainuse)
      noflinch_score = float64(noflinch_score  +  50)*(item.Goalcoefficient)
      noflinch_number = noflinch_number +1
       } 
    


  
  


     if  acceptfact_from_client := gjson.Get(json, "acceptfactandseektruth").String();acceptfact_from_client=="yes"{
      //fmt.Println(brainuse)
      acceptfactandseektruth_score = float64(acceptfactandseektruth_score  +  10)*(item.Goalcoefficient)
      acceptfactandseektruth_number = acceptfactandseektruth_number +1
       } 
   






     if  attackactively_from_client := gjson.Get(json, "attackactively").String();attackactively_from_client=="yes"{
      //fmt.Println(brainuse)
      attackactively_score=  float64(attackactively_score  +  10)*item.Goalcoefficient
      attackactively_number= attackactively_number +1
       } 





    
     if  acceptfact_from_client := gjson.Get(json, "acceptpain").String();acceptfact_from_client=="yes"{
      //fmt.Println(brainuse)
      acceptpain_score =  float64(acceptpain_score  +  10)*item.Goalcoefficient
      acceptpain_number = acceptpain_number +1
      pain_coeffient =  pain_coeffient + 0.5
       } 




if  makeuseofthings := gjson.Get(json, "makeuseofthings").String();makeuseofthings=="yes"{
makeuseofthethingsuhavelearned_score =  float64( makeuseofthethingsuhavelearned_score + 5)*item.Goalcoefficient
makeuseofthethingsuhavelearned_number = makeuseofthethingsuhavelearned_number +1
 }





if  doanimportantthingearly := gjson.Get(json, "doanimportantthingearly").String();doanimportantthingearly =="yes"{
doanimportantthingearly_score =  float64(doanimportantthingearly_score + 10)*item.Goalcoefficient
doanimportantthingearly_number = doanimportantthingearly_number + 1
 }
 

if  markataskimmediately := gjson.Get(json, "markataskimmediately").String();markataskimmediately =="yes"{
markataskimmediately_score = float64(markataskimmediately_score + 1 )*item.Goalcoefficient
markataskimmediately_number = markataskimmediately_number + 1
 }





if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
alwaysprofit_score = float64(alwaysprofit_score + 5)*item.Goalcoefficient
alwaysprofit_number = alwaysprofit_number +1
 }


// if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
// alwaysprofit_score = alwaysprofit_score + 5
//  }




if  learnnewthings := gjson.Get(json, "learnnewthings").String();learnnewthings=="yes"{
learnnewthings_score = float64(learnnewthings_score +5 )*item.Goalcoefficient
learnnewthings_number = learnnewthings_number + 1
 }


if  serviceforgoal := gjson.Get(json, "serviceforgoal").String();serviceforgoal=="yes"{
serviceforgoal_score  = float64(serviceforgoal_score  + 20  )*item.Goalcoefficient
 }


if  onlystartatask := gjson.Get(json, "onlystartatask").String();onlystartatask=="yes"{
onlystartatask_score  =  float64(onlystartatask_score  + 10)*item.Goalcoefficient
 }





if  battlewithlowerbrain := gjson.Get(json, "battlewithlowerbrain").String();battlewithlowerbrain=="yes"{
battlewithlowerbrain_score = float64(battlewithlowerbrain_score +5 )*item.Goalcoefficient
battlewithlowerbrainnumber = battlewithlowerbrainnumber + 1
 }


 if  conquerthefear := gjson.Get(json, "conquerthefear").String();conquerthefear=="yes"{
  conquerthefear_score = float64(conquerthefear_score +50 )*item.Goalcoefficient
  conquerthefear_number = conquerthefear_number + 1
   }



   if  setarecord := gjson.Get(json, "setarecord").String();setarecord=="yes"{
    setarecord_score = float64(setarecord_score +50 )*item.Goalcoefficient
    setarecord_number = setarecord_number + 1
     }





if  atomadifficulttask := gjson.Get(json, "atomadifficulttask").String();atomadifficulttask=="yes"{
atomadifficulttask_score = float64(atomadifficulttask_score +5 )*item.Goalcoefficient
atomadifficulttask_number = atomadifficulttask_number+1
 }






if  patience := gjson.Get(json, "patience").String();patience=="yes"{
patience_score =  float64(patience_score + 10)*item.Goalcoefficient
patiencenumber = patiencenumber + 1
 }


 if  solveakeyproblem := gjson.Get(json, "solveakeyproblem").String();solveakeyproblem=="yes"{
  solveakeyproblem_score =  float64(solveakeyproblem_score + 50)*item.Goalcoefficient
  solveakeyproblem_number = solveakeyproblem_number + 1
   }





if  difficultthings := gjson.Get(json, "difficultthings").String();difficultthings=="yes"{
difficultthings_score =  float64(difficultthings_score +20)*item.Goalcoefficient
difficultthings_number = difficultthings_number + 1
difficult_coeffient =  difficult_coeffient + 0.5
 }



if  threeminutes := gjson.Get(json, "threeminutes").String();threeminutes=="yes"{
threeminutes_score = float64(threeminutes_score +5 )*item.Goalcoefficient
threeminutes_number = threeminutes_number + 1
 }

 if  dfs := gjson.Get(json, "depthfirstsearch").String();dfs=="yes"{
  dfs_score = float64(dfs_score +20 )*item.Goalcoefficient
  dfs_number = dfs_number + 1
   }




if  getlesson:= gjson.Get(json, "getlesson").String();getlesson=="yes"{
getlesson_score = float64(getlesson_score +5  )*item.Goalcoefficient

 }







if  learntechuse := gjson.Get(json, "learntechuse").String();learntechuse=="yes"{
learntechuse_score = float64(learntechuse_score +5)*item.Goalcoefficient
 }


 }




total_score:=acceptfactandseektruth_score+dfs_score+useprinciple_score+attackactively_score+solveakeyproblem_score+acceptpain_score+buildframeandprinciple_score+conquerthefear_score+setarecord_score+taskcount_score+doanimportantthingearly_score+atomadifficulttask_score+onlystartatask_score+markataskimmediately_score+challengetag_score +atomtag_score+ brainuse_score+alwaysprofit_score + makeuseofthethingsuhavelearned_score + battlewithlowerbrain_score + patience_score + learnnewthings_score+difficultthings_score+threeminutes_score+getlesson_score+learntechuse_score + serviceforgoal_score

fmt.Println("--------之前的成绩----")
fmt.Println(total_score)




fmt.Println("-------voted to shanghai -----------")
//----------------------------------------------------------plan obey part------------------
//plan obey coeffient
//the part to judge how self-reglation you r
var tasksforstatistics []Tasks
//email:="yangming1"
db.Where("Email= ?", email).Not("status", []string{"unfinished","unfinish","giveup","g"}).Find(&tasksforstatistics)
// var alleverydays = Sort_tasksbyday(tasks)
// var tasksbydays []Everyday

// if counts == -1{
//   tasksbydays = alleverydays[1:2]      
// }else{
//   tasksbydays = alleverydays[0:counts]
// }
//fmt.Println(tasksbydays)
var plannedtask_count = 0
//var plannedtask_yesterday_count = 0
var plannedtask_same_with_finished_count = 0
//var plannedtask_same_with_finished_yesterday_count = 0
//loc, _ := time.LoadLocation("Asia/Shanghai")
//https://stackoverflow.com/questions/37697285/how-to-get-yesterday-date-in-golang
//yesterdaytime :=  time.Now().In(loc).AddDate(0, 0,-1).Format("060102")
//todaytime :=  time.Now().In(loc).AddDate(0,0,0).Format("060102")
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",date).Count(&plannedtask_count)
fmt.Println("------------------i am writting----------------------")
fmt.Println(plannedtask_count)
fmt.Println(date)
//db.Table("tasks").Where("Email= ?", email).Where("plantime =?",yesterdaytime).Count(&plannedtask_yesterday_count)
db.Table("tasks").Where("Email= ?", email).Where("plantime =?",date).Where("finishtime =?",date).Count(&plannedtask_same_with_finished_count)
//db.Table("tasks").Where("Email= ?", email).Where("plantime =?",yesterdaytime).Where("finishtime =?",yesterdaytime).Count(&plannedtask_same_with_finished_yesterday_count)
var planobey_coffient = 0.0
if plannedtask_count !=0 {
  planobey_coffient = float64(plannedtask_same_with_finished_count)/float64(plannedtask_count)
  fmt.Println(planobey_coffient)
  if plannedtask_count ==0{
    planobey_coffient = 0.2
  }
}
 fmt.Println(plannedtask_same_with_finished_count)
fmt.Println(planobey_coffient)




//--------the plan priority task order finishd----------------
var priority_execute_coffient = Task_execute_priority_table_review(date,email)
// if plannedtask_count !=0 {
//   planobey_coffient = float64(plannedtask_same_with_finished_count)/float64(plannedtask_count)
//   fmt.Println(planobey_coffient)
//   if plannedtask_count ==0{
//     planobey_coffient = 0.2
//   }
// }
//  fmt.Println(plannedtask_same_with_finished_count)
// fmt.Println(planobey_coffient)







fmt.Println("------------------------------------Do you plan for tomorrow?-------------------------------------")
loc, _ := time.LoadLocation("Asia/Shanghai")
yesterday :=  time.Now().In(loc).AddDate(0, 0, -1).Format("060102")
db.Table("tasks").Where("Email= ?", email).Where("task = ?","make plan for tomorrow on "+yesterday).Where("plantime =?",yesterday).Where("status =?","finish").Count(&count_makeplanfortomorrow)

fmt.Println("---------------the date is ------------------")

fmt.Println("make plan for tomorrow on "+date)
//if u dnoot plan fro tomorrow the score will * 0.75

// float compute  https://www.digitalocean.com/community/tutorials/how-to-do-math-in-go-with-operators
var makeplanfortomorrow_coffient = 0.5
if count_makeplanfortomorrow == 0{
  //total_score = total_score/4*3
  makeplanfortomorrow_coffient = 0.75
  //plancoffient = 2
}else{
  //total_score = total_score *1
  makeplanfortomorrow_coffient = 1.0
  //plancoffient = 4
}
fmt.Println(makeplanfortomorrow_coffient)

if priority_execute_coffient < 0.1 {
  priority_execute_coffient = 0.1
}

var noflinch_coefficient = 0.5 
if noflinch_number !=0{
  noflinch_coefficient = noflinch_coefficient*float64(noflinch_number)*2.0+ noflinch_coefficient
}

//-----------------------------------everygoal score---------------------------------
total_score = self_discipline_score+noflinch_coefficient*atom_coffient*priority_execute_coffient*float64(total_score)*makeplanfortomorrow_coffient*planobey_coffient*pain_coeffient*difficult_coeffient

fmt.Println("-------------alll kinds of coeffient is fellowing------------")
fmt.Printf("the pain coeffient %f\n",pain_coeffient)
fmt.Printf("the plan obey coeffient %f\n",planobey_coffient)
fmt.Printf("the diffcult coeffient %f\n",difficult_coeffient)
fmt.Printf("make plan for tomorrow  coeffient %f\n",makeplanfortomorrow_coffient)
fmt.Printf("total score is  %f\n",total_score)
fmt.Printf("task execute coeffient  %f\n",priority_execute_coffient)
fmt.Printf("atom coeffient  %f\n",atom_coffient)
fmt.Printf("no flinch coeffient  %f\n",noflinch_coefficient)
fmt.Println("-------------alll kinds of coeffient is above------------")

if math.IsNaN(total_score){
  total_score = 0.0
}

review := &Reviewdatadetail{Self_discipline_score:self_discipline_score,Totalscore:total_score,Executeability_score:executeability_score,Noflinch:noflinch_score,Setarecord:setarecord_score,Conquerthefear:conquerthefear_score,Depthfirstsearch:dfs_score,Useprinciple:useprinciple_score,Attackactively:attackactively_score,Solveakeyproblem:solveakeyproblem_score,Acceptpain:acceptpain_score,Acceptfactandseektruth:acceptfactandseektruth_score,Buildframeandprinciple:buildframeandprinciple_score,Challengethings:challengetag_score,Markataskimmediately:markataskimmediately_score,Doanimportantthingearly:doanimportantthingearly_score,Alwaysprofit:alwaysprofit_score,Atomadifficulttask:atomadifficulttask_score,Onlystartatask:onlystartatask_score,Thenumberoftasks_score:taskcount_score,Difficultthings:difficultthings_score,Threeminutes:threeminutes_score,Getlesson:getlesson_score,Learntechuse:learntechuse_score,Patience:patience_score,Serviceforgoal_score:serviceforgoal_score,Usebrain:brainuse_score,Battlewithlowerbrain:battlewithlowerbrain_score,Learnnewthings:learnnewthings_score,Makeuseofthingsuhavelearned:makeuseofthethingsuhavelearned_score}

color.Red("We have red")

fmt.Println(markataskimmediately_number)
reviewfortimecount_from_client := Reviewfortimescount{Self_discipline_number:self_discipline_number,Email:email,Noflinch:noflinch_number,Challengethings:challengetag_number,Conquerthefear:conquerthefear_number,Setarecord:setarecord_number,Depthfirstsearch:dfs_number,Date:date,Useprinciple:useprinciple_number,Attackactively:attackactively_number,Acceptpain:acceptpain_number,Solveakeyproblem:solveakeyproblem_number,Acceptfactandseektruth:acceptfactandseektruth_number,Atomadifficulttask:atomadifficulttask_number,Serviceforgoal_score:serviceforgoal_number,Doanimportantthingearly:doanimportantthingearly_number,Makeuseofthingsuhavelearned:makeuseofthethingsuhavelearned_number,Difficultthings:difficultthings_number,Learnnewthings:learnnewthings_number,Threeminutes:threeminutes_number,Alwaysprofit:alwaysprofit_number,Markataskimmediately:markataskimmediately_number,Usebrain:usebrainnumber,Battlewithlowerbrain:battlewithlowerbrainnumber,Buildframeandprinciple:buildframeandprinciplenumber,Patience:patiencenumber}

//https://stackoverflow.com/questions/8270816/converting-go-struct-to-json

reviewjson, err := json.Marshal(review)
if err!=nil{
fmt.Println("----------there is an error ----------")
fmt.Println(err)
}

reviewstring := string(reviewjson)
fmt.Println("-------------i am pritning reviewstring---------------")
fmt.Printf("%+v\n", review)
//fmt.Println(reviewjson)
fmt.Println(reviewstring)
fmt.Println("-------------i am pritning reviewstring---------------")


fmt.Println(total_score)

total_score_float_precision := fmt.Sprintf("%.4f", total_score)
set_status, err1 := redis.Int(redisDBcon.Do("SET", "initscore",total_score_float_precision))

fmt.Println(set_status)
if err1 != nil{
  fmt.Println(err)
}


r, err := redis.Int(redisDBcon.Do("PUBLISH", "channel_1",total_score_float_precision))

fmt.Println(r)
if err != nil{
  fmt.Println(err)
}


var reviewday Reviewofday
var reviewfortimecount Reviewfortimescount
//这里可能存在bug，我需要非常注意
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewday)
db.Model(&reviewday).Update("Details", reviewstring)
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewfortimecount)
db.Model(&reviewfortimecount).Updates(reviewfortimecount_from_client)
// db.Model(&reviewfortimecount).Update("Patiencenumber", patiencenumber)
// db.Model(&reviewfortimecount).Update("Battlewithlowerbrain", battlewithlowerbrainnumber)
// db.Model(&reviewfortimecount).Update("Usebrain", usebrainnumber)
if (total_score <= 0.0) {
  total_score = 0.01
}
return total_score
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
