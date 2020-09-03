package models
import(
  "fmt"

  "os"
  //"net/http"
	//"strconv"
//"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
)



// createTodo add a new todo
func Freewriting(c *gin.Context) {
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  //email:=emailcookie.Value
  freewriting := c.PostForm("freewriting")
  fmt.Println(freewriting)
  f, err := os.OpenFile("freewriting/freewriting.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
      panic(err)
  }
  defer f.Close()

  if _, err = f.WriteString(freewriting); err != nil {
      panic(err)
  }
  c.JSON(200, gin.H{
			"status": "posted",
		})
	}
