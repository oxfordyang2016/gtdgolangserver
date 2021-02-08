package main

/*GO*/
//https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
//https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
//the design mode learn from  https://github.com/beego/samples/blob/master/todo/models/task.go
//the import package learn from https://golang.org/doc/code.html
//please attention use  the things u donnot famialr with
//gopath https://github.com/golang/go/wiki/SettingGOPATH

/*database*/
//in order to keep db from losing ,i using db backup https://www.eversql.com/how-to-transfer-a-mysql-database-between-two-servers/

import (
	//"net/http"
	//"github.com/yangming/stringutil"
	//"fmt"
	// "os"
	// "io"
	//"./math"

	//"./math"

	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//"github.com/gin-contrib/sessions"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// the  . https://www.golang-book.com/books/intro/11
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/gtdgolangserver/models"
)

var db *gorm.DB

//about init https://stackoverflow.com/questions/24790175/when-is-the-init-function-run

//about init https://stackoverflow.com/questions/24790175/when-is-the-init-function-run

func main() {
	//golang https://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html
	//golang import var
	//fmt.Println(Modeltest)

	// f, err := os.Create("/tmp/dat2")

	// f, _ := os.Create("engine.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.Use(cors.Default())
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
	// 	if err, ok := recovered.(string); ok {
	// 		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	// 	}
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// }))

	v1 := router.Group("/v1")

	router.LoadHTMLGlob("templates/*.html")
	//i using https://gitlab.com/gin-gonic/gin/issues/358 to server static css
	v1.StaticFile("/show.css", "./static/css/show.css")
	v1.StaticFile("/user.css", "./static/css/user.css")
	v1.StaticFile("/taskmanage.css", "./static/css/taskmanage.css")
	v1.StaticFile("/moment.js", "./static/js/moment.js")
	v1.StaticFile("/tasks.js", "./static/js/tasks.js")

	v1.StaticFile("/echart.js", "./static/js/echart.js")
	v1.StaticFile("/finance.js", "./static/js/finance.js")
	v1.StaticFile("/jquery.js", "./static/js/jquery.js")
	v1.StaticFile("/chosen.js", "./static/js/chosen.js")
	v1.StaticFile("/chosen.css", "./static/css/chosen.css")
	v1.StaticFile("/style.css", "./static/css/style.css")
	v1.StaticFile("/index.js", "./static/js/index.js")
	v1.StaticFile("/balancealgo.js", "./static/js/balance.js")
	v1.StaticFile("/background.png", "./static/images/background.png")

	//test api
	v1.GET("/test", Test)
	v1.GET("/time", Clock)
	//user system
	v1.GET("/", User)
	v1.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})
	v1.GET("/location", Canvas)
	v1.GET("/emailverify", EmailGenerateCode)
	v1.POST("/login", Login)
	v1.POST("/register", Register)

	//operate system command
	//for blog update
	v1.GET("/blogupdate", Blogupdate)

	//json API
	v1.GET("/inboxjson", Inboxjson)
	v1.GET("/todayjson", Todaytaskjson)
	v1.GET("/taskmanager", Taskmanager)
	v1.GET("/yesterdayjson", Yesterdaytaskjson)
	v1.GET("/tomorrowjson", Tomorrowtaskjson)
	v1.GET("/unfinishedtasksjson", Unfinishedtaskjson)
	v1.GET("/readinglist", Readinglistjson)
	v1.GET("/habitlist", Habitlistjson)
	v1.GET("/healthlist", Healthlistjson)
	v1.GET("/financelist", Financelistjson)
	v1.GET("/projectsjson", Projectsjson)
	// 这里是评价组
	v1.GET("/reviewsjson", Reviewsjson)
	v1.GET("/reviewdaydatajson", Reviewalgorithmjson)
	v1.GET("/reviewdaydatajsonforyangming", Reviewalgorithmjsonforyangming)
	v1.GET("/reviewdaydatajsonforios", Reviewalgorithmjsonforios)
	// 这里是获取今天的分数
	v1.GET("/reviewscoreoftoday", Reviewscore_today)

	//平衡算法系统
	v1.GET("/balancealgo", Balancealgo)
	v1.GET("/rebalance", Rebalancealgo)
	//tasks system
	v1.POST("/createtask", CreatetaskbyJSON)
	//v1.POST("/gtdcli",Createtask)
	v1.POST("/createtaskfromsiri", Createtaskfromsiri)
	//v1.POST("/gtdclifromios",Createtaskfromios)
	v1.POST("/update", Update)
	v1.POST("/giveuptasksbatch", Giveuptasksbatch)
	v1.POST("/tomorrowtasksbatch", Tomorrowtasksbatch)
	v1.POST("/updateforios", Updateforios)
	v1.GET("/map", Googlemapservice)
	v1.GET("/deadlinesystemjson", Deadlinesystem)
	v1.POST("/taskexecute", Taskexecutelogfun)
	//principle system
	v1.POST("/createprinciplelib", Createprinciplelib)
	v1.POST("/createprincipledetail", Createprincipledetail)
	v1.POST("/connectpcodewithtasktag", Connectpcodewithtasktag)
	// v1.POST("/Createprincipledetail",Createprincipledetail)
	v1.GET("/getprinciple", Principlesystem)
	v1.GET("/getprincipledetails", Principlesdetailsystem)
	v1.GET("/searchwithprinciplecode", Searchwithprinciplecode)

	//goalsystem
	v1.POST("/createproject", Createprojectofgoal)
	v1.POST("/updateproject", Updateprojectofgoal)
	v1.POST("/creategoal", Creategoal)
	v1.POST("/updategoal", Updategoal)
	v1.GET("/comparegoal", Goalcompare)
	v1.GET("/searchwithgoalcode", Searchwithgoalcode)
	v1.GET("/goaljson", Goalsystem)
	v1.GET("/goalsgraph", Goalsjson)
	v1.GET("/goallist", Goallist)
	v1.GET("/goalincludeproject", Goalincludeprojectforserver)
	//web page
	//v1.GET("/mainboard",Mainboard)
	v1.GET("/inbox", Inbox)
	v1.GET("/project", Project)
	v1.GET("/everyday", Everydays)
	v1.GET("/pride", Finished)
	v1.GET("/place", Placebased)
	v1.GET("/freewriting", Freewriting)
	v1.GET("/todaytasks", Todaytaskweb)
	//get the tree map
	v1.GET("/review", Review)
	v1.GET("/reviewfrommaster", Reviewfromyangming)
	v1.GET("/reviewfortimes", Reviewforstastics)
	v1.GET("/errorlogjson", Errorlog)
	v1.GET("/problemsjson", Problemssystem)
	v1.GET("/questionsjson", Questionssystem)

	v1.GET("/reviewgraphforios", Reviewforios)
	v1.GET("/search", Search)
	v1.GET("/searchwithtag", Searchwithtags)
	//finance
	v1.GET("/mybalancejson", Getmywealth)
	//v1.POST("/createfee",CreatefeebyJSON)
	router.Run(":8081")

}
