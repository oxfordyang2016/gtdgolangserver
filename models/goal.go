
package models
import(
"fmt"
//"time"
//"github.com/bradfitz/slice"
//"encoding/json"
"net/http"
"github.com/jinzhu/gorm"
//"strconv"
//"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"

//"github.com/tidwall/gjson"
)

type(
Goals  struct{
	Name                    string
	Allprojectsingoal       []Projects
	}

Goalfordbs  struct{
	gorm.Model
	Name                   string  `json:"name"`
	//ID uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	//i will use email+ab(2 alphebet table),such as yang756260386@gmail.comab
	Goalcode             string    `json:"goalcode"`
	Email                 string   `json:"email"`
	}



Goalsincludetasks struct{
	Name  string
	Alltasksingoal    []Tasks

}	

)





func Goalsjson(c *gin.Context) {

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
	   //loc, _ := time.LoadLocation("Asia/Shanghai")
	   //today :=  time.Now().In(loc).Format("060102")
	   //tomorrow :=  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
       //db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime in (?)", []string{today, tomorrow}).Order("id desc").Find(&tasks)
	   statusbool:=c.Query("statusbool")  
	   if statusbool =="yes"{
		db.Where("Email= ?", email).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
	   }else{
		db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("goal", []string{"no goal",""}).Order("id desc").Find(&tasks)
	   }
	    
	   client:= c.Request.Header.Get("client")
	   querytype:=c.Query("type")
	   
	   

	   fmt.Println("+++++++client is++++++++")
	   fmt.Println(querytype)
	   fmt.Println(client)
	  fmt.Println("+++++++client is++++++++")
	  //use maps to aviod to design complex algorithm

	   
	   
	   alltasks_ingoal:=make(map[string] []Tasks)

	
    
	   for _,item :=range tasks{
		   if item.Goal == ""{
			alltasks_ingoal["no goal"]=append(alltasks_ingoal["no goal"],item)
		   }else{
			alltasks_ingoal[item.Goal]=append(alltasks_ingoal[item.Goal],item)
		   }
		
	 }
	 var allgoalsonlyincludetasks []Goalsincludetasks
	 for k,v:= range alltasks_ingoal{
		 allgoalsonlyincludetasks = append(allgoalsonlyincludetasks,Goalsincludetasks{k,v})

	 }
	 fmt.Println("========i am here 1========")
   if querytype == "noproject"{
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoalsonlyincludetasks})
	fmt.Println("========i am here========")
	return
   }
	 
   fmt.Println("========i am here 3========")

	   var goals []string 
	   for k :=range(alltasks_ingoal){
        goals = append(goals,k)
	   }
	   fmt.Println(len(goals))
	   allclassgoals := make(map[string] []Projects) 
	    for _,key := range(goals){
		    allclassproject:=make(map[string] []Tasks)
		    for _,item := range(alltasks_ingoal[key]){
			  //projects -- goal ,firsr to generate project and then goals
			  allclassproject[item.Project]=append(allclassproject[item.Project],item)
		  }
		  var allprojects []Projects
		  for k,v := range allclassproject{
			 allprojects =append(allprojects,Projects{k,v})
	 
		  }
		  allclassgoals[key] = allprojects

	  }
	
	  var allgoals []Goals
	  for k,v := range allclassgoals{
		allgoals =append(allgoals,Goals{k,v})

	 }


	   //fmt.Println(allclassproject["gtd1"])
      c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoals})
  
		}
  
  