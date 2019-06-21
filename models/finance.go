package models
import(
  //"fmt"
  //"os/exec"
  //"os"
  //"net/http"
	//"strconv"
//"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
)








func Getmywealth (c *gin.Context){

  var balance = 1520
  c.JSON(200, gin.H{
						"status": "blog had updated",
						"blance":balance,
                })
}

//record spending   of everyday
func Create_fee(c *gin.Context){
  var balance = 1520
  c.JSON(200, gin.H{
						"status": "blog had updated",
						"blance":balance,
                })
}

//record the debt   


//record the income 









