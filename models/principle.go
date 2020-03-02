
package models
import(
"fmt"
// "time"
// "sort"
//"github.com/bradfitz/slice"
//"encoding/json"
"net/http"
// "github.com/fatih/color"
"github.com/jinzhu/gorm"
//"strconv"
//"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"
"github.com/tidwall/gjson"
//"github.com/tidwall/gjson"
)

type(
// Goals  struct{
// 	Name                    string
// 	Priority                int
// 	Allprojectsingoal       []Projects
// 	}

Principledetails  struct{
	gorm.Model
	Principlecode                   string  `json:"PrincipleName"`
	Email                     string   `json:"email"`
	//ID uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	//i will use email+ab(2 alphebet table),such as yang756260386@gmail.comab
	Detailitem             string    `json:"Detailitem"`
	}



	// Goalfordbs  struct{
	// 	gorm.Model
	// 	Name                   string  `json:"name"`
	// 	//ID uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	// 	//i will use email+ab(2 alphebet table),such as yang756260386@gmail.comab
	// 	Goalcode             string    `json:"goalcode"`
	// 	Priority             int    `json:"priority"`
	// 	Email                 string   `json:"email"`
	// 	Goalstatus            string   `json:"goalstatus"`
	// 	Plantime            string   `json:"plantime"`
	// 	Finishtime            string   `json:"finishtime"`
	// 	Chinesename            string   `json:"chinesename"`
	// 	Timerange               int   `json:"timerange"`   //制定目标的时间范围
	// 	}





	//该原则系统将会被加载到标签系统上面去
	Principlefordbs  struct{
		gorm.Model
		Principlename                   string  `json:"principlename"`
		//ID uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
		//i will use email+ab(2 alphebet table),such as yang756260386@gmail.comab
		Principlecode             string    `json:"Principlecode"`
		// Priority             int    `json:"priority"`
		Email                     string   `json:"email"`
		Describe                  string   `json:"describe"`
		// Goalstatus            string   `json:"goalstatus"`
		// Plantime            string   `json:"plantime"`
		// Finishtime            string   `json:"finishtime"`
		Chinesename            string   `json:"chinesename"`
		// Timerange               int   `json:"timerange"`   //制定目标的时间范围
		}





// Goalsincludetasks struct{
// 	Name  string
// 	Priority                int
// 	Devotedtime int 
// 	Alltasksingoal    []Tasks

// }	

)

func Createprinciplelib(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie,_:=c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email:=emailcookie.Value
	principlename := gjson.Get(reqBody, "principlename").String()
	// priority:= gjson.Get(reqBody, "priority").Int()
	
	var principlesforemail []Principlefordbs
	// var goalsforemail []Goalfordbs 
	// var principlecountforsamegoal = 0
	var goalcountforsamegoal = 0
	// var goalcount = 0
	var principlecount = 0

	alphabet :=[26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    db.Where("Email= ?", email).Where("Name=?",principlename).Find(&principlesforemail).Count(&principlecount)
    if goalcountforsamegoal >0{
      c.JSON(200, gin.H{
        "status":  "posted",
        "message": "the goal had ben created,not repeated youself",
      })
      return
    }else{
   db.Where("Email= ?", email).Find(&principlesforemail).Count(&principlecount)
   if principlecount ==0{
	principlefromclient := Principlefordbs{Principlename:principlename,Email:email,Principlecode:"aaa"}
    db.Create(&principlefromclient).Scan(&principlefromclient)
   }else{
	principlecode := principlesforemail[len(principlesforemail)-1].Principlecode
    a,b,c := string(principlecode[0]),string(principlecode[1]),string(principlecode[2])
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
	//这里原来引号出了问题。。。。多出了一个空格
    principlefromclient := Principlefordbs{Principlename:principlename,Email:email,Principlecode:fmt.Sprintf("%s%s%s",a,b,c)}
	db.Create(&principlefromclient).Scan(&principlefromclient)

   } 
  }
  c.JSON(200, gin.H{
	"result":"u have created task",
  })
}



func Createprincipledetail(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie,_:=c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email:=emailcookie.Value
	principledetail := gjson.Get(reqBody, "principledetail").String()
	principlecode := gjson.Get(reqBody, "principlecode").String()
	// priority:= gjson.Get(reqBody, "priority").Int()
	//这里原来引号出了问题。。。。多出了一个空格
    principledetailsfromclient := Principledetails{Detailitem:principledetail,Email:email,Principlecode:principlecode}
	db.Create(&principledetailsfromclient).Scan(&principledetailsfromclient)
	c.JSON(200, gin.H{
		"result":"u have created single principle",
	  })
}













func Updateprinciple(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie,_:=c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email:=emailcookie.Value
	goal := gjson.Get(reqBody, "goal").String()
	fmt.Printf("---goal is------%s-----\n",goal)
	goalcode := gjson.Get(reqBody, "goalcode").String()
       fmt.Println(goalcode)	
       goalstatus := gjson.Get(reqBody, "goalstatus").String()
	finishtime := gjson.Get(reqBody, "finishtime").String()
	plantime := gjson.Get(reqBody, "plantime").String()
	priority:= gjson.Get(reqBody, "priority").Int()
	timerange:= gjson.Get(reqBody, "timerange").Int()
fmt.Println(timerange)	
var goalindb  Goalfordbs
	db.Where("Email= ?", email).Where("Goalcode= ?",goalcode).Find(&goalindb)
	if priority != -1  {db.Model(&goalindb).Update("Priority", int(priority)) }
	if goal != "unspecified"{if goal!= "nocontent"{db.Model(&goalindb).Update("Name", goal)}}
	if timerange !=0{db.Model(&goalindb).Update("Timerange", int(timerange)) }
	if goalstatus !="unspecified"{db.Model(&goalindb).Update("Goalstatus", goalstatus)}
	if finishtime !="unspecified"{db.Model(&goalindb).Update("Finishtime", finishtime)}
	if plantime !="unspecified"{db.Model(&goalindb).Update("Plantime", plantime)}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK })
}





// func Goalcompare(c *gin.Context) {
//   //i use email as identifier
// //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
//   emailcookie,err:=c.Request.Cookie("email")
//   //fmt.Println(emailcookie.Value)
//   var email string
//    if err!=nil{
//      email = c.Request.Header.Get("email")
//    }else{
//      fmt.Println(emailcookie.Value)
//      email =emailcookie.Value
//    }
  
// fmt.Println(email)
  

// type Result struct {
// 	Name string
// 	Goalcode string
// 	Priority int
	
// }

// loc, _ := time.LoadLocation("Asia/Shanghai")
// var result []Result
// today :=  time.Now().In(loc).Format("060102")
//    db.Raw(`SELECT name,goalcode,priority  FROM goalfordbs  WHERE email ="`+email+`"`+" and goalstatus not in ("+`"giveup","g","finished","finish"`+`) and `+ ` name   NOT IN (SELECT goal  FROM tasks  WHERE finishtime=` +`"`+today+`"`+` and email =`+`"`+email+`"`+`);`).Scan(&result)
//    color.Red("red")
//    fmt.Println(result)
//    c.JSON(200, gin.H{
//       "undevotedgoals":result,
//     })

// }




// //使用目标的goalcode
// func Searchwithgoalcode(c *gin.Context) {
// 	//i use email as identifier
//   //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
// 	emailcookie,err:=c.Request.Cookie("email")
// 	//fmt.Println(emailcookie.Value)
// 	var email string
// 	 if err!=nil{
// 	   email = c.Request.Header.Get("email")
// 	 }else{
// 	   fmt.Println(emailcookie.Value)
// 	   email =emailcookie.Value
// 	 }
	
// 	goalcode :=  c.Query("goalcode")
// 	tasktag :=  c.Query("tasktag")
	
//     fmt.Println(email,tasktag)
	
  
//   type Simpletask struct {
// 	  Id string
// 	  Task string
// 	  Status string
// 	  Parentid string `json:"parentid"`
//   }
  
//   type Simplegoal struct {
// 	Goalcode string
// 	Name string
// 	// Taskstag string
// }

// //   loc, _ := time.LoadLocation("Asia/Shanghai")
//      var result []Simpletask
// //   today :=  time.Now().In(loc).Format("060102")
//      querystring := `SELECT id,task,status,parentid  FROM tasks  WHERE  email ="`+email+`"`+"and status not in ("+`"giveup","g","finished","finish"`+`) and `+ ` goal  IN (SELECT name  FROM goalfordbs  WHERE goalcode =` +`"`+goalcode+`"`+` and email =`+`"`+email+`"`+`);`
	 
// 	 if tasktag != ""&&tasktag != "notag"{
// 		querystring = `SELECT id,task,status,parentid  FROM tasks  WHERE  email ="`+email+`"`+" and tasktags REGEXP "+"'"+`"`+tasktag+`"`+":[ ]{0,1}"+`"yes"`+"'"+" and status not in ("+`"giveup","g","finished","finish"`+`) and `+ ` goal  IN (SELECT name  FROM goalfordbs  WHERE goalcode =` +`"`+goalcode+`"`+` and email =`+`"`+email+`"`+`);`
// 	 }
// 	 var goalinfo Simplegoal
// 	 goalquery := `SELECT name,goalcode  FROM goalfordbs  WHERE goalcode =` +`"`+goalcode+`"`+` and email =`+`"`+email+`"`+`;`
// 	 fmt.Println(querystring)
// 	 db.Raw(goalquery).Scan(&goalinfo)
// 	 db.Raw(querystring).Scan(&result)
// 	 color.Red("red")
// 	 fmt.Println(result)
// 	 fmt.Println("-----------")
// 	 fmt.Printf("%#v----\n",goalinfo.Goalcode)
// 	 c.JSON(200, gin.H{
// 		 "taskstag":tasktag,
// 		 "goalinfo":goalinfo,
// 		"goalcode_unfinishedtask":result,
// 	  })
  
//   }



// //   `and tasktags REGEXP "+"'"+`"`+keywords+`"`+":[ ]{0,1}"+`"yes"``












func Principlesystem(c *gin.Context) {
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



 
	//var goals []Tasks
	var principlelibs []Principlefordbs
	//db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
//db.Where("email =  ?", email).Not("goalstatus", []string{"finished"}).Order("id").Find(&goals)
db.Where("email =  ?", email).Order("id").Find(&principlelibs)
fmt.Println(principlelibs)
	c.JSON(200, gin.H{
		"principlelibs":principlelibs,
	  })
  
  }
  
  
  
//   func Goalreviewfortoday(c *gin.Context) {
// 	//i use email as identifier
//   //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
// 	emailcookie,_:=c.Request.Cookie("email")
// 	fmt.Println(emailcookie.Value)
// 	email:=emailcookie.Value
// 	//fmt.Println(cookie1.Value)
  
// 	//var goals []Tasks
// 	var goals []Goalfordbs
// 	//db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
// 	db.Where("email =  ?", email).Order("id").Find(&goals)
// 	c.JSON(200, gin.H{
// 		"goals":goals,
// 	  })
  
//   }
  





//    func Get_goal_coffient(goal string,email string) float64{
// 	   var total_priotity = 0
// 	   var goals_belonged2email  []Goalfordbs
// 	   db.Where("Email= ?", email).Find(&goals_belonged2email)
// 	   for _,item :=range goals_belonged2email{  
// 	   total_priotity = total_priotity + item.Priority}
// 	   var querygoal Goalfordbs
//        db.Where("Email= ?", email).Where("Name = ?",goal).Find(&querygoal)
// 	   goal_coffient  :=  float64(querygoal.Priority)/float64(total_priotity)
// 	   fmt.Println("-------------------------goal coeffient-------------------------")
// 	   fmt.Println("hahhahh")
// 	   fmt.Println(goal_coffient)
// 	   return goal_coffient
//    }










// func Goalsjson(c *gin.Context) {
   
//  	//the algorithm can be upgrade
// 				//i use email as identifier
// 			  //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
// 		emailcookie,_:=c.Request.Cookie("email")
// 		fmt.Println(emailcookie.Value)
// 		email:=emailcookie.Value
// 		var tasks []Tasks
// 		//fmt.Println(cookie1.Value)
// 				//email:="yangming1"
// 	   // db.Where("Email= ?", email).Find(&tasks)
// 	   loc, _ := time.LoadLocation("Asia/Shanghai")
// 	   today :=  time.Now().In(loc).Format("060102")
// 	   tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
//        //db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime in (?)", []string{today, tomorrow}).Order("id desc").Find(&tasks)
// 	   statusbool:=c.Query("statusbool")
// 	   //https://blog.csdn.net/u013474436/article/details/71516649
// 	   time_range := c.Request.Header.Get("time-range")
// 	//    fmt.Println(time_range)
// 	//    fmt.Println("status bool is ",statusbool)
// 	//    get   all days in this week or month
// 	   week_day := Getweekday()  
// 	   month_day := Getmonthday() 
// 	   if statusbool =="yes"{
// 		fmt.Println("******************")
// 		  if time_range == "week"{
			 
// 			db.Where("Email= ?", email).Where("finishtime in (?)", week_day).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		  }else if  time_range == "month"{
// 			db.Where("Email= ?", email).Where("finishtime in (?)", month_day).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		  }else{
// 			db.Where("Email= ?", email).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		  }
		
// 	   }else{
// 		fmt.Println("----------------------")
// 		if time_range == "week"{
// 		db.Where("Email= ?", email).Where("finishtime in (?)",week_day).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today,tomorrow}).Order("id desc").Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		}else if time_range == "month"{
// 			db.Where("Email= ?", email).Where("finishtime in (?)",month_day).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today,tomorrow}).Order("id desc").Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		}else{
// 			db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today,tomorrow}).Order("id desc").Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
// 		}  
// 	}
		
	   
// 	   client:= c.Request.Header.Get("client")
// 	   querytype:=c.Query("type")
	   
	   

// 	   fmt.Println("+++++++client is++++++++")
// 	   fmt.Println(querytype)
// 	   fmt.Println(client)
// 	  fmt.Println("+++++++client is++++++++")
// 	  //use maps to aviod to design complex algorithm

	   
	   
// 	   alltasks_ingoal:=make(map[string] []Tasks)
// 	   devotedtime_for_goal:=make(map[string] int)
	
    
// 	   for _,item :=range tasks{
// 		   if item.Goal == ""{
// 			alltasks_ingoal["no goal"]=append(alltasks_ingoal["no goal"],item)
// 			devotedtime_for_goal["no goal"] = devotedtime_for_goal["no goal"] + item.Devotedtime
// 		   }else{
// 			alltasks_ingoal[item.Goal]=append(alltasks_ingoal[item.Goal],item)
// 			devotedtime_for_goal[item.Goal] = devotedtime_for_goal[item.Goal] + item.Devotedtime
// 		   }
		
// 	 }
// 	 var allgoalsonlyincludetasks []Goalsincludetasks
// 	 for k,v:= range alltasks_ingoal{
// 		var goal Goalfordbs
// 		var	goalcountforsamegoal int
// 		var goal_level = 0
// 		// find product with id 1
// 		// db.First(&goal, "Name = ?", "L1212") 
// 		db.Where("Email= ?", email).Where("Name=?",k).Find(&goal).Count(&goalcountforsamegoal)
// 		if goalcountforsamegoal == 0{
// 		  print(goal_level)
// 		  goal_level = 0
// 		}else{
// 			goal_level = goal.Priority
// 			if k=="no goal"{
// 				goal_level = 0
// 			}

// 		}
// 		 allgoalsonlyincludetasks = append(allgoalsonlyincludetasks,Goalsincludetasks{Name:k,Priority:goal_level,Devotedtime:devotedtime_for_goal[k],Alltasksingoal:v})

// 	 }
// 	 fmt.Println("========i am here 1========")
//    if querytype == "noproject"{
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoalsonlyincludetasks})
// 	fmt.Println("========i am here========")
// 	return
//    }
	 
//    fmt.Println("========i am here 3========")

// 	   var goals []string 
// 	   for k :=range(alltasks_ingoal){
//         goals = append(goals,k)
// 	   }
// 	   fmt.Println(len(goals))
// 	   allclassgoals := make(map[string] []Projects) 
// 	    for _,key := range(goals){
// 		    allclassproject:=make(map[string] []Tasks)
// 		    for _,item := range(alltasks_ingoal[key]){
// 			  //projects -- goal ,firsr to generate project and then goals
// 			  allclassproject[item.Project]=append(allclassproject[item.Project],item)
// 		  }
// 		  var allprojects []Projects
// 		  for k,v := range allclassproject{
// 			 allprojects =append(allprojects,Projects{k,v})
	 
// 		  }
// 		  allclassgoals[key] = allprojects

// 	  }
	
// 	  var allgoals []Goals
// 	  for k,v := range allclassgoals{
// 		var goal Goalfordbs
// 		var	goalcountforsamegoal int
// 		var goal_level = 0
// 		// find product with id 1
// 		// db.First(&goal, "Name = ?", "L1212") 

// 		db.Where("Email= ?", email).Where("Name=?",k).Find(&goal).Count(&goalcountforsamegoal)
                

//                 fmt.Println(goal.Name)
//                 fmt.Println(goal.Priority)

// 		if goalcountforsamegoal == 0{
// 		  print(goal_level)
// 		  goal_level = 0
// 		}else{
// 			goal_level = goal.Priority
// 			if k=="no goal"{
// 				goal_level = 0
// 			}

// 		}
// 		allgoals =append(allgoals,Goals{k,goal_level,v})

// 	 }
	  
// 	 sort.Slice(allgoals, func(i, j int) bool {
// 		return allgoals[i].Priority > allgoals[j].Priority
// 	  })

// 	   //fmt.Println(allclassproject["gtd1"])
//       c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoals})
  
// 		}
  
  
