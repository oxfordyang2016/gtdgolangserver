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
	//"./math"
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




	router := gin.Default()

v1 := router.Group("/v1")


  router.LoadHTMLGlob("templates/*.html")
	//i using https://gitlab.com/gin-gonic/gin/issues/358 to server static css
v1.StaticFile("/show.css", "./static/css/show.css")
v1.StaticFile("/user.css", "./static/css/user.css")

v1.StaticFile("/moment.js", "./static/js/moment.js")

v1.StaticFile("/echart.js", "./static/js/echart.js")
v1.StaticFile("/finance.js", "./static/js/finance.js")
v1.StaticFile("/jquery.js", "./static/js/jquery.js")
v1.StaticFile("/style.css", "./static/css/style.css")
 v1.StaticFile("/index.js", "./static/js/index.js")
v1.StaticFile("/background.png", "./static/images/background.png")

	 //test api
  v1.GET("/test", Test)
   v1.GET("/time",Clock)
	//user system
  v1.GET("/", User)
	v1.POST("/login",Login)
	v1.POST("/register",Register)



  //operate system command 
  //for blog update
  v1.GET("/blogupdate", Blogupdate) 





  //json API
  v1.GET("/inboxjson",Inboxjson)
  v1.GET("/todayjson",Todaytaskjson)
  v1.GET("/yesterdayjson",Yesterdaytaskjson)
  v1.GET("/tomorrowjson",Tomorrowtaskjson)
  v1.GET("/unfinishedtasksjson",Unfinishedtaskjson)
  v1.GET("/readinglist",Readinglistjson) 
  v1.GET("/habitlist",Habitlistjson)
  v1.GET("/healthlist",Healthlistjson)
  v1.GET("/financelist",Financelistjson)
  v1.GET("/projectsjson",Projectsjson)
  // 这里是评价组
  v1.GET("/reviewsjson",Reviewsjson)
  v1.GET("/reviewdaydatajson", Reviewalgorithmjson)  
  v1.GET("/reviewdaydatajsonforios", Reviewalgorithmjsonforios)
  // 这里是获取今天的分数
  v1.GET("/reviewscoreoftoday", Reviewscore_today)
	


        //tasks system
        v1.POST("/createtask",CreatetaskbyJSON)
        //v1.POST("/gtdcli",Createtask)
  v1.POST("/createtaskfromsiri",Createtaskfromsiri) 
        //v1.POST("/gtdclifromios",Createtaskfromios)
	      v1.POST("/update",Update)
        v1.POST("/updateforios",Updateforios)	
        v1.GET("/map",Googlemapservice)
        v1.GET("/location",Canvas)
        v1.POST("/creategoal",Creategoal)
        v1.POST("/updategoal",Updategoal)
        //web page
	      //v1.GET("/mainboard",Mainboard)
        v1.GET("/inbox",Inbox)
        v1.GET("/project",Project)
        v1.GET("/everyday",Everydays)
        v1.GET("/pride",Finished)
        v1.GET("/place",Placebased)
        v1.GET("/freewriting",Freewriting)
        //get the tree map
        v1.GET("/review",Review)
        v1.GET("/reviewfortimes",Reviewforstastics)
        v1.GET("/errorlogjson",Errorlog)
        v1.GET("/problemsjson",Problemssystem)
        v1.GET("/questionsjson",Questionssystem)
        v1.GET("/goaljson",Goalsystem)
        v1.GET("/goalsgraph",Goalsjson)
        v1.GET("/reviewgraphforios",Reviewforios)
        v1.GET("/search",Search)
        v1.GET("/searchwithtag",Searchwithtags)
        //finance
        v1.GET("/mybalancejson",Getmywealth)
        router.Run(":8081")

}
