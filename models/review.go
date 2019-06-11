package models

import (
"fmt"
"encoding/json"
"net/http"
"github.com/jinzhu/gorm"
"strconv"
//"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"

"github.com/tidwall/gjson"

)


//from now on, i only to modify this file to add gtd review standard,that includes, detail struct and if control and totalscore and review struct instance data

//json is that it will be changed to this string json in db
//https://stackoverflow.com/questions/26327391/json-marshalstruct-returns
type Reviewdatadetail struct{
Totalscore    int `json:"totalscore"`
Averagescoreofhistory   int `json:"averagescoreofhistory"`
Patience      int  `json:"patience"`
Attackactively      int  `json:"attackactively"`
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
Alwaysprofit       int     `json:"alwaysprofit"` 
Markataskimmediately int   `json:"markataskimmediately"`
Doanimportantthingearly int  `json:"doanimportantthingearly"`
Buildframeandprinciple    int `json:"buildframeandprinciple"`
Acceptfactandseektruth    int `json:"acceptfactandseektruth"`
Acceptpain                int `json:"acceptpain"`
Solveakeyproblem                int `json:"solveakeyproblem"`
}




type Reviewofday  struct {
  gorm.Model
  Date string   `json:"date"`
  Email    string   `json:"email"`     
  Details string `json:"details" sql:"type:text;"`

   }


 










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
}







func Reviewalgorithmjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  var reviewdays []Reviewofday
  db.Where("email =  ?", email).Order("date").Find(&reviewdays)
  //if u set the len,u will get the size of slice
  if len(reviewdays)<33{
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }else{
    reviewdays = reviewdays[len(reviewdays)-31:]
  
    c.JSON(200, gin.H{
      //"reviewdata":review30days,
      "reviewdata":reviewdays,
    })
  }
 

}


func Errorlog(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

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



func Searchwithtags(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)
  var keywords = c.Query("keywords")
  var search []Tasks
  //var s string = "12312sf"
  querystring:= "select * from tasks where status not in ('finished','finish','giveup','g') and  email =" +`"`+ email +`" `+ " and tasktags REGEXP "+"'"+`"`+keywords+`"`+":[ ]{0,1}"+`"yes"`+"'"
  //qurystring = fmt.Sprintf("select * from tasks where tasktags REGEXP '%s %s %s",s,"123123")
 // select * from tasks where tasktags REGEXP  '"hardtag":"yes"'\G;
  //db.Where("email =  ?", email).Where("task LIKE ?", "%"+keywords+"%").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&search)
  fmt.Println(querystring)
  db.Raw(querystring).Scan(&search)
  c.JSON(200, gin.H{
      "search":search,
    })

}










func Goalsystem(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  //var goals []Tasks
  var goals []Goalfordbs
  //db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
  db.Where("email =  ?", email).Order("id").Find(&goals)
  c.JSON(200, gin.H{
      "goals":goals,
    })

}



func Goalreviewfortoday(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  //var goals []Tasks
  var goals []Goalfordbs
  //db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
  db.Where("email =  ?", email).Order("id").Find(&goals)
  c.JSON(200, gin.H{
      "goals":goals,
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










func Reviewforios(c *gin.Context) {


c.HTML(http.StatusOK, "reviewalgoforios.html",nil)


}


//this api was used to prepare the data of review
func Reviewforstastics(c *gin.Context){
  //get 7 days review datas 
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  count_need_bystastics_from_client := c.Query("days")
  counts, _:= strconv.Atoi(count_need_bystastics_from_client)


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

if counts == -1{
  tasksbydays = alleverydays[1:2]
}else{
  tasksbydays = alleverydays[0:counts]
}
//fmt.Println(tasksbydays)

var goal_devotedtime = make(map[string]int)
var alltasks_count = 0
var all_time_u_had_devoted_inthe_time_range = 0
var  alltime_goal_oriented = 0 
for  _,item :=range tasksbydays{
  alltasks_count =  alltasks_count+len(item.Alldays)
  for _,item1 := range item.Alldays{
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
}
fmt.Printf("the task length is %d",len(tasksbydays))
fmt.Printf("theses task counts is %d",alltasks_count)
fmt.Printf("u had devoted %d  minutes in the time range",all_time_u_had_devoted_inthe_time_range)
fmt.Printf("u had devoted %d  minutes in the time range for goal",alltime_goal_oriented)
  var reviewsfortimescount []Reviewfortimescount
  db.Where("email =  ?", email).Order("date").Find(&reviewsfortimescount)
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
      "goaltime":goal_devotedtime,
      "reviewdata":reviewdata,
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








//compute total_scores of someday

func Compute_singleday(date string,email string) string{
//https://tour.golang.org/basics/10
fmt.Println("------------ i am here to compute the single day---------------------------")
var tasks []Tasks
//email := "yang756260386@gmail.com"
var brainuse_score,makeuseofthethingsuhavelearned_score,difficultthings_score,threeminutes_score,getlesson_score,learntechuse_score,battlewithlowerbrain_score,patience_score,learnnewthings_score int = 0,0,0,0,0,0,0,0,0 
var serviceforgoal_score,onlystartatask_score = 0,0
var atomadifficulttask_score,alwaysprofit_score = 0,0
var doanimportantthingearly_score,markataskimmediately_score = 0,0
var challengetag_score = 0
db.Where("Email= ?", email).Where("finishtime =  ?", date).Order("id desc").Find(&tasks)

var taskcount_score int

var countoffinishedtasks int

var countofgivenuptasks int


//for times stastics
var patiencenumber = 0 
var battlewithlowerbrainnumber = 0
var usebrainnumber = 0
var buildframeandprinciple_score =0 
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
var acceptfactandseektruth_score = 0
var acceptfactandseektruth_number = 0
var acceptpain_score = 0
var acceptpain_number = 0
var solveakeyproblem_score = 0
var solveakeyproblem_number = 0
var attackactively_number = 0
var attackactively_score = 0

var useprinciple_number = 0
var useprinciple_score = 0

db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Not("status", []string{"unfinished","unfinish"}).Count(&countoffinishedtasks)


db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Where("status =?","giveup").Count(&countofgivenuptasks)




fmt.Println("-----=======-------++++++++++++++=-----======----------------")
fmt.Println(countofgivenuptasks)



taskcount_score =  2  * (countoffinishedtasks - countofgivenuptasks) +countofgivenuptasks*1

 for _,item :=range tasks{
fmt.Println("------------i had been into loop----------------")

var jsonoftasktags = item.Tasktags
if  challengetag := gjson.Get(jsonoftasktags, "challengetag").String();challengetag=="yes"{
challengetag_score  = challengetag_score + 5
}




var json = item.Reviewdatas
fmt.Println(json)
fmt.Println("------------i had been into loop----------------")
if  brainuse := gjson.Get(json, "brainuse").String();brainuse=="yes"{
fmt.Println(brainuse)
brainuse_score = brainuse_score +5
usebrainnumber = usebrainnumber +1
 } 

if  buildframeandprinciple_from_client := gjson.Get(json, "buildframeandprinciple").String();buildframeandprinciple_from_client=="yes"{
  //fmt.Println(brainuse)
  buildframeandprinciple_score = buildframeandprinciple_score +5
  buildframeandprinciplenumber = buildframeandprinciplenumber +1
   } 

   if  useprinciple_from_client := gjson.Get(json, "useprinciple").String();useprinciple_from_client=="yes"{
    //fmt.Println(brainuse)
    useprinciple_score = useprinciple_score +10
    useprinciple_number = useprinciple_number +1
     } 





   if  acceptfact_from_client := gjson.Get(json, "acceptfactandseektruth").String();acceptfact_from_client=="yes"{
    //fmt.Println(brainuse)
    acceptfactandseektruth_score =  acceptfactandseektruth_score  +  10
    acceptfactandseektruth_number = acceptfactandseektruth_number +1
     } 


     if  attackactively_from_client := gjson.Get(json, "attackactively").String();attackactively_from_client=="yes"{
      //fmt.Println(brainuse)
      attackactively_score=  attackactively_score  +  10
      attackactively_number= attackactively_number +1
       } 






     if  acceptfact_from_client := gjson.Get(json, "acceptpain").String();acceptfact_from_client=="yes"{
      //fmt.Println(brainuse)
      acceptpain_score =  acceptpain_score  +  10
      acceptpain_number = acceptpain_number +1
       } 




if  makeuseofthings := gjson.Get(json, "makeuseofthings").String();makeuseofthings=="yes"{
makeuseofthethingsuhavelearned_score = makeuseofthethingsuhavelearned_score + 5
makeuseofthethingsuhavelearned_number = makeuseofthethingsuhavelearned_number +1
 }


if  doanimportantthingearly := gjson.Get(json, "doanimportantthingearly").String();doanimportantthingearly =="yes"{
doanimportantthingearly_score = doanimportantthingearly_score + 10
doanimportantthingearly_number = doanimportantthingearly_number + 1
 }
 

if  markataskimmediately := gjson.Get(json, "markataskimmediately").String();markataskimmediately =="yes"{
markataskimmediately_score = markataskimmediately_score + 1
markataskimmediately_number = markataskimmediately_number + 1
 }





if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
alwaysprofit_score = alwaysprofit_score + 5
alwaysprofit_number = alwaysprofit_number +1
 }


// if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
// alwaysprofit_score = alwaysprofit_score + 5
//  }




if  learnnewthings := gjson.Get(json, "learnnewthings").String();learnnewthings=="yes"{
learnnewthings_score = learnnewthings_score +5
learnnewthings_number = learnnewthings_number + 1
 }


if  serviceforgoal := gjson.Get(json, "serviceforgoal").String();serviceforgoal=="yes"{
serviceforgoal_score  = serviceforgoal_score  + 20
 }


if  onlystartatask := gjson.Get(json, "onlystartatask").String();onlystartatask=="yes"{
onlystartatask_score  = onlystartatask_score  + 10
 }





if  battlewithlowerbrain := gjson.Get(json, "battlewithlowerbrain").String();battlewithlowerbrain=="yes"{
battlewithlowerbrain_score = battlewithlowerbrain_score +5
battlewithlowerbrainnumber = battlewithlowerbrainnumber + 1
 }


if  atomadifficulttask := gjson.Get(json, "atomadifficulttask").String();atomadifficulttask=="yes"{
atomadifficulttask_score = atomadifficulttask_score +5
atomadifficulttask_number = atomadifficulttask_number+1
 }






if  patience := gjson.Get(json, "patience").String();patience=="yes"{
patience_score = patience_score + 10
patiencenumber = patiencenumber + 1
 }


 if  solveakeyproblem := gjson.Get(json, "solveakeyproblem").String();solveakeyproblem=="yes"{
  patience_score = solveakeyproblem_score + 50
  solveakeyproblem_number = solveakeyproblem_number + 1
   }





if  difficultthings := gjson.Get(json, "difficultthings").String();difficultthings=="yes"{
difficultthings_score = difficultthings_score +10
difficultthings_number = difficultthings_number + 1
 }



if  threeminutes := gjson.Get(json, "threeminutes").String();threeminutes=="yes"{
threeminutes_score = threeminutes_score +5
threeminutes_number = threeminutes_number + 1
 }



if  getlesson:= gjson.Get(json, "getlesson").String();getlesson=="yes"{
getlesson_score = getlesson_score +5

 }



if  learntechuse := gjson.Get(json, "learntechuse").String();learntechuse=="yes"{
learntechuse_score = learntechuse_score +5

 }


}

total_score:=acceptfactandseektruth_score+useprinciple_score+attackactively_score+solveakeyproblem_score+acceptpain_score+buildframeandprinciple_score+taskcount_score+doanimportantthingearly_score+atomadifficulttask_score+onlystartatask_score+markataskimmediately_score+challengetag_score + brainuse_score+alwaysprofit_score + makeuseofthethingsuhavelearned_score + battlewithlowerbrain_score + patience_score + learnnewthings_score+difficultthings_score+threeminutes_score+getlesson_score+learntechuse_score + serviceforgoal_score
review := &Reviewdatadetail{Totalscore:total_score,Useprinciple:useprinciple_score,Attackactively:attackactively_score,Solveakeyproblem:solveakeyproblem_score,Acceptpain:acceptpain_score,Acceptfactandseektruth:acceptfactandseektruth_score,Buildframeandprinciple:buildframeandprinciple_score,Challengethings:challengetag_score,Markataskimmediately:markataskimmediately_score,Doanimportantthingearly:doanimportantthingearly_score,Alwaysprofit:alwaysprofit_score,Atomadifficulttask:atomadifficulttask_score,Onlystartatask:onlystartatask_score,Thenumberoftasks_score:taskcount_score,Difficultthings:difficultthings_score,Threeminutes:threeminutes_score,Getlesson:getlesson_score,Learntechuse:learntechuse_score,Patience:patience_score,Serviceforgoal_score:serviceforgoal_score,Usebrain:brainuse_score,Battlewithlowerbrain:battlewithlowerbrain_score,Learnnewthings:learnnewthings_score,Makeuseofthingsuhavelearned:makeuseofthethingsuhavelearned_score}
reviewfortimecount_from_client := Reviewfortimescount{Email:email,Date:date,Useprinciple:useprinciple_number,Attackactively:attackactively_number,Acceptpain:acceptpain_number,Solveakeyproblem:solveakeyproblem_number,Acceptfactandseektruth:acceptfactandseektruth_number,Atomadifficulttask:atomadifficulttask_number,Serviceforgoal_score:serviceforgoal_number,Doanimportantthingearly:doanimportantthingearly_number,Makeuseofthingsuhavelearned:makeuseofthethingsuhavelearned_number,Difficultthings:difficultthings_number,Learnnewthings:learnnewthings_number,Threeminutes:threeminutes_number,Alwaysprofit:alwaysprofit_number,Markataskimmediately:markataskimmediately_number,Usebrain:usebrainnumber,Battlewithlowerbrain:battlewithlowerbrainnumber,Buildframeandprinciple:buildframeandprinciplenumber,Patience:patiencenumber}

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

var reviewday Reviewofday
var reviewfortimecount Reviewfortimescount
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewday)
db.Model(&reviewday).Update("Details", reviewstring)
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewfortimecount)
db.Model(&reviewfortimecount).Updates(reviewfortimecount_from_client)
// db.Model(&reviewfortimecount).Update("Patiencenumber", patiencenumber)
// db.Model(&reviewfortimecount).Update("Battlewithlowerbrain", battlewithlowerbrainnumber)
// db.Model(&reviewfortimecount).Update("Usebrain", usebrainnumber)
return reviewstring
}






