package models

import (
	"fmt"
	"sort"
	"time"

	//"github.com/bradfitz/slice"
	//"encoding/json"
	"net/http"

	"github.com/fatih/color"
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

type (
	Goals struct {
		Name              string
		Priority          int
		Goalcode          string
		Allprojectsingoal []Projects
	}

	Goalfordbs struct {
		gorm.Model
		Name string `json:"name"`
		//ID uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
		//i will use email+ab(2 alphebet table),such as yang756260386@gmail.comab
		Goalcode    string `json:"goalcode"`
		Priority    int    `json:"priority"`
		Email       string `json:"email"`
		Goalstatus  string `json:"goalstatus"`
		Plantime    string `json:"plantime"`
		Finishtime  string `json:"finishtime"`
		Chinesename string `json:"chinesename"`
		Timerange   int    `json:"timerange"` //制定目标的时间范围，计划几个月完成
	}

	Goalsincludetasks struct {
		Name           string
		Goalcode       string
		Priority       int
		Devotedtime    int
		Alltasksingoal []Tasks
	}
)

func Creategoal(g *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := g.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie, _ := g.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	goal := gjson.Get(reqBody, "goal").String()
	priority := gjson.Get(reqBody, "priority").Int()
	color.Red("--------ts========")
	fmt.Println(priority)
	var goalsforemail []Goalfordbs
	var goalcountforsamegoal = 0
	var goalcount = 0
	color.Red(goal)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	today := time.Now().In(loc).Format("060102")
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	db.Where("Email= ?", email).Where("Name=?", goal).Find(&goalsforemail).Count(&goalcountforsamegoal)
	if goalcountforsamegoal > 0 {
		g.JSON(902, gin.H{
			"status":  "posted",
			"message": "the goal had ben created,not repeated youself",
		})
		return
	} else {
		db.Where("Email= ?", email).Find(&goalsforemail).Count(&goalcount)
		if goalcount == 0 {

			goalfromclient := Goalfordbs{Name: goal, Plantime: today, Finishtime: "unspecified", Email: email, Goalstatus: "unfinished", Goalcode: "aaa", Priority: int(priority)}
			db.Create(&goalfromclient).Scan(&goalfromclient)
			g.JSON(200, gin.H{
				"status":  "posted",
				"message": "the goal had ben created,not repeated youself",
			})
		} else {
			goalcode := goalsforemail[len(goalsforemail)-1].Goalcode
			a, b, c := string(goalcode[0]), string(goalcode[1]), string(goalcode[2])
			if c != "z" {
				c = alphabet[getindex(c)+1]
			} else {
				if b != "z" {
					c = "a"
					b = alphabet[getindex(b)+1]
				} else {
					c = "a"
					b = "a"
					a = alphabet[getindex(a)+1]
				}
			}
			//这里原来引号出了问题。。。。多出了一个空格
			goalfromclient := Goalfordbs{Name: goal, Plantime: today, Priority: int(priority), Finishtime: "unspecified", Goalstatus: "unfinished", Email: email, Goalcode: fmt.Sprintf("%s%s%s ", a, b, c)}
			db.Create(&goalfromclient).Scan(&goalfromclient)
			g.JSON(200, gin.H{
				"status":  "posted",
				"message": "the goal had ben created,not repeated youself",
			})
		}
	}
}

func Updategoal(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	goal := gjson.Get(reqBody, "goal").String()
	fmt.Printf("---goal is------%s-----\n", goal)
	goalcode := gjson.Get(reqBody, "goalcode").String()
	fmt.Println(goalcode)
	goalstatus := gjson.Get(reqBody, "goalstatus").String()
	finishtime := gjson.Get(reqBody, "finishtime").String()
	plantime := gjson.Get(reqBody, "plantime").String()
	priority := gjson.Get(reqBody, "priority").Int()
	timerange := gjson.Get(reqBody, "timerange").Int()

	//完成时间使用今天
	loc, _ := time.LoadLocation("Asia/Shanghai")
	today := time.Now().In(loc).Format("060102")
	fmt.Println(timerange)
	fmt.Println(plantime)
	var goalindb Goalfordbs
	db.Where("Email= ?", email).Where("Goalcode= ?", goalcode).Find(&goalindb)
	if priority != -1 {
		db.Model(&goalindb).Update("Priority", int(priority))
	}

	//这里很容易出错，如果你真想更新目标的话，你可以更新权重，但是此时你不能更新状态
	if goal != "unspecified" && goalstatus == "unspecified" && goalstatus != "" {
		// ""  giveup  0
		color.Yellow("-----能不能给点面子------")
		if goal != "nocontent" {
			color.Yellow("-----能不能给点面子------")
			color.Yellow(goal)
			color.Yellow(goalstatus)
			fmt.Println(len(goal))
			// color.Yellow(goalstatus)
			db.Model(&goalindb).Update("Name", goal)
		}
	}

	if timerange != 0 {
		db.Model(&goalindb).Update("Timerange", int(timerange))
	}

	if goalstatus != "unspecified" {
		if goalstatus == "f" {
			goalstatus = "finished"
		}
		if goalstatus == "g" {
			goalstatus = "giveup"
		}
		db.Model(&goalindb).Update("Goalstatus", goalstatus)
		if finishtime != "unspecified" {
			db.Model(&goalindb).Update("Finishtime", today)
		}

	}

	// if plantime != "unspecified" {
	// 	db.Model(&goalindb).Update("Plantime", plantime)
	// }
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func Goalcompare(c *gin.Context) {
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
	emailcookie, err := c.Request.Cookie("email")
	//fmt.Println(emailcookie.Value)
	var email string
	if err != nil {
		email = c.Request.Header.Get("email")
	} else {
		fmt.Println(emailcookie.Value)
		email = emailcookie.Value
	}

	fmt.Println(email)

	type Result struct {
		Name     string
		Goalcode string
		Priority int
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	var result []Result
	today := time.Now().In(loc).Format("060102")
	db.Raw(`SELECT name,goalcode,priority  FROM goalfordbs  WHERE email ="` + email + `"` + " and goalstatus not in (" + `"giveup","g","finished","finish"` + `) and ` + ` name   NOT IN (SELECT goal  FROM tasks  WHERE finishtime=` + `"` + today + `"` + ` and email =` + `"` + email + `"` + `);`).Scan(&result)
	color.Red("red")
	fmt.Println(result)
	c.JSON(200, gin.H{
		"undevotedgoals": result,
	})

}

//使用目标的goalcode
func Searchwithgoalcode(c *gin.Context) {
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
	emailcookie, err := c.Request.Cookie("email")
	//fmt.Println(emailcookie.Value)
	var email string
	if err != nil {
		email = c.Request.Header.Get("email")
	} else {
		fmt.Println(emailcookie.Value)
		email = emailcookie.Value
	}

	goalcode := c.Query("goalcode")
	tasktag := c.Query("tasktag")

	fmt.Println(email, tasktag)

	type Simpletask struct {
		Id       string
		Task     string
		Status   string
		Parentid string `json:"parentid"`
	}

	type Simplegoal struct {
		Goalcode string
		Name     string
		// Taskstag string
	}

	//   loc, _ := time.LoadLocation("Asia/Shanghai")
	var result []Simpletask
	//   today :=  time.Now().In(loc).Format("060102")
	querystring := `SELECT id,task,status,parentid  FROM tasks  WHERE  email ="` + email + `"` + "and status not in (" + `"giveup","g","finished","finish"` + `) and ` + ` goal  IN (SELECT name  FROM goalfordbs  WHERE goalcode =` + `"` + goalcode + `"` + ` and email =` + `"` + email + `"` + `);`

	if tasktag != "" && tasktag != "notag" {
		querystring = `SELECT id,task,status,parentid  FROM tasks  WHERE  email ="` + email + `"` + " and tasktags REGEXP " + "'" + `"` + tasktag + `"` + ":[ ]{0,1}" + `"yes"` + "'" + " and status not in (" + `"giveup","g","finished","finish"` + `) and ` + ` goal  IN (SELECT name  FROM goalfordbs  WHERE goalcode =` + `"` + goalcode + `"` + ` and email =` + `"` + email + `"` + `);`
	}
	var goalinfo Simplegoal
	goalquery := `SELECT name,goalcode  FROM goalfordbs  WHERE goalcode =` + `"` + goalcode + `"` + ` and email =` + `"` + email + `"` + `;`
	fmt.Println(querystring)
	db.Raw(goalquery).Scan(&goalinfo)
	db.Raw(querystring).Scan(&result)
	color.Red("red")
	fmt.Println(result)
	fmt.Println("-----------")
	fmt.Printf("%#v----\n", goalinfo.Goalcode)
	c.JSON(200, gin.H{
		"taskstag":                tasktag,
		"goalinfo":                goalinfo,
		"goalcode_unfinishedtask": result,
	})

}

//   `and tasktags REGEXP "+"'"+`"`+keywords+`"`+":[ ]{0,1}"+`"yes"``

func Goalsystem(c *gin.Context) {
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie

	emailcookie, err := c.Request.Cookie("email")
	var email string
	if err != nil {
		email = c.Request.Header.Get("email")
	} else {
		fmt.Println(emailcookie.Value)
		email = emailcookie.Value
	}

	goaltype := c.Query("goaltype")

	//var goals []Tasks
	var goals []Goalfordbs
	//db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
	//db.Where("email =  ?", email).Not("goalstatus", []string{"finished"}).Order("id").Find(&goals)
	if goaltype == "unfinished" {
		db.Where("email =  ?", email).Where("goalstatus=?", "unfinished").Order("id").Find(&goals)
	} else {
		db.Where("email =  ?", email).Order("id").Find(&goals)
	}

	fmt.Println(goals)
	c.JSON(200, gin.H{
		"goals": goals,
	})

}

func Goalreviewfortoday(c *gin.Context) {
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	//fmt.Println(cookie1.Value)

	//var goals []Tasks
	var goals []Goalfordbs
	//db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
	db.Where("email =  ?", email).Order("id").Find(&goals)
	c.JSON(200, gin.H{
		"goals": goals,
	})

}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Goalincludeprojectforserver(c *gin.Context) {
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	c.JSON(200, gin.H{
		"goalincludeproject": Goalincludeproject(email),
	})

}

func Goalincludeproject(email string) map[string][]string {
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
	// emailcookie, _ := c.Request.Cookie("email")
	// fmt.Println(emailcookie.Value)
	// email := emailcookie.Value
	//fmt.Println(cookie1.Value)
	type Result struct {
		Project string
		Goal    string
	}
	//   type Goals struct{
	// 	  Goals string
	//   }
	var result []Result
	//from server
	var result1fromprojectofgoals []Result
	var goals []Goalfordbs
	var sql2 = `SELECT DISTINCT  project ,goal   FROM tasks  WHERE email ="` + email + `"` + " and" + ` goal  IN (SELECT name  FROM goalfordbs  WHERE   email =` + `"` + email + `"` + ")"

	//var goals []Tasksdb.Distinct("name", "age").Order("name, age desc").Find(&results)
	// var sql = fmt.Sprintf("SELECT  DISTINCT  project ,goal  FROM tasks WHERE email ='%s'",email)
	var sql1 = fmt.Sprintf("SELECT  DISTINCT  name FROM goalfordbs WHERE email ='%s'", email)
	fmt.Println(sql2)

	var sql3 = fmt.Sprintf("SELECT  DISTINCT  project,goal FROM projectofgoals WHERE status=unfinished and email ='%s'", email)

	db.Raw(sql2).Scan(&result)
	db.Raw(sql3).Scan(&result1fromprojectofgoals)
	db.Raw(sql1).Scan(&goals)
	fmt.Println(result)
	fmt.Println(sql1)
	fmt.Println(goals)

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

	allprojects_ingoal := make(map[string][]string)

	for _, item := range result {
		//    if item.Goal == ""{
		// 	alltasks_ingoal["no goal"]=append(alltasks_ingoal["no goal"],item)
		// 	devotedtime_for_goal["no goal"] = devotedtime_for_goal["no goal"] + item.Devotedtime

		// 	alltasks_ingoal[item.Goal]=append(alltasks_ingoal[item.Goal],item)
		// 	devotedtime_for_goal[item.Goal] = devotedtime_for_goal[item.Goal] + item.Devotedtime
		//    }
		// contains(arr,"plum")
		// if  (item.Goal != "")
		// _, found := Find(goals, item.Goal)

		// if found{
		// 	fmt.Println(item.Goal)
		if item.Goal != "" {
			allprojects_ingoal[item.Goal] = append(allprojects_ingoal[item.Goal], item.Project)
			// }

		}
	}

	//第二次循环用来检查
	for _, item := range result1fromprojectofgoals {
		//    if item.Goal == ""{
		// 	alltasks_ingoal["no goal"]=append(alltasks_ingoal["no goal"],item)
		// 	devotedtime_for_goal["no goal"] = devotedtime_for_goal["no goal"] + item.Devotedtime

		// 	alltasks_ingoal[item.Goal]=append(alltasks_ingoal[item.Goal],item)
		// 	devotedtime_for_goal[item.Goal] = devotedtime_for_goal[item.Goal] + item.Devotedtime
		//    }
		// contains(arr,"plum")
		// if  (item.Goal != "")
		// _, found := Find(goals, item.Goal)

		// if found{
		// 	fmt.Println(item.Goal)
		if item.Goal != "" {
			_, found := Find(allprojects_ingoal[item.Goal], item.Project)
			if !found {
				allprojects_ingoal[item.Goal] = append(allprojects_ingoal[item.Goal], item.Project)
			}
			// }

		}
	}
	return allprojects_ingoal
	// db.Where("email =  ?", email).Select("goal", "project").Scan(&tasks)
	//db.Where("email =  ?", email).Where("project =  ?", "goal").Not("status", []string{"finished","f","finish","giveup","g"}).Order("id").Find(&goals)
	// db.Where("email =  ?", email).Order("id").Find(&goals)
	// c.JSON(200, gin.H{
	// 	"goals": allprojects_ingoal,
	// })

}

func Get_goal_coffient(goal string, email string) float64 {
	var total_priotity = 0
	var goals_belonged2email []Goalfordbs
	db.Where("Email= ?", email).Find(&goals_belonged2email)
	for _, item := range goals_belonged2email {
		total_priotity = total_priotity + item.Priority
	}
	var querygoal Goalfordbs
	db.Where("Email= ?", email).Where("Name = ?", goal).Find(&querygoal)
	goal_coffient := float64(querygoal.Priority) / float64(total_priotity)
	fmt.Println("-------------------------goal coeffient-------------------------")
	fmt.Println("hahhahh")
	fmt.Println(goal_coffient)
	return goal_coffient
}

func Goalsjson(c *gin.Context) {
	/*
	   【A】这里生成goal的逻辑
	   1.获取所有没有完成的任务
	   2.将任务当中的goal抽取出来，组成以goal为key，任务数组为value的字典
	   3.loop 3中的每个key，然后提取任务的projetc生成三级goal
	   4.最后补充前面没有的goal，










	*/

	//the algorithm can be upgrade
	//i use email as identifier
	//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	email := emailcookie.Value
	var tasks []Tasks
	//fmt.Println(cookie1.Value)
	//email:="yangming1"
	// db.Where("Email= ?", email).Find(&tasks)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	today := time.Now().In(loc).Format("060102")
	tomorrow := time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
	//db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime in (?)", []string{today, tomorrow}).Order("id desc").Find(&tasks)
	statusbool := c.Query("statusbool")
	//https://blog.csdn.net/u013474436/article/details/71516649
	time_range := c.Request.Header.Get("time-range")
	//    fmt.Println(time_range)
	//    fmt.Println("status bool is ",statusbool)
	//    get   all days in this week or month
	week_day := Getweekday()
	month_day := Getmonthday()
	if statusbool == "yes" {
		fmt.Println("******************")
		if time_range == "week" {

			db.Where("Email= ?", email).Where("finishtime in (?)", week_day).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		} else if time_range == "month" {
			db.Where("Email= ?", email).Where("finishtime in (?)", month_day).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		} else {
			db.Where("Email= ?", email).Where("status in (?)", []string{"finished", "finish"}).Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		}

	} else {
		fmt.Println("----------------------")
		if time_range == "week" {
			db.Where("Email= ?", email).Where("finishtime in (?)", week_day).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today, tomorrow}).Order("id desc").Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		} else if time_range == "month" {
			db.Where("Email= ?", email).Where("finishtime in (?)", month_day).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today, tomorrow}).Order("id desc").Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		} else {
			db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Not("plantime", []string{today, tomorrow}).Order("id desc").Not("goal", []string{"no goal", ""}).Order("id desc").Find(&tasks)
		}
	}

	client := c.Request.Header.Get("client")
	querytype := c.Query("type")

	fmt.Println("+++++++client is++++++++")
	fmt.Println(querytype)
	fmt.Println(client)
	fmt.Println("+++++++client is++++++++")
	//use maps to aviod to design complex algorithm

	alltasks_ingoal := make(map[string][]Tasks)
	devotedtime_for_goal := make(map[string]int)

	//这里在组建以goal为目标key然后以数组内容task的字典
	for _, item := range tasks {
		if item.Goal == "" {
			alltasks_ingoal["no goal"] = append(alltasks_ingoal["no goal"], item)
			devotedtime_for_goal["no goal"] = devotedtime_for_goal["no goal"] + item.Devotedtime
		} else {
			alltasks_ingoal[item.Goal] = append(alltasks_ingoal[item.Goal], item)
			devotedtime_for_goal[item.Goal] = devotedtime_for_goal[item.Goal] + item.Devotedtime
		}

	}

	//这里参考上面的结构体就是上面的目标里面包含任务的那种
	var allgoalsonlyincludetasks []Goalsincludetasks
	for k, v := range alltasks_ingoal {
		var goal Goalfordbs
		var goalcountforsamegoal int
		var goal_level = 0
		// find product with id 1
		// db.First(&goal, "Name = ?", "L1212")
		//搜索具体的目标
		db.Where("Email= ?", email).Where("Name=?", k).Find(&goal).Count(&goalcountforsamegoal)
		if goalcountforsamegoal == 0 {
			print(goal_level)
			goal_level = 0
		} else {
			goal_level = goal.Priority
			if k == "no goal" {
				goal_level = 0
			}

		}
		allgoalsonlyincludetasks = append(allgoalsonlyincludetasks, Goalsincludetasks{Name: k, Goalcode: goal.Goalcode, Priority: goal_level, Devotedtime: devotedtime_for_goal[k], Alltasksingoal: v})

	}
	fmt.Println("========i am here 1========")
	if querytype == "noproject" {
		sort.Slice(allgoalsonlyincludetasks, func(i, j int) bool {
			if allgoalsonlyincludetasks[i].Priority == allgoalsonlyincludetasks[j].Priority {
				return allgoalsonlyincludetasks[i].Name < allgoalsonlyincludetasks[j].Name
			}
			return allgoalsonlyincludetasks[i].Priority > allgoalsonlyincludetasks[j].Priority
		})
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoalsonlyincludetasks})
		fmt.Println("========i am here========")
		return
	}

	fmt.Println("========i am here 3========")

	//这里提取所有的project！！！！
	//获取所有目标字符串的数组

	//alltasks_ingoal  目标里包含的所有任务

	var goals []string
	for k := range alltasks_ingoal {
		goals = append(goals, k)
	}
	fmt.Println(len(goals))

	//-------这里的代码非常的危险-------------
	//在末尾进行检测是否空的goals被检测到
	//在这里查询新创建的目标来不及创建目标的那种
	//查询所有对应的goal

	type Result struct {
		Name       string
		Goalcode   string
		Priority   int
		Goalstatus string
	}

	//这里因该在补充goals
	var goalsfinal []Result
	db.Raw(`SELECT name,goalcode,priority,goalstatus  FROM goalfordbs  WHERE email ="` + email + `"` + " and goalstatus not in (" + `"giveup","g","finished","finish"` + `)`).Scan(&goalsfinal)
	color.Red("red")

	for i := 0; i < len(goalsfinal); i++ {
		_, found := Find(goals, goalsfinal[i].Name)
		if !found {
			// var d []Projects
			// d := make([]Projects, 0)
			// goalmapproject[goals[i].Name] = d
			goals = append(goals, goalsfinal[i].Name)
			// allprojects_ingoal[item.Goal] = append(allprojects_ingoal[item.Goal], item.Project)
		}
	}

	//如果这里还有空的project也需要处理的！！！

	//这里可能包含具体的projects
	allclassgoals := make(map[string][]Projects)
	//这里的goals不包含所有的goal！！！！
	for _, singlegoal := range goals {
		allclassproject := make(map[string][]Tasks)
		//这里原来决定了他们提取的key所包含的project
		for _, item := range alltasks_ingoal[singlegoal] {
			//projects -- goal ,firsr to generate project and then goals
			allclassproject[item.Project] = append(allclassproject[item.Project], item)
		}

		//从project的表载入其他未完成的project，但是里面不包含空任务
		/*
			1.获取到目标包含的project
			2.检查project是否在其中，不再其中生成空的
		*/

		var projectsofgoalfromptojectstable = Getprojectofgoal(email)
		//从已经生的这里获取name
		// projectnames := reflect.ValueOf(allclassproject).MapKeys()
		projectnames := make([]string, 0, len(allclassproject))
		for k := range allclassproject {
			projectnames = append(projectnames, k)
		}

		color.Red("77777777777777777777")
		fmt.Println(projectsofgoalfromptojectstable)
		for _, singleproject := range projectsofgoalfromptojectstable[singlegoal] {
			//检测singleproject 是否在allclassproject当中，如果不在就加空的数组
			_, found := Find(projectnames, singleproject)
			if !found {
				var k []Tasks
				allclassproject[singleproject] = k
			}
		}

		//这里是在整理一个目标的所有project
		var allprojects []Projects
		for k, v := range allclassproject {
			allprojects = append(allprojects, Projects{k, v})

		}

		sort.Slice(allprojects, func(i, j int) bool {
			return allprojects[i].Name > allprojects[j].Name
		})

		color.Red("--------在这里再次检测-------是否allproject为nil")
		fmt.Println(singlegoal)
		fmt.Println(allprojects)
		//这里已经实在目标层面了，就是将最终的project塞进去
		if len(allprojects) == 0 {
			d := make([]Projects, 0)
			allclassgoals[singlegoal] = d
		} else {
			allclassgoals[singlegoal] = allprojects
		}

	}

	var allgoals []Goals
	for k, v := range allclassgoals {
		var goal Goalfordbs
		var goalcountforsamegoal int
		var goal_level = 0
		// find product with id 1
		// db.First(&goal, "Name = ?", "L1212")

		db.Where("Email= ?", email).Where("Name=?", k).Find(&goal).Count(&goalcountforsamegoal)

		fmt.Println(goal.Name)
		fmt.Println(goal.Priority)

		if goalcountforsamegoal == 0 {
			print(goal_level)
			goal_level = 0
		} else {
			goal_level = goal.Priority
			if k == "no goal" {
				goal_level = 0
			}

		}
		allgoals = append(allgoals, Goals{k, goal_level, goal.Goalcode, v})

	}

	// var goalsfinal []Goalfordbs
	// db.Where("Email= ?", email).Where("").Find(&goalsfinal)
	//提取所有的keys
	// exsiting_goals := reflect.ValueOf(goalmapproject).MapKeys()
	// exsiting_goals := make([]string, len(goalmapproject))
	var exsiting_goals []string
	for k := 0; k < len(allgoals); k++ {
		exsiting_goals = append(exsiting_goals, allgoals[k].Name)
	}

	for i := 0; i < len(goalsfinal); i++ {
		_, found := Find(exsiting_goals, goalsfinal[i].Name)
		if !found {
			// var d []Projects
			d := make([]Projects, 0)
			// goalmapproject[goals[i].Name] = d
			allgoals = append(allgoals, Goals{goalsfinal[i].Name, goalsfinal[i].Priority, goalsfinal[i].Goalcode, d})
			// allprojects_ingoal[item.Goal] = append(allprojects_ingoal[item.Goal], item.Project)
		}
	}

	//------------------------这里写的代码非常的危险---------

	sort.Slice(allgoals, func(i, j int) bool {
		if allgoals[i].Priority == allgoals[j].Priority {
			return allgoals[i].Name < allgoals[j].Name
		} else {
			return allgoals[i].Priority > allgoals[j].Priority
		}

	})

	//fmt.Println(allclassproject["gtd1"])
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "goals": allgoals})

}
