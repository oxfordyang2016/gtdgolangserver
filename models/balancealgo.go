package models

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	// "math"
	//"github.com/jinzhu/gorm"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tidwall/gjson"
	//"github.com/gin-contrib/sessions"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/lib/pq"
)

type Tasktagfortimescount struct {
	gorm.Model
	Email string `json:"email"`
	//Details string `json:"details" sql:"type:text;"`
	Date                 string `json:"date"`
	Importanttag         int    `json:"importanttag" gorm:"default:0"`
	Feartag              int    `json:"feartag" gorm:"default:0"`
	Principletag         int    `json:"principletag" gorm:"default:0"`
	Atomtag              int    `json:"atomtag" gorm:"default:0"`
	Dfstag               int    `json:"dfstag" gorm:"default:0"`
	Hardtag              int    `json:"hardtag" gorm:"default:0"`
	Easytag              int    `json:"easytag" gorm:"default:0"`
	Urgenttag            int    `json:"urgenttag" gorm:"default:0"`
	Challengetag         int    `json:"challengetag" gorm:"default:0"`
	Uncomfortabletag     int    `json:"uncomfortabletag" gorm:"default:0"`
	Troublesometag       int    `json:"troublesometag" gorm:"default:0"`
	Keyproblemtag        int    `json:"keyproblemtag" gorm:"default:0"`
	Tenminutestasktag    int    `json:"tenminutestasktag" gorm:"default:0"`
	Thirtyminutestasktag int    `json:"thirtyminutestasktag" gorm:"default:0"`
	Threeminutestasktag  int    `json:"threeminutestasktag" gorm:"default:0"`
}

func Balancealgo(c *gin.Context) {
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	Allmarked_factors(emailcookie.Value)
	c.HTML(http.StatusOK, "rebalancealgo.html", nil)

}

func Rebalancealgo(c *gin.Context) {
	emailcookie, _ := c.Request.Cookie("email")
	fmt.Println(emailcookie.Value)
	var k = Allmarked_factors(emailcookie.Value)
	// var k = Allmarked_factors("yang756260386@gmail.com")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "allmarked": k})

}

// 哪些目标是已经被投入过的
func Allmarked_factors(email string) map[string][]string {
	var tags_reviewalgo = []string{"attackactively", "makeuseofthings", "depthfirstsearch", "markataskimmediately", "alwaysprofit", "brainuse", "useprinciple", "patience", "solveakeyproblem", "learnnewthings", "battlewithlowerbrain", "difficultthings", "acceptfactandseektruth", "dealwithahardtask", "threeminutes", "serviceforgoal", "buildframeandprinciple", "atomadifficulttask", "noflinch", "learntechuse", "doanimportantthingearly", "acceptpain", "conquerthefear", "importanttag", "feartag", "principletag", "atomtag", "dfstag", "hardtag", "easytag", "urgenttag", "challengetag", "uncomfortabletag", "troublesometag", "keyproblemtag", "tenminutestasktag", "thirtyminutestasktag", "threeminutestasktag"}
	fmt.Println(tags_reviewalgo)
	type Result struct {
		Name     string
		Goalcode string
		Priority int
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var result []Result
	today := time.Now().In(loc).Format("060102")

	db.Raw(`SELECT name,goalcode,priority  FROM goalfordbs  WHERE email ="` + email + `"` + ` and ` + ` name   IN (SELECT goal  FROM tasks  WHERE  status in (` + `"giveup","g","finished","finish"` + `)` + `and  finishtime=` + `"` + today + `"` + ` and email =` + `"` + email + `"` + `);`).Scan(&result)
	var allgoals_finished = []string{}

	for i, _ := range result {
		allgoals_finished = append(allgoals_finished, result[i].Name)
	}
	color.Red("-------yangming-------")
	fmt.Println(allgoals_finished)

	// db.Raw(`SELECT *  FROM reviewfortimescounts WHERE plantime ="` + today + `" and email ="` + email + `"`).Scan(&result)

	var timescountsofemail Reviewfortimescount

	db.Where("Email= ?", email).Where("Date= ?", today).First(&timescountsofemail)

	// structValue := FooBar{Foo: "foo", Bar: 10}
	fields := reflect.TypeOf(timescountsofemail)
	values := reflect.ValueOf(timescountsofemail)

	num := fields.NumField()

	var allmarked_review = []string{}
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)

		fmt.Print("Type:", reflect.TypeOf(field.Type), ",", field.Name, "=", value, "\n")
		switch value.Kind() {
		case reflect.Int:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "wecan\n")
			if value.Int() > 0 {
				allmarked_review = append(allmarked_review, field.Name)
			}

		}
		// if field.Type == int {
		// 	if value > 0 {
		// 		allmarked_review = append(allmarked_review, field.Name)
		// 	}
		// }
	}

	allmarked_factors := make(map[string][]string)
	allmarked_factors["goal"] = allgoals_finished
	allmarked_factors["review"] = allmarked_review
	allmarked_factors["tag"] = ChecktagsFinished(email, []string{today})
	return allmarked_factors
}

func ChecktagsFinished(email string, dates []string) []string {
	// 这个是用来检查已经完成的tag在具体的时间范围内
	var tasks []Tasks
	//email:="yangming1"
	//use http://doc.gorm.io/crud.html#query
	//the next line is for all tasks
	//db.Where("Email= ?", email).Order("id desc").Find(&tasks)
	//the next line is for all unfinished task
	db.Where("Email= ?", email).Where("plantime in (?)", dates).Where("status in (?)", []string{"finish", "finished"}).Order("id desc").Find(&tasks)

	// var Importanttag = 0
	// var Feartag = 0
	// var Principletag = 0
	// var Atomtag = 0
	// var Dfstag = 0
	// var Hardtag = 0
	// var Easytag = 0
	// var Urgenttag = 0
	// var Challengetag = 0
	// var Uncomfortabletag = 0
	// var Troublesometag = 0
	// var Keyproblemtag = 0
	// var Tenminutestasktag = 0
	// var Thirtyminutestasktag = 0
	// var Threeminutestasktag = 0
	var alltag_tobedeal = []string{"importanttag", "importanttag", "feartag", "principletag", "atomtag", "dfstag", "hardtag", "easytag", "urgenttag", "challengetag", "uncomfortabletag", "troublesometag", "keyproblemtag", "tenminutestasktag", "thirtyminutestasktag", "threeminutestasktag"}
	var allmarked_tag = []string{}
	for i := 0; i < len(tasks); i++ {
		var thistask = tasks[i]
		var tags = thistask.Tasktags
		color.Yellow("---------分割线--------")

		for i := 0; i < len(alltag_tobedeal); i++ {
			thistag := gjson.Get(tags, alltag_tobedeal[i])
			if thistag.String() == "yes" {
				_, found := Find(allmarked_tag, alltag_tobedeal[i])
				if !found {
					allmarked_tag = append(allmarked_tag, alltag_tobedeal[i])
				}

			}
		}

	}

	return allmarked_tag
}
