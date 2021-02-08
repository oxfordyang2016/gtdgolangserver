
package models
import(
"fmt"
"bytes"
// "time"
// "sort"
//"github.com/bradfitz/slice"
"encoding/json"
"net/http"
// "github.com/fatih/color"
// "github.com/jinzhu/gorm"
// //"strconv"
// //"github.com/jinzhu/gorm"
// "github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"
// "github.com/tidwall/gjson"
//"github.com/tidwall/gjson"
)

func push2scheduler(url string,user string,taskid string,jobid string,task string ,time [6]int,typeofaction string){
values := map[string]interface{}{"taskid":taskid,"user":user,"task":task,"time":time,"jobid":jobid,"typeofaction":typeofaction}
jsonValue, _ := json.Marshal(values)

resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
if (resp!=nil){
	fmt.Println("ok")
}
// fmt.Println(resp)
}





